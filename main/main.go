package main

import (
	"context"
	"fmt"
	"feishu2mkdocs/core"
)

func main() {
	client := core.NewClient("cli_a877a088c478500c", "vVUcx8YentVNvUYM7UA75d1LbIR7eAgX")
	nodes, err := client.GetWikiNodeList(context.Background(), "7565233961329246212", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nodes)
}
