package list_test

import (
	"strconv"
	"testing"

	"github.com/bastianccm/list"

	"github.com/stretchr/testify/assert"
)

type kv struct {
	k, v string
}

func TestList(t *testing.T) {
	assert.True(t, list.Contains([]string{"a", "b", "c"}, "b"))
	assert.False(t, list.Contains([]string{"a", "b", "c"}, "x"))

	assert.Equal(t, []string{"1", "2", "3"}, list.Map([]int{1, 2, 3}, strconv.Itoa))

	res, err := list.TryMap([]string{"1", "2", "3"}, strconv.Atoi)
	assert.Equal(t, []int{1, 2, 3}, res)
	assert.NoError(t, err)

	res, err = list.TryMap([]string{"1", "2", "abc"}, strconv.Atoi)
	assert.Equal(t, []int{1, 2, 0}, res)
	assert.Error(t, err)

	assert.Equal(t, "concatted: 123", list.Reduce([]int{1, 2, 3}, "concatted: ", func(i int, s string) string {
		return s + strconv.Itoa(i)
	}))
	assert.Equal(t, map[string][]string{"a": {"b", "c"}}, list.Reduce([]kv{{"a", "b"}, {"a", "c"}}, make(map[string][]string), func(e kv, s map[string][]string) map[string][]string {
		s[e.k] = append(s[e.k], e.v)
		return s
	}))

	assert.Equal(t, []kv{{"a", "b"}, {"a", "c"}}, list.ReduceMap(map[string][]string{"a": {"b", "c"}}, nil, func(key string, value []string, initial []kv) []kv {
		return append(initial, list.Map(value, func(v string) kv { return kv{key, v} })...)
	}))

	assert.Equal(t, []int{1, 2, 3, 4, 5}, list.Sort([]int{3, 5, 1, 2, 4}, func(l, r int) bool {
		return l < r
	}))

	assert.Equal(t, []int{3, 1, 2}, list.Filter([]int{3, 5, 1, 2, 4}, func(item int) bool { return item <= 3 }))
}
