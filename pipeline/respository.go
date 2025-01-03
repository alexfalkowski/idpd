package pipeline

import (
	"bytes"
	"hash/fnv"
	"sync"

	"github.com/alexfalkowski/go-service/crypto/rand"
	"github.com/alexfalkowski/go-service/encoding/gob"
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
	// The mux is to make sure we don't accidentally corrupt.
	// The cache and enc are there to make sure we don't maintain pointers in memory.
	InMemoryRepository struct {
		enc   *gob.Encoder
		cache *cache.LRU[ID, []byte]
		mu    sync.Mutex
	}
)

// NewRepository for pipeline.
func NewRepository(enc *gob.Encoder) Repository {
	// No need to check for err, as we have valid arguments.
	c, _ := cache.New[ID, []byte](1024, hash)

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

	id, err := r.generateID()
	if err != nil {
		return nil, err
	}

	p.ID = id

	var b bytes.Buffer
	if err := r.enc.Encode(&b, p); err != nil {
		return nil, err
	}

	r.cache.Add(id, b.Bytes())

	return p, nil
}

// Update a pipeline.
func (r *InMemoryRepository) Update(id ID, p *Pipeline) (*Pipeline, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.cache.Get(id)
	if !ok {
		return nil, ErrPipelineNotFound
	}

	p.ID = id

	var b bytes.Buffer
	if err := r.enc.Encode(&b, p); err != nil {
		return nil, err
	}

	r.cache.Add(id, b.Bytes())

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

	r.cache.Remove(id)

	return p, nil
}

func (r *InMemoryRepository) get(id ID) (*Pipeline, error) {
	b, ok := r.cache.Get(id)
	if !ok {
		return nil, ErrPipelineNotFound
	}

	var p Pipeline
	if err := r.enc.Decode(bytes.NewBuffer(b), &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *InMemoryRepository) generateID() (ID, error) {
	id, err := rand.GenerateLetters(6)
	if err != nil {
		return "", err
	}

	return ID(id), nil
}

func hash(id ID) uint32 {
	h := fnv.New32a()
	h.Write([]byte(id))

	return h.Sum32()
}
