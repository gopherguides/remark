package htm

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gopherguides/remark"
)

func Print(elems ...remark.Element) (string, error) {
	p := &Printer{
		custom: map[string]TagPrinter{},
	}
	return p.Print(elems...)
}

type TagPrinter func(t remark.Tag) (string, error)

type htmlable interface {
	HTML() template.HTML
}

type TagDecorator func(t remark.Tag) (remark.Tag, error)

type Printer struct {
	custom     map[string]TagPrinter
	decorators map[string]TagDecorator
}

func (hp *Printer) Print(elems ...remark.Element) (string, error) {
	if hp.custom == nil {
		hp.custom = map[string]TagPrinter{}
	}
	if hp.decorators == nil {
		hp.decorators = map[string]TagDecorator{}
	}

	var err error
	var lines []string
	for _, e := range elems {
		if els, ok := e.(remark.Elements); ok {
			l, err := hp.Print(els...)
			if err != nil {
				return "", err
			}
			lines = append(lines, l)
			continue
		}
		if s, ok := e.(remark.String); ok {
			if len(s) > 0 {
				lines = append(lines, string(s))
			}
			continue
		}

		tag, ok := e.(remark.Tag)
		if !ok {
			continue
		}

		if td, ok := hp.decorators[tag.Name()]; ok {
			tag, err = td(tag)
			if err != nil {
				return "", err
			}
		}

		if fn, ok := hp.custom[tag.Name()]; ok {
			s, err := fn(tag)
			if err != nil {
				return "", err
			}
			lines = append(lines, s)
			continue
		}

		var lns []string
		switch t := e.(type) {
		case htmlable:
			lines = append(lines, string(t.HTML()))
			continue
		case remark.Elemental:
			s, err := hp.Print(t.Elements()...)
			if err != nil {
				return "", err
			}
			lns = append(lns, s)
		}

		opts := tag.Options()

		t := tags.New(e.Name(), opts)
		t.Append(strings.Join(lns, ""))
		lines = append(lines, t.String())
	}
	return strings.Join(lines, ""), nil
}

func (h *Printer) Set(name string, fn TagPrinter) {
	if h.custom == nil {
		h.custom = map[string]TagPrinter{}
	}
	h.custom[name] = fn
}

func (h *Printer) Decorate(name string, fn TagDecorator) {
	if h.decorators == nil {
		h.decorators = map[string]TagDecorator{}
	}
	h.decorators[name] = fn
}
