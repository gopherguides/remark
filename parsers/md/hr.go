package md

import "github.com/gopherguides/remark"

func NewHR() (*remark.HR, error) {
	return &remark.HR{
		Attributes: remark.Attributes{},
	}, nil
}
