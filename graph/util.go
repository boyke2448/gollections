package graph

func FindAdjecentVertices[T comparable](vertex T, edges [][]T) []T {
	adjecentVertices := make([]T, 0)
	for _, edge := range edges {
		if edge[0] == vertex {
			adjecentVertices = append(adjecentVertices, edge[1])
		} else if edge[1] == vertex {
			adjecentVertices = append(adjecentVertices, edge[0])
		}
	}
	return adjecentVertices
}

func IsConnected[T comparable](vertexA, vertexB T, edges [][]T) bool {
	for _, edge := range edges {
		if (edge[0] == vertexA && edge[1] == vertexB) || (edge[0] == vertexB && edge[1] == vertexA) {
			return true
		}
	}
	return false
}
