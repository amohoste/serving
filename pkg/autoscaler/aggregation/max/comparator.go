package max

type comparator func(int32, int32) bool

func less(i int32, j int32) bool {
	return i < j
}
func greater(i int32, j int32) bool {
	return i > j
}