package htm

import "github.com/gopherguides/remark"

type Document struct {
	*remark.Generic
	Data remark.Metadata
}

func (d *Document) Metadata() remark.Metadata {
	if d.Data == nil {
		d.Data = remark.Metadata{}
	}
	return d.Data
}
