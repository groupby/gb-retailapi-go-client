/*
GroupBy Retail

GroupBy Retail API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
	"fmt"
)

// ValueFilterValueFilterType the model 'ValueFilterValueFilterType'
type ValueFilterValueFilterType string

// List of ValueFilter.ValueFilterType
const (
	TEXTUAL ValueFilterValueFilterType = "TEXTUAL"
	NUMERIC ValueFilterValueFilterType = "NUMERIC"
)

// All allowed values of ValueFilterValueFilterType enum
var AllowedValueFilterValueFilterTypeEnumValues = []ValueFilterValueFilterType{
	"TEXTUAL",
	"NUMERIC",
}

func (v *ValueFilterValueFilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ValueFilterValueFilterType(value)
	for _, existing := range AllowedValueFilterValueFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ValueFilterValueFilterType", value)
}

// NewValueFilterValueFilterTypeFromValue returns a pointer to a valid ValueFilterValueFilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewValueFilterValueFilterTypeFromValue(v string) (*ValueFilterValueFilterType, error) {
	ev := ValueFilterValueFilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ValueFilterValueFilterType: valid values are %v", v, AllowedValueFilterValueFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ValueFilterValueFilterType) IsValid() bool {
	for _, existing := range AllowedValueFilterValueFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ValueFilter.ValueFilterType value
func (v ValueFilterValueFilterType) Ptr() *ValueFilterValueFilterType {
	return &v
}

type NullableValueFilterValueFilterType struct {
	value *ValueFilterValueFilterType
	isSet bool
}

func (v NullableValueFilterValueFilterType) Get() *ValueFilterValueFilterType {
	return v.value
}

func (v *NullableValueFilterValueFilterType) Set(val *ValueFilterValueFilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableValueFilterValueFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableValueFilterValueFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueFilterValueFilterType(val *ValueFilterValueFilterType) *NullableValueFilterValueFilterType {
	return &NullableValueFilterValueFilterType{value: val, isSet: true}
}

func (v NullableValueFilterValueFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueFilterValueFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

