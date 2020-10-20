package remark

import "testing"

func Test_Heading_String(t *testing.T) {
	t.Parallel()
	compare(t, "heading")
}
