package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Tags_Body(t *testing.T) {
	r := require.New(t)

	html := NewGeneric("html")
	html.Children = append(html.Children, NewGeneric("head"))
	body := NewGeneric("body")

	inc := NewGeneric("include")
	inc.Attributes["src"] = "foo.go"

	body.Children = append(body.Children, inc, String("hello"))
	html.Children = append(html.Children, body)

	b, ok := html.Children.Body()
	r.True(ok)
	r.NotNil(b)
}

func Test_Tags_Overview(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create a tag with metadata that has overview
	section := NewGeneric("section")
	section.Data = Metadata{"overview": "true"}

	// Add some children including heading and string content
	heading := &Heading{
		Generic: NewGeneric("h1"),
		Level:   1,
	}
	heading.Generic.Children = append(heading.Generic.Children, String("Title"))

	section.Children = append(section.Children, heading)
	section.Children = append(section.Children, String("This is overview content."))
	section.Children = append(section.Children, String("More content here."))

	tags := Tags{section}
	overview := tags.Overview()

	r.Contains(overview, "This is overview content.")
	r.Contains(overview, "More content here.")
	r.NotContains(overview, "Title") // Headings should be excluded
}

func Test_Tags_Overview_NoOverview(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	// Create tags without overview metadata
	section := NewGeneric("section")
	section.Children = append(section.Children, String("Regular content"))

	tags := Tags{section}
	overview := tags.Overview()

	r.Equal("", overview)
}

func Test_Tags_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tags := Tags{
		String("First content"),
		String("Second content"),
		NewGeneric("div"),
	}

	result := tags.String()
	expected := "First content Second content <div></div>"
	r.Equal(expected, result)
}

func Test_Tags_String_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tags := Tags{}
	result := tags.String()
	r.Equal("", result)
}

func Test_Tags_FindFirst(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	div := NewGeneric("div")
	span := NewGeneric("span")
	div.Children = append(div.Children, span)

	tags := Tags{div}

	// Test finding existing tag
	found, ok := tags.FindFirst("span")
	r.True(ok)
	r.Equal(span, found)

	// Test finding non-existing tag
	notFound, ok := tags.FindFirst("nonexistent")
	r.False(ok)
	r.Nil(notFound)
}

func Test_Tags_FindAll(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	div1 := NewGeneric("div")
	div2 := NewGeneric("div")
	span1 := NewGeneric("span")
	span2 := NewGeneric("span")

	div1.Children = append(div1.Children, span1)
	div2.Children = append(div2.Children, span2)

	tags := Tags{div1, div2}

	// Test finding all divs
	divs := tags.FindAll("div")
	r.Len(divs, 2)
	r.Contains(divs, div1)
	r.Contains(divs, div2)

	// Test finding all spans (nested)
	spans := tags.FindAll("span")
	r.Len(spans, 2)
	r.Contains(spans, span1)
	r.Contains(spans, span2)

	// Test finding non-existing tags
	notFound := tags.FindAll("nonexistent")
	r.Len(notFound, 0)
}
