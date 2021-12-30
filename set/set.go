// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package set

import "github.com/vence722/gcoll/collection"

type Set[T comparable] interface {
	collection.Collection[T]
}
