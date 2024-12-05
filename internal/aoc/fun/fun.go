package fun

func Map[S ~[]T, T, E any](slice S, function func(T) E) []E {
	mapped := make([]E, len(slice))
	for i := range slice {
		mapped[i] = function(slice[i])
	}

	return mapped
}

func MapInPlace[S ~[]T, T any](slice S, function func(T) T) {
	for i := range slice {
		slice[i] = function(slice[i])
	}
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	filtered := make([]T, 0, len(slice))
	for i := range slice {
		if predicate(slice[i]) {
			filtered = append(filtered, slice[i])
		}
	}

	return filtered
}

func Fold[T any](slice []T, pairFunction func(a, b T) T, zero T) T {
	if len(slice) == 0 {
		return zero
	}

	return pairFunction(slice[0], Fold(slice[1:], pairFunction, zero))
}
