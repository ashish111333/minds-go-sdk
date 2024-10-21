package minds

import (
	"log"

	"github.com/ashish111333/minds-go-sdk/api"
	"github.com/ashish111333/minds-go-sdk/datasources"
)

type Mind struct {
	ModeName    string
	Provider    string
	CreatedAt   string
	UpdatedAt   string
	Parameters  map[string]interface{}
	Datasources *datasources.DataSources
}

func (mind *Mind) Update(name, modelName, promptTemplate string, params map[string]interface{}) {

}

func (mind *Mind) AddDatasource() {

}

func (mind *Mind) DeleteDatasource() {

}

func (mind *Mind) Completion() {

}

type Minds struct {
	api api.RestApi

	project string
}

func NewMinds() *Minds {

	return &Minds{}
}

func (minds *Minds) Create() {

}

func (minds *Minds) Drop(name string) {
	_, err := minds.api.Delete("/projects/"+minds.project+"/minds/"+name, nil)
	if err != nil {
		log.Printf("failed to delete mind : %v \n", err)
	}

}

func (minds *Minds) List() {
	data, err := minds.api.Get("/projects/"+minds.project+"/minds", nil)
	if err != nil {
		log.Printf("failed to get minds list : %v \n", err)
	}

}

func () Get(name string) {

}

func () checkDatasource() {

}
