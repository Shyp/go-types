package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/kevinburke/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func ExamplePrefixUUID() {
	p, _ := NewPrefixUUID("usr_6740b44e-13b9-475d-af06-979627e0e0d6")
	fmt.Println(p.Prefix)
	fmt.Println(p.UUID.String())
	fmt.Println(p.String())
	// Output: usr_
	// 6740b44e-13b9-475d-af06-979627e0e0d6
	// usr_6740b44e-13b9-475d-af06-979627e0e0d6
}

func TestUUIDString(t *testing.T) {
	u := uuid.NewV4()
	pfx := PrefixUUID{
		Prefix: "job_",
		UUID:   u,
	}
	assertEquals(t, pfx.String(), fmt.Sprintf("job_%s", u))
}

func TestNewPrefixUUID(t *testing.T) {
	pfx, err := NewPrefixUUID("usr_6740b44e-13b9-475d-af06-979627e0e0d6")
	assertNotError(t, err, "")
	assertEquals(t, pfx.Prefix, "usr_")
	assertEquals(t, pfx.UUID.String(), "6740b44e-13b9-475d-af06-979627e0e0d6")

	pfx, err = NewPrefixUUID("6740b44e-13b9-475d-af06-979627e0e0d6")
	assertNotError(t, err, "")
	assertEquals(t, pfx.Prefix, "")
	assertEquals(t, pfx.UUID.String(), "6740b44e-13b9-475d-af06-979627e0e0d6")
}

func TestGenerateUUID(t *testing.T) {
	id := GenerateUUID("job_")
	assertEquals(t, id.Prefix, "job_")
	assert(t, len(id.String()) > 20, "")
}

var unmarshalTests = []struct {
	in         string
	prefix     string
	uuidString string
	err        error
}{
	{"usr_6740b44e-13b9-475d-af06-979627e0e0d6", "usr_", "6740b44e-13b9-475d-af06-979627e0e0d6", nil},
	{"6740b44e-13b9-475d-af06-979627e0e0d6", "", "6740b44e-13b9-475d-af06-979627e0e0d6", nil},
	{"", "", "", errors.New("types: Could not parse \"\" as a UUID with a prefix")},
	{"foo", "", "", errors.New("types: Could not parse \"foo\" as a UUID with a prefix")},
	// Has one extra char.
	{"6740b44e-13b9-475d-af069-79627e0e0d6", "", "", errors.New("uuid: incorrect UUID format 6740b44e-13b9-475d-af069-79627e0e0d6")},
}

func TestUUIDUnmarshal(t *testing.T) {
	for _, tt := range unmarshalTests {
		var pfxu PrefixUUID
		err := json.Unmarshal([]byte(fmt.Sprintf(`"%s"`, tt.in)), &pfxu)
		if tt.err != nil {
			assertError(t, err, "")
			assertEquals(t, err.Error(), tt.err.Error())
		} else {
			assertNotError(t, err, "")
			assertEquals(t, pfxu.Prefix, tt.prefix)
			assertEquals(t, pfxu.UUID.String(), tt.uuidString)
		}
	}
}

func TestUUIDMarshal(t *testing.T) {
	u, _ := uuid.FromString("6740b44e-13b9-475d-af06-979627e0e0d6")
	pfx := &PrefixUUID{
		Prefix: "usr_",
		UUID:   u,
	}
	b, err := json.Marshal(pfx)
	assertNotError(t, err, "")
	assertEquals(t, string(b), "\"usr_6740b44e-13b9-475d-af06-979627e0e0d6\"")
}

func TestScan(t *testing.T) {
	var pu PrefixUUID
	err := pu.Scan([]byte("pik_6740b44e-13b9-475d-af06-979627e0e0d6"))
	assertNotError(t, err, "scanning byte array")
	assertEquals(t, pu.Prefix, "pik_")
	assertEquals(t, pu.UUID.String(), "6740b44e-13b9-475d-af06-979627e0e0d6")

	err = pu.Scan("pik_6740b44e-13b9-475d-af06-979627e0e0d6")
	assertNotError(t, err, "scanning string")
	assertEquals(t, pu.Prefix, "pik_")
	assertEquals(t, pu.UUID.String(), "6740b44e-13b9-475d-af06-979627e0e0d6")

	err = pu.Scan([]byte("6740b44e-13b9-475d-af06-979627e0e0d6"))
	assertNotError(t, err, "scanning byte array")
	assertEquals(t, pu.Prefix, "")

	err = pu.Scan([]byte{0x67, 0x40, 0xb4, 0x4e, 0x13, 0xb9, 0x47, 0x5d, 0xaf, 0x6, 0x97, 0x96, 0x27, 0xe0, 0xe0, 0xd6})
	assertNotError(t, err, "scanning byte array")
	assertEquals(t, pu.Prefix, "")
	assertEquals(t, pu.UUID.String(), "6740b44e-13b9-475d-af06-979627e0e0d6")

	err = pu.Scan(7)
	assertError(t, err, "scanning a number")
	assertEquals(t, err.Error(), "types: can't scan value of unknown type 7 into a PrefixUUID")
}

func TestSetBSON(t *testing.T) {
	var pu PrefixUUID
	err := pu.SetBSON(bson.Raw{Data: []byte{0x25, 0x0, 0x0, 0x0, 0x31, 0x36, 0x38, 0x65, 0x37, 0x31, 0x36, 0x61, 0x2d, 0x37, 0x39, 0x34, 0x61, 0x2d, 0x31, 0x31, 0x65, 0x37, 0x2d, 0x39, 0x35, 0x30, 0x62, 0x2d, 0x34, 0x63, 0x33, 0x32, 0x37, 0x35, 0x39, 0x32, 0x34, 0x32, 0x61, 0x35, 0x0}})
	assertNotError(t, err, "setting BSON")
	assertEquals(t, pu.String(), "168e716a-794a-11e7-950b-4c32759242a5")
}
