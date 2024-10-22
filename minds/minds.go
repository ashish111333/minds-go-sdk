package minds

import (
	"fmt"

	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
	"github.com/ashish111333/minds-go-sdk/minds"
)

type Mind struct {
	ModeName    string                   `json:"model_name"`
	Provider    string                   `json:"provider"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
	Parameters  map[string]interface{}   `json:"parameters"`
	Datasources *datasources.DataSources `json:"datasources"`
}

func (mind *Mind) Update() {

}

func (mind *Mind) AddDatasource() {

}

func (mind *Mind) DeleteDatasource() {

}string

func (mind *Mind) Completion() {

}

type Minds struct {
	api api.RestApi

	project string
}

func NewMinds() *Minds {

	return &Minds{}
}

func (minds *Minds) Create(mindConfig *Mind, replace bool) (*Mind, error) {
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

}

func (minds *Minds) Drop(name string) error {
	_, err := minds.api.Delete("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		return fmt.Errorf("failed to create mind: %w", err)
	}
	return nil

}

func (minds *Minds) List() (*[]Mind, error) {
	_, err := minds.api.Get("/projects/"+minds.project+"/minds", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create minds: %w", err)
	}

}

func (minds *Minds) Get(name string) *Mind {

	resp,err:=minds.api.Get("/projects/"+minds.project+"/minds/"+name,nil)
	
	
	
}

func (minds *Minds) checkDatasource() {

}
