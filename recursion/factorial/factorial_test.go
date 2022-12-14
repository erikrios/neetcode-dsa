package factorial

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
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
			name:     "it should return 2, when input is 2",
			input:    2,
			expected: 2,
		},
		{
			name:     "it should return 6, when input is 3",
			input:    3,
			expected: 6,
		},
		{
			name:     "it should return 24, when input is 4",
			input:    4,
			expected: 24,
		},
		{
			name:     "it should return 120, when input is 5",
			input:    5,
			expected: 120,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := Factorial(testCase.input)
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
