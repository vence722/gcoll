// gcoll
// @description gcoll is a collection library of the most frequently used data structures in Go programing language
// @authors     Vence Lin(vence722@gmail.com)

package stereotype

// The Comparable interface
type Comparable interface {
	CompareTo(x any) int
}

type IntComparable int

func (this IntComparable) CompareTo(that any) int {
	return simpleCompare(this, that.(IntComparable))
}

type Int64Comparable int64

func (this Int64Comparable) CompareTo(that any) int {
	return simpleCompare(this, that.(Int64Comparable))
}

type FloatComparable float32

func (this FloatComparable) CompareTo(that any) int {
	return simpleCompare(this, that.(FloatComparable))
}

type Float64Comparable float64

func (this Float64Comparable) CompareTo(that any) int {
	return simpleCompare(this, that.(Float64Comparable))
}

func simpleCompare(v1, v2 any) int {
	switch v1.(type) {
	case IntComparable:
		if v1.(IntComparable) < v2.(IntComparable) {
			return -1
		} else if v1.(IntComparable) > v2.(IntComparable) {
			return 1
		}
		return 0
	case Int64Comparable:
		if v1.(Int64Comparable) < v2.(Int64Comparable) {
			return -1
		} else if v1.(Int64Comparable) > v2.(Int64Comparable) {
			return 1
		}
		return 0
	case FloatComparable:
		if v1.(FloatComparable) < v2.(FloatComparable) {
			return -1
		} else if v1.(FloatComparable) > v2.(FloatComparable) {
			return 1
		}
		return 0
	case Float64Comparable:
		if v1.(Float64Comparable) < v2.(Float64Comparable) {
			return -1
		} else if v1.(Float64Comparable) > v2.(Float64Comparable) {
			return 1
		}
		return 0
	}
	return 0
}
