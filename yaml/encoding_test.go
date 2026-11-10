package yaml_test

import (
	"strings"
	"testing"

	"github.com/go-4devs/encoding/internal/test"
	"github.com/go-4devs/encoding/internal/test/require"
	"github.com/go-4devs/encoding/yaml"
)

const data = `
string: "string data"
time: 2020-05-24T09:15:00Z
struct:
    id: 42
`

func TestDecode(t *testing.T) {
	t.Parallel()
	test.RunDecode(t, yaml.Decode, data)
}

func TestDecodeExtraDocument(t *testing.T) {
	t.Parallel()

	r := strings.NewReader("a: 1\n---\nb: 2\n")

	var got map[string]int

	require.EqualError(t, yaml.Decode(r, &got), "yaml: extra document")
	require.Equal(t, map[string]int{"a": 1}, got)
}

func TestEncode(t *testing.T) {
	t.Parallel()
	test.RunEncode(t, yaml.Encode, yaml.Decode)
}
