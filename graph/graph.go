package graph

type IGraph[T comparable] interface {
	GetEdges() []Edge[T]
	FindAdjecentVertices(vertex T) []T
	IsConnected(vertexA, vertexB T) bool
}

type graph[T comparable] struct {
	edges []Edge[T]
}

func New[T comparable](edges []Edge[T]) IGraph[T] {
	return &graph[T]{edges: edges}
}

func (g *graph[T]) GetEdges() []Edge[T] {
	return g.edges
}
