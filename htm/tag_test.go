package htm

import (
	"strings"
	"testing"

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

func Test_Parser_NewTag_IMG(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	rpm.AppendFile(t, cab, "assets/foo.png")
	p := testParser(t, cab)

	node := AttrNode(t, "img", map[string]string{
		"src": "assets/foo.png",
	})

	tag, err := p.NewTag(node)
	r.NoError(err)
	r.NotNil(tag)
}
