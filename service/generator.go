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

	nodeMap := NewNodeMap()
	
	nodes, _ := c.GetWikiNodeListAll(context.Background(), cfg.Feishu.SpaceId, &parentNodeToken)

	nodeMap.NodeMapBuildFromFlatNodes(nodes, cfg.Output.DocsDir)
	
	parser := core.NewParser(cfg.Output)

	for _, node := range nodes {
		blocks, _ := c.GetDocumentBlockAll(context.Background(), node.ObjToken)
		md, err := parser.ParseDocxsContent(node, blocks)

		if err != nil {
			return err
		}

		path := nodeMap.NodeMeta[*node.NodeToken].Path
		dir := nodeMap.NodeMeta[*node.NodeToken].Dir

		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("创建目录失败:", err)
        	return err
    	}

		if err := os.WriteFile(path, []byte(md), 0644); err != nil {
			fmt.Println("写入文件失败:", err)
			return err
		}

		fmt.Println( *node.NodeToken + ": 输出已保存到 " + path)
	}

	return nil
}

