package pipeline

// Get pipeline by id.
func (s *Service) Get(id ID) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Get(id)
}
