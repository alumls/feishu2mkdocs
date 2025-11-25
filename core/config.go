package core

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Future output options may include:
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
	YamlPath string `yaml:"yaml_path"`
}

// NewConfig: 创建一个Config实例。该函数只保留了Config的必要接口，其它接口将以默认值填充。
//
// 参数：
//   - appId: 应用的 App ID
//   - appSecret: 应用的 App Secret
//   - spaceId: 知识库的 Space ID
func NewConfig(appId, appSecret, spaceId string) *Config {
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
// ReadFromConfigFile: 从配置文件中读取配置。
//
// 参数：
//   - path: 配置文件的路径
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
