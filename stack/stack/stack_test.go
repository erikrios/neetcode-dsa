package stack

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("it should not nil, when stack is instantiated", func(t *testing.T) {
		stack := New[string]()
		assertNotNil(t, stack)
	})
}

func TestPush(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "it should return []int{1} when Push function is called",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "it should return []int{1, 2} when Push function is called",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "it should return []int{1, 2, 3} when Push function is called",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New[int]()

			for _, v := range testCase.input {
				stack.Push(v)
			}

			if !reflect.DeepEqual(stack.items, testCase.expected) {
				assertValue(t, stack.items, testCase.expected)
			}
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		name        string
		input       []string
		expected    string
		expectedErr error
	}{
		{
			name:        "it should return error, when stack is empty",
			input:       []string{},
			expected:    "",
			expectedErr: StackEmptyErr,
		},
		{
			name:        "it should return last items, when stack is not empty",
			input:       []string{"hello", "world", "erik"},
			expected:    "erik",
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New[string]()

			for _, v := range testCase.input {
				stack.Push(v)
			}

			got, err := stack.Pop()
			assertError(t, err, testCase.expectedErr)
			assertValue(t, got, testCase.expected)
		})
	}
}

func TestPeek(t *testing.T) {
	testCases := []struct {
		name        string
		input       []float64
		expected    float64
		expectedErr error
	}{
		{
			name:        "it should return error, when stack is empty",
			input:       []float64{},
			expected:    0.0,
			expectedErr: StackEmptyErr,
		},
		{
			name:        "it should return last items, when stack is not empty",
			input:       []float64{1.0, 2.5, 0.88},
			expected:    0.88,
			expectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			stack := New[float64]()

			for _, v := range testCase.input {
				stack.Push(v)
			}

			got, err := stack.Peek()
			assertError(t, err, testCase.expectedErr)
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

func assertNotNil(t testing.TB, got any) {
	t.Helper()
	if got == nil {
		t.Fatalf("%#v is nil", got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("wamt %#v, but got %#v", want, got)
	}
}
