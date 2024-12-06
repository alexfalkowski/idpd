package api

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
)

// DeletePipelineResponse a map of meta and the updated pipeline.
type DeletePipelineResponse struct {
	Meta     map[string]string `json:"meta,omitempty"`
	Pipeline *Pipeline         `json:"pipeline,omitempty"`
}

func (s *Service) deletePipeline(ctx context.Context) (any, error) {
	request := hc.Request(ctx)
	id := s.service.ID(request.PathValue("id"))

	p, err := s.service.Delete(id)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &DeletePipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}
