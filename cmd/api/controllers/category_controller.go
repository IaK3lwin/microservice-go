package controllers

import (
	"log"
	usecases "microservice/internal/use_cases/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

// uma estrutura que vai esperar os campos que vem do request
type responseCategory struct {
	// essa estrutura representa a requisição vinda do client,
	//e a gente informa ao gin que esese campo é obrigatório
	//usando o binding passando um required, ou seja, internamente o gin valida essa requsiição
	//usando um validatorpersonalizado que ele cria
	Name string `json:"nome" binding:"required"`
}

func CategoryCreate(context *gin.Context) {

	//criando o body da requisiçã
	var bodyCateegory responseCategory

	// recebendo os dados da rwsuisiçãop em json
	err := context.BindJSON(&bodyCateegory)

	//tratar possivel errro de validação
	if err != nil {
		log.Println(err)                         //logando
		context.JSON(http.StatusBadRequest, err) //mostrando o erro como response
		return                                   // parando a execução por ai
	}

	//usecase rules do categories
	usecase := usecases.NewUseCaseCategory()

	//criando o nosso category no db
	err = usecase.CreateCategory(bodyCateegory.Name)

	//validando erro
	if err != nil {
		context.JSON(http.StatusOK, err)
		return
	}

	//se não gouver erro
	context.JSON(http.StatusCreated, "criado com sucesso!")

}
