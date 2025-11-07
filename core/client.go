package core

import (
	"context"
	"feishu2mkdocs/utils"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Client struct {
	larkClient *lark.Client
}

// 创建一个 Client 实例
// appId: 应用的 App ID
// appSecret: 应用的 App Secret
func NewClient(appId, appSecret string) *Client {
	return &Client{
		larkClient: lark.NewClient(appId, appSecret),
	}
}

func (c *Client) GetWikiNodeList(ctx context.Context, spaceId, token string) ([]*larkwiki.Node, error) {

	// 创建请求对象
	req := larkwiki.NewListSpaceNodeReqBuilder().
		SpaceId(spaceId).
		PageToken(token).
		Build()

	// 发送请求
	resp, err := c.larkClient.Wiki.V2.SpaceNode.List(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, err
	}

	// 打印测试
	fmt.Println(utils.PrettyPrint(resp))

	// TODO: 需要分页继续获得完整的NodeList，经过测试，子页面不会包含在List中，需要通过haschild判断后递归获取
	nodes := resp.Data.Items

	return nodes, nil
}
