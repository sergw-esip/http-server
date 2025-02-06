package application

import (
	"fmt"
	"log"
	"net/http"

	"httpServer/internal/adapters"
)

type App struct {
	Handler *adapters.HTTPHandler
}

func NewApp(handler *adapters.HTTPHandler) *App {
	return &App{Handler: handler}
}

func (a *App) RunService(addr string) error {
	http.HandleFunc("/create", a.Handler.CreateEntity)
	http.HandleFunc("/delete", a.Handler.DeleteEntity)
	http.HandleFunc("/list", a.Handler.ListEntities)

	fmt.Println("Server started at", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
	return nil // this line won't be reached due to log.Fatal
}
