package matrix

type Matrix interface {
	Get(x int, y int) interface{}
	Set(x int, y int, value interface{})
	Size() (int, int)
	Resize(x int, y int)
}
