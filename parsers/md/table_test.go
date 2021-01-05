package md

import (
	"html/template"
	"testing"

	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_Table_HTML(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	table := Table{
		Elements: Elements{
			TableHeader{
				Elements: Elements{
					TableCell{
						Elements: Elements{String("hooligan")},
					},
					TableCell{
						Elements: Elements{String("hippo")},
					},
				},
			},
			TableRow{
				Elements: Elements{
					TableCell{
						Elements: Elements{String("snow")},
					},
					TableCell{
						Elements: Elements{
							&remark.Image{
								Attributes: remark.Attributes{
									"src": "./foo.png",
								},
							},
						},
					},
				},
			},
		},
	}

	exp := template.HTML(`<table><th><td>hooligan</td><td>hippo</td></th><tr><td>snow</td><td><img src="./foo.png"></img></td></tr></table>`)
	r.Equal(exp, table.HTML())
}
