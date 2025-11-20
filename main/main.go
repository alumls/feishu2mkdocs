package main

import (
	"context"
	"feishu2mkdocs/core"
	"feishu2mkdocs/utils"
	//"feishu2mkdocs/service"
	"os"
	"fmt"
)

func main() {
	parentNodeToken := ""
	docxToken := "GOWNd16b4oGagxxbxvZcK1pEnZe"
	cfg, _ := core.ReadFromConfigFile("config.yaml")
	spaceId := cfg.Feishu.SpaceId
	client := core.NewClient(cfg.Feishu.AppId, cfg.Feishu.AppSecret)
	nodes, _ := client.GetWikiNodeListAll(context.Background(), spaceId, &parentNodeToken)
	blocks, _ := client.GetDocumentBlockAll(context.Background(), &docxToken)
	nodeMap := core.NewNodeMap()
	nodeMap.BuildFromFlatNodes(nodes, cfg.Output.DocsDir)
	
	fmt.Println(utils.PrettyPrint(cfg.Output))

	fmt.Println(utils.PrettyPrint(blocks))

	parser := core.NewParser(cfg.Output, nodeMap)
	markdown, _ := parser.ParseDocxsContent(nodes[0], blocks)
	
	parseroutput := markdown

	er := os.WriteFile("md.txt", []byte(parseroutput), 0644)
    if er != nil {
        fmt.Println("写入文件失败:", er)
        return
    }

    fmt.Println("输出已保存到 md.txt")
}