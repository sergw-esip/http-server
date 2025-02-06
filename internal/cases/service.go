package cases

import (
	"httpServer/internal/entities"
	"httpServer/internal/ports/http"
)

type EntityUseCases struct {
	Store http.EntityStore
}

func NewEntityUseCases(store http.EntityStore) *EntityUseCases {
	return &EntityUseCases{Store: store}
}

func (uc *EntityUseCases) CreateEntity(entity *entities.Entity) error {
	return uc.Store.Create(entity)
}

func (uc *EntityUseCases) DeleteEntity(id int) error {
	return uc.Store.Delete(id)
}

func (uc *EntityUseCases) ListEntities() ([]entities.Entity, error) {
	return uc.Store.List()
}
