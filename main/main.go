package main

import (
	"context"
	"feishu2mkdocs/core"
	//"feishu2mkdocs/utils"
	"os"
	"fmt"
)

func main() {
	parentNodeToken := ""
	spaceId := "7565233961329246212"
	client := core.NewClient("cli_a877a088c478500c", "vVUcx8YentVNvUYM7UA75d1LbIR7eAgX")
	nodes, _ := client.GetWikiNodeList(context.Background(), spaceId, &parentNodeToken)
	blocks, _ := client.GetDocumentBlockAll(context.Background(), "GOWNd16b4oGagxxbxvZcK1pEnZe")

	//fmt.Println(utils.PrettyPrint(nodes))

	//fmt.Println(utils.PrettyPrint(blocks))

	config := core.OutputConfig{}
	parser := core.NewParser(config)
	markdown := parser.ParseDocxsContent(nodes[0], blocks)
	
	parseroutput := markdown

	er := os.WriteFile("md.txt", []byte(parseroutput), 0644)
    if er != nil {
        fmt.Println("写入文件失败:", er)
        return
    }

    fmt.Println("输出已保存到 md.txt")
}
