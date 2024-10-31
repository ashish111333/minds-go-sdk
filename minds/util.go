package minds

import (
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
)

// checks if the given datasource is instance of DatabaseConfig or DataSource struct,
// if not returns error unknown Datasource
func checkDatasource(ds interface{}, dss *datasources.DataSources) (string, error) {

	var name string
	dsStruct, ok := ds.(datasources.DataSource)
	if ok {
		name = dsStruct.Name
		return name, nil
	}
	dsConfigStruct, ok := ds.(datasources.DatabaseConfig)
	if ok {
		ds, err := dss.Get(dsConfigStruct.Name)
		if err != nil {
			return "", fmt.Errorf("failed to get Datasource : %w", err)
		}
		err = dss.Create(&ds.DatabaseConfig, false)
		if err != nil {
			return "", fmt.Errorf("failed to create Datasource: %w", err)
		}
		name = dsConfigStruct.Name

	} else {
		return "", fmt.Errorf("unknown datasource")
	}

	return name, nil

}

// returns Minds client
func NewMindsClient(api *api.RestApi) *Mind {

	return &Mind{
		api:     api,
		project: "mindsdb",
		dss:     datasources.NewDatasourcesClient(api),
	}
}
