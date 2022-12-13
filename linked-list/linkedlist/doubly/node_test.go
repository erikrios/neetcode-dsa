package doubly

import "testing"

func TestNewNode(t *testing.T) {
	node := NewNode[complex64](128.5, nil, nil)
	assertNotNil(t, node)
}

func assertNotNil(t testing.TB, got any) {
	t.Helper()
	if got == nil {
		t.Fatalf("%#v is nil", got)
	}
}
