package pipeline

import (
	"errors"
)

var (
	// ErrInvalidPipelineName for pipeline.
	ErrInvalidPipelineName = errors.New("invalid pipeline name")

	// ErrMissingJobs for pipeline.
	ErrMissingJobs = errors.New("missing jobs")

	// ErrInvalidJobName for pipeline.
	ErrInvalidJobName = errors.New("invalid job name")

	// ErrMissingSteps for pipeline.
	ErrMissingSteps = errors.New("missing steps")
)

// IsInvalidArgument for pipeline.
func IsInvalidArgument(err error) bool {
	errs := []error{ErrInvalidPipelineName, ErrMissingJobs, ErrInvalidJobName, ErrMissingSteps}

	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}

type (
	// Job of the pipeline.
	Job struct {
		Name  string
		Steps []string
	}

	// Pipeline to be executed.
	Pipeline struct {
		Name string
		Jobs []*Job
		ID   uint64
	}
)

// Valid job or error.
func (j *Job) Valid() error {
	if j.Name == "" {
		return ErrInvalidJobName
	}

	if len(j.Steps) == 0 {
		return ErrMissingSteps
	}

	return nil
}

// Valid pipeline or error.
func (p *Pipeline) Valid() error {
	if p.Name == "" {
		return ErrInvalidPipelineName
	}

	if len(p.Jobs) == 0 {
		return ErrMissingSteps
	}

	for _, j := range p.Jobs {
		if err := j.Valid(); err != nil {
			return err
		}
	}

	return nil
}
