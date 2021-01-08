package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UnorderedList_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var ul UnorderedList
	r.Equal("ul", ul.Name())
}

// func Test_UnorderedList_HTML(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)
//
// 	ul := UnorderedList{
// 		Attributes: Attributes{},
// 		Elems: Elements{
// 			String("hello"),
// 		},
// 	}
// 	ul.Set("k1", "one")
// 	ul.Set("k2", "two")
// 	ul.Set("k3", "three")
//
// 	exp := template.HTML(`<ul k1="one" k2="two" k3="three">hello</ul>`)
// 	act := ul.HTML()
// 	r.Equal(exp, act)
// }
