package main

import (
	"hofill/services"
)

func main() {
	services.SetupRoutes()
	services.PopulateRepository()
}
