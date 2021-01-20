package htm

import (
	"html/template"
	"strings"

	"github.com/gobuffalo/tags/v3"
	"github.com/gopherguides/remark"
)

func Print(elems ...remark.Tag) (string, error) {
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

func (hp *Printer) Print(elems ...remark.Tag) (string, error) {
	if hp.custom == nil {
		hp.custom = map[string]TagPrinter{}
	}
	if hp.decorators == nil {
		hp.decorators = map[string]TagDecorator{}
	}

	var err error
	var lines []string
	for _, tag := range elems {
		if tag.TagName() == "metadata" {
			continue
		}
		if h, ok := tag.(htmlable); ok {
			lines = append(lines, string(h.HTML()))
			continue
		}

		if doc, ok := tag.(*Document); ok {
			s, err := hp.Print(doc.GetChildren()...)
			if err != nil {
				return "", err
			}
			lines = append(lines, s)
			continue
		}

		if s, ok := tag.(remark.String); ok {
			lines = append(lines, s.String())
			continue
		}

		if td, ok := hp.decorators[tag.TagName()]; ok {
			tag, err = td(tag)
			if err != nil {
				return "", err
			}
		}

		if fn, ok := hp.custom[tag.TagName()]; ok {
			s, err := fn(tag)
			if err != nil {
				return "", err
			}
			lines = append(lines, s)
			continue
		}

		t := tags.New(tag.TagName(), tag.Options())

		s, err := hp.Print(tag.GetChildren()...)
		if err != nil {
			return "", err
		}
		t.Append(s)

		lines = append(lines, t.String())
	}

	return strings.Join(lines, " "), nil
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
