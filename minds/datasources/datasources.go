package datasources

import "github.com/ashish111333/minds-go-sdk/minds"

type DatabaseConfig struct {
	Name           string            `json:"name"`
	Engine         string            `json:"engine"`
	Description    string            `json:"descriptiong"`
	ConnectionData map[string]string `json:"connection_data,omitempty"`
}

type DataSource struct {
	DatabaseConfig
}

type DataSources struct {
	client minds.Client
}

func () create() {

}
func () list() {

}
func () get() {

}

func () drop() {

}
