package entities

import (
	"fmt"
	"time"
)

// categoria com seus campos
type Category struct {
	Id          uint      `json:"id"`
	Nome        string    `json:"nome"`
	DataCreated time.Time `json:"create_data"`
	DataUpdate  time.Time `json:"update_data"`
}

// create category recommended
func CreateCategory(Nome string) (*Category, error) {
	category := Category{
		Nome:        Nome,
		DataCreated: time.Now(),
		DataUpdate:  time.Now(),
	}

	//validar
	if err := category.isValid(); err != nil {
		//retorna um erro
		return nil, err
	}

	return &category, nil
}

//Valida o nome

func (categ Category) isValid() error {
	//verificar o nome
	if len(categ.Nome) < 3 {
		return fmt.Errorf("lenght of the name is invalid please insert age less then %v", len(categ.Nome))
	}

	return nil
}
