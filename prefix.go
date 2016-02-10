package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/nu7hatch/gouuid"
)

type PrefixUUID struct {
	UUID   *uuid.UUID
	Prefix string
}

func (u *PrefixUUID) String() string {
	return u.Prefix + u.UUID.String()
}

func (pu *PrefixUUID) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	if len(s) < 36 {
		return fmt.Errorf("types: could not read \"%s\" as a UUID", s)
	}
	uuidPart := s[len(s)-36:]
	u, err := uuid.ParseHex(uuidPart)
	if err != nil {
		return err
	}
	pu.Prefix = s[:len(s)-36]
	pu.UUID = u
	return nil
}

func (pu PrefixUUID) MarshalJSON() ([]byte, error) {
	if pu.UUID == nil {
		return []byte{}, errors.New("no UUID to convert to JSON")
	}
	return json.Marshal(pu.String())
}

// Scan implements the Scanner interface. Note only the UUID gets scanned/set
// here, we can't determine the prefix from the database. `value` should be
// a [16]byte
func (pu *PrefixUUID) Scan(value interface{}) error {
	if value == nil {
		return errors.New("types: cannot scan null into a PrefixUUID")
	}
	bits, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("types: can't scan value %v into a PrefixUUID", value)
	}
	u, err := uuid.Parse(bits)
	if err != nil {
		return err
	}
	pu.UUID = u
	return nil
}

// Value implements the driver.Valuer interface.
func (pu PrefixUUID) Value() (driver.Value, error) {
	return pu.UUID[:], nil
}
