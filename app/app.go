package app

import (
	"log"
	"time"

	v1 "github.com/aparovysnik/go-currency-converter/api/v1"
	"github.com/aparovysnik/go-currency-converter/config"
	"github.com/aparovysnik/go-currency-converter/middleware"
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/aparovysnik/go-currency-converter/services"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
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

	//Read config
	config := config.Load()

	//Init repositories
	projectRepository := repositories.NewProjectRepository(db)
	conversionRateRepository := repositories.NewConversionRateRepository(db)

	//Init services
	projectService := services.NewProjectService(projectRepository)
	converterService := services.NewCurrencyConverter()

	//Init controllers
	v1.InitHealthController(ginEngine)
	v1.InitProjectController(ginEngine, projectService)

	authGroup := ginEngine.Group("/")
	authGroup.Use(middleware.Authenticate(projectRepository))
	{
		v1.InitCurrencyConversion(authGroup, converterService, config)
	}

	//Setup poll job
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every("30s").Do(services.UpdateConversionRate, conversionRateRepository, config)
	scheduler.StartAsync()

	//Serve API
	ginEngine.Run()
}
