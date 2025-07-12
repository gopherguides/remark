package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_String_Options(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")
	opts := s.Options()

	r.NotNil(opts)
	r.Len(opts, 0)
}

func Test_String_Attrs(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")
	attrs := s.Attrs()

	r.NotNil(attrs)
	r.Len(attrs, 0)
	r.Equal(Attributes{}, attrs)
}

func Test_String_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")
	result := s.String()

	r.Equal("test content", result)
}

func Test_String_TagName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")
	name := s.TagName()

	r.Equal("string", name)
}

func Test_String_GetChildren(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")
	children := s.GetChildren()

	r.NotNil(children)
	r.Len(children, 0)
	r.Equal(Tags{}, children)
}

func Test_String_AsTag(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test content")

	// Test that String implements Tag interface
	var tag Tag = s
	r.NotNil(tag)
	r.Equal("string", tag.TagName())
	r.Equal("test content", tag.String())
}

func Test_String_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("")

	r.Equal("", s.String())
	r.Equal("string", s.TagName())
	r.Len(s.GetChildren(), 0)
}

func Test_String_WithSpecialChars(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	s := String("test <>&\"' content")

	r.Equal("test <>&\"' content", s.String())
}
