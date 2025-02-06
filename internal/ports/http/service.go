package http

import "httpServer/internal/entities"

type EntityStore interface {
	Create(entity *entities.Entity) error
	Delete(id int) error
	List() ([]entities.Entity, error)
	Get(id int) (entities.Entity, bool)
}
