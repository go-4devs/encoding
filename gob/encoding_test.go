package gob_test

import (
	"testing"

	"github.com/go-4devs/encoding/gob"
	"github.com/go-4devs/encoding/internal/test"
)

func TestEncode(t *testing.T) {
	t.Parallel()
	test.RunEncode(t, gob.Encode, gob.Decode)
}

func TestMarshal(t *testing.T) {
	t.Parallel()
	test.RunMarshal(t, gob.Marshal, gob.Unmarshal)
}
