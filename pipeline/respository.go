package pipeline

import (
	"sync"
)

type (
	// Repository for pipeline.
	Repository interface {
		// Create a pipeline.
		Create(p *Pipeline) (*Pipeline, error)
	}

	// InMemoryRepository for pipeline.
	InMemoryRepository struct {
		pipelines []*Pipeline
		counter   uint64
		mu        sync.Mutex
	}
)

// NewRepository for pipeline.
func NewRepository() Repository {
	r := &InMemoryRepository{
		pipelines: make([]*Pipeline, 0),
	}

	return r
}

// Create a pipeline and set the identifier.
func (r *InMemoryRepository) Create(p *Pipeline) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.pipelines = append(r.pipelines, p)

	r.counter++
	p.ID = r.counter

	return p, nil
}
