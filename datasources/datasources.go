package datasources

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ashish111333/minds-go-sdk/api"
)

type DatabaseConfig struct {
	Name           string            `json:"name"`
	Engine         string            `json:"engine"`
	Description    string            `json:"description"`
	ConnectionData map[string]string `json:"connection_data,omitempty"`
	Tables         []string          `json:"tables"`
}

type DataSource struct {
	DatabaseConfig
}

type DataSources struct {
	Api *api.RestApi
}

// creates a new instance for Datasources struct
func NewDatasources(api *api.RestApi) *DataSources {

	return &DataSources{
		Api: api,
	}

}

// creates a new datasource
func (d *DataSources) Create(DsConfig *DatabaseConfig, replace bool) error {
	name := DsConfig.Name
	if replace {
		_, err := d.Get(name)
		if err != nil {
			log.Printf("failed to replace Datasource: %v \n", err)
		}
		err = d.Drop(name)
		if err != nil {
			return fmt.Errorf("failed to get Datasource")
		}
	}
	_, err := d.Api.Post("/datasources", DsConfig)
	if err != nil {
		log.Printf("failed to create datasource: %v \n", err)
		return err
	}
	return nil
}

func (d *DataSources) List() ([]DataSource, error) {
	resp, err := d.Api.Get("/datasources", nil)
	if err != nil {
		log.Printf("failed to get list of datasources: %v \n", err)
		return nil, err
	}
	defer resp.Body.Close()
	var data []map[string]interface{}
	var datasources []DataSource

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Printf("failed to decode json: %v \n", err)
		return nil, err
	}
	for _, v := range data {
		ds := DataSource{
			DatabaseConfig: DatabaseConfig{
				Name:           v["name"].(string),
				Engine:         v["engine"].(string),
				Description:    v["description"].(string),
				ConnectionData: v["connection_data"].(map[string]string),
				Tables:         v["tables"].([]string),
			},
		}
		datasources = append(datasources, ds)
	}
	return datasources, nil

}

func (d *DataSources) Get(name string) (*DataSource, error) {
	resp, err := d.Api.Get("/datasources"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get datsource: %w", err)
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json : %w", err)
	}
	if data["engine"] == "" {
		return nil, fmt.Errorf("wrong type of datasource: %ws", err)
	}
	return &DataSource{
		DatabaseConfig: DatabaseConfig{
			Name:           data["name"].(string),
			Engine:         data["engine"].(string),
			Description:    data["description"].(string),
			ConnectionData: data["connection_data"].(map[string]string),
			Tables:         data["tables"].([]string),
		},
	}, nil

}

func (d *DataSources) Drop(name string) error {
	_, err := d.Api.Delete("/datasources/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Datasource :%w", err)
	}
	return nil
}
