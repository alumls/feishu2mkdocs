package main

import (
	//"context"
	"feishu2mkdocs/core"
	"feishu2mkdocs/utils"
	//"os"
	"fmt"
)

func main() {
	config, err := core.ReadFromConfigFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(utils.PrettyPrint(config))
}