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

// NavigationTypeDto Type of navigation.
type NavigationTypeDto string

// List of NavigationTypeDto
const (
	VALUE NavigationTypeDto = "Value"
	RANGE NavigationTypeDto = "Range"
)

// All allowed values of NavigationTypeDto enum
var AllowedNavigationTypeDtoEnumValues = []NavigationTypeDto{
	"Value",
	"Range",
}

func (v *NavigationTypeDto) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NavigationTypeDto(value)
	for _, existing := range AllowedNavigationTypeDtoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NavigationTypeDto", value)
}

// NewNavigationTypeDtoFromValue returns a pointer to a valid NavigationTypeDto
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNavigationTypeDtoFromValue(v string) (*NavigationTypeDto, error) {
	ev := NavigationTypeDto(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NavigationTypeDto: valid values are %v", v, AllowedNavigationTypeDtoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NavigationTypeDto) IsValid() bool {
	for _, existing := range AllowedNavigationTypeDtoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NavigationTypeDto value
func (v NavigationTypeDto) Ptr() *NavigationTypeDto {
	return &v
}

type NullableNavigationTypeDto struct {
	value *NavigationTypeDto
	isSet bool
}

func (v NullableNavigationTypeDto) Get() *NavigationTypeDto {
	return v.value
}

func (v *NullableNavigationTypeDto) Set(val *NavigationTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigationTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigationTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigationTypeDto(val *NavigationTypeDto) *NullableNavigationTypeDto {
	return &NullableNavigationTypeDto{value: val, isSet: true}
}

func (v NullableNavigationTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigationTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

