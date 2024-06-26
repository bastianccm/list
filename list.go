package list

import (
	"slices"
	"sort"
)

// Contains checks if the needle element exists in a given haystack
func Contains[T comparable](haystack []T, needle T) bool {
	return slices.Contains(haystack, needle)
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

func PassEntry[E comparable](e E) E {
	return e
}

func SortEntries[S ~[]E, E any, C comparable](list S, comparator func(l, r C) int, getter func(E) C) S {
	slices.SortFunc(list, func(l, r E) int { return comparator(getter(l), getter(r)) })
	return list
}

// Filter returns all items which pass the filterFunc
func Filter[T any](list []T, filterFunc func(item T) bool) []T {
	res := make([]T, 0, len(list))
	for _, item := range list {
		if filterFunc(item) {
			res = append(res, item)
		}
	}
	return res
}

// Keys returns all map-keys for the provided map m
func Keys[T comparable, U any](m map[T]U) []T {
	var keys = make([]T, len(m))
	var i = 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// Find finds an element in the haystick by given predicate
func Find[T any](haystack []T, predicate func(T) bool) (T, bool) {
	for _, candidate := range haystack {
		if predicate(candidate) {
			return candidate, true
		}
	}
	return *new(T), false
}
