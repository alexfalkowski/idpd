package pipeline

import (
	"slices"
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

		// Delete a pipeline.
		Delete(id ID) (*Pipeline, error)
	}

	// InMemoryRepository for pipeline.
	//
	// The counter is used as a basic id generator, though an incrementing number is nit recommend as it easy to guess.
	// The mux is to make sure we don't accidentally corrupt or increment incorrectly.
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

	i, err := r.pipeline(id)
	if err != nil {
		return nil, err
	}

	return r.pipelines[i], nil
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

	i, err := r.pipeline(id)
	if err != nil {
		return nil, err
	}

	pipeline.ID = id
	r.pipelines[i] = pipeline

	return pipeline, nil
}

// Delete a pipeline.
func (r *InMemoryRepository) Delete(id ID) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	i, err := r.pipeline(id)
	if err != nil {
		return nil, err
	}

	p := r.pipelines[i]
	r.pipelines = slices.Delete(r.pipelines, i, i+1)

	return p, nil
}

func (r *InMemoryRepository) pipeline(id ID) (int, error) {
	i := slices.IndexFunc(r.pipelines, func(p *Pipeline) bool { return p.ID == id })
	if i == -1 {
		return 0, ErrPipelineNotFound
	}

	return i, nil
}
