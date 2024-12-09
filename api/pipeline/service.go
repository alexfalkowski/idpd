package pipeline

import (
	"net/http"

	"github.com/alexfalkowski/go-service/net/http/rest"
	"github.com/alexfalkowski/go-service/net/http/status"
	"github.com/alexfalkowski/idpd/pipeline"
)

// Register routes.
func Register(service *Service) {
	rest.Post("/pipelines", service.createPipeline)
	rest.Get("/pipelines/{id}", service.getPipeline)
	rest.Put("/pipelines/{id}", service.updatePipeline)
	rest.Delete("/pipelines/{id}", service.deletePipeline)
	rest.Post("/pipelines/{id}/trigger", service.triggerPipeline)
}

// Service for pipeline.
type Service struct {
	service *pipeline.Service
}

// NewService for pipeline.
func NewService(service *pipeline.Service) *Service {
	return &Service{service: service}
}

func (s *Service) toPipeline(pl *Pipeline) *pipeline.Pipeline {
	p := &pipeline.Pipeline{
		Name: pl.Name,
		ID:   pipeline.ID(pl.ID),
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
		ID:   uint64(pl.ID),
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
