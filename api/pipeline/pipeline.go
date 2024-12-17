package pipeline

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
		ID   uint32 `json:"id,omitempty"`
	}
)
