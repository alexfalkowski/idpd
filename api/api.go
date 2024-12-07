package api

import (
	"context"
	"net/http"

	"github.com/alexfalkowski/go-service/meta"
	"github.com/alexfalkowski/go-service/net/http/content"
	hc "github.com/alexfalkowski/go-service/net/http/context"
	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/idpd/pipeline"
)

// Register routes.
func Register(service *Service) {
	rest.Post("/pipeline", service.createPipeline)
	rest.Get("/pipeline/{id}", service.getPipeline)
}

// Service for v1.
type Service struct {
	service *pipeline.Service
}

// NewService for v1.
func NewService(service *pipeline.Service) *Service {
	return &Service{service: service}
}

func (s *Service) getPipeline(ctx context.Context) (any, error) {
	req := hc.Request(ctx)

	id, err := s.service.ID(req.PathValue("id"))
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

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

func (s *Service) toPipeline(pl *Pipeline) *pipeline.Pipeline {
	p := &pipeline.Pipeline{
		Name: pl.Name,
		ID:   pl.ID,
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

func (s *Service) fromPipeline(pl *pipeline.Pipeline) *Pipeline {
	p := &Pipeline{
		Name: pl.Name,
		ID:   pl.ID,
	}

	p.Jobs = make([]*Job, len(pl.Jobs))

	for i, j := range pl.Jobs {
		job := &Job{
			Name: j.Name,
		}

		job.Steps = make([]string, len(j.Steps))
		copy(job.Steps, j.Steps)

		p.Jobs[i] = job
	}

	return p
}

func (s *Service) handleError(err error) error {
	if pipeline.IsInvalidArgument(err) {
		return status.Error(http.StatusBadRequest, err.Error())
	}

	if pipeline.IsNotFound(err) {
		return status.Error(http.StatusNotFound, err.Error())
	}

	return err
}
