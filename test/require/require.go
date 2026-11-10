package require

import (
	"reflect"
	"testing"
)

// NoError stops the test when err is not nil.
func NoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

// Equal stops the test when expected and actual differ.
func Equal(t *testing.T, expected, actual any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected %#v, got %#v", expected, actual)
	}
}

// EqualError stops the test when err is nil or its text is not msg.
func EqualError(t *testing.T, err error, msg string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error %q, got nil", msg)
	}

	if err.Error() != msg {
		t.Fatalf("expected error %q, got %q", msg, err.Error())
	}
}
