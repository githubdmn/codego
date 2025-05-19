package greeting

import "testing"

func TestHello(t *testing.T) {
    expected := "Hello, Bob!"
    actual := Hello("Bob")
    if actual != expected {
        t.Errorf("Expected %q, got %q", expected, actual)
    }
}

