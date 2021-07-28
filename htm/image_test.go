package htm

import (
	"testing"

	"github.com/gopherguides/remark"
	"github.com/gopherguides/remark/internal/rpm"
	"github.com/stretchr/testify/require"
)

func Test_Parser_NewImage(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	cab := rpm.ModuleFS(t)
	rpm.AppendFile(t, cab, "assets/foo.png")
	p := testParser(t, cab)

	ats := map[string]string{
		"name": "FOO",
		"src":  "assets/foo.png",
	}

	node := AttrNode(t, "img", ats)

	img, err := p.NewImage(node)
	r.NoError(err)
	r.NotNil(img)
	r.Equal("img", img.Name)

	exp := remark.Attributes(ats)
	r.Equal(exp, img.Attributes)
}

func Test_Parser_NewImage_Errors(t *testing.T) {
	t.Parallel()

	table := []struct {
		name    string
		control bool
		atrs    remark.Attributes
	}{
		{name: "no src", atrs: nil},
		{name: "src doesn't exist", atrs: remark.Attributes{
			"src": "oops.png",
		}},
		{name: "control", atrs: remark.Attributes{
			"src": "http://foo.png",
		}, control: true},
	}

	cab := rpm.ModuleFS(t)
	rpm.AppendFile(t, cab, "assets/foo.png")
	p := testParser(t, cab)

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			r := require.New(t)

			node := AttrNode(t, "img", tt.atrs)

			_, err := p.NewImage(node)
			if tt.control {
				r.NoError(err)
				return
			}
			r.Error(err)
		})
	}

}
