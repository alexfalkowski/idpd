package pipeline

import (
	"context"
	"os/exec"
	"strings"
)

type (
	// Command runs a specific command.
	Command interface {
		// Exec a command and output the result or error.
		Exec(ctx context.Context, cmd string) (string, error)
	}

	// OSCommand uses os/exec.
	OSCommand struct{}
)

// NewCommand for pipeline.
func NewCommand() Command {
	return &OSCommand{}
}

// Exec using os/exec.
func (r *OSCommand) Exec(ctx context.Context, cmd string) (string, error) {
	s := strings.Split(cmd, " ")
	c := exec.CommandContext(ctx, s[0], s[1:]...) // #nosec G204

	o, err := c.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(o), "\n"), nil
}
