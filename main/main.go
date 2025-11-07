package main

import (
	"context"
	"feishu2mkdocs/core"
	"feishu2mkdocs/utils"
	"fmt"
)

func main() {
	parentNodeToken := "Tfehw6XAciPJn6kcZEgc59fEn7e"
	client := core.NewClient("cli_a877a088c478500c", "vVUcx8YentVNvUYM7UA75d1LbIR7eAgX")
	nodes, err := client.GetWikiNodeList(context.Background(), "7565233961329246212", &parentNodeToken)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(utils.PrettyPrint(nodes))
}
