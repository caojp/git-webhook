package git

import (
	"fmt"
	"git-webhook/config"
	"os"
)

// AuthStrategy 认证策略接口
type AuthStrategy interface {
	Apply() error
}

// BasicAuth 基本认证
type BasicAuth struct {
	Username, Password string
}

func (a *BasicAuth) Apply() error {
	return os.Setenv("GIT_ASKPASS", fmt.Sprintf("%s:%s", a.Username, a.Password))
}

// SSHAuth SSH 认证
type SSHAuth struct {
	KeyPath string
}

func (s *SSHAuth) Apply() error {
	return os.Setenv("GIT_SSH_COMMAND", fmt.Sprintf("ssh -i %s", s.KeyPath))
}

func GetAuthStrategy(cfg *config.Config) (AuthStrategy, error) {
	if cfg.GitUsername != "" && cfg.GitPassword != "" {
		return &BasicAuth{cfg.GitUsername, cfg.GitPassword}, nil
	} else if cfg.SSHKeyPath != "" {
		return &SSHAuth{cfg.SSHKeyPath}, nil
	}
	return nil, fmt.Errorf("no authentication method provided")
}
