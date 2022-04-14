package list_test

import (
	"github.com/stretchr/testify/assert"
	"linked-list/list"
	"testing"
)

func TestAtWhenExists(t *testing.T) {
	tests := []struct {
		name  string
		array []int
		index int
	}{
		{name: "with several nodes and first index", array: []int{3, 5, 7, 2}, index: 0},
		{name: "with several nodes and last index", array: []int{3, 5, 7, 2}, index: 3},
		{name: "with several nodes and other index", array: []int{3, 5, 7, 2}, index: 2},
		{name: "with one node", array: []int{4}, index: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := list.FromArray(test.array)
			res, err := l.At(test.index)
			assert.Nil(t, err)
			assert.Equal(t, test.array[test.index], res)
		})
	}
}

func TestAtWhenDoesntExist(t *testing.T) {
	arr := []int{1, 2, 3}
	index := 3

	l := list.FromArray[int](arr)
	_, err := l.At(index)
	assert.Error(t, err, "element not found")
}

func TestAsArray(t *testing.T) {
	arr := []int{1, 3, 5, 2}
	l := list.FromArray(arr)
	assert.Equal(t, arr, l.AsArray())
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name           string
		array          []int
		index          int
		value          int
		expectedResult []int
	}{
		{name: "in the middle", array: []int{2, 3, 2, 6}, index: 3, value: 4, expectedResult: []int{2, 3, 2, 4, 6}},
		{name: "at the end", array: []int{1, 3, 4}, index: 3, value: 6, expectedResult: []int{1, 3, 4, 6}},
		{name: "at the beginning", array: []int{2, 7, 8}, index: 0, value: 1, expectedResult: []int{1, 2, 7, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := list.FromArray(test.array)
			err := l.Insert(test.index, test.value)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedResult, l.AsArray())
		})
	}
}

func TestInsertWhenOutOfRange(t *testing.T) {
	arr := []int{1, 2, 1}
	index := 4
	val := 3

	l := list.FromArray[int](arr)
	err := l.Insert(index, val)
	assert.Error(t, err, "index is out of range")
	assert.Equal(t, arr, l.AsArray())
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name           string
		array          []int
		index          int
		expectedResult []int
	}{
		{name: "in the middle", array: []int{2, 3, 2, 6}, index: 2, expectedResult: []int{2, 3, 6}},
		{name: "at the end", array: []int{1, 3, 4}, index: 2, expectedResult: []int{1, 3}},
		{name: "at the beginning", array: []int{2, 7, 8}, index: 0, expectedResult: []int{7, 8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := list.FromArray(test.array)
			err := l.Delete(test.index)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedResult, l.AsArray())
		})
	}
}

func TestDeleteWhenOutOfRange(t *testing.T) {
	arr := []int{1, 2, 1}
	index := 4

	l := list.FromArray[int](arr)
	err := l.Delete(index)
	assert.Error(t, err, "index is out of range")
	assert.Equal(t, arr, l.AsArray())

}
