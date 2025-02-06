package http

import (
	"httpServer/internal/entities"
	"sync"
)

type Server struct {
	store map[int]entities.Entity
	mu    sync.Mutex
	idSeq int
}

func NewServer() *Server {
	return &Server{
		store: make(map[int]entities.Entity),
		mu:    sync.Mutex{},
		idSeq: 0,
	}
}

func (s *Server) Create(entity *entities.Entity) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.idSeq++
	entity.ID = s.idSeq
	s.store[entity.ID] = *entity
	return nil
}

func (s *Server) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, id)
	return nil
}

func (s *Server) List() ([]entities.Entity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entities := make([]entities.Entity, 0, len(s.store))
	for _, entity := range s.store {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (s *Server) Get(id int) (entities.Entity, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	entity, ok := s.store[id]
	return entity, ok
}
