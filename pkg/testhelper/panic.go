package testhelper

import "testing"

// ShouldPanic helps to test the panic case
// Refer: https://gist.github.com/wrunk/4afea3d85cc9feb7fd8fcef5a8a98b5e
func ShouldPanic(t *testing.T, f func()) {
	defer func() { recover() }()
	f()
	t.Errorf("should have panicked")
}
