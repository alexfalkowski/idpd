package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/idpd/pipeline"
)

// TriggerPipelineResponse a map of meta and the pipeline with the output in steps.
type TriggerPipelineResponse struct {
	Meta     meta.Map  `json:"meta,omitempty"`
	Pipeline *Pipeline `json:"pipeline,omitempty"`
}

// TriggerPipeline for the api.
func (s *Service) TriggerPipeline(ctx context.Context) (*TriggerPipelineResponse, error) {
	req := hc.Request(ctx)
	id := pipeline.ID(req.PathValue("id"))

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
