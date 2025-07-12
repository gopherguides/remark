package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Snippet_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippet := Snippet{
		Content:  "fmt.Println(\"Hello, World!\")",
		File:     "main.go",
		Language: "go",
		Name:     "hello",
		Start:    1,
		End:      3,
	}

	html := snippet.HTML()
	r.Equal(template.HTML("fmt.Println(\"Hello, World!\")"), html)
}

func Test_Snippet_HTML_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippet := Snippet{
		Content:  "",
		File:     "empty.go",
		Language: "go",
		Name:     "empty",
		Start:    0,
		End:      0,
	}

	html := snippet.HTML()
	r.Equal(template.HTML(""), html)
}

func Test_Snippet_HTML_WithHTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippet := Snippet{
		Content:  "<div>Hello &amp; World</div>",
		File:     "index.html",
		Language: "html",
		Name:     "greeting",
		Start:    1,
		End:      1,
	}

	html := snippet.HTML()
	r.Equal(template.HTML("<div>Hello &amp; World</div>"), html)
}

func Test_Snippet_Fields(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippet := Snippet{
		Content:  "package main",
		File:     "/path/to/main.go",
		Language: "golang",
		Name:     "package_declaration",
		Start:    1,
		End:      1,
	}

	r.Equal("package main", snippet.Content)
	r.Equal("/path/to/main.go", snippet.File)
	r.Equal("golang", snippet.Language)
	r.Equal("package_declaration", snippet.Name)
	r.Equal(1, snippet.Start)
	r.Equal(1, snippet.End)
}

func Test_Snippets_Map(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippets := Snippets{
		"hello": Snippet{
			Content:  "fmt.Println(\"Hello\")",
			Language: "go",
			Name:     "hello",
		},
		"world": Snippet{
			Content:  "fmt.Println(\"World\")",
			Language: "go",
			Name:     "world",
		},
	}

	r.Len(snippets, 2)

	hello, exists := snippets["hello"]
	r.True(exists)
	r.Equal("fmt.Println(\"Hello\")", hello.Content)
	r.Equal("hello", hello.Name)

	world, exists := snippets["world"]
	r.True(exists)
	r.Equal("fmt.Println(\"World\")", world.Content)
	r.Equal("world", world.Name)
}

func Test_Snippets_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	snippets := Snippets{}
	r.Len(snippets, 0)

	_, exists := snippets["nonexistent"]
	r.False(exists)
}
