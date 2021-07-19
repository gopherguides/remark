package md

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_Parser_NewInclude(t *testing.T) {
	r := require.New(t)

	p := NewParser("")
	src := `<include src="foo.go"></include>`

	n, err := html.Parse(strings.NewReader(src))
	r.NoError(err)
	r.NotNil(n)

	i, err := p.NewInclude(n)
	r.NoError(err)
	r.NotNil(i)

	r.Equal("include", i.Name)
}
