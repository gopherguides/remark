package htm

import (
	"testing"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_Parser_Metadata(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	node := ElementNode(t, "details")
	node.FirstChild = TextNode(t, "a:A\nB:b\n")

	p := testParser(t, rpm.ModuleFS(t))

	md, err := p.parseMetadata(node)
	r.NoError(err)

	a, ok := md["a"]
	r.True(ok)
	r.Equal("A", a)

	b, ok := md["b"]
	r.True(ok)
	r.Equal("b", b)
}
