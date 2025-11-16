package core

import (
	"context"
	//"feishu2mkdocs/utils"
	//"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
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

func (c *Client) GetWikiNodeList(ctx context.Context, spaceId string, parentNodetoken *string) ([]*larkwiki.Node, error) {
	// 创建请求对象
	req := larkwiki.NewListSpaceNodeReqBuilder().
		SpaceId(spaceId).
		ParentNodeToken(*parentNodetoken).
		Build()

	// 发送请求
	resp, err := c.larkClient.Wiki.V2.SpaceNode.List(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, err
	}

	// 打印测试
	//fmt.Println(utils.PrettyPrint(resp))

	nodes := resp.Data.Items
	previousPageToken := ""

	for *resp.Data.HasMore && previousPageToken != *resp.Data.PageToken {
		previousPageToken = *resp.Data.PageToken
		req := larkwiki.NewListSpaceNodeReqBuilder().
			SpaceId(spaceId).
			PageToken(*resp.Data.PageToken).
			Build()

		resp, err := c.larkClient.Wiki.V2.SpaceNode.List(context.Background(), req)

		if err != nil {
			return nil, err
		}

		nodes = append(nodes, resp.Data.Items...)
	}

	// 打印测试
	//fmt.Println(utils.PrettyPrint(nodes))

	return nodes, nil
}

func(c *Client) GetWikiNodeListAll(ctx context.Context, spaceId string, parentNodeToken *string) ([]*larkwiki.Node, error) {
	currentNodes, err := c.GetWikiNodeList(ctx, spaceId, parentNodeToken)
	resultNodes := currentNodes
	for _, node := range currentNodes {
		if *node.HasChild {
			nodes, err := c.GetWikiNodeListAll(ctx, spaceId, node.NodeToken)
			if err != nil {
				return nil, err
			}
			resultNodes = append(resultNodes, nodes...)
		}
	}
	if err != nil {
		return nil, err
	}
	return resultNodes, nil
}

func (c *Client) GetDocumentBlockAll(ctx context.Context, documentId string) ([]*larkdocx.Block, error) {

	req := larkdocx.NewListDocumentBlockReqBuilder().
		PageSize(500).
		DocumentId(documentId).
		DocumentRevisionId(-1).
		Build()

	resp, err := c.larkClient.Docx.V1.DocumentBlock.List(context.Background(), req)

	if err != nil {
		return nil, err
	}

	blocks := resp.Data.Items
	previousPageToken := ""

	for *resp.Data.HasMore && previousPageToken != *resp.Data.PageToken {
		previousPageToken = *resp.Data.PageToken
		req := larkdocx.NewListDocumentBlockReqBuilder().
			DocumentId(documentId).
			PageToken(*resp.Data.PageToken).
			DocumentRevisionId(-1).
			Build()

		resp, err := c.larkClient.Docx.V1.DocumentBlock.List(context.Background(), req)

		if err != nil {
			return nil, err
		}

		blocks = append(blocks, resp.Data.Items...)
	}

	// 打印测试
	//fmt.Println(utils.PrettyPrint(blocks))

	return blocks, nil
}
