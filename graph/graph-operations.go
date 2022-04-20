package graph

func (g *graph[T]) FindAdjecentVertices(vertex T) []T {
	adjecentVertices := make([]T, 0)
	for _, edge := range g.GetEdges() {
		if edge.VertexA == vertex {
			adjecentVertices = append(adjecentVertices, edge.VertexB)
		} else if edge.VertexB == vertex {
			adjecentVertices = append(adjecentVertices, edge.VertexA)
		}
	}
	return adjecentVertices
}

func (g *graph[T]) IsConnected(vertexA, vertexB T) bool {
	for _, edge := range g.GetEdges() {
		if (edge.VertexA == vertexA && edge.VertexB == vertexB) || (edge.VertexA == vertexB && edge.VertexB == vertexA) {
			return true
		}
	}
	return false
}
