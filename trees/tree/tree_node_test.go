package tree

import "testing"

func TestNew(t *testing.T) {
	t.Run("it should not nill, when tree node is initiated", func(t *testing.T) {
		treeNode := New(5)
		assertNotNil(t, treeNode)
	})
}

func assertNotNil(t testing.TB, got any) {
	t.Helper()
	if got == nil {
		t.Fatalf("%#v is nil", got)
	}
}
