package remark

import (
	"fmt"
	"strings"
)

type Element interface {
	fmt.Stringer
}

type Elements []Element

func (elems Elements) String() string {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		lines = append(lines, e.String())
	}
	return strings.Join(lines, "\n\n")
}
