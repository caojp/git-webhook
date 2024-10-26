package git

import (
	"git-webhook/config"
	"os/exec"
)

type PullOperation struct {
	cfg *config.Config
}

func NewPullOperation(cfg *config.Config) *PullOperation {
	return &PullOperation{cfg: cfg}
}

func (p *PullOperation) Execute() error {
	var cmd *exec.Cmd
	if p.cfg.UseTag {
		cmd = exec.Command("git", "-C", p.cfg.ProjectPath, "fetch", "--tags")
	} else {
		cmd = exec.Command("git", "-C", p.cfg.ProjectPath, "pull", "origin", p.cfg.Branch)
	}
	return cmd.Run()
}
