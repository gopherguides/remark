package rpm

import "golang.org/x/net/html"

type Attrs map[string]string

func (attrs Attrs) ToAttrs() []html.Attribute {
	var ats []html.Attribute

	for k, v := range attrs {
		ats = append(ats, html.Attribute{
			Key: k,
			Val: v,
		})
	}
	return ats
}
