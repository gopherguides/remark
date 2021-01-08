package remark

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OrderedList_Name(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	var ol OrderedList
	r.Equal("ol", ol.Name())
}

// func Test_OrderedList_HTML(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)
//
// 	ol := OrderedList{
// 		Attributes: Attributes{},
// 		Elems: Elements{
// 			String("hello"),
// 		},
// 	}
// 	ol.Set("k1", "one")
// 	ol.Set("k2", "two")
// 	ol.Set("k3", "three")
//
// 	exp := template.HTML(`<ol k1="one" k2="two" k3="three">hello</ol>`)
// 	act := ol.HTML()
// 	r.Equal(exp, act)
// }
