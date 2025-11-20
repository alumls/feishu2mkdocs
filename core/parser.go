package core

import (
	"fmt"
	"reflect"
	"strings"

	"feishu2mkdocs/utils"

	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

// TODO: 由于feishu和mkdocs的解析器差异，有些语法并不能在mkdocs中得到支持，但是目前Parser没有做这一层的处理，如果出些不支持的语法，应该给出相应的警告。

type Parser struct {
	ImgTokens []string
	blockMap  map[string]*larkdocx.Block
}

// NewParser: 创建一个Parser实例。
//
// 参数：
//   - config: 输出配置
func NewParser(config OutputConfig) *Parser {
	return &Parser{
		ImgTokens: make([]string, 0),
		blockMap:  make(map[string]*larkdocx.Block),
	}
}

// ParseDocxsContent: 解析文档内容。该函数往往作为外部调用，用于解析知识库节点下的文档内容。
//
// 参数：
//   - node: 知识库节点
//   - blocks: 文档块
func (p *Parser) ParseDocxsContent(node *larkwiki.Node, blocks []*larkdocx.Block) (string, error) {
	if utils.IsNilPointer(node) {
		return "", fmt.Errorf(
			"ParseDocxsContent error: node pointer is nil (node=%+v)",
			node,
		)
	}

	for _, block := range blocks {
		if utils.IsNilPointer(block) {
			return "", fmt.Errorf(
				"ParseDocxsContent error: block pointer is nil (block=%+v)",
				block,
			)
		}
		if utils.IsNilPointer(block.BlockId) {
			return "", fmt.Errorf(
				"ParseDocxsContent error: block.BlockId pointer is nil (blockId=%+v)",
				block.BlockId,
			)
		}
		p.blockMap[*block.BlockId] = block
	}

	// TODO: 目前存在生成的md文件有多余空行的问题，可以在这里对解析后的字符串进行处理，移除多余空行。

	if utils.IsNilPointer(node.ObjToken) {
		return "", fmt.Errorf(
			"ParseDocxsContent error: node.ObjToken pointer is nil (nodeObjToken=%+v)",
			node.ObjToken,
		)
	}
	entryBlock := p.blockMap[*node.ObjToken]
	return p.ParseDocxBlock(entryBlock, 0), nil
}

func (p *Parser) ParseDocxBlock(b *larkdocx.Block, indentLevel int) string {
	buf := new(strings.Builder)
	buf.WriteString(strings.Repeat("    ", indentLevel))
	switch *b.BlockType {
	case DocxBlockTypePage:
		buf.WriteString(p.ParseDocxBlockPage(b))
	case DocxBlockTypeText:
		buf.WriteString(p.ParseDocxBlockText(b.Text))
	case DocxBlockTypeHeading1:
		buf.WriteString(p.ParseDocxBlockHeading(b, 1))
	case DocxBlockTypeHeading2:
		buf.WriteString(p.ParseDocxBlockHeading(b, 2))
	case DocxBlockTypeHeading3:
		buf.WriteString(p.ParseDocxBlockHeading(b, 3))
	case DocxBlockTypeHeading4:
		buf.WriteString(p.ParseDocxBlockHeading(b, 4))
	case DocxBlockTypeHeading5:
		buf.WriteString(p.ParseDocxBlockHeading(b, 5))
	case DocxBlockTypeHeading6:
		buf.WriteString(p.ParseDocxBlockHeading(b, 6))
	case DocxBlockTypeHeading7:
		buf.WriteString(p.ParseDocxBlockHeading(b, 7))
	case DocxBlockTypeHeading8:
		buf.WriteString(p.ParseDocxBlockHeading(b, 8))
	case DocxBlockTypeHeading9:
		buf.WriteString(p.ParseDocxBlockHeading(b, 9))
	case DocxBlockTypeOrdered:
		buf.WriteString(p.ParseDocxBlockOrdered(b, indentLevel))
	case DocxBlockTypeBullet:
		buf.WriteString(p.ParseDocxBlockBullet(b, indentLevel))
	case DocxBlockTypeTodo:
		buf.WriteString(p.ParseDocxBlockTodo(b, indentLevel))
	case DocxBlockTypeCode:
		buf.WriteString(p.ParseDocxBlockCode(b))
	case DocxBlockTypeQuoteContainer:
		buf.WriteString(p.ParseDocxBlockQuote(b))
	case DocxBlockTypeCallout:
		buf.WriteString(p.ParseDocxBlockCallout(b))
	case DocxBlockTypeDivider:
		buf.WriteString("---\n")
	case DocxBlockTypeImage:
		// TODO: 这里需要做解析图片函数的入口。
	case DocxBlockTypeFile:
		// TODO: 这里需要做解析文件的入口。
	default:
	}
	return buf.String()
}

func (p *Parser) ParseDocxBlockPage(b *larkdocx.Block) string {
	buf := new(strings.Builder)

	buf.WriteString("# ")
	buf.WriteString(p.ParseDocxBlockText(b.Page))
	buf.WriteString("\n")

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, 0))
		buf.WriteString("\n")
	}

	return buf.String()
}

func (p *Parser) ParseDocxBlockText(b *larkdocx.Text) string {
	buf := new(strings.Builder)
	numElem := len(b.Elements)
	for _, e := range b.Elements {
		inline := numElem > 1
		buf.WriteString(p.ParseDocxTextElement(e, inline))
	}
	buf.WriteString("\n")
	return buf.String()
}

func (p *Parser) ParseDocxBlockHeading(b *larkdocx.Block, headingLevel int) string {
	buf := new(strings.Builder)

	buf.WriteString(strings.Repeat("#", headingLevel))
	buf.WriteString(" ")

	headingText := reflect.ValueOf(b).Elem().FieldByName(fmt.Sprintf("Heading%d", headingLevel))
	buf.WriteString(p.ParseDocxBlockText(headingText.Interface().(*larkdocx.Text)))

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, 0))
	}

	return buf.String()
}

func (p *Parser) ParseDocxBlockOrdered(b *larkdocx.Block, indentLevel int) string {
	buf := new(strings.Builder)

	parent := p.blockMap[*b.ParentId]
	order := 1
	for idx, child := range parent.Children {
		if child == *b.BlockId {
			for i := idx-1; i >=0; i-- {
				if *p.blockMap[parent.Children[i]].BlockType == DocxBlockTypeOrdered {
					order++
				} else {
					break
				}
			}
		}
	}

	buf.WriteString(fmt.Sprintf("%d. ", order))
	buf.WriteString(p.ParseDocxBlockText(b.Ordered))

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, indentLevel+1))
	}

	return buf.String()
}

func (p *Parser) ParseDocxBlockBullet(b *larkdocx.Block, indentLevel int) string {
	buf := new(strings.Builder)

	buf.WriteString("- ")
	buf.WriteString(p.ParseDocxBlockText(b.Bullet))

	for _, childId := range b.Children{
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, indentLevel + 1))
	}

	return buf.String()
}

func (p *Parser) ParseDocxBlockTodo(b *larkdocx.Block, indentLevel int) string {
	buf := new(strings.Builder)

	if *b.Todo.Style.Done {
		buf.WriteString("- [x] ")
	} else{
		buf.WriteString("- [ ] ")
	}
	buf.WriteString(p.ParseDocxBlockText(b.Todo))

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, indentLevel + 1))
	}

	return buf.String()
}

func (p *Parser) ParseDocxBlockCode(b *larkdocx.Block) string {
	buf := new(strings.Builder)

	buf.WriteString("```" + DocxCodeLang2MdStr[*b.Code.Style.Language] + "\n")
	buf.WriteString(strings.TrimSpace(p.ParseDocxBlockText(b.Code)))
	buf.WriteString("\n```\n")

	return buf.String()
}

func (p *Parser) ParseDocxBlockQuote(b *larkdocx.Block) string {
	buf := new(strings.Builder)

	buf.WriteString("> ")

	startIndex := buf.Len()

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, 0))
		buf.WriteString("\n")
	}

	s := buf.String()

	newPart := s[startIndex: len(s)-2]

	processed := new(strings.Builder)
	for i := 0; i < len(newPart); i++ {
		processed.WriteByte(newPart[i])
		if newPart[i] == '\n' && i < len(newPart)-1 {
			processed.WriteString("> ")
		}
	}
	processed.WriteString("\n")

	return s[:startIndex] + processed.String()
}

func (p *Parser) ParseDocxBlockCallout(b *larkdocx.Block) string {
	buf := new(strings.Builder)

	calloutType := "note"

	if newCalloutType, ok := DocxCalloutEmoji2MdStr[*b.Callout.EmojiId]; ok {
		calloutType = newCalloutType
	}

	buf.WriteString("!!! "+ calloutType + "\n\n    ")

	startIndex := buf.Len()

	for _, childId := range b.Children {
		childBlock := p.blockMap[childId]
		buf.WriteString(p.ParseDocxBlock(childBlock, 0))
		buf.WriteString("\n")
	}

	s := buf.String()

	newPart := s[startIndex: len(s)-2]

	processed := new(strings.Builder)
	for i := 0; i < len(newPart); i++ {
		processed.WriteByte(newPart[i])
		if newPart[i] == '\n' && i < len(newPart)-1 {
			processed.WriteString("    ")
		}
	}
	processed.WriteString("\n")

	return s[:startIndex] + processed.String()
}

func (p *Parser) ParseDocxTextElement(e *larkdocx.TextElement, inline bool) string {
	buf := new(strings.Builder)
	if e.TextRun != nil {
		buf.WriteString(p.ParseDocxTextElementTextRun(e.TextRun))
	}
	if e.Equation != nil {
		symbol := "$$"
		if inline {
			symbol = "$"
		}
		buf.WriteString(symbol + strings.TrimSuffix(*e.Equation.Content, "\n") + symbol)
	}
	return buf.String()
}

func (p *Parser) ParseDocxTextElementTextRun(tr *larkdocx.TextRun) string {
	buf := new(strings.Builder)
	postWrite := ""
	if style := tr.TextElementStyle; style != nil {
		if *style.Bold {
			buf.WriteString("**")
			postWrite = "**" + postWrite
		}
		if *style.Italic {
			buf.WriteString("*")
			postWrite = "*" + postWrite
		}
		if *style.Strikethrough {
			buf.WriteString("~~")
			postWrite = "~~" + postWrite
		}
		if *style.Underline {
			buf.WriteString("^^")
			postWrite = "^^" + postWrite
		}
		if *style.InlineCode {
			buf.WriteString("`")
			postWrite = "`" + postWrite
		}
		if style.TextColor != nil {
			buf.WriteString("==")
			postWrite = "==" + postWrite
		}
		if style.Link != nil {
			buf.WriteString("[")
			postWrite = fmt.Sprintf("](%s)", utils.UnescapeURL(*style.Link.Url)) + postWrite
		}
	}
	buf.WriteString(*tr.Content)
	buf.WriteString(postWrite)
	return buf.String()
}

