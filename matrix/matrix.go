package matrix

type Matrix interface {
	Get(x int, y int) interface{}
	Set(x int, y int, value interface{}) error
	Size() (int, int)
	Resize(x int, y int, initVal interface{}) error
	RemoveAt(index int) error
}
