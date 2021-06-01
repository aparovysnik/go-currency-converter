package v1

import (
	"net/http"

	"github.com/aparovysnik/go-currency-converter/api/v1/models"
	"github.com/aparovysnik/go-currency-converter/config"
	"github.com/aparovysnik/go-currency-converter/services"
	"github.com/gin-gonic/gin"
)

type CurrencyConversion interface {
	Convert(*gin.Context) error
}

type currencyConversion struct {
	service services.CurrencyConverter
	config  config.Config
}

func InitCurrencyConversion(ginGroup *gin.RouterGroup, service services.CurrencyConverter,
	config config.Config) {
	controller := currencyConversion{
		service: service,
		config:  config,
	}
	ginGroup.GET("/currency/convert", controller.Convert)
}

func (controller *currencyConversion) Convert(c *gin.Context) {
	reqBody := new(models.ConvertCurrencyRequest)
	err := c.Bind(reqBody)

	if err != nil || !reqBody.IsValid(controller.config) {
		c.JSON(400, models.ErrorResponse{
			BaseResponse: models.BaseResponse{
				Status: 400,
			},
			Message: "Invalid request",
		})
		return
	}

	result, err := controller.service.Convert(reqBody.Amount, reqBody.FromCurrency, reqBody.ToCurrency)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			BaseResponse: models.BaseResponse{
				Status: http.StatusInternalServerError,
			},
			Message: "Something went wrong.",
		})
		return
	}

	c.JSON(200, models.ConvertCurrencyResponse{
		BaseResponse: models.BaseResponse{
			Status: 200,
		},
		Amount: result,
	})
}
