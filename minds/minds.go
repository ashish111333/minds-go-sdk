package minds

import (
	"github.com/ashish111333/minds-go-sdk/datasources"
)

type Mind struct {
	ModeName    string
	provider    string
	CreatedAt   string
	UpdatedAt   string
	datasources *datasources.DataSources
}

func NewMind() {

}

func (mind *Mind) Update() {

}

func (mind *Mind) AddDatasource() {

}

func (mind *Mind) DeleteDatasource() {

}

func (mind *Mind) Completion() {

}
