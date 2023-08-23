package grid

import "reflect"

type Grid [4][4]int

func (g Grid) Equals(other Grid) bool {
	return reflect.DeepEqual(g, other)
}
