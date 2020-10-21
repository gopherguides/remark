package remark

import "testing"

func Test_HR_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "hr")
}
