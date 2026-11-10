package xml_test

import (
	"testing"

	"github.com/go-4devs/encoding/test"
	"github.com/go-4devs/encoding/xml"
)

func TestEncode(t *testing.T) {
	t.Parallel()
	test.RunEncode(t, xml.Encode, xml.Decode)
}
