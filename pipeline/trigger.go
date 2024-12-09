package pipeline

import (
	"context"
	"fmt"
)

// Update an existing pipeline.
//
// This assumes that we will just run each job and each step serially. The first error we exit.
func (s *Service) Trigger(ctx context.Context, id ID) (*Pipeline, error) {
	if err := id.Valid(); err != nil {
		return nil, err
	}

	pipeline, err := s.Get(id)
	if err != nil {
		return nil, err
	}

	for _, job := range pipeline.Jobs {
		outputs := make([]string, len(job.Steps))

		for i, step := range job.Steps {
			output, err := s.cmd.Exec(ctx, step)
			if err != nil {
				return nil, fmt.Errorf("pipeline %s: job %s failed: step %s: %w", pipeline.Name, job.Name, step, err)
			}

			outputs[i] = output
		}

		job.Steps = outputs
	}

	return pipeline, nil
}
