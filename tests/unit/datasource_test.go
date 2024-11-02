package unit

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ashish111333/minds-go-sdk/client"
	"github.com/ashish111333/minds-go-sdk/datasources"
)

const apiKey string = "123456789abc"

var exampleDatasource = datasources.DataSource{DatabaseConfig: datasources.DatabaseConfig{

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

func compareDs(Ds1 datasources.DataSource, Ds2 datasources.DataSource) bool {

	return reflect.DeepEqual(Ds1, Ds2)
}
func TestGetDatasource(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("expected GET request,got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/datasources/example_ds") {
			t.Fatal("wrong request url")
		}
		w.Header().Set("content-type", "application/json")
		data, err := json.Marshal(exampleDatasource)
		if err != nil {
			t.Fatalf("failed to encode to json")
		}
		w.Write(data)
	}))
	defer mockServer.Close()
	client, err := client.NewClient(apiKey, mockServer.URL)
	if err != nil {
		t.Fatal("failed to create Client: ", err)
	}
	ds, err := client.Datasources.Get("example_ds")
	if err != nil {
		t.Fatal("failed to get datasource")
	}

	if !compareDs(*ds, exampleDatasource) {
		t.Fatal("dataources recieved doesn't match")
	}

}

func TestDeleteDatasource(t *testing.T) {

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Fatalf("expected delete method, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/datasources/ds_name") {
			t.Fatal("wrong Request Url")
		}
	}))
	defer mockServer.Close()
	client, err := client.NewClient(apiKey, mockServer.URL)
	if err != nil {
		t.Fatal("failed to create client: ", err)
	}
	err = client.Datasources.Drop("ds_name")
	if err != nil {
		t.Fatal("failed to delete datasource")
	}

}

func TestCreateDatasource(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))

}
