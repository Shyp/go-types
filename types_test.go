package types

import (
	"encoding/json"
	"testing"

	"github.com/Shyp/goshyp/Godeps/_workspace/src/github.com/letsencrypt/boulder/test"
)

func TestString(t *testing.T) {
	var ns NullString
	str := []byte("\"foo\"")
	err := json.Unmarshal(str, &ns)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, ns.Valid, true)
	test.AssertEquals(t, ns.String, "foo")
}

func TestNullString(t *testing.T) {
	var ns NullString
	str := []byte("null")
	err := json.Unmarshal(str, &ns)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, ns.Valid, false)
}

func TestStringMarshal(t *testing.T) {
	ns := NullString{
		Valid:  true,
		String: "foo bar",
	}
	b, err := json.Marshal(ns)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, string(b), "\"foo bar\"")
}

func TestStringMarshalNull(t *testing.T) {
	ns := NullString{
		Valid:  false,
		String: "",
	}
	b, err := json.Marshal(ns)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, string(b), "null")
}
