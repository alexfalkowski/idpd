package pipeline

// Create a new pipeline.
func (s *Service) Create(p *Pipeline) (*Pipeline, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Create(p)
}
