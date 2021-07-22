package remark

import "github.com/gopherguides/remark/meta"

var _ Tag = &Include{}

type Include struct {
	*Generic
	Info *meta.Info
}

func (i Include) String() string {
	return i.Generic.Children.String()
}
