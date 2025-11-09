package core

// Option:
// 	- StartHeading
// 	- MaxHeading
// 	- TagsDetectEnable
// 	- LogOutput

type Config struct{
	Feishu FeishuConfig `json:"feishu"`
	Output OutputConfig `json:"output"`
}

type FeishuConfig struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	SpaceId   string `json:"space_id"`
}

type OutputConfig struct {
	
}

//Create Config
//appId: 应用的 App ID
//appSecret: 应用的 App Secret
//spaceId: 知识库的 Space ID
func NewConfig(appId, appSecret, spaceId string) *Config {
	return &Config{
		Feishu: FeishuConfig{
			AppId:     appId,
			AppSecret: appSecret,
			SpaceId: spaceId,
		},
	}
}