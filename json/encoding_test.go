package json_test

import (
	"testing"

	"github.com/go-4devs/encoding/json"
	"github.com/go-4devs/encoding/test"
	"github.com/stretchr/testify/require"
)

const data = `{"String":"string data","Time":"2020-05-24T09:15:00Z","Struct":{"ID":42}}`

func TestDecode(t *testing.T) {
	t.Parallel()
	test.RunDecode(t, json.Decode, data)
}

func TestEncode(t *testing.T) {
	t.Parallel()
	test.RunEncode(t, json.Encode, json.Decode)
}

func TestCanSafeSet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		r    rune
		safe bool
	}{
		{r: 'a', safe: true},
		{r: ' ', safe: true},
		{r: '\u007f', safe: true},
		{r: '"', safe: false},
		{r: '\\', safe: false},
		{r: 0, safe: false},
	}

	for _, tt := range tests {
		require.Equal(t, tt.safe, json.CanSafeSet(tt.r))
	}
}
