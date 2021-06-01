package app

import (
	"log"
	"time"

	v1 "github.com/aparovysnik/go-currency-converter/api/v1"
	"github.com/aparovysnik/go-currency-converter/config"
	_ "github.com/aparovysnik/go-currency-converter/docs"
	"github.com/aparovysnik/go-currency-converter/middleware"
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/aparovysnik/go-currency-converter/services"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Currency Converter API
// @version 0.1
// @description An API to convert currencies

// @contact.name A Parovysnik
// @contact.email a.parovysnik@gmail.com

// @host localhost:17249
// @BasePath /
// @schemes http
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
	converterService := services.NewCurrencyConverter(conversionRateRepository)

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
	scheduler.Every("1h").Do(services.UpdateConversionRate, conversionRateRepository, config)
	scheduler.StartAsync()

	//Add Swagger
	swagUrl := ginSwagger.URL("http://localhost:17249/swagger/doc.json") // The url pointing to API definition
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagUrl))

	//Serve API
	ginEngine.Run()
}
