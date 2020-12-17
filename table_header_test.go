package remark

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TableHeader_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	th := TableHeader{
		Elements: Elements{
			TableCell{
				Elements: Elements{String("hooligan")},
			},
			TableCell{
				Elements: Elements{String("hippo")},
			},
		},
	}

	exp := template.HTML(`<th><td>hooligan</td><td>hippo</td></th>`)
	r.Equal(exp, th.HTML())
}
