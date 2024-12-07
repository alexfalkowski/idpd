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

// ID from raw id.
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
