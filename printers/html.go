package printers

import (
	"fmt"
	"html/template"
	"io"

	"github.com/gopherguides/remark"
)

type TagPrinter func(t remark.Tag) (template.HTML, error)

type HTML struct {
	Printer TagPrinter
	custom  map[string]TagPrinter
}

func (h *HTML) Set(name string, fn TagPrinter) {
	if h.custom == nil {
		h.custom = map[string]TagPrinter{}
	}
	h.custom[name] = fn
}

func (h *HTML) Print(w io.Writer, elems ...remark.Element) error {
	fn := h.Printer
	if fn == nil {
		fn = func(t remark.Tag) (template.HTML, error) {
			return t.HTML(), nil
		}
	}
	if h.custom == nil {
		h.custom = map[string]TagPrinter{}
	}

	for _, e := range elems {
		t, ok := e.(remark.Tag)
		if !ok {
			continue
		}
		fmt.Printf(">>>TODO printers/html.go:41: t.Name() (%T): %v \n", t.Name(), t.Name())
		if f, ok := h.custom[t.Name()]; ok {
			h, err := f(t)
			if err != nil {
				return err
			}
			w.Write([]byte(h))
			return nil
		}

		h, err := fn(t)
		if err != nil {
			return err
		}
		w.Write([]byte(h))
	}
	return nil
}
