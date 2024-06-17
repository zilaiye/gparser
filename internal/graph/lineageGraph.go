package graph

type LineageGraph struct {
	Vertex map[string]LineageNode
	Edge   map[string][]LineageNode
}

func NewLineageGraph() *LineageGraph {
	return new(LineageGraph)
}
