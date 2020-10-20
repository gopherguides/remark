package md

import "github.com/gomarkdown/markdown/parser"

func NewParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs |
		parser.SuperSubscript | parser.OrderedListStart | parser.Footnotes
	return parser.NewWithExtensions(extensions)

}
