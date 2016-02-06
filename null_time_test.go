package types

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/letsencrypt/boulder/test"
)

func TestTime(t *testing.T) {
	var nt NullTime
	str := []byte("\"2015-08-03T22:43:19.000Z\"")
	err := json.Unmarshal(str, &nt)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, nt.Valid, true)
	test.AssertEquals(t, nt.Time.Year(), 2015)
	test.AssertEquals(t, nt.Time.Second(), 19)
}

func TestNullTime(t *testing.T) {
	var nt NullTime
	str := []byte("null")
	err := json.Unmarshal(str, &nt)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, nt.Valid, false)
}

func TestNullTimeMarshal(t *testing.T) {
	tim, _ := time.Parse("2006-01-02", "2016-01-01")
	nt := NullTime{
		Valid: true,
		Time:  tim,
	}
	bits, err := json.Marshal(nt)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, string(bits), "\"2016-01-01T00:00:00Z\"")
}

func TestNullTimeNullMarshal(t *testing.T) {
	nt := NullTime{
		Valid: false,
		Time:  time.Time{},
	}
	bits, err := json.Marshal(nt)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, string(bits), "null")
}
