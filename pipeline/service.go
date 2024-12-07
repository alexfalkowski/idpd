package pipeline

// Service coordinates all the pipeline activities.
type Service struct {
	repo Repository
}

// NewService for pipeline.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// Create a new pipeline.
func (s *Service) Create(p *Pipeline) (*Pipeline, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}

	p, err := s.repo.Create(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
