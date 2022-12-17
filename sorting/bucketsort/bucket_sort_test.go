package bucketsort

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
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
			name:     "it should return []int{0, 1, 2, 3, 4}, when the input is []int{4, 2, 3, 1, 0}",
			input:    []int{4, 2, 3, 1, 0},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "it should return []int{0, 1, 2, 3, 4}, when the input is []int{0, 1, 2, 3, 4}",
			input:    []int{0, 1, 2, 3, 4},
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "it should return []int{0, 1, 2, 2, 3, 3, 4}, when the input is []int{2, 2, 3, 4, 3, 1, 0}",
			input:    []int{2, 2, 3, 4, 3, 1, 0},
			expected: []int{0, 1, 2, 2, 3, 3, 4},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			BucketSort(testCase.input)
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
