package fibonacci

import (
	"reflect"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "it should return 1, when input is 1",
			input:    1,
			expected: 1,
		},
		{
			name:     "it should return 1, when input is 2",
			input:    1,
			expected: 1,
		},
		{
			name:     "it should return 2, when input is 3",
			input:    3,
			expected: 2,
		},
		{
			name:     "it should return 3, when input is 4",
			input:    4,
			expected: 3,
		},
		{
			name:     "it should return 5, when input is 5",
			input:    5,
			expected: 5,
		},
		{
			name:     "it should return 8, when input is 6",
			input:    6,
			expected: 8,
		},
		{
			name:     "it should return 13, when input is 7",
			input:    7,
			expected: 13,
		},
		{
			name:     "it should return 21, when input is 8",
			input:    8,
			expected: 21,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := Fibonacci(testCase.input)
			assertValue(t, got, testCase.expected)
		})
	}
}

func assertValue(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wamt %#v, but got %#v", want, got)
	}
}
