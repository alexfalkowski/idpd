package pipeline

import (
	"sync"
)

type (
	// Repository for pipeline.
	Repository interface {
		// Get a pipeline.
		Get(id uint64) (*Pipeline, error)

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

// Get a pipeline.
func (r *InMemoryRepository) Get(id uint64) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range r.pipelines {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, ErrPipelineNotFound
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
