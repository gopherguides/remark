package md

import (
	"testing"

	"github.com/gopherguides/remark/internal/rpm"
	"github.com/gopherguides/remark/px"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	mod := &px.Module{
		FS:      rpm.ModuleFS(t),
		ModName: "module.md",
	}

	parser := NewParser(mod)

	tags, err := parser.Parse()
	r.NoError(err)
	r.NotNil(tags)
	for _, t := range tags {
		r.Equal("section", t.TagName())
	}

	r.Len(tags, 1)

	// sinc, ok := tags[8].(*Section)
	// r.True(ok)
	// x := sinc.String()
	// r.Contains(x, "<h2>Get On Gopher Slack</h2>")
	// r.NotContains(x, `<include src="resources.md">`)

	// sec, ok := tags[0].(*Section)
	// r.True(ok)
	// r.Equal("section", sec.TagName())
	// r.Len(sec.GetChildren(), 5)
	// r.Equal("Intro to Go: Getting Started", sec.Title)

	// h1 := sec.GetChildren()[0]
	// r.Equal("h1", h1.TagName())
	// r.Len(h1.GetChildren(), 1)
	// r.Equal("<h1>Intro to Go: Getting Started</h1>", h1.String())

	// p, ok := sec.GetChildren()[1].(*remark.Generic)
	// r.True(ok)
	// r.Equal("p", p.TagName())
	// r.Len(p.GetChildren(), 1)
}
