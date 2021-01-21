package md

import "github.com/gopherguides/remark"

type Section struct {
	*remark.Generic
}

func (s *Section) Metadata() remark.Metadata {
	if s.Data == nil {
		s.Data = remark.Metadata{}
	}
	return s.Data
}
