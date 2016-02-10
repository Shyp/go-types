package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

const version = "0.2"

type NullString struct {
	sql.NullString
	Valid	bool
	String	string
}

func (ns *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		ns.Valid = false
		return nil
	}
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	ns.Valid = true
	ns.String = s
	return nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid == false {
		return []byte("null"), nil
	}
	s, err := json.Marshal(ns.String)
	if err != nil {
		return []byte{}, err
	}
	return s, nil
}

func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	ns.String, ns.Valid = value.(string)
	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}
