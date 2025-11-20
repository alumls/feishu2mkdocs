package service

import (
	"context"
	"feishu2mkdocs/core"
	"os"

	//"feishu2mkdocs/utils"
	"fmt"
	//larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Generator struct {
	Client *core.Client
	Config *core.Config
	Parser *core.Parser
	NodeMap *core.NodeMap
}

func NewGenerator(c *core.Client, cfg *core.Config) *Generator {
	nodeMap := core.NewNodeMap()
	return &Generator{
		Client: c,
		Config: cfg,
		Parser: core.NewParser(cfg.Output, nodeMap),
		NodeMap: nodeMap,
	}
}

func (g *Generator) GenerateWikiContent() error {
	parentNodeToken := ""

	nodes, _ := g.Client.GetWikiNodeListAll(context.Background(), g.Config.Feishu.SpaceId, &parentNodeToken)

	g.NodeMap.BuildFromFlatNodes(nodes, g.Config.Output.DocsDir)

	for _, node := range nodes {
		blocks, _ := g.Client.GetDocumentBlockAll(context.Background(), node.ObjToken)
		md, err := g.Parser.ParseDocxsContent(node, blocks)

		if err != nil {
			return err
		}

		path := g.NodeMap.NodeMeta[*node.NodeToken].Path
		dir := g.NodeMap.NodeMeta[*node.NodeToken].Dir

		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("创建目录失败:", err)
			return err
		}

		if err := os.WriteFile(path, []byte(md), 0644); err != nil {
			fmt.Println("写入文件失败:", err)
			return err
		}

		fmt.Println(*node.NodeToken + ": 输出已保存到 " + path)
	}

	return nil
}

func (g *Generator) GenerateWikiNav(cfg *core.Config, yamlDir string) error {
	return nil
}
