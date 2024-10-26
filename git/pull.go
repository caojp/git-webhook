package git

import (
	"fmt"
	"git-webhook/config"
	"git-webhook/logger"
	"os/exec"
)

type PullOperation struct {
	cfg *config.Config
}

func NewPullOperation(cfg *config.Config) *PullOperation {
	return &PullOperation{cfg: cfg}
}

// Execute performs a git fetch or pull based on the configuration.
func (p *PullOperation) Execute() error {
	var cmd *exec.Cmd

	// 根据配置选择 fetch tags 或 pull 指定分支
	if p.cfg.UseTag {
		logger.Logger.Println("Starting git fetch --tags operation")
		cmd = exec.Command("git", "fetch", "--tags")
	} else {
		logger.Logger.Printf("Starting git pull operation on branch %s\n", p.cfg.Branch)
		cmd = exec.Command("git", "pull", "origin", p.cfg.Branch)
	}

	// 切换到项目目录
	cmd.Dir = p.cfg.ProjectPath

	// 打印执行的命令
	logger.Logger.Printf("Executing command: %s\n", cmd.String())

	// 运行命令并捕获可能的错误
	if err := cmd.Run(); err != nil {
		logger.Logger.Printf("Command execution failed: %v\n", err)
		return fmt.Errorf("git operation failed: %w", err)
	}

	logger.Logger.Println("Git operation completed successfully")
	return nil
}
