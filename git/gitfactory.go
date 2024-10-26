package git

import (
	"git-webhook/config"
	"git-webhook/logger"
	"os"
)

func GetGitOperation(cfg *config.Config) (GitOperation, error) {
	if _, err := os.Stat(cfg.ProjectPath); os.IsNotExist(err) {
		logger.Logger.Printf("Project path %s does not exist. Cloning repository...", cfg.ProjectPath)
		return NewCloneOperation(cfg), nil
	} else {
		logger.Logger.Printf("Project path %s exists. Pulling updates...", cfg.ProjectPath)
		return NewPullOperation(cfg), nil
	}
}
