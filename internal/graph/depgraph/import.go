package depgraph

type Import struct {
	id string
}

func (i *Import) Id() string {
	return i.id
}
