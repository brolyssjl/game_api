package utils

import (
	"database/sql/driver"
	"strconv"
)

type NullInt64 struct {
	Val     int64
	IsValid bool
}

func NewNullInt64(val interface{}) NullInt64 {
	ni := NullInt64{}
	ni.Set(val)
	return ni
}

func (ni NullInt64) Value() (driver.Value, error) {
	if !ni.IsValid {
		return nil, nil
	}
	return ni.Val, nil
}

func (ni *NullInt64) Set(val interface{}) {
	ni.Val, ni.IsValid = val.(int64)
}

func (ni *NullInt64) Scan(value interface{}) error {
	ni.Val, ni.IsValid = value.(int64)
	return nil
}

func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.IsValid {
		return []byte(strconv.FormatInt(0, 10)), nil
	}

	return []byte(strconv.FormatInt(ni.Val, 10)), nil
}
