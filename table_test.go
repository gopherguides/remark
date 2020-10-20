package remark

import "testing"

func Test_Table_String(t *testing.T) {
	t.Parallel()

	compare(t, "table")
}
