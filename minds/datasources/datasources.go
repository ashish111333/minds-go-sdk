package datasources

import (
	"log"

	"github.com/ashish111333/minds-go-sdk/minds"
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
	Api minds.RestApi
}

func (d *DataSources) create(DsConfig *DatabaseConfig, replace bool) {
	name = DsConfig.Name

	d.Api.Post("/datasources", DsConfig)

}
func (d *DataSources) list() {

}
func (d *DataSources) get(name string) DataSource {
	log.Printf("Making Get request Url:%s", "/datasources"+name)
	data, err := d.Api.Get("/datasources"+name, nil)
	if err != nil {
		log.Printf("failed to get datasource :", err)
	}

}

func (d *DataSources) drop(name string) {

	log.Printf("Making Delete request Url:%s", "/datasources/"+name)
	_, err := d.Api.Delete("/datasources/"+name, nil)
	if err != nil {
		log.Printf("failed to delete Datasource: ", err)
	}

}
