package main

// Intersections Intersect проверяет, есть ли пересечения между двумя слайсами, и возвращает пересечения.
func Intersections(slice1, slice2 []int) (bool, []int) {
	set := make(map[int]struct{})
	for _, v := range slice1 {
		set[v] = struct{}{}
	}

	var cross []int
	found := false
	for _, v := range slice2 {
		if _, ok := set[v]; ok {
			cross = append(cross, v)
			found = true
		}
	}
	if len(cross) == 0 {
		return false, []int{}
	}
	return found, cross
}
