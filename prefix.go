package types

import (
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
