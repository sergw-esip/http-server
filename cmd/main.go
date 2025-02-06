package main

import (
	"httpServer/internal/adapters"
	"httpServer/internal/application"
	"httpServer/internal/cases"
	"log"
)

func main() {
	// Infrastructure setup (adapters)
	store := adapters.NewInMemoryEntityStore()

	// Application setup (cases + ports)
	useCases := cases.NewEntityUseCases(store)

	// Transport/Presentation setup (adapters)
	handler := adapters.NewHTTPHandler(useCases)

	// Application setup
	app := application.NewApp(handler)

	// Run the service
	if err := app.RunService(":8080"); err != nil {
		log.Fatalf("Failed to run service: %v", err)
	}
}
