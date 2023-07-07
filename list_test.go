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

	assert.Equal(t, []string{"a", "b", "c"}, list.Keys(map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
	}))
	assert.Equal(t, []int{10, 20, 30}, list.Keys(map[int]struct{}{
		10: {},
		20: {},
		30: {},
	}))

	v, found := list.Find([]int{3, 1, 2}, func(in int) bool { return in == 1 })
	assert.Equal(t, 1, v)
	assert.True(t, found)

	v, found = list.Find([]int{3, 1, 2}, func(in int) bool { return in == 7 })
	assert.Equal(t, 0, v)
	assert.False(t, found)

	type testStruct struct {
		num int
	}
	vts, found := list.Find([]testStruct{{num: 3}, {num: 1}, {num: 2}}, func(in testStruct) bool { return in.num == 1 })
	assert.Equal(t, 1, vts.num)
	assert.True(t, found)

	vts, found = list.Find([]testStruct{{num: 3}, {num: 1}, {num: 2}}, func(in testStruct) bool { return in.num == 7 })
	assert.Equal(t, 0, vts.num)
	assert.False(t, found)
}
