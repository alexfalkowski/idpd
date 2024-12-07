package api

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
)

// GetPipelineResponse a map of meta and the pipeline.
type GetPipelineResponse struct {
	Meta     map[string]string `json:"meta,omitempty"`
	Pipeline *Pipeline         `json:"pipeline,omitempty"`
}

func (s *Service) getPipeline(ctx context.Context) (any, error) {
	req := hc.Request(ctx)
	id := s.service.ID(req.PathValue("id"))

	p, err := s.service.Get(id)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &GetPipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}
