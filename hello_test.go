package hello

import "testing"

func TestHello(t *testing.T) {
  expected := "Hello, world."
  if got := Hello(); got != expected {
    t.Errorf("Hello test failed")
  }
}
