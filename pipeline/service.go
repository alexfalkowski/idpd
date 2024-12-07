package pipeline

import (
	"strconv"
)

type (
	// Service coordinates all the pipeline activities.
	Service struct {
		repo Repository
	}

	// ID for service.
	ID uint64
)

// NewService for pipeline.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Get pipeline by id.
func (s *Service) Get(id ID) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Get(id)
}

// Create a new pipeline.
func (s *Service) Create(p *Pipeline) (*Pipeline, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Create(p)
}

// Update an existing pipeline.
func (s *Service) Update(id ID, p *Pipeline) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	if err := p.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Update(id, p)
}

// Update an existing pipeline.
func (s *Service) Delete(id ID) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Delete(id)
}

// ID from a string.
func (s *Service) ID(id string) ID {
	i, _ := strconv.ParseUint(id, 10, 64)

	return ID(i)
}

// Valid or error.
func (i ID) Valid() error {
	if i == 0 {
		return ErrInvalidID
	}

	return nil
}
