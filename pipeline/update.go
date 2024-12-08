package pipeline

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
