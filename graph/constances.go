package graph

import (
	"errors"
)

const (
	ITERATE_METHOD_DFS string = "DFS"
	ITERATE_METHOD_BFS string = "BFS"
)

var (
	ERR_VERTEX_KEY_EXISTS      error = errors.New("Vertex key exists")
	ERR_VERTEX_KEY_NOT_EXISTS  error = errors.New("Vertex key not exists")
	ERR_EDGE_EXISTS            error = errors.New("Edge exists")
	ERR_EDGE_NOT_EXISTS        error = errors.New("Edge not exists")
	ERR_NEIGHBOR_IS_NULL       error = errors.New("Neighbor is null")
	ERR_NEIGHBOR_NOT_EXISTS    error = errors.New("Neighbor not exists")
	ERR_ROOT_VERTEX_NOT_EXISTS error = errors.New("Root vertex not exists")
	ERR_INDEX_OUT_OF_BOUND     error = errors.New("Index out of bound")
)
