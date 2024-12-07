package pipeline

import (
	"sync"
)

type (
	// Repository for pipeline.
	Repository interface {
		// Get a pipeline.
		Get(id ID) (*Pipeline, error)

		// Create a pipeline.
		Create(p *Pipeline) (*Pipeline, error)

		// Update a pipeline.
		Update(id ID, p *Pipeline) (*Pipeline, error)
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
func (r *InMemoryRepository) Get(id ID) (*Pipeline, error) {
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
	p.ID = ID(r.counter)

	return p, nil
}

// Update a pipeline.
func (r *InMemoryRepository) Update(id ID, pipeline *Pipeline) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, p := range r.pipelines {
		if p.ID == id {
			p = pipeline

			return p, nil
		}
	}

	return nil, ErrPipelineNotFound
}
