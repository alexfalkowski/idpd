package pipeline

// Service coordinates all the pipeline activities.
//
// This is entry point and the type to be used when interacting with this package.
type Service struct {
	repo Repository
	cmd  Command
}

// NewService for pipeline.
func NewService(repo Repository, cmd Command) *Service {
	return &Service{repo: repo, cmd: cmd}
}
