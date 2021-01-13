package remark

type String string

func (s String) Name() string {
	return "string"
}

func (s String) String() string {
	return string(s)
}
