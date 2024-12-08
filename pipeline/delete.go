package pipeline

// Update an existing pipeline.
func (s *Service) Delete(id ID) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	return s.repo.Delete(id)
}
