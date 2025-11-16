package service

import (
	"context"
	"os"
	"feishu2mkdocs/core"
	//"feishu2mkdocs/utils"
	"fmt"

	//larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

func GenerateWikiContent(c *core.Client, cfg *core.Config) error {
	parentNodeToken := ""

	wikiMap := NewWikiMap()
	
	nodes, _ := c.GetWikiNodeListAll(context.Background(), cfg.Feishu.SpaceId, &parentNodeToken)

	//fmt.Println(utils.PrettyPrint(nodes))

	wikiMap.WikiMapBuildFromFlatNodes(nodes, cfg.Output.DocsDir)
	
	parser := core.NewParser(cfg.Output)

	for _, node := range nodes {
		blocks, _ := c.GetDocumentBlockAll(context.Background(), *node.ObjToken)
		// fmt.Println(node)
		// fmt.Println(blocks)
		md := parser.ParseDocxsContent(node, blocks)
		path := wikiMap.WikiMapNode[*node.NodeToken].Path
		dir := wikiMap.WikiMapNode[*node.NodeToken].Dir

		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("创建目录失败:", err)
        	return err
    	}

		if err := os.WriteFile(path, []byte(md), 0644); err != nil {
			fmt.Println("写入文件失败:", err)
			return err
		}

		fmt.Println("输出已保存到" + path)
	}

	return nil
}

