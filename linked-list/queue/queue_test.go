package queue

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	q := New[string]()
	assertNotNil(t, q)
}

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected bool
	}{
		{
			name:     "it should return true, when queue is empty",
			input:    []string{},
			expected: true,
		},
		{
			name:     "it should return false, when queue is not empty",
			input:    []string{"hello", "world", "erik"},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := New[string]()

			for _, v := range testCase.input {
				q.Enqueue(v)
			}

			got := q.IsEmpty()

			assertValue(t, got, testCase.expected)
		})
	}
}

func TestEnqueue(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "it should return empty slice, when queue is empty",
			expected: []string{},
		},
		{
			name:     "it should return [hello, world, erik] when queue have particular data",
			expected: []string{"hello", "world", "erik"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := New[string]()

			for _, v := range testCase.expected {
				q.Enqueue(v)
			}

			assertValue(t, q.ToList(), testCase.expected)
		})
	}
}

func TestDequeue(t *testing.T) {
	testCases := []struct {
		name               string
		input              []int
		deqTimes           int
		expectedReturnVals []int
		expectedReturnErrs []error
		expectedList       []int
	}{
		{
			name:               "it should return err, when dequeing the empty queue",
			input:              []int{},
			deqTimes:           1,
			expectedReturnVals: []int{0},
			expectedReturnErrs: []error{ErrEmptyQueue},
			expectedList:       []int{},
		},
		{
			name:               "it should return expected values, when dequeue is valid.",
			input:              []int{1, 2, 3, 4, 5},
			deqTimes:           2,
			expectedReturnVals: []int{1, 2},
			expectedReturnErrs: []error{nil, nil},
			expectedList:       []int{3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := New[int]()

			for _, v := range testCase.input {
				q.Enqueue(v)
			}

			for i := 0; i < testCase.deqTimes; i++ {
				got, gotErr := q.Dequeue()
				assertValue(t, got, testCase.expectedReturnVals[i])
				assertValue(t, gotErr, testCase.expectedReturnErrs[i])
			}

			gotList := q.ToList()
			assertValue(t, gotList, testCase.expectedList)
		})
	}

}

func TestToList(t *testing.T) {
	testCases := []struct {
		name     string
		expected []string
	}{
		{
			name:     "it should return empty slice, when queue is empty",
			expected: []string{},
		},
		{
			name:     "it should return [hello, world, erik] when queue have particular data",
			expected: []string{"hello", "world", "erik"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			q := New[string]()

			for _, v := range testCase.expected {
				q.Enqueue(v)
			}

			assertValue(t, q.ToList(), testCase.expected)
		})
	}
}

func assertNotNil(t testing.TB, got any) {
	t.Helper()
	if got == nil {
		t.Fatalf("%#v is nil", got)
	}
}

func assertValue(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but got %#v", want, got)
	}
}
