package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Metadata_TagName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{}
	r.Equal("metadata", md.TagName())
}

func Test_Metadata_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	html := md.HTML()
	r.Equal(template.HTML(""), html)
}

func Test_Metadata_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	result := md.String()
	r.Equal("", result)
}

func Test_Metadata_GetChildren(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	children := md.GetChildren()
	r.NotNil(children)
	r.Len(children, 0)
	r.Equal(Tags{}, children)
}

func Test_Metadata_Options(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	opts := md.Options()
	r.NotNil(opts)
	r.Len(opts, 0)
}

func Test_Metadata_Attrs(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	attrs := md.Attrs()
	r.NotNil(attrs)
	r.Len(attrs, 0)
	r.Equal(Attributes{}, attrs)
}

func Test_Metadata_AsTag(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	md := Metadata{
		"title":  "Test Title",
		"author": "Test Author",
	}

	// Test that Metadata implements Tag interface
	var tag Tag = md
	r.NotNil(tag)
	r.Equal("metadata", tag.TagName())
}
