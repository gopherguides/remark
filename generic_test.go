package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Generic_Append(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	generic := NewGeneric("div")
	child := String("test content")

	generic.Append(child)

	r.Len(generic.Children, 1)
	r.Equal(child, generic.Children[0])
}

func Test_Generic_GetChildren(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	generic := NewGeneric("div")
	child1 := String("test1")
	child2 := String("test2")

	generic.Children = append(generic.Children, child1, child2)

	children := generic.GetChildren()
	r.Len(children, 2)
	r.Equal(child1, children[0])
	r.Equal(child2, children[1])
}

func Test_Generic_TagName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	generic := NewGeneric("div")
	r.Equal("div", generic.TagName())
}

func Test_Generic_Metadata(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	generic := NewGeneric("div")

	// Test initial metadata
	metadata := generic.Metadata()
	r.NotNil(metadata)
	r.Equal(Metadata{}, metadata)

	// Test that it initializes if nil
	generic.Data = nil
	metadata = generic.Metadata()
	r.NotNil(metadata)
	r.Equal(Metadata{}, metadata)
}

func Test_Generic_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	t.Run("with name and children", func(t *testing.T) {
		generic := NewGeneric("div")
		generic.Set("class", "container")
		generic.Children = append(generic.Children, String("Hello World"))

		result := generic.String()
		r.Contains(result, "<div")
		r.Contains(result, `class="container"`)
		r.Contains(result, "Hello World")
		r.Contains(result, "</div>")
	})

	t.Run("with name but no children", func(t *testing.T) {
		generic := NewGeneric("br")
		generic.Set("class", "break")

		result := generic.String()
		r.Contains(result, "<br")
		r.Contains(result, `class="break"`)
	})

	t.Run("without name", func(t *testing.T) {
		generic := &Generic{
			Name:       "",
			Attributes: Attributes{},
			Children:   Tags{String("just text")},
		}

		result := generic.String()
		r.Equal("just text", result)
	})
}

func Test_NewGeneric(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	t.Run("normal case", func(t *testing.T) {
		generic := NewGeneric("DIV")
		r.Equal("div", generic.Name)
		r.NotNil(generic.Attributes)
		r.NotNil(generic.Data)
		r.Len(generic.Children, 0) // Children is initialized as nil slice
	})

	t.Run("with whitespace", func(t *testing.T) {
		generic := NewGeneric("  SPAN  ")
		r.Equal("span", generic.Name)
	})

	t.Run("empty name", func(t *testing.T) {
		generic := NewGeneric("")
		r.Equal("", generic.Name)
		r.NotNil(generic.Attributes)
		r.NotNil(generic.Data)
		r.Len(generic.Children, 0) // Children is initialized as nil slice
	})
}
