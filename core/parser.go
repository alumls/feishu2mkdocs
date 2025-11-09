package core

import (
	"fmt"
	"reflect"
	"strings"

	"feishu2mkdocs/utils"

	larkdocx "github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	larkwiki "github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Parser struct {
	ImgTokens []string
	blockMap  map[string]*larkdocx.Block
}

func NewParser(config OutputConfig) *Parser {
	return &Parser{
		ImgTokens: make([]string, 0),
		blockMap:  make(map[string]*larkdocx.Block),
	}
}

func (p *Parser) ParseDocxsContent(node *larkwiki.Node, blocks []*larkdocx.Block) string {
	for _, block := range blocks {
		p.blockMap[*block.BlockId] = block
	}

	entryBlock := p.blockMap[*node.ObjToken]
	return p.ParseDocxBlock(entryBlock, 0)
}

func (p *Parser) ParseDocxBlock(b *larkdocx.Block, indentLevel int) string {
	buf := new(strings.Builder)
	buf.WriteString(strings.Repeat("\t", indentLevel))
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
	case DocxBlockTypeQuote:
	case DocxBlockTypeCallout:
	case DocxBlockTypeDivider:
	case DocxBlockTypeImage:
	case DocxBlockTypeFile:
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

