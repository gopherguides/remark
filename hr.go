package remark

import "github.com/gobuffalo/tags/v3"

type HR struct {
	Attributes
}

func (HR) Name() string {
	return "hr"
}

func (g HR) String() string {
	t := tags.New(g.Name(), g.Options())
	return t.String()
}
