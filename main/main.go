package main

import (
	"context"
	"feishu2mkdocs/core"
	"feishu2mkdocs/utils"
	"fmt"
)

func main() {
	//parentNodeToken := ""
	client := core.NewClient("cli_a877a088c478500c", "vVUcx8YentVNvUYM7UA75d1LbIR7eAgX")
	nodes, err := client.GetDocumentBlockAll(context.Background(), "GOWNd16b4oGagxxbxvZcK1pEnZe")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(utils.PrettyPrint(nodes))
}
