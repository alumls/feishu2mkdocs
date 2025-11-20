package service

import (
	"feishu2mkdocs/utils"
	"fmt"

	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type NodeMap struct {
	NodeMeta map[string]*NodeMeta
}

type NodeMeta struct {
	Dir string
	Path string
	ChildNodeTokens []string
	IsShortCut bool
	Node *larkwiki.Node
}

func NewNodeMap() *NodeMap {
	return &NodeMap{
		NodeMeta: make(map[string]*NodeMeta),
	}
}

func (m *NodeMap) NodeMapAddNode (node *larkwiki.Node, path string, isShortCut bool) error {
	if utils.IsNilPointer(node) {
		return fmt.Errorf(
			"NodeMapAddNode error: node pointer is nil (node=%+v)",
			node,
		)
	}
	m.NodeMeta[*node.NodeToken] = &NodeMeta{
		ChildNodeTokens: make([]string, 0),
		IsShortCut: isShortCut,
		Node: node,
	}
	return nil
}

func (m *NodeMap) NodeMapBuildFromFlatNodes(nodes []*larkwiki.Node, docsRoot string) error {
	for _, node := range nodes {
		isShortCut := false
		if *node.ObjType == "shortcut" {
			isShortCut = true
		}
		err := m.NodeMapAddNode(node, "", isShortCut)
		if err != nil {
			return err
		}
	}

	for _, node := range nodes {
		if *node.ParentNodeToken != "" {
			err := m.NodeMapNodeAddChild(*node.ParentNodeToken, *node.NodeToken)
			if err != nil {
				return err
			}
		}
	}

	for _, node := range nodes {
		rootPath, err := m.NodeMapNodeResolveRootPath(*node.NodeToken, docsRoot)
		if err != nil {
			return err
		}
		fileName, err := m.NodeMapNodeResolveFileName(*node.NodeToken)
		if err != nil {
			return err
		}

		if *node.HasChild {
			m.NodeMeta[*node.NodeToken].Path = rootPath + "/" + fileName + "/" + fileName + ".md"
			m.NodeMeta[*node.NodeToken].Dir = rootPath + "/" + fileName
		} else {
			m.NodeMeta[*node.NodeToken].Path = rootPath + "/" + fileName + ".md"
			m.NodeMeta[*node.NodeToken].Dir = rootPath
		}
	}
	return nil
}

func (m *NodeMap) NodeMapNodeResolveRootPath(nodeToken string, docsRoot string) (string, error){
	if _, ok := m.NodeMeta[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	if *m.NodeMeta[nodeToken].Node.ParentNodeToken == "" {
		return docsRoot, nil
	}

	parentNodeToken := *m.NodeMeta[nodeToken].Node.ParentNodeToken
	parentNodeTitle := *m.NodeMeta[parentNodeToken].Node.Title
	if parentNodeTitle == "" {
		parentNodeTitle = "untitled-" + parentNodeToken
	}

	rootPath, _ := m.NodeMapNodeResolveRootPath(parentNodeToken, docsRoot)

	return  rootPath + "/" + utils.SanitizeFileName(parentNodeTitle), nil
}

func (m *NodeMap) NodeMapNodeResolveFileName(nodeToken string) (string, error){
	if _, ok := m.NodeMeta[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	title := *m.NodeMeta[nodeToken].Node.Title
	if title == "" {
		title = "untitled-" + nodeToken
	}
	return utils.SanitizeFileName(title), nil
}

func (m *NodeMap) NodeMapNodeAddChild(nodeToken string, childNodeToken string) error{
	if _, ok := m.NodeMeta[nodeToken]; !ok {
		return fmt.Errorf("missing Node: %s", nodeToken)
	}
	m.NodeMeta[nodeToken].ChildNodeTokens = append(m.NodeMeta[nodeToken].ChildNodeTokens, childNodeToken)
	return nil
}
