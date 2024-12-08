package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/idpd/pipeline"
)

// TriggerPipelineResponse a map of meta and the pipeline with the output in steps.
type TriggerPipelineResponse struct {
	Meta     map[string]string `json:"meta,omitempty"`
	Pipeline *Pipeline         `json:"pipeline,omitempty"`
}

func (s *Service) triggerPipeline(ctx context.Context) (any, error) {
	req := hc.Request(ctx)
	id := pipeline.NewID(req.PathValue("id"))

	p, err := s.service.Trigger(ctx, id)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &TriggerPipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}
