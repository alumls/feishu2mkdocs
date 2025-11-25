package main

import (
	"feishu2mkdocs/core"
	"feishu2mkdocs/service"
	"fmt"
)

func main() {
	cfg, err := core.ReadFromConfigFile("config.yaml")
	if err != nil {
		fmt.Println("读取配置失败:", err)
		return
	}

	client := core.NewClient(cfg.Feishu.AppId, cfg.Feishu.AppSecret)

	gen := service.NewGenerator(client, cfg)

	if err := gen.GenerateWikiContent(); err != nil {
		fmt.Println("GenerateWikiContent 失败:", err)
		return
	}
	fmt.Println("GenerateWikiContent 完成")

	if err := gen.GenerateWikiNav(); err != nil {
		fmt.Println("GenerateWikiNav 失败:", err)
		return
	}
	fmt.Println("GenerateWikiNav 完成")
}
