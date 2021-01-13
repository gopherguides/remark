package remark

import (
	"fmt"
	"strings"
)

type Element interface {
	Name() string
}

type Elemental interface {
	Elements() Elements
}

type Elements []Element

func (elems Elements) Name() string {
	return ""
}

func (elems Elements) String() string {
	lines := make([]string, 0, len(elems))
	for _, e := range elems {
		if s, ok := e.(fmt.Stringer); ok {
			lines = append(lines, s.String())
		}
	}
	return strings.Join(lines, "")
}
