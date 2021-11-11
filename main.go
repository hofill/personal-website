package main

import (
	"hofill/repository"
	"hofill/services"
)

func main() {
	var repo repository.WriteUpRepository = &repository.InMemoryRepository{}
	services.PopulateRepository(repo)
	services.SetupRoutes(repo)
}
