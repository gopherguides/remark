package printers

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gopherguides/remark"
)

type TagPrinter func(t remark.Tag) (string, error)

type htmlable interface {
	HTML() template.HTML
}

type HTML struct {
	custom map[string]TagPrinter
}

func (hp *HTML) Print(elems ...remark.Element) (string, error) {
	if hp.custom == nil {
		hp.custom = map[string]TagPrinter{}
	}
	var lines []string
	for _, e := range elems {
		if fn, ok := hp.custom[e.Name()]; ok {
			if t, ok := e.(remark.Tag); ok {
				s, err := fn(t)
				if err != nil {
					return "", err
				}
				lines = append(lines, s)
			}
			continue
		}

		var lns []string
		switch t := e.(type) {
		case htmlable:
			lines = append(lines, string(t.HTML()))
			continue
		case remark.String:
			ss := string(t)
			if len(ss) > 0 {
				lines = append(lines, ss)
			}
			continue
		case remark.Elemental:
			s, err := hp.Print(t.Elements()...)
			if err != nil {
				return "", err
			}
			lns = append(lns, s)
		}

		opts := tags.Options{}
		if o, ok := e.(remark.Tag); ok {
			opts = o.Options()
		}

		t := tags.New(e.Name(), opts)
		t.Append(strings.Join(lns, ""))
		lines = append(lines, t.String())
	}
	return strings.Join(lines, ""), nil
}

func (h *HTML) Set(name string, fn TagPrinter) {
	if h.custom == nil {
		h.custom = map[string]TagPrinter{}
	}
	h.custom[name] = fn
}
