package remark

import "testing"

func Test_SuperScript_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "super_script")
}
