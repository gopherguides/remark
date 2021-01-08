package printers

import (
	"bytes"
	"html/template"
	"testing"

	"github.com/gobuffalo/tags"
	"github.com/gopherguides/remark"
	"github.com/stretchr/testify/require"
)

func Test_HTML_ZeroValue(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	tag := &remark.Generic{
		TagName: "strong",
		Elems: remark.Elements{
			remark.String("hi"),
		},
	}

	bb := &bytes.Buffer{}

	p := &HTML{}
	err := p.Print(bb, tag)
	r.NoError(err)
	r.Equal(string(tag.HTML()), bb.String())
}

func Test_HTML_Custom(t *testing.T) {
	t.Parallel()

	table := []struct {
		Tag string
		Exp string
	}{
		{Tag: "strong", Exp: `<stronger><strong>hi</strong></stronger>`},
		{Tag: "em", Exp: `<em>hi</em>`},
	}

	fn := func(tag remark.Tag) (template.HTML, error) {
		if tag.Name() != "strong" {
			return tag.HTML(), nil
		}

		tg := tags.New("stronger", nil)
		tg.Append(tag.HTML())
		return tg.HTML(), nil
	}

	p := &HTML{
		Printer: fn,
	}

	for _, tt := range table {
		tt := tt
		t.Run(tt.Tag, func(st *testing.T) {
			r := require.New(st)

			tag := &remark.Generic{
				TagName: "strong",
				Elems: remark.Elements{
					remark.String("hi"),
				},
			}

			bb := &bytes.Buffer{}

			err := p.Print(bb, tag)
			r.NoError(err)

			exp := `<stronger><strong>hi</strong></stronger>`
			r.Equal(exp, bb.String())
		})
	}
}

func Test_HTML_Custom_Map(t *testing.T) {
	t.Parallel()

	table := []struct {
		Tag string
		Exp string
	}{
		{Tag: "strong", Exp: `<stronger><strong>hi</strong></stronger>`},
		{Tag: "em", Exp: `<em>hi</em>`},
		{Tag: "strike", Exp: `<strike>BACK!</strike>`},
	}

	fn := func(tag remark.Tag) (template.HTML, error) {
		if tag.Name() != "strong" {
			return tag.HTML(), nil
		}

		tg := tags.New("stronger", nil)
		tg.Append(tag.HTML())
		return tg.HTML(), nil
	}

	p := &HTML{
		Printer: fn,
	}

	p.Set("strike", func(tag remark.Tag) (template.HTML, error) {
		return template.HTML(`<strike>BACK!</strike>`), nil
	})

	for _, tt := range table {
		tt := tt
		t.Run(tt.Tag, func(st *testing.T) {
			r := require.New(st)

			tag := &remark.Generic{
				TagName: tt.Tag,
				Elems: remark.Elements{
					remark.String("hi"),
				},
			}

			bb := &bytes.Buffer{}

			err := p.Print(bb, tag)
			r.NoError(err)

			r.Equal(tt.Exp, bb.String())
		})
	}
}
