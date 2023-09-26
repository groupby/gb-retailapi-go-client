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

// ZoneDtoType Define type of content which is can be stored in zone.
type ZoneDtoType string

// List of ZoneDtoType
const (
	CONTENT ZoneDtoType = "Content"
	RICH_CONTENT ZoneDtoType = "Rich_Content"
	PRODUCTS ZoneDtoType = "Products"
	GENERATED_CONTENT ZoneDtoType = "Generated_Content"
)

// All allowed values of ZoneDtoType enum
var AllowedZoneDtoTypeEnumValues = []ZoneDtoType{
	"Content",
	"Rich_Content",
	"Products",
	"Generated_Content",
}

func (v *ZoneDtoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZoneDtoType(value)
	for _, existing := range AllowedZoneDtoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZoneDtoType", value)
}

// NewZoneDtoTypeFromValue returns a pointer to a valid ZoneDtoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneDtoTypeFromValue(v string) (*ZoneDtoType, error) {
	ev := ZoneDtoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZoneDtoType: valid values are %v", v, AllowedZoneDtoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneDtoType) IsValid() bool {
	for _, existing := range AllowedZoneDtoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ZoneDtoType value
func (v ZoneDtoType) Ptr() *ZoneDtoType {
	return &v
}

type NullableZoneDtoType struct {
	value *ZoneDtoType
	isSet bool
}

func (v NullableZoneDtoType) Get() *ZoneDtoType {
	return v.value
}

func (v *NullableZoneDtoType) Set(val *ZoneDtoType) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneDtoType) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneDtoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneDtoType(val *ZoneDtoType) *NullableZoneDtoType {
	return &NullableZoneDtoType{value: val, isSet: true}
}

func (v NullableZoneDtoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneDtoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

