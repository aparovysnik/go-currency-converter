package main

import (
	"github.com/aparovysnik/go-currency-converter/repositories"
	"github.com/aparovysnik/go-currency-converter/router"
)

func main() {
	r := router.Init()
	repositories.SetupDB()
	r.Run()
}
