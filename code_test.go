package remark

import (
	"testing"
)

func Test_Code_String(t *testing.T) {
	t.Parallel()
	compare(t, "code")
}
