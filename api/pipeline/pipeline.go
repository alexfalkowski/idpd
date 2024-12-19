package pipeline

import (
	"errors"
)

// ErrMissingPipeline when the payload is invalid.
var ErrMissingPipeline = errors.New("missing pipeline")

type (
	// Job of the pipeline.
	Job struct {
		Name  string   `json:"name,omitempty"`
		Steps []string `json:"steps,omitempty"`
	}

	// Pipeline to be executed.
	Pipeline struct {
		Name string `json:"name,omitempty"`
		ID   string `json:"id,omitempty"`
		Jobs []*Job `json:"jobs,omitempty"`
	}
)
