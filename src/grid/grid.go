package grid

import "reflect"

type Grid [][]int

func (g Grid) Equals(other Grid) bool {
	return reflect.DeepEqual(g, other)
}
