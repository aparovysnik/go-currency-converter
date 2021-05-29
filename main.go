package main

import (
	"github.com/aparovysnik/go-currency-converter/router"
)

func main() {
	r := router.Init()
	r.Run()
}
