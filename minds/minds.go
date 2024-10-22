package minds

import (
	"encoding/json"
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"

	"github.com/ashish111333/minds-go-sdk/datasources"
)

type Mind struct {
	ModeName    string                 `json:"model_name"`
	Provider    string                 `json:"provider"`
	CreatedAt   string                 `json:"created_at"`
	UpdatedAt   string                 `json:"updated_at"`
	Parameters  map[string]interface{} `json:"parameters"`
	Datasources interface{}            `json:"datasources"`
}

func (mind *Mind) Update(mindConfig *Mind) {
	var data map[string]interface{}
	if mindConfig.Datasources != nil {

		_, ok := mindConfig.Datasources.([]string)
		_, ok = mindConfig.Datasources.(datasources.DatabaseConfig)

	}

}

func (mind *Mind) AddDatasource() {

}

func (mind *Mind) DeleteDatasource() {

}

func (mind *Mind) Completion() {

}

type Minds struct {
	Api     api.RestApi
	Project string
}

func NewMinds(api *api.RestApi) *Minds {

	return &Minds{
		Api:     *api,
		Project: "mindsdb",
	}
}

func (minds *Minds) Create(mindConfig *Mind, replace bool) (*Mind, error) {
	if replace {

	}
	var ds_names []string
	if mindConfig.Datasources != nil {

	}
	if mindConfig.Parameters == nil {

	}

	resp, err := minds.Api.Post("/projects/"+minds.Project+"/minds", mindConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create mind : %w", err)
	}
	defer resp.Body.Close()

}

func (minds *Minds) Drop(name string) error {
	_, err := minds.Api.Delete("/projects/"+minds.Project+"/minds/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to create mind: %w", err)
	}
	return nil

}

func (minds *Minds) List() (*[]Mind, error) {
	resp, err := minds.Api.Get("/projects/"+minds.Project+"/minds", nil)
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

	return &minds_slice, nil
}

func (minds *Minds) Get(name string) (*Mind, error) {
	resp, err := minds.Api.Get("/projects/"+minds.Project+"/minds/"+name, nil)
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

func (minds *Minds) checkDatasource(ds interface{}) string {

	var name string
	ds_val, ok := ds.(datasources.DataSource)
	if ok {
		name = ds_val.Name
		return name
	}
	dc_val, ok := ds.(datasources.DatabaseConfig)
	if ok {

		name = dc_val.Name

	}

}
