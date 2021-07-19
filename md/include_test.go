package md

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func Test_Parser_NewInclude(t *testing.T) {
	r := require.New(t)

	p := NewParser(week00)

	n := &html.Node{
		Data: "include",
		Attr: []html.Attribute{
			{Key: "src", Val: "assignments/assignment01/assignment.md"},
		},
	}

	i, err := p.NewInclude(n)
	r.NoError(err)
	r.NotNil(i)

	r.Equal("include", i.Name)
}
