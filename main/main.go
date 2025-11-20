package main

import (
	//"context"
	"feishu2mkdocs/core"
	//"feishu2mkdocs/utils"
	"feishu2mkdocs/service"
	//"os"
	"fmt"
)

func main() {
	config, err := core.ReadFromConfigFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(utils.PrettyPrint(config))
	client := core.NewClient(config.Feishu.AppId, config.Feishu.AppSecret)
	nodeMap := service.NewNodeMap()
	nodeMap.GenerateWikiContent(client, config)
}