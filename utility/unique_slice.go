package utility

func UniqueSliceElements[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]struct{}, len(inputSlice))
	for _, element := range inputSlice {
		if _, ok := seen[element]; !ok {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = struct{}{}
		}
	}
	return uniqueSlice
}

func UniqueStringSliceElements(inputSlice []string) []string {
	uniqueSlice := make([]string, 0, len(inputSlice))
	seen := make(map[string]struct{}, len(inputSlice))
	for _, element := range inputSlice {
		if _, ok := seen[element]; !ok {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = struct{}{}
		}
	}
	return uniqueSlice
}
