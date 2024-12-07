package pipeline

import (
	"strconv"
)

// Service coordinates all the pipeline activities.
type Service struct {
	repo Repository
}

// NewService for pipeline.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Get(id uint64) (*Pipeline, error) {
	if id == 0 {
		return nil, ErrInvalidID
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
func (s *Service) ID(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}
