package utils

import (
	"fmt"
	"path/filepath"
	"strings"
)

type NavNode struct {
	File     string
	Children map[string]*NavNode
	Entries  []NavEntry
}

type NavEntry struct {
	Title string
	Token string
}

func NewNavNode() *NavNode {
	return &NavNode{
		Children: make(map[string]*NavNode),
	}
}

func SplitPath(path string) []string {
	// Normalize both Windows and POSIX separators to '/' then split
	normalized := strings.ReplaceAll(path, "\\\\", "/")
	normalized = strings.ReplaceAll(normalized, string(filepath.Separator), "/")
	parts := strings.Split(normalized, "/")
	// filter out empty parts (leading/trailing slashes)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func Rel(path string, rootDir string) (string, error) {
	rel, err := filepath.Rel(rootDir, path)
	return rel, err
}

func InsertNavNode(root *NavNode, path string, title string) error {
	parts := SplitPath(path)

	if len(parts) == 0 {
		return fmt.Errorf("empty path")
	}

	cur := root

	// walk directories for all parts except the last one
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if cur.Children[part] == nil {
			cur.Children[part] = NewNavNode()
			// record created folder in order (NavEntry with empty token)
			cur.Entries = append(cur.Entries, NavEntry{Title: part, Token: ""})
		} else {
			// ensure entry exists for existing child
			found := false
			for _, e := range cur.Entries {
				if e.Title == part {
					found = true
					break
				}
			}
			if !found {
				cur.Entries = append(cur.Entries, NavEntry{Title: part, Token: ""})
			}
		}
		cur = cur.Children[part]
	}

	// If a child with the given title already exists under current node, fail
	if cur.Children[title] != nil {
		return fmt.Errorf("title %s already exists in path %s", title, path)
	}

	// create node for the title and assign file path
	cur.Children[title] = NewNavNode()
	cur.Children[title].File = path

	// record insertion order for this parent node
	// ensure there isn't already an entry with the same title
	exists := false
	for _, e := range cur.Entries {
		if e.Title == title {
			exists = true
			break
		}
	}
	if !exists {
		cur.Entries = append(cur.Entries, NavEntry{Title: title, Token: path})
	}

	return nil
}

func BuildYaml(node *NavNode) []any {
	var out []any

	// Respect insertion order stored in Entries. If Entries is empty, fall back to map order.
	if len(node.Entries) > 0 {
		for _, entry := range node.Entries {
			name := entry.Title
			child, ok := node.Children[name]
			if !ok {
				// ignore stale entry
				continue
			}
			if len(child.Children) == 0 && child.File != "" {
				out = append(out, map[string]string{name: child.File})
			} else {
				out = append(out, map[string]any{name: BuildYaml(child)})
			}
		}
		return out
	}

	// fallback: unordered traversal if Entries not populated
	for name, child := range node.Children {
		if len(child.Children) == 0 && child.File != "" {
			out = append(out, map[string]string{name: child.File})
		} else {
			out = append(out, map[string]any{name: BuildYaml(child)})
		}
	}
	return out
}
