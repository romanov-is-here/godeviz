package dto

type Node struct {
	Name        string    `json:"name"`
	IsHome      bool      `json:"isHome"`
	Color       string    `json:"color,omitempty"`
	FanInCount  int       `json:"fanInCount"`
	FanOutCount int       `json:"fanOutCount"`
	Imports     []*Import `json:"imports"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Color  string `json:"color,omitempty"`
}

type Graph struct {
	Nodes map[string]*Node `json:"nodes"`
	Edges map[string]*Edge `json:"edges"`
}

type Import struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
