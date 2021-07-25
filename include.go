package remark

var _ Tag = &Include{}

type Include struct {
	*Generic
}

func (i Include) String() string {
	s := i.Children.String()
	return s
}
