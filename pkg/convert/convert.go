package convert

// MapToValueSlice returns the slice of map values.
//
// The order in slice is nondeterministic.
func MapToValueSlice[T comparable, S any](m map[T]S) []S {
	s := make([]S, 0, len(m))

	for k := range m {
		s = append(s, m[k])
	}

	return s
}

// MapToKeySlice returns the slice of map keys.
//
// The order in slice is nondeterministic.
func MapToKeySlice[T comparable, S any](m map[T]S) []T {
	s := make([]T, 0, len(m))

	for k := range m {
		s = append(s, k)
	}

	return s
}

// SliceToSet creates a set from slice.
func SliceToSet[T comparable](s []T) map[T]bool {
	m := make(map[T]bool, len(s))
	for k := range s {
		m[s[k]] = true
	}
	return m
}
