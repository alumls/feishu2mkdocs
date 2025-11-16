package service

import (
	"feishu2mkdocs/utils"
	"fmt"

	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type WikiMap struct {
	WikiMapNode map[string]*WikiMeta
}

type WikiMeta struct {
	Dir string
	Path string
	ChildNodeTokens []string
	IsShortCut bool
	Node *larkwiki.Node
}

func NewWikiMap() *WikiMap {
	return &WikiMap{
		WikiMapNode: make(map[string]*WikiMeta),
	}
}

func (m *WikiMap) WikiMapAddNode (node *larkwiki.Node, path string, isShortCut bool) {
	m.WikiMapNode[*node.NodeToken] = &WikiMeta{
		ChildNodeTokens: make([]string, 0),
		IsShortCut: isShortCut,
		Node: node,
	}
}

func (m *WikiMap) WikiMapBuildFromFlatNodes(nodes []*larkwiki.Node, docsRoot string) error {
	for _, node := range nodes {
		if node.NodeToken == nil {
			continue
		}
		isShortCut := false
		if *node.ObjType == "shortcut" {
			isShortCut = true
		}
		m.WikiMapAddNode(node, "", isShortCut)
	}

	for _, node := range nodes {
		if *node.ParentNodeToken != "" {
			m.WikiMapNodeAddChild(*node.ParentNodeToken, *node.NodeToken)
		}
	}

	for _, node := range nodes {
		rootPath, err := m.WikiMapNodeResolveRootPath(*node.NodeToken, docsRoot)
		if err != nil {
			return err
		}
		fileName, err := m.WikiMapNodeResolveFileName(*node.NodeToken)
		if err != nil {
			return err
		}

		if *node.HasChild {
			m.WikiMapNode[*node.NodeToken].Path = rootPath + "/" + fileName + "/" + fileName + ".md"
			m.WikiMapNode[*node.NodeToken].Dir = rootPath + "/" + fileName
		} else {
			m.WikiMapNode[*node.NodeToken].Path = rootPath + "/" + fileName + ".md"
			m.WikiMapNode[*node.NodeToken].Dir = rootPath
		}
	}
	return nil
}

func (m *WikiMap) WikiMapNodeResolveRootPath(nodeToken string, docsRoot string) (string, error){
	if _, ok := m.WikiMapNode[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	if *m.WikiMapNode[nodeToken].Node.ParentNodeToken == "" {
		return docsRoot, nil
	}

	parentNodeToken := *m.WikiMapNode[nodeToken].Node.ParentNodeToken
	parentNodeTitle := *m.WikiMapNode[parentNodeToken].Node.Title
	if parentNodeTitle == "" {
		parentNodeTitle = "untitled-" + parentNodeToken
	}

	rootPath, _ := m.WikiMapNodeResolveRootPath(parentNodeToken, docsRoot)

	return  rootPath + "/" + utils.SanitizeFileName(parentNodeTitle), nil
}

func (m *WikiMap) WikiMapNodeResolveFileName(nodeToken string) (string, error){
	if _, ok := m.WikiMapNode[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	title := *m.WikiMapNode[nodeToken].Node.Title
	if title == "" {
		title = "untitled-" + nodeToken
	}
	return utils.SanitizeFileName(title), nil
}

func (m *WikiMap) WikiMapNodeAddChild(nodeToken string, childNodeToken string) error{
	if _, ok := m.WikiMapNode[nodeToken]; !ok {
		return fmt.Errorf("missing Node: %s", nodeToken)
	}
	m.WikiMapNode[nodeToken].ChildNodeTokens = append(m.WikiMapNode[nodeToken].ChildNodeTokens, childNodeToken)
	return nil
}
