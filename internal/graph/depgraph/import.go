package depgraph

type Import struct {
	text string
}

func (i *Import) Id() string {
	return i.text
}
