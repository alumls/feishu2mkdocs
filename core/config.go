package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Future Option:
// 	- StartHeading
// 	- MaxHeading
// 	- TagsDetectEnable
// 	- LogOutput

type Config struct {
	Feishu FeishuConfig `yaml:"feishu"`
	Output OutputConfig `yaml:"output"`
}

type FeishuConfig struct {
	AppId     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	SpaceId   string `yaml:"space_id"`
}

type OutputConfig struct {
	DocsDir string `yaml:"docs_dir"`
}

// Create Config
// appId: 应用的 App ID
// appSecret: 应用的 App Secret
// spaceId: 知识库的 Space ID
func NewConfig(appId, appSecret, spaceId, docsDir string) *Config {
	return &Config{
		Feishu: FeishuConfig{
			AppId:     appId,
			AppSecret: appSecret,
			SpaceId:   spaceId,
		},
		Output: OutputConfig{
			DocsDir: "docs",
		},
	}
}

func ReadFromConfigFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
