package test

import (
	"bytes"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gitoa.ru/go-4devs/encoding"
)

type Data struct {
	String string
	Time   time.Time
	Struct SData
}

type SData struct {
	ID int64
}

func expected() Data {
	return Data{
		String: "string data",
		Time:   time.Date(2020, time.May, 24, 9, 15, 0, 0, time.UTC),
		Struct: SData{ID: 42},
	}
}

type DecodeSuite struct {
	suite.Suite

	decode encoding.Decode
	data   io.Reader
}

// RunDecode run test by provider.
func RunDecode(t *testing.T, decode encoding.Decode, data string) {
	t.Helper()

	caseSuite := DecodeSuite{
		decode: decode,
		data:   bytes.NewBufferString(data),
		Suite: suite.Suite{
			Assertions: nil,
		},
	}

	suite.Run(t, &caseSuite)
}

func (ds *DecodeSuite) TestDecode() {
	var d Data

	ds.Require().NoError(ds.decode(ds.data, &d))
	ds.Require().Equal(expected(), d)
}
