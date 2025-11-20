package core

import (
	"context"
	"fmt"
	"feishu2mkdocs/utils"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Client struct {
	larkClient *lark.Client
}

// NewClient: 创建一个 Client 实例。
//
// 参数：
//   - appId: 应用的 App ID
//   - appSecret: 应用的 App Secret
func NewClient(appId, appSecret string) *Client {
	return &Client{
		larkClient: lark.NewClient(appId, appSecret),
	}
}

// GetWikiNodeList: 获取知识库节点列表。获取的节点列表只有父节点下的第一级子节点，如果需要获取包含父节点下的所有子节点，可以使用GetWikiNoteListAll方法。
//
// 参数：
//   - ctx: 上下文对象
//   - spaceId: 知识库的 Space ID
//   - parentNodetoken: 父节点的 Node Token。如果为空字符串，则指向知识库根目录。
func (c *Client) GetWikiNodeList(ctx context.Context, spaceId string, parentNodeToken *string) ([]*larkwiki.Node, error) {
	// 创建请求对象
	if utils.IsNilPointer(parentNodeToken) {
		return nil, fmt.Errorf(
			"GetWikiNodeList error: parent node token is nil (node=%+v)",
			parentNodeToken,
		)
	}

	req := larkwiki.NewListSpaceNodeReqBuilder().
		SpaceId(spaceId).
		ParentNodeToken(*parentNodeToken).
		Build()

	// 发送请求
	resp, err := c.larkClient.Wiki.V2.SpaceNode.List(context.Background(), req)

	// 处理错误
	if err != nil {
		return nil, err
	}

	// 打印测试
	//fmt.Println(utils.PrettyPrint(resp))

	//分页处理
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

// GetWikiNodeListAll: 获取知识库中父节点下所有节点列表。
//
// 参数：
//   - ctx: 上下文对象
//   - spaceId: 知识库的 Space ID
//   - parentNodetoken: 父节点的 Node Token。如果为空字符串，则指向知识库根目录。
func (c *Client) GetWikiNodeListAll(ctx context.Context, spaceId string, parentNodeToken *string) ([]*larkwiki.Node, error) {
	if utils.IsNilPointer(parentNodeToken) {
		return nil, fmt.Errorf(
			"GetWikiNodeListAll error: parent node token is nil (node=%+v)",
			parentNodeToken,
		)
	}

	localNodes, err := c.GetWikiNodeList(ctx, spaceId, parentNodeToken)

	if err != nil {
		return nil, err
	}

	//TODO: 这里的递归得到的结果为BFS结果，后续处理起来可能会比较麻烦，可以优化一下。

	allNodes := localNodes

	for _, node := range localNodes {
		if *node.HasChild {
			nodes, err := c.GetWikiNodeListAll(ctx, spaceId, node.NodeToken)
			if err != nil {
				return nil, err
			}
			allNodes = append(allNodes, nodes...)
		}
	}

	return allNodes, nil
}

// GetDocumentBlockAll: 获取文档所有块。
//
// 参数：
//   - ctx: 上下文对象
//   - documentId: 文档的 Document ID（注意区分 Document ID 和 Node Token ）.
func (c *Client) GetDocumentBlockAll(ctx context.Context, documentId *string) ([]*larkdocx.Block, error) {
	if utils.IsNilPointer(documentId) {
		return nil, fmt.Errorf(
			"GetDocumentBlockAll error: document id is nil (documentId=%+v)",
			documentId,
		)
	}

	req := larkdocx.NewListDocumentBlockReqBuilder().
		PageSize(500).
		DocumentId(*documentId).
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
			DocumentId(*documentId).
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
