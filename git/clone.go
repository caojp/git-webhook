package git

import (
	"git-webhook/config"
	"os/exec"
)

type CloneOperation struct {
	cfg *config.Config
}

func NewCloneOperation(cfg *config.Config) *CloneOperation {
	return &CloneOperation{cfg: cfg}
}

func (c *CloneOperation) Execute() error {
	var cmd *exec.Cmd
	if c.cfg.UseTag {
		cmd = exec.Command("git", "clone", "--branch", "latest", "--depth", "1", c.cfg.RepoURL, c.cfg.ProjectPath)
	} else {
		cmd = exec.Command("git", "clone", "-b", c.cfg.Branch, c.cfg.RepoURL, c.cfg.ProjectPath)
	}
	return cmd.Run()
}
