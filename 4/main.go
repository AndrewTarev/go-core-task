package main

// Difference возвращает элементы, которые есть в первом слайсе, но отсутствуют во втором.
func Difference(slice1, slice2 []string) []string {
	slice2Map := make(map[string]struct{})
	for _, item := range slice2 {
		slice2Map[item] = struct{}{}
	}

	var result []string
	for _, item := range slice1 {
		if _, found := slice2Map[item]; !found {
			result = append(result, item)
		}
	}

	if len(result) == 0 {
		return []string{}
	}

	return result
}
