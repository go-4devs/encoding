package yaml_test

import (
	"testing"

	"github.com/go-4devs/encoding/test"
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

func TestEncode(t *testing.T) {
	t.Parallel()
	test.RunEncode(t, yaml.Encode, yaml.Decode)
}
