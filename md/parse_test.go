package md

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Parse(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	fp := week00
	b, err := ioutil.ReadFile(fp)
	r.NoError(err)

	tags, err := Parse(filepath.Dir(fp), b)
	r.NoError(err)
	r.NotNil(tags)
	for _, t := range tags {
		r.Equal("section", t.TagName())
	}

	r.Len(tags, 9)

	sinc, ok := tags[8].(*Section)
	r.True(ok)
	x := sinc.String()
	r.Contains(x, "<h2>Get On Gopher Slack</h2>")
	r.NotContains(x, `<include src="resources.md">`)

	sec, ok := tags[0].(*Section)
	r.True(ok)
	r.Equal("section", sec.TagName())
	r.Len(sec.GetChildren(), 5)
	r.Equal("Intro to Go: Getting Started", sec.Title)

	h1 := sec.GetChildren()[0]
	r.Equal("h1", h1.TagName())
	r.Len(h1.GetChildren(), 1)
	r.Equal("<h1>Intro to Go: Getting Started</h1>", h1.String())

	p, ok := sec.GetChildren()[1].(*remark.Generic)
	r.True(ok)
	r.Equal("p", p.TagName())
	r.Len(p.GetChildren(), 1)
}

func Test_Parse_LineEndings(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Test content with horizontal rules using different line endings
	contentLF := "# Section 1\n\nContent 1\n\n---\n\n# Section 2\n\nContent 2"
	contentCRLF := "# Section 1\r\n\r\nContent 1\r\n\r\n---\r\n\r\n# Section 2\r\n\r\nContent 2"

	// Parse with LF line endings
	tagsLF, err := Parse(".", []byte(contentLF))
	r.NoError(err)
	r.Len(tagsLF, 2)

	// Parse with CRLF line endings - should produce same result
	tagsCRLF, err := Parse(".", []byte(contentCRLF))
	r.NoError(err)
	r.Len(tagsCRLF, 2)

	// Both should have same number of sections
	r.Equal(len(tagsLF), len(tagsCRLF))
}
