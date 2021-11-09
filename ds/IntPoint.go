package ds

type IntPoint struct {
	X, Y int
}

func (it IntPoint) IsEqual(other IntPoint) bool {
	return it.X == other.X && it.Y == other.Y
}
