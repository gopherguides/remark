package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Code_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := Code{
		Generic: &Generic{},
	}

	r.Equal("code", code.Name())
}

func Test_Code_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := Code{
		Generic: &Generic{
			Attributes: Attributes{
				"class": "highlight",
				"id":    "test-code",
			},
			Children: Tags{
				String("fmt.Println(\"Hello, World!\")"),
			},
		},
	}

	html := code.HTML()
	htmlStr := string(html)

	r.Contains(htmlStr, `<code`)
	r.Contains(htmlStr, `class="highlight"`)
	r.Contains(htmlStr, `id="test-code"`)
	r.Contains(htmlStr, `language="plain"`)
	r.Contains(htmlStr, `fmt.Println("Hello, World!")`)
	r.Contains(htmlStr, `</code>`)
}

func Test_Code_String(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := Code{
		Generic: &Generic{
			Attributes: Attributes{
				"class": "highlight",
			},
			Children: Tags{
				String("fmt.Println(&quot;Hello&quot;)"),
			},
		},
	}

	result := code.String()
	// String() should unescape HTML entities
	r.Contains(result, `fmt.Println("Hello")`)
	r.Contains(result, `<code`)
	r.Contains(result, `class="highlight"`)
}

func Test_Code_Empty(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	code := Code{
		Generic: &Generic{
			Attributes: Attributes{},
			Children:   Tags{},
		},
	}

	html := code.HTML()
	htmlStr := string(html)

	r.Contains(htmlStr, `<code`)
	r.Contains(htmlStr, `language="plain"`)
	r.Contains(htmlStr, `</code>`)
}
