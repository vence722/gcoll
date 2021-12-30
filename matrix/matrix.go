package matrix

type Matrix interface {
	Get(x int, y int) any
	Set(x int, y int, value any) error
	Size() (int, int)
	Resize(x int, y int, initVal any) error
	RemoveAt(index int) error
}
