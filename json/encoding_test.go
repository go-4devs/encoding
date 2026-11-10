package json_test

import (
	"strings"
	"testing"

	"github.com/go-4devs/encoding/internal/require"
	"github.com/go-4devs/encoding/json"
	"github.com/go-4devs/encoding/test"
)

const data = `{"String":"string data","Time":"2020-05-24T09:15:00Z","Struct":{"ID":42}}`

func TestDecode(t *testing.T) {
	t.Parallel()
	test.RunDecode(t, json.Decode, data)
}

func TestDecodeTwoValues(t *testing.T) {
	t.Parallel()

	r := strings.NewReader(`{"A":1}{"A":2}`)

	var first, second struct{ A int }

	require.NoError(t, json.Decode(r, &first))
	require.NoError(t, json.Decode(r, &second))
	require.Equal(t, 1, first.A)
	require.Equal(t, 2, second.A)
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
		{r: 'é', safe: true},
		{r: '"', safe: false},
		{r: '\\', safe: false},
		{r: 0, safe: false},
		{r: -1, safe: false},
	}

	for _, tt := range tests {
		require.Equal(t, tt.safe, json.CanSafeSet(tt.r))
	}
}
