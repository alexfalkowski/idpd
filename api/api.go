package api

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/content"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/idpd/pipeline"
)

// Register routes.
func Register(service *Service) {
	rest.Post("/pipeline", service.createPipeline)
}

// Service for v1.
type Service struct {
	service *pipeline.Service
}

// NewService for v1.
func NewService(service *pipeline.Service) *Service {
	return &Service{service: service}
}

func (s *Service) createPipeline(ctx context.Context) (any, error) {
	var req CreatePipelineRequest
	if err := content.Decode(ctx, &req); err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	p := s.convert(req.Pipeline)

	p, err := s.service.Create(p)
	if err != nil {
		if pipeline.IsInvalidArgument(err) {
			return nil, status.Error(http.StatusBadRequest, err.Error())
		}

		return nil, err
	}

	req.Pipeline.ID = p.ID

	res := &CreatePipelineResponse{
		Meta:     meta.CamelStrings(ctx, ""),
		Pipeline: req.Pipeline,
	}

	return res, nil
}

func (s *Service) convert(pl *Pipeline) *pipeline.Pipeline {
	p := &pipeline.Pipeline{
		Name: pl.Name,
	}

	p.Jobs = make([]*pipeline.Job, len(pl.Jobs))

	for i, j := range pl.Jobs {
		job := &pipeline.Job{
			Name: j.Name,
		}

		job.Steps = make([]string, len(j.Steps))
		copy(job.Steps, j.Steps)

		p.Jobs[i] = job
	}

	return p
}
