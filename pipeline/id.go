package pipeline

import (
	"strconv"
)

// ID for service.
//
// A common pattern in DDD is to make sure we have defined types and allow us to extend what we mean by an id.
type ID uint64

// NewID from a string.
//
// We ignore the error as 0 is an invalid id.
func NewID(id string) ID {
	i, _ := strconv.ParseUint(id, 10, 64)

	return ID(i)
}

// Valid or error if ID is 0.
func (i ID) Valid() error {
	if i == 0 {
		return ErrInvalidID
	}

	return nil
}

// String of the ID.
func (i ID) String() string {
	return strconv.FormatUint(uint64(i), 10)
}
