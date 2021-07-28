package htm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NodeString(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	n := AttrNode(t, "div", map[string]string{
		"id": "1",
	})
	n.Namespace = "foo"

	const exp = `{"attrs":{"id":"1"},"data":"div","namespace":"foo"}`
	act := NodeString(n)
	r.Equal(exp, act)
}
