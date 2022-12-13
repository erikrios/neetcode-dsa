package doubly

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	testCases := []struct {
		name  string
		input []float32
	}{
		{
			name:  "it should return empty slice, when the linked list is empty",
			input: []float32{},
		},
		{
			name:  "it should return []int{1.1}, when the linked list is {1.1}",
			input: []float32{1.1},
		},
		{
			name:  "it should return []int{1.1, 2.2}, when the linked list is {1.1, 2.2}",
			input: []float32{1.1, 2.2},
		},
		{
			name:  "it should return []int{1.1, 2.2, 3.3}, when the linked list is {1.1, 2.2, 3.3}",
			input: []float32{1.1, 2.2, 3.3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll := New[float32]()

			for _, v := range testCase.input {
				ll.Insert(v)
			}

			got := ll.ToList()

			assertValue(t, got, testCase.input)
		})
	}
}

func TestInsertAt(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		val      int
		index    int
		expected []int
	}{
		{
			name:     "it should return []int{1}, when input is []int{} and index is 0",
			index:    0,
			val:      1,
			input:    []int{},
			expected: []int{1},
		},
		{
			name:     "it should return []int{2, 1}, when input is []int{1} and index is 0",
			index:    0,
			val:      2,
			input:    []int{1},
			expected: []int{2, 1},
		},
		{
			name:     "it should return []int{1}, when input is []int{2} and index is 1",
			index:    1,
			val:      2,
			input:    []int{},
			expected: []int{2},
		},
		{
			name:     "it should return []int{1, 2, 3, 4}, when input is []int{1, 2, 3} and index is 10",
			index:    10,
			val:      4,
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "it should return []int{1, 2, 3, 4}, when input is []int{1, 2, 3} and index is 3",
			index:    3,
			val:      4,
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "it should return []int{1, 2, 3, 4, 5}, when input is []int{1, 2, 4, 5} and index is 2",
			index:    2,
			val:      3,
			input:    []int{1, 2, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll := New[int]()
			for _, v := range testCase.input {
				ll.Insert(v)
			}

			ll.InsertAt(testCase.index, testCase.val)

			got := ll.ToList()
			assertValue(t, got, testCase.expected)
		})
	}
}

func TestFind(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		initialValues []string
		expected      bool
	}{
		{
			name:          "it should return false, when input is not present in the linked list",
			input:         "hello",
			initialValues: []string{"halo", "hola", "world"},
			expected:      false,
		},
		{
			name:          "it should return true, when input is present in the linked list",
			input:         "hello",
			initialValues: []string{"halo", "hola", "hello", "world"},
			expected:      true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll := New[string]()

			for _, v := range testCase.initialValues {
				ll.Insert(v)
			}

			got := ll.Find(testCase.input)
			assertValue(t, got, testCase.expected)
		})
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name          string
		initialValues []int
		expected      []int
	}{
		{
			name:          "it should return []int{}, when initialValues is []int{}",
			initialValues: []int{},
			expected:      []int{},
		},
		{
			name:          "it should return []int{}, when initialValues is []int{5}",
			initialValues: []int{5},
			expected:      []int{},
		},
		{
			name:          "it should return []int{5}, when initialValues is []int{5, 4}",
			initialValues: []int{5, 4},
			expected:      []int{5},
		},
		{
			name:          "it should return []int{5, 4, 3, 2}, when initialValues is []int{5, 4, 3, 2, 1}",
			initialValues: []int{5, 4, 3, 2, 1},
			expected:      []int{5, 4, 3, 2},
		},
	}

	for _, testCase := range testCases {
		ll := New[int]()

		for _, v := range testCase.initialValues {
			ll.Insert(v)
		}

		ll.Delete()

		got := ll.ToList()

		assertValue(t, got, testCase.expected)
	}
}

func TestToList(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "it should return empty slice, when the linked list is empty",
			input: []int{},
		},
		{
			name:  "it should return []int{1}, when the linked list is {1}",
			input: []int{1},
		},
		{
			name:  "it should return []int{1, 2}, when the linked list is {1, 2}",
			input: []int{1, 2},
		},
		{
			name:  "it should return []int{1, 2, 3}, when the linked list is {1, 2, 3}",
			input: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ll := New[int]()

			for _, v := range testCase.input {
				ll.Insert(v)
			}

			got := ll.ToList()

			assertValue(t, got, testCase.input)
		})
	}
}

func assertValue(t testing.TB, got any, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but got %#v", want, got)
	}
}
