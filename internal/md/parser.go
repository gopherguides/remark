package md

import "github.com/gomarkdown/markdown/parser"

func NewParser() *parser.Parser {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	return parser.NewWithExtensions(extensions)
}
