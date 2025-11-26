package service

import (
	"context"
	"feishu2mkdocs/core"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Generator struct {
	Client  *core.Client
	Config  *core.Config
	Parser  *core.Parser
	NodeMap *core.NodeMap
}

func NewGenerator(c *core.Client, cfg *core.Config) *Generator {
	nodeMap := core.NewNodeMap()
	return &Generator{
		Client:  c,
		Config:  cfg,
		Parser:  core.NewParser(cfg.Output, nodeMap),
		NodeMap: nodeMap,
	}
}

func (g *Generator) GenerateWikiContent() error {
	parentNodeToken := ""

	nodes, _ := g.Client.GetWikiNodeListAll(context.Background(), g.Config.Feishu.SpaceId, &parentNodeToken)

	g.NodeMap.BuildFromFlatNodes(nodes, g.Config.Output.DocsDir)

	for _, node := range nodes {
		if node.ObjType != nil && *node.ObjType != "docx" {
			continue
		}
		blocks, _ := g.Client.GetDocumentBlockAll(context.Background(), node.ObjToken)
		md, err := g.Parser.ParseDocxsContent(node, blocks)

		if err != nil {
			return err
		}

		path := g.NodeMap.Meta[*node.NodeToken].Path
		dir := g.NodeMap.Meta[*node.NodeToken].Dir

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

func (g *Generator) GenerateWikiNav() error {
	// 1) 构建临时节点结构以表示父子关系
	type navNode struct {
		Token    string
		Title    string
		Path     string // 相对于 docs dir 的路径（用于 nav 链接）
		Parent   string
		Children []string
	}

	nodes := make(map[string]*navNode)
	// 保持顺序：NodeOrder 给出的顺序
	for _, nodeToken := range g.NodeMap.Entries {
		meta := g.NodeMap.Meta[nodeToken]
		rel, _ := filepath.Rel(g.Config.Output.DocsDir, meta.Path)

		// 一些 NodeMeta 字段可能不同名，请确保 NodeMap 中存在 Title 和 ParentNodeToken 字段
		title := meta.Node.Title
		parent := meta.Node.ParentNodeToken

		nodes[nodeToken] = &navNode{
			Token:  nodeToken,
			Title:  *title,
			Path:   rel,
			Parent: *parent,
		}
	}

	// 建立子关系
	for tok, n := range nodes {
		if n.Parent != "" {
			if p, ok := nodes[n.Parent]; ok {
				p.Children = append(p.Children, tok)
			}
		}
	}

	// 2) 递归构造 nav 数据结构，符合 mkdocs 的 nav 要求
	var buildEntry func(string) interface{}
	buildEntry = func(token string) interface{} {
		n := nodes[token]
		if n == nil {
			return nil
		}
		if len(n.Children) == 0 {
			// 叶子节点 -> 标题: path
			return map[string]interface{}{n.Title: n.Path}
		}
		// 有子节点 -> 标题: [ ...children... ]
		childrenList := make([]interface{}, 0, len(n.Children))
		for _, childTok := range n.Children {
			entry := buildEntry(childTok)
			if entry != nil {
				childrenList = append(childrenList, entry)
			}
		}
		return map[string]interface{}{n.Title: childrenList}
	}

	navList := make([]interface{}, 0)
	for _, tok := range g.NodeMap.Entries {
		meta := g.NodeMap.Meta[tok]
		if *meta.Node.ParentNodeToken == "" {
			entry := buildEntry(tok)
			if entry != nil {
				navList = append(navList, entry)
			}
		}
	}

	// 3) 把 nav 插回到目标 yaml 文件（如果存在则替换 nav 字段，否则新建文件）
	yamlPath := g.Config.Output.YamlPath
	if yamlPath == "" {
		return fmt.Errorf("yaml path not configured in output")
	}

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(yamlPath), 0755); err != nil {
		return err
	}

	var root map[string]interface{}
	if _, err := os.Stat(yamlPath); err == nil {
		// 文件存在 -> 读取并解析
		b, err := os.ReadFile(yamlPath)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(b, &root); err != nil {
			// 解析失败也可以覆盖新建，但这里返回错误以便用户确认
			return err
		}
	} else {
		root = make(map[string]interface{})
		// 保留默认 site_name，若需要可改
		if _, ok := root["site_name"]; !ok {
			root["site_name"] = "feishu2mkdocs"
		}
	}

	// 更新 nav 字段
	root["nav"] = navList

	out, err := yaml.Marshal(root)
	if err != nil {
		return err
	}

	if err := os.WriteFile(yamlPath, out, 0644); err != nil {
		return err
	}

	fmt.Println("Nav 已写入到:", yamlPath)
	return nil
}
