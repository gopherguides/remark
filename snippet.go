package remark

type Snippet struct {
	Content  string
	File     string
	Language string
	Name     string
	Start    int
	End      int
}

type Snippets map[string]Snippet
