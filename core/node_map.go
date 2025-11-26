package core

import (
	"feishu2mkdocs/utils"
	"fmt"

	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type NodeMap struct {
	Meta map[string]*NodeMeta
	FileNameCount map[string]int
	Entries []string
}

type NodeMeta struct {
	Dir string
	FileName string
	Path string
	ChildNodeTokens []string
	IsShortCut bool
	Node *larkwiki.Node
}

func NewNodeMap() *NodeMap {
	return &NodeMap{
		Meta: make(map[string]*NodeMeta),
		FileNameCount: make(map[string]int),
		Entries: make([]string, 0),
	}
}

func (m *NodeMap) AddNode (node *larkwiki.Node, path string, isShortCut bool) error {
	if utils.IsNilPointer(node) {
		return fmt.Errorf(
			"NodeMapAddNode error: node pointer is nil (node=%+v)",
			node,
		)
	}
	m.Meta[*node.NodeToken] = &NodeMeta{
		ChildNodeTokens: make([]string, 0),
		IsShortCut: isShortCut,
		Node: node,
	}
	m.Entries = append(m.Entries, *node.NodeToken)
	return nil
}

func (m *NodeMap) BuildFromFlatNodes(nodes []*larkwiki.Node, docsRoot string) error {
	for _, node := range nodes {
		isShortCut := false
		if *node.ObjType == "shortcut" {
			isShortCut = true
		}
		err := m.AddNode(node, "", isShortCut)
		if err != nil {
			return err
		}
	}

	for _, node := range nodes {
		if *node.ParentNodeToken != "" {
			err := m.NodeAddChild(*node.ParentNodeToken, *node.NodeToken)
			if err != nil {
				return err
			}
		}
	}

	for _, node := range nodes {
		rootPath, err := m.NodeResolveRootPath(*node.NodeToken, docsRoot)
		if err != nil {
			return err
		}
		fileName, err := m.NodeResolveFileName(*node.NodeToken)
		if err != nil {
			return err
		}	

		m.Meta[*node.NodeToken].FileName = fileName
		if *node.HasChild {
			m.Meta[*node.NodeToken].Path = rootPath + "/" + fileName + "/" + fileName + ".md"
			m.Meta[*node.NodeToken].Dir = rootPath + "/" + fileName
		} else {
			m.Meta[*node.NodeToken].Path = rootPath + "/" + fileName + ".md"
			m.Meta[*node.NodeToken].Dir = rootPath
		}
	}
	return nil
}

func (m *NodeMap) NodeResolveRootPath(nodeToken string, docsRoot string) (string, error){
	if _, ok := m.Meta[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	if *m.Meta[nodeToken].Node.ParentNodeToken == "" {
		return docsRoot, nil
	}

	parentNodeToken := *m.Meta[nodeToken].Node.ParentNodeToken
	parentNodeTitle := *m.Meta[parentNodeToken].Node.Title
	if parentNodeTitle == "" {
		parentNodeTitle = "untitled-" + parentNodeToken
	}

	rootPath, _ := m.NodeResolveRootPath(parentNodeToken, docsRoot)

	return  rootPath + "/" + utils.SanitizeFileName(parentNodeTitle), nil
}

func (m *NodeMap) NodeResolveFileName(nodeToken string) (string, error){
	if _, ok := m.Meta[nodeToken]; !ok {
		return "", fmt.Errorf("missing Node: %s", nodeToken)
	}
	title := *m.Meta[nodeToken].Node.Title
	if title == "" {
		title = "untitled-" + nodeToken
	}
	return utils.SanitizeFileName(title), nil
}

func (m *NodeMap) NodeAddChild(nodeToken string, childNodeToken string) error{
	if _, ok := m.Meta[nodeToken]; !ok {
		return fmt.Errorf("missing Node: %s", nodeToken)
	}
	m.Meta[nodeToken].ChildNodeTokens = append(m.Meta[nodeToken].ChildNodeTokens, childNodeToken)
	return nil
}
