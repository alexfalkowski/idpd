package pipeline

import (
	"context"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/idpd/pipeline"
)

// DeletePipelineResponse a map of meta and the updated pipeline.
type DeletePipelineResponse struct {
	Meta     meta.Map  `json:"meta,omitempty"`
	Pipeline *Pipeline `json:"pipeline,omitempty"`
}

// DeletePipeline for the api.
func (s *Service) DeletePipeline(ctx context.Context) (*DeletePipelineResponse, error) {
	req := hc.Request(ctx)
	id := pipeline.ID(req.PathValue("id"))

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
