package list

import "sort"

// Contains checks if the needle element exists in a given haystack
func Contains[T comparable](haystack []T, needle T) bool {
	for _, candidate := range haystack {
		if candidate == needle {
			return true
		}
	}
	return false
}

// Map transforms the list by running the f func on each element in list
func Map[T, U any](list []T, f func(T) U) []U {
	res := make([]U, len(list))
	for i, e := range list {
		res[i] = f(e)
	}
	return res
}

// TryMap tries to run f on each element in list, and returns an error if an error happens
func TryMap[T, U any](list []T, f func(T) (U, error)) ([]U, error) {
	res := make([]U, len(list))
	var err error
	for i, e := range list {
		res[i], err = f(e)
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

// Reduce maps and reduces f on every list item with the initial U
func Reduce[T, U any](list []T, initial U, f func(T, U) U) U {
	for _, item := range list {
		initial = f(item, initial)
	}
	return initial
}

// ReduceMap maps and reduces f on every list item with the initial U
func ReduceMap[K comparable, V, U any](m map[K]V, initial U, f func(K, V, U) U) U {
	for key, value := range m {
		initial = f(key, value, initial)
	}
	return initial
}

// Sort uses sort.Slice to sort a given list (in-place)
func Sort[T any](list []T, sortFunc func(l, r T) bool) []T {
	sort.Slice(list, func(i, j int) bool {
		return sortFunc(list[i], list[j])
	})
	return list
}
