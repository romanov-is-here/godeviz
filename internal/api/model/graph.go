package model

type Node struct {
	Name   string `json:"name"`
	IsHome bool   `json:"isHome"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type Graph struct {
	Nodes map[string]*Node `json:"nodes"`
	Edges map[string]*Edge `json:"edges"`
}
