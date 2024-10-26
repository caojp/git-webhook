package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Secret      string `yaml:"secret"`
	ProjectPath string `yaml:"project_path"`
	LogFilePath string `yaml:"log_file_path"`
	Port        int    `yaml:"port"`
	RepoURL     string `yaml:"repo_url"`
	UseTag      bool   `yaml:"use_tag"`
	Branch      string `yaml:"branch"`
	GitUsername string `yaml:"git_username"`
	GitPassword string `yaml:"git_password"`
	SSHKeyPath  string `yaml:"ssh_key_path"`
	SSLCertFile string `yaml:"ssl_cert_file"`
	SSLKeyFile  string `yaml:"ssl_key_file"`
}

// LoadConfig 加载 YAML 配置文件
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
