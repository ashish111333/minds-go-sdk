package minds

import (
	"encoding/json"
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"

	"github.com/ashish111333/minds-go-sdk/datasources"
)

var defaultPromptTemplate = ""

type Mind struct {
	api            *api.RestApi
	project        string
	dss            *datasources.DataSources
	PromptTemplate string
	Name           string                  `json:"name"`
	ModelName      string                  `json:"model_name"`
	Provider       string                  `json:"provider"`
	CreatedAt      string                  `json:"created_at"`
	UpdatedAt      string                  `json:"updated_at"`
	Parameters     *map[string]interface{} `json:"parameters"`
	Datasources    []interface{}           `json:"datasources"`
}
type MindConfig struct {
	Name           string
	ModelName      string
	Provider       string
	PromptTemplate string
	Parameters     *map[string]interface{}
	Datasources    []interface{}
}

// updates a mind
func (mind *Mind) Update(mindConfig *MindConfig) error {
	data := make(map[string]interface{})
	if mindConfig.Datasources != nil {
		dsNames := []interface{}{}
		for _, ds := range mindConfig.Datasources {
			Name, err := checkDatasource(ds, mind.dss)
			if err != nil {
				return fmt.Errorf("failed to update mind: %w", err)
			}
			dsNames = append(dsNames, Name)

		}
		mindConfig.Datasources = dsNames

	}
	if mindConfig.ModelName != "" {
		data["model_name"] = mindConfig.ModelName
	}
	if mindConfig.Provider != "" {
		data["provider"] = mindConfig.Provider
	}
	if mindConfig.Parameters == nil {
		data["parameters"] = make(map[string]interface{})
	}

	resp, err := mind.api.Post("/projects/"+mind.project+"/minds/"+mind.Name, data)
	if err != nil {
		return fmt.Errorf("failed to update mind: %w", err)
	}
	defer resp.Body.Close()

	return nil
}

// adds datasource to a mind
func (mind *Mind) AddDatasource(ds interface{}) error {

	dsName, err := checkDatasource(ds, mind.dss)
	if err != nil {
		return fmt.Errorf("failed to check datasource: %w", err)
	}
	data := make(map[string]interface{})
	data["name"] = dsName
	resp, err := mind.api.Post("/projects/"+mind.project+"/minds/"+"/datasources", data)
	if err != nil {
		return fmt.Errorf("failed to add datasource to mind : %w", err)
	}
	defer resp.Body.Close()
	return nil
}

// deletes a datasource from a mind
func (mind *Mind) DeleteDatasource(ds interface{}) error {

	var dsName string
	dsStruct, ok := ds.(datasources.DataSource)
	if ok {
		dsName = dsStruct.Name
	}
	name, ok := ds.(string)
	if ok {
		dsName = name
	}
	if dsName == "" {
		return fmt.Errorf("unknown datasource")
	}

	resp, err := mind.api.Delete("/projects/"+mind.project+"minds"+"/datasources/"+dsName, nil)
	if err != nil {
		return fmt.Errorf("failed to delete datasource for mind : %w", err)
	}
	defer resp.Body.Close()
	return nil
}

func (mind *Mind) Completion(message string, stream bool) {

}

// Minds
type Minds struct {
	api     *api.RestApi
	project string
	dss     *datasources.DataSources
}

// takes MindConfig of type Mind and creates a Mind from it,
// Datasources can be type string,Datasource or DatabaseConfig any other type is rejected
func (minds *Minds) Create(mindConfig *Mind, replace bool) (*Mind, error) {
	if replace {

	}
	dsNames := []interface{}{}
	if mindConfig.Datasources != nil {
		for _, ds := range mindConfig.Datasources {
			Name, err := checkDatasource(ds, minds.dss)
			if err != nil {
				return nil, fmt.Errorf("failed to check datasource: %w", err)
			}
			dsNames = append(dsNames, Name)
		}
		mindConfig.Datasources = dsNames
	}
	if mindConfig.Parameters == nil {
		mindConfig.Parameters = &map[string]interface{}{}
	}
	if mindConfig.PromptTemplate == "" {
		mindConfig.PromptTemplate = defaultPromptTemplate
	}

	resp, err := minds.api.Post("/projects/"+minds.project+"/minds", mindConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create mind : %w", err)
	}
	defer resp.Body.Close()
	md, err := minds.Get(mindConfig.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get mind: %w", err)
	}
	return md, nil

}

// deletes the given mind
func (minds *Minds) Drop(name string) error {
	resp, err := minds.api.Delete("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mind: %w", err)
	}
	defer resp.Body.Close()
	return nil
}

// returns the List of Minds created by user
func (minds *Minds) List() (*[]Mind, error) {
	resp, err := minds.api.Get("/projects/"+minds.project+"/minds", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get  minds list : %w", err)
	}
	defer resp.Body.Close()
	var minds_list []Mind
	err = json.NewDecoder(resp.Body).Decode(&minds_list)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json :%w", err)
	}
	return &minds_list, nil
}

// returns Mind takes Name as argument err if it doesn't exist
func (minds *Minds) Get(name string) (*Mind, error) {
	resp, err := minds.api.Get("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get mind: %w ", err)
	}
	defer resp.Body.Close()
	var md Mind
	err = json.NewDecoder(resp.Body).Decode(&md)
	if err != nil {
		return nil, fmt.Errorf("failed to create mind: %w", err)
	}
	md.api = minds.api
	return &md, nil
}
