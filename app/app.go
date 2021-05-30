package app

import (
	"log"

	v1 "github.com/aparovysnik/go-currency-converter/api/v1"
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	//init DB
	db, err := repositories.SetupDB()

	if err != nil {
		log.Fatalf("DB setup error: %v", err)
	}

	//Init gin
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	//Init repositories
	projectRepository := repositories.NewProjectRepository(db)

	//Init controllers
	v1.InitHealthController(ginEngine)
	v1.InitProjectController(ginEngine, projectRepository)

	//Serve API
	ginEngine.Run()
}
