package datasources

import (
	"testing"
)

const apiKey string = "123456789abc"

var exampleDatasource = DataSource{DatabaseConfig: DatabaseConfig{

	Name:        "example_ds",
	Engine:      "postgres",
	Description: "minds example datasource",
	ConnectionData: map[string]string{
		"user":     "demo_user",
		"password": "demo_password",
		"host":     "samples.mindsdb.com",
		"port":     "5432",
		"database": "demo",
		"schema":   "demo_data",
	},
}}

func compareDs(Ds1 DataSource, Ds2 DataSource) bool {

}
func TestGetDatasource(t *testing.T) {

}
