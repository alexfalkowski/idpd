package pipeline

// Service coordinates all the pipeline activities.
type Service struct {
	repo Repository
	cmd  Command
}

// NewService for pipeline.
func NewService(repo Repository, cmd Command) *Service {
	return &Service{repo: repo, cmd: cmd}
}
