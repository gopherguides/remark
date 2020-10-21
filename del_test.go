package remark

import "testing"

func Test_Del_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "del")
}
