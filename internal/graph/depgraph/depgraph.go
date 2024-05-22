package depgraph

//func (g *DepGraph) Hits() Hits {
//	hitmap := make(map[string]*Hit)
//
//	for _, pi := range g.Packs {
//		h := &Hit{pack: pi}
//		hitmap[pi.Id()] = h
//	}
//
//	for _, pi := range g.Packs {
//		for _, imp := range pi.Imports() {
//			h, ok := hitmap[imp.Id()]
//			if !ok {
//				fmt.Println("HIT NOT FOUND [", imp.Id(), "] FROM ", pi.Name())
//				continue
//			}
//			h.add(imp.Id())
//		}
//	}
//
//	hsorted := make([]*Hit, 0, len(hitmap))
//
//	for _, value := range hitmap {
//		hsorted = append(hsorted, value)
//	}
//
//	sort.Slice(hsorted, func(i, j int) bool {
//		return hsorted[i].Count() > hsorted[j].Count()
//	})
//
//	return Hits{
//		Hits: hsorted,
//	}
//}
