package model

type Node struct {
	Name   string `json:"name"`
	IsHome bool   `json:"isHome"`
	Color  string `json:"color,omitempty"`
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
