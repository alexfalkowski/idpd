package pipeline

import (
	"errors"
)

var (
	// ErrPipelineNotFound for pipeline.
	ErrPipelineNotFound = errors.New("pipeline not found")

	// ErrInvalidPipelineName for pipeline.
	ErrInvalidPipelineName = errors.New("invalid pipeline name")

	// ErrMissingJobs for pipeline.
	ErrMissingJobs = errors.New("missing jobs")

	// ErrInvalidJobName for pipeline.
	ErrInvalidJobName = errors.New("invalid job name")

	// ErrMissingSteps for pipeline.
	ErrMissingSteps = errors.New("missing steps")

	// ErrInvalidID for pipeline.
	ErrInvalidID = errors.New("invalid id")
)

// IsNotFound for pipeline.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrPipelineNotFound)
}

// IsInvalidArgument for pipeline.
func IsInvalidArgument(err error) bool {
	errs := []error{
		ErrInvalidPipelineName, ErrMissingJobs,
		ErrInvalidJobName, ErrMissingSteps, ErrInvalidID,
	}

	for _, e := range errs {
		if errors.Is(err, e) {
			return true
		}
	}

	return false
}

type (
	// Job of the pipeline.
	//
	// Each job just has a list of commands that will run on the host.
	// A good job definition would have some sort of workflow definition to allow different patterns.
	// An example can be found at https://circleci.com/docs/workflows/
	Job struct {
		Name  string
		Steps []string
	}

	// Pipeline to be executed.
	//
	// A pipeline has a list of jobs.
	Pipeline struct {
		Name string
		ID   ID
		Jobs []*Job
	}
)

// Valid job or error if name is blank or empty steps.
func (j *Job) Valid() error {
	if j.Name == "" {
		return ErrInvalidJobName
	}

	if len(j.Steps) == 0 {
		return ErrMissingSteps
	}

	return nil
}

// Valid pipeline or error if name is blank, empty jobs or invalid jobs.
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
