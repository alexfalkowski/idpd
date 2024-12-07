package api

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/content"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/go-service/net/http/status"
)

type (
	// UpdatePipelineRequest with a definition.
	UpdatePipelineRequest struct {
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}

	// UpdatePipelineResponse a map of meta and the updated pipeline.
	UpdatePipelineResponse struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Pipeline *Pipeline         `json:"pipeline,omitempty"`
	}
)

func (s *Service) updatePipeline(ctx context.Context) (any, error) {
	var req UpdatePipelineRequest
	if err := content.Decode(ctx, &req); err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	p := s.toPipeline(req.Pipeline)
	request := hc.Request(ctx)
	id := s.service.ID(request.PathValue("id"))

	p, err := s.service.Update(id, p)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &CreatePipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}