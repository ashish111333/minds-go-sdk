package datasources

import (
	"encoding/json"
	"fmt"

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
func NewDatasourcesClient(api *api.RestApi) *DataSources {

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
			return fmt.Errorf("failed to replace Datasource:%w", err)
		}
		err = d.Drop(name)
		if err != nil {
			return fmt.Errorf("failed to get Datasource")
		}
	}
	_, err := d.Api.Post("/datasources", DsConfig)
	if err != nil {
		return fmt.Errorf("failed to create datasource:%w", err)
	}
	return nil
}

func (d *DataSources) List() ([]DataSource, error) {
	resp, err := d.Api.Get("/datasources", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get list of datasources:%w", err)
	}
	defer resp.Body.Close()
	var datasources []DataSource

	err = json.NewDecoder(resp.Body).Decode(&datasources)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json:%w", err)
	}
	return datasources, nil

}

// Get's Datasource takes name as argument
func (d *DataSources) Get(name string) (*DataSource, error) {
	resp, err := d.Api.Get("/datasources/"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get datsource: %w", err)
	}
	defer resp.Body.Close()
	var datasource DataSource
	err = json.NewDecoder(resp.Body).Decode(&datasource)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json : %w", err)
	}
	if datasource.Engine == "" {
		return nil, fmt.Errorf("wrong type of datasource: %ws", err)
	}
	return &datasource, nil

}

func (d *DataSources) Drop(name string) error {
	_, err := d.Api.Delete("/datasources/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Datasource :%w", err)
	}
	return nil
}
