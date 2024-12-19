package pipeline

// ID for service.
//
// A common pattern in DDD is to make sure we have defined types and allow us to extend what we mean by an id.
type ID string

// Valid or error if ID is not the correct length.
func (i ID) Valid() error {
	if len(i) != 6 {
		return ErrInvalidID
	}

	return nil
}
