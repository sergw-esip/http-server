package adapters

import (
	"sync"

	"httpServer/internal/entities"
	"httpServer/internal/errors"
	"httpServer/internal/ports/http"
)

type InMemoryEntityStore struct {
	store map[int]entities.Entity
	mu    sync.Mutex
	idSeq int
}

func NewInMemoryEntityStore() *InMemoryEntityStore {
	return &InMemoryEntityStore{
		store: make(map[int]entities.Entity),
		mu:    sync.Mutex{},
		idSeq: 0,
	}
}

func (s *InMemoryEntityStore) Create(entity *entities.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.idSeq++
	entity.ID = s.idSeq
	s.store[entity.ID] = *entity
	return nil
}

func (s *InMemoryEntityStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.store[id]
	if !ok {
		return errors.ErrEntityNotFound
	}

	delete(s.store, id)
	return nil
}

func (s *InMemoryEntityStore) List() ([]entities.Entity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entities := make([]entities.Entity, 0, len(s.store))
	for _, entity := range s.store {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (s *InMemoryEntityStore) Get(id int) (entities.Entity, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entity, ok := s.store[id]
	return entity, ok
}

var _ http.EntityStore = (*InMemoryEntityStore)(nil)
