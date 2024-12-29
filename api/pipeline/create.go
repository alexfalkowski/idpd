package pipeline

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/go-service/structs"
)

type (
	// CreatePipelineRequest with a definition.
	CreatePipelineRequest struct {
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}

	// CreatePipelineResponse a map of meta and the new pipeline.
	CreatePipelineResponse struct {
		Meta     meta.Map  `json:"meta,omitempty"`
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}
)

// Valid returns an error if we have missing pipeline.
func (c *CreatePipelineRequest) Valid() error {
	if structs.IsZero(c) {
		return ErrMissingPipeline
	}

	return nil
}

// CreatePipeline for the api.
func (s *Service) CreatePipeline(ctx context.Context, req *CreatePipelineRequest) (*CreatePipelineResponse, error) {
	if err := req.Valid(); err != nil {
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
