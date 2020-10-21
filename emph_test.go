package remark

import "testing"

func Test_Emph_string(t *testing.T) {
	t.Parallel()
	compareStringer(t, "emph")
}
