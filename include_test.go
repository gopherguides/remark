package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Include_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	include := Include{
		Generic: &Generic{
			Attributes: Attributes{
				"src": "test.md",
			},
			Children: Tags{
				String("This is included content"),
				String("More content"),
			},
		},
	}

	result := include.String()
	r.Equal("This is included content More content", result)
}

func Test_Include_String_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	include := Include{
		Generic: &Generic{
			Attributes: Attributes{
				"src": "test.md",
			},
			Children: Tags{},
		},
	}

	result := include.String()
	r.Equal("", result)
}

func Test_Include_String_WithNilGeneric(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	include := Include{
		Generic: nil,
	}

	// This should panic or handle gracefully
	r.Panics(func() {
		_ = include.String()
	})
}

func Test_Include_AsTag(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	include := Include{
		Generic: &Generic{
			Name:       "include",
			Attributes: Attributes{"src": "test.md"},
			Children:   Tags{String("content")},
		},
	}

	// Test that Include implements Tag interface
	var tag Tag = &include
	r.NotNil(tag)
}

func Test_Include_WithComplexChildren(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	child1 := NewGeneric("p")
	child1.Children = append(child1.Children, String("Paragraph content"))

	child2 := NewGeneric("div")
	child2.Set("class", "container")
	child2.Children = append(child2.Children, String("Div content"))

	include := Include{
		Generic: &Generic{
			Attributes: Attributes{
				"src": "complex.md",
			},
			Children: Tags{child1, child2},
		},
	}

	result := include.String()
	r.Contains(result, "Paragraph content")
	r.Contains(result, "Div content")
	r.Contains(result, "<p>")
	r.Contains(result, "<div")
	r.Contains(result, `class="container"`)
}
