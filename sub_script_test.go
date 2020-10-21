package remark

import "testing"

func Test_SubScript_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "sub_script")
}
