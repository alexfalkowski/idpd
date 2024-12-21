package pipeline

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/go-service/structs"
	"github.com/alexfalkowski/idpd/pipeline"
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

// Valid returns an error if we have missing pipeline.
func (u *UpdatePipelineRequest) Valid() error {
	if structs.IsZero(u) {
		return ErrMissingPipeline
	}

	return nil
}

// UpdatePipeline for the api.
func (s *Service) UpdatePipeline(ctx context.Context, req *UpdatePipelineRequest) (*UpdatePipelineResponse, error) {
	if err := req.Valid(); err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	p := s.toPipeline(req.Pipeline)
	request := hc.Request(ctx)
	id := pipeline.ID(request.PathValue("id"))

	p, err := s.service.Update(id, p)
	if err != nil {
		return nil, s.handleError(err)
	}

	res := &UpdatePipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: s.fromPipeline(p),
	}

	return res, nil
}
