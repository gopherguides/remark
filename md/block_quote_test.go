package md

import (
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_BlockQuote_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "block_quote")
}

func Test_BlockQuote_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	bq := &remark.BlockQuote{
		Elements: Elements{
			String("ala mode"),
		},
	}

	exp := template.HTML(`<blockquote>ala mode</blockquote>`)
	r.Equal(exp, bq.HTML())
}
