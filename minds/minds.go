package minds

import (
	"encoding/json"
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"

	"github.com/ashish111333/minds-go-sdk/datasources"
)

type Mind struct {
	api            *api.RestApi
	project        string
	dss            *datasources.DataSources
	PromptTemplate string
	Name           string                 `json:"name"`
	ModeName       string                 `json:"model_name"`
	Provider       string                 `json:"provider"`
	CreatedAt      string                 `json:"created_at"`
	UpdatedAt      string                 `json:"updated_at"`
	Parameters     map[string]interface{} `json:"parameters"`
	Datasources    interface{}            `json:"datasources"`
}

func (mind *Mind) Update(mindConfig *Mind) error {
	var data map[string]interface{}
	if mindConfig.Datasources != "" {
	}
	if mindConfig.ModeName != "" {
		data["model_name"] = mindConfig.ModeName
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

func (mind *Mind) AddDatasource(ds interface{}) error {

	ds_name, err := mind.CheckDatasource(ds)
	if err != nil {
		fmt.Errorf("failed to check datasource: %w", err)
	}
	var data map[string]interface{}
	data["name"] = ds_name
	resp, err := mind.api.Post("/projects/"+mind.project+"/minds/"+"/datasources", data)
	if err != nil {
		return fmt.Errorf("failed to add datasource to mind : %w", err)
	}
	defer resp.Body.Close()

}

func (mind *Mind) DeleteDatasource(ds interface{}) {

}

func (mind *Mind) Completion() {

}

func (minds *Mind) Create(mindConfig *Mind, replace bool) (*Mind, error) {
	if replace {

	}
	var ds_names []string
	if mindConfig.Datasources != nil {

	}
	if mindConfig.Parameters == nil {

	}

	resp, err := minds.api.Post("/projects/"+minds.project+"/minds", mindConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create mind : %w", err)
	}
	defer resp.Body.Close()

}

func (minds *Mind) Drop(name string) error {
	resp, err := minds.api.Delete("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to create mind: %w", err)
	}

	return nil

}

func (minds *Mind) List() ([]Mind, error) {
	resp, err := minds.api.Get("/projects/"+minds.project+"/minds", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get  minds list : %w", err)
	}
	defer resp.Body.Close()
	var minds_slice []Mind
	var data []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	for _, md := range data {

		mind := Mind{
			ModeName:    md["model_name"].(string),
			Provider:    md["provider"].(string),
			CreatedAt:   md["created_at"].(string),
			UpdatedAt:   md["updated_at"].(string),
			Parameters:  md["parameters"].(map[string]interface{}),
			Datasources: md["datasources"].([]string),
		}
		minds_slice = append(minds_slice, mind)
	}

	return minds_slice, nil
}

func (minds *Mind) Get(name string) (*Mind, error) {
	resp, err := minds.api.Get("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get mind: %w ", err)
	}
	defer resp.Body.Close()
	var mind_data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(mind_data)
	if err != nil {
		return nil, fmt.Errorf("failed to create mind: %w", err)
	}
	return &Mind{
		ModeName:    mind_data["model_name"].(string),
		Provider:    mind_data["provider"].(string),
		CreatedAt:   mind_data["created_at"].(string),
		UpdatedAt:   mind_data["updated_at"].(string),
		Parameters:  mind_data["parameters"].(map[string]interface{}),
		Datasources: mind_data["datasources"].([]string),
	}, nil

}

func (minds *Mind) CheckDatasource(ds interface{}) (string, error) {
	var name string
	ds_val, ok := ds.(datasources.DataSource)
	if ok {
		name = ds_val.Name
		return name, nil
	}
	dc_val, ok := ds.(datasources.DatabaseConfig)
	if ok {
		ds, err := minds.dss.Get(dc_val.Name)
		if err != nil {
			return "", fmt.Errorf("failed to get Datasource : %w", err)
		}
		err = minds.dss.Create(&ds.DatabaseConfig, false)
		if err != nil {
			return "", fmt.Errorf("failed to create Datasource: %w", err)
		}
		name = dc_val.Name

	} else {

		return "", fmt.Errorf("unknown datasource")
	}

	return name, nil
}

// util function for client
func NewMindsClient(api *api.RestApi) *Mind {

	return &Mind{
		api:     api,
		project: "mindsdb",
		dss:     datasources.NewDatasources(api),
	}
}
