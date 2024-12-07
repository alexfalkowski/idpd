package api

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/content"
	"github.com/alexfalkowski/go-service/net/http/status"
)

type (
	// CreatePipelineRequest with a definition.
	CreatePipelineRequest struct {
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}

	// CreatePipelineResponse a map of meta and the new pipeline.
	CreatePipelineResponse struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Pipeline *Pipeline         `json:"pipeline,omitempty"`
	}
)

func (s *Service) createPipeline(ctx context.Context) (any, error) {
	var req CreatePipelineRequest
	if err := content.Decode(ctx, &req); err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	p := s.toPipeline(req.Pipeline)

	p, err := s.service.Create(p)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &CreatePipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}
