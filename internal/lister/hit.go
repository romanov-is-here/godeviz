package lister

type Hit struct {
	pack *PackageInfo
	from []string
}

type Hits struct {
	Hits []*Hit
}

func (h *Hit) Count() int {
	return len(h.from)
}

func (h *Hit) Name() string {
	return h.pack.Name()
}

func (h *Hit) add(from string) {
	h.from = append(h.from, from)
}
