package usecases

import (
	"fmt"
	"log"
	"microservice/internal/entities"
)

func (categUseCase useCaseCategory) CreateCategory(nome string) error {

	//criando uma caregora
	category, err := entities.CreateCategory(nome)

	//validndo erro
	if err != nil {

		return fmt.Errorf("Houve um erro na camada de use case, tentativa de criar categoria mas algo deu erro: ")
	}

	//TODO: persistir dados no banxo de dados
	log.Println(category)

	return nil
}
