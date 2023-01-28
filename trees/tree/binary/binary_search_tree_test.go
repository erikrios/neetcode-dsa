package binary

import (
	"reflect"
	"testing"
	"trees/tree"
)

func TestNew(t *testing.T) {
	t.Run("it should not nill, when binary tree node is initiated", func(t *testing.T) {
		binaryTree := New(5)
		assertNotNil(t, binaryTree)
	})
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name     string
		target   int
		tree     *BinarySearchTree
		expected bool
	}{
		{
			name:   "it should return false, when target is not exists",
			target: 5,
			tree: &BinarySearchTree{
				Root: &tree.TreeNode{
					Val: 9,
					Left: &tree.TreeNode{
						Val:   3,
						Left:  &tree.TreeNode{Val: 2},
						Right: &tree.TreeNode{Val: 4},
					},
					Right: &tree.TreeNode{
						Val:   11,
						Left:  &tree.TreeNode{Val: 10},
						Right: &tree.TreeNode{Val: 12},
					},
				},
			},
			expected: false,
		},
		{
			name:   "it should return true, when target is exists",
			target: 12,
			tree: &BinarySearchTree{
				Root: &tree.TreeNode{
					Val: 9,
					Left: &tree.TreeNode{
						Val:   3,
						Left:  &tree.TreeNode{Val: 2},
						Right: &tree.TreeNode{Val: 4},
					},
					Right: &tree.TreeNode{
						Val:   11,
						Left:  &tree.TreeNode{Val: 10},
						Right: &tree.TreeNode{Val: 12},
					},
				},
			},
			expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.tree.Search(testCase.target)
			assertValue(t, got, testCase.expected)
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
		t.Fatalf("wamt %#v, but got %#v", want, got)
	}
}
