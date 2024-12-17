package pipeline

import (
	"bytes"
	"sync"

	"github.com/alexfalkowski/go-service/encoding/gob"
	"github.com/cespare/xxhash/v2"
	cache "github.com/elastic/go-freelru"
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
	// The cache and enc are there to make sure we don't maintain pointers in memory.
	InMemoryRepository struct {
		enc     *gob.Encoder
		cache   *cache.LRU[string, []byte]
		counter uint64
		mu      sync.Mutex
	}
)

// NewRepository for pipeline.
func NewRepository(enc *gob.Encoder) Repository {
	// No need to check for err, as we have valid arguments.
	c, _ := cache.New[string, []byte](1024, hash)

	return &InMemoryRepository{enc: enc, cache: c}
}

// Get a pipeline.
func (r *InMemoryRepository) Get(id ID) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.get(id)
}

// Create a pipeline and set the identifier.
func (r *InMemoryRepository) Create(p *Pipeline) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	id := ID(r.counter)

	p.ID = id

	var b bytes.Buffer
	if err := r.enc.Encode(&b, p); err != nil {
		return nil, err
	}

	r.cache.Add(id.String(), b.Bytes())

	return p, nil
}

// Update a pipeline.
func (r *InMemoryRepository) Update(id ID, p *Pipeline) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.get(id)
	if err != nil {
		return nil, err
	}

	p.ID = id

	var b bytes.Buffer
	if err := r.enc.Encode(&b, p); err != nil {
		return nil, err
	}

	r.cache.Add(id.String(), b.Bytes())

	return p, nil
}

// Delete a pipeline.
func (r *InMemoryRepository) Delete(id ID) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, err := r.get(id)
	if err != nil {
		return nil, err
	}

	r.cache.Remove(id.String())

	return p, nil
}

func (r *InMemoryRepository) get(id ID) (*Pipeline, error) {
	b, ok := r.cache.Get(id.String())
	if !ok {
		return nil, ErrPipelineNotFound
	}

	var p Pipeline
	if err := r.enc.Decode(bytes.NewBuffer(b), &p); err != nil {
		return nil, err
	}

	return &p, nil
}

//nolint:gosec
func hash(s string) uint32 {
	return uint32(xxhash.Sum64String(s))
}
