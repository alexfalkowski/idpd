package api

type (

	// Job of the pipeline.
	Job struct {
		Name  string   `json:"name,omitempty"`
		Steps []string `json:"steps,omitempty"`
	}

	// Pipeline to be executed.
	Pipeline struct {
		Name string `json:"name,omitempty"`
		Jobs []*Job `json:"jobs,omitempty"`
		ID   uint64 `json:"id,omitempty"`
	}

	// CreatePipeline with a definition.
	CreatePipelineRequest struct {
		Pipeline *Pipeline `json:"pipeline,omitempty"`
	}

	// CreatePipelineResponse a map of meta and the new pipeline.
	CreatePipelineResponse struct {
		Meta     map[string]string `json:"meta,omitempty"`
		Pipeline *Pipeline         `json:"pipeline,omitempty"`
	}
)
