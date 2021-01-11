package remark

type HR struct {
	Attributes
}

func (HR) Name() string {
	return "hr"
}
