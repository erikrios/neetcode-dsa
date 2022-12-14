package insertion

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "it should return []int{}, when the input is []int{}",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "it should return []int{1, 2, 3, 4, 5}, when the input is []int{5, 4, 2, 3, 1}",
			input:    []int{5, 4, 2, 3, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "it should return []int{1, 2, 3, 4, 5}, when the input is []int{1, 2, 3, 4, 5}",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "it should return []int{1, 2, 2, 3, 4, 5}, when the input is []int{5, 2, 3, 4, 3, 1}",
			input:    []int{5, 2, 3, 4, 3, 1},
			expected: []int{1, 2, 3, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			InsertionSort(testCase.input)
			assertValue(t, testCase.input, testCase.expected)
		})
	}
}

func assertValue(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but got %#v", want, got)
	}
}
