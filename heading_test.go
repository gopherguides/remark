package remark

import "testing"

func Test_Heading_String(t *testing.T) {
	t.Parallel()
	compareStringer(t, "heading")
}
