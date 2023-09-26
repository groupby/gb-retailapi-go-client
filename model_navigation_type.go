/*
GroupBy Retail

GroupBy Retail API

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
	"fmt"
)

// NavigationType Type of navigation.
type NavigationType string

// List of NavigationType
const (
	VALUE NavigationType = "VALUE"
	RANGE NavigationType = "RANGE"
)

// All allowed values of NavigationType enum
var AllowedNavigationTypeEnumValues = []NavigationType{
	"VALUE",
	"RANGE",
}

func (v *NavigationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NavigationType(value)
	for _, existing := range AllowedNavigationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NavigationType", value)
}

// NewNavigationTypeFromValue returns a pointer to a valid NavigationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNavigationTypeFromValue(v string) (*NavigationType, error) {
	ev := NavigationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NavigationType: valid values are %v", v, AllowedNavigationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NavigationType) IsValid() bool {
	for _, existing := range AllowedNavigationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NavigationType value
func (v NavigationType) Ptr() *NavigationType {
	return &v
}

type NullableNavigationType struct {
	value *NavigationType
	isSet bool
}

func (v NullableNavigationType) Get() *NavigationType {
	return v.value
}

func (v *NullableNavigationType) Set(val *NavigationType) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigationType) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigationType(val *NavigationType) *NullableNavigationType {
	return &NullableNavigationType{value: val, isSet: true}
}

func (v NullableNavigationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

