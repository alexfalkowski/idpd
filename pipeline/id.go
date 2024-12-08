package pipeline

import (
	"strconv"
)

// ID for service.
type ID uint64

// NewID from a string.
func NewID(id string) ID {
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
