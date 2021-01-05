package md

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Heading_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "heading")
}

func Test_Heading_HTML(t *testing.T) {
	t.Parallel()

	table := []struct {
		Level int
		Exp   template.HTML
	}{
		{Level: 0, Exp: `<h1>Hello</h1>`},
		{Level: 1, Exp: `<h1>Hello</h1>`},
		{Level: 2, Exp: `<h2>Hello</h2>`},
	}

	for _, tt := range table {
		tt := tt
		t.Run(fmt.Sprintf("%d/%s", tt.Level, tt.Exp), func(st *testing.T) {
			r := require.New(st)

			h := &remark.Heading{
				Level: tt.Level,
				Elements: Elements{
					String("Hello"),
				},
			}

			got := h.HTML()
			r.Equal(tt.Exp, got)
		})
	}
}
