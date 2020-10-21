package remark

import "testing"

func Test_Link_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "link")
}
