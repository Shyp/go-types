package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/letsencrypt/boulder/test"
	"github.com/nu7hatch/gouuid"
)

func TestUUIDString(t *testing.T) {
	u, _ := uuid.NewV4()
	pfx := PrefixUUID{
		Prefix: "job_",
		UUID:   u,
	}
	test.AssertEquals(t, pfx.String(), fmt.Sprintf("job_%s", u))
}

func TestNewUUIDPrefix(t *testing.T) {
	u, _ := uuid.NewV4()
	pfx, err := NewPrefixUUID("job_", u.String())
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, pfx.Prefix, "job_")
	test.AssertEquals(t, pfx.UUID.String(), u.String())
}

var unmarshalTests = []struct {
	in         string
	prefix     string
	uuidString string
	err        error
}{
	{"usr_6740b44e-13b9-475d-af06-979627e0e0d6", "usr_", "6740b44e-13b9-475d-af06-979627e0e0d6", nil},
	{"6740b44e-13b9-475d-af06-979627e0e0d6", "", "6740b44e-13b9-475d-af06-979627e0e0d6", nil},
	{"", "", "", errors.New("types: could not read \"\" as a UUID")},
	{"foo", "", "", errors.New("types: could not read \"foo\" as a UUID")},
	{"6740b44e-13b9-475d-af069-79627e0e0d6", "", "", errors.New("Invalid UUID string")},
}

func TestUUIDUnmarshal(t *testing.T) {
	for _, tt := range unmarshalTests {
		var pfxu PrefixUUID
		err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", tt.in)), &pfxu)
		if tt.err != nil {
			test.AssertError(t, err, "")
			test.AssertEquals(t, err.Error(), tt.err.Error())
		} else {
			test.AssertNotError(t, err, "")
			test.AssertEquals(t, pfxu.Prefix, tt.prefix)
			test.AssertEquals(t, pfxu.UUID.String(), tt.uuidString)
		}
	}
}

func TestUUIDMarshal(t *testing.T) {
	u, _ := uuid.ParseHex("6740b44e-13b9-475d-af06-979627e0e0d6")
	pfx := &PrefixUUID{
		Prefix: "usr_",
		UUID:   u,
	}
	b, err := json.Marshal(pfx)
	test.AssertNotError(t, err, "")
	test.AssertEquals(t, string(b), "\"usr_6740b44e-13b9-475d-af06-979627e0e0d6\"")

	pfx = &PrefixUUID{
		Prefix: "usr_",
		UUID:   nil,
	}
	_, err = json.Marshal(pfx)
	test.AssertEquals(t, err.Error(), "json: error calling MarshalJSON for type *types.PrefixUUID: no UUID to convert to JSON")
}
