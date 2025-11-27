package service

import(
	"os"

	"feishu2mkdocs/core"
	"gopkg.in/yaml.v3"
)

// LoadConfig: 从配置文件中读取配置。
//
// 参数：
//   - path: 配置文件的路径
func LoadConfig(path string) (*core.Config, error) {
	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var cfg core.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
