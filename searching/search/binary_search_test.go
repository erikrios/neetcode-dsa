package search

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "it should return -1, when the target is 1 & nums is []",
			target:   1,
			nums:     []int{},
			expected: -1,
		},
		{
			name:     "it should return -1, when the target is 2 & nums is [1, 3, 3, 4, 5]",
			target:   2,
			nums:     []int{1, 3, 3, 4, 5},
			expected: -1,
		},
		{
			name:     "it should return -1, when the target is 1 & nums is [2, 3, 4, 5]",
			target:   1,
			nums:     []int{2, 3, 4, 5},
			expected: -1,
		},
		{
			name:     "it should return -1, when the target is 9 & nums is [2, 3, 4, 5]",
			target:   9,
			nums:     []int{2, 3, 4, 5},
			expected: -1,
		},
		{
			name:     "it should return 2, when the target is 3 & nums is [1, 2, 3, 4, 5]",
			target:   3,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 2,
		},
		{
			name:     "it should return 3, when the target is 4 & nums is [1, 2, 3, 4, 5]",
			target:   4,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 3,
		},
		{
			name:     "it should return 0, when the target is 1 & nums is [1, 2, 3, 4, 5]",
			target:   1,
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := BinarySearch(testCase.target, testCase.nums)
			assertValue(t, got, testCase.expected)
		})
	}
}

func assertValue(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but got %#v", want, got)
	}
}
