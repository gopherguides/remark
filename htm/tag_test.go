package htm

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_HTMLBody(t *testing.T) {
	r := require.New(t)

	src := `<include src="foo.go"></include>`
	root, err := html.Parse(strings.NewReader(src))

	r.NoError(err)
	r.NotNil(root)

	body := HTMLBody(root)

	r.NotNil(body)
	r.Equal("body", body.Data)
}

func Test_Parser_NewTag_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	p := testParser(t, rpm.ModuleFS(t))
	node := TextNode(t, "hello!")

	tag, err := p.NewTag(node)
	r.NoError(err)

	s, ok := tag.(remark.String)
	r.True(ok)
	r.Equal("hello!", s.String())
}

func Test_Parser_NewTag(t *testing.T) {
	t.Parallel()

	img := AttrNode(t, "img", map[string]string{
		"src": "assets/foo.png",
	})
	code := AttrNode(t, "code", map[string]string{
		"src": "src/main.go",
	})

	cab := rpm.ModuleFS(t)
	rpm.AppendFile(t, cab, "assets/foo.png")
	rpm.AppendFile(t, cab, "src/main.go")

	p := testParser(t, cab)
	type pass struct {
		name string
		node *html.Node
		exp  remark.Tag
	}

	table := []pass{
		{name: "b", node: ElementNode(t, "b"), exp: &remark.Generic{}},
		{name: "code", node: code, exp: &remark.CodeBlock{}},
		{name: "default", node: ElementNode(t, "foo"), exp: &remark.Generic{}},
		{name: "document", node: DocNode(t), exp: &Document{}},
		{name: "em", node: ElementNode(t, "em"), exp: &remark.Generic{}},
		{name: "i", node: ElementNode(t, "i"), exp: &remark.Generic{}},
		{name: "img", node: img, exp: &remark.Image{}},
		{name: "metadata", node: ElementNode(t, "details"), exp: remark.Metadata{}},
		{name: "string", node: TextNode(t, "hello!"), exp: remark.String("hello!")},
		{name: "strong", node: ElementNode(t, "strong"), exp: &remark.Generic{}},
	}

	for _, h := range HEADS.Allowed() {
		table = append(table, pass{name: h, node: ElementNode(t, h), exp: &remark.Heading{}})
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			tag, err := p.NewTag(tt.node)
			r.NoError(err)
			r.NotNil(tag)

			exp := reflect.TypeOf(tt.exp)
			act := reflect.TypeOf(tag)

			r.True(act.AssignableTo(exp), fmt.Sprintf("%v != %v", act, exp))
		})
	}
}
