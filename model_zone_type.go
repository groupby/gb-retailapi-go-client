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

// ZoneType the model 'ZoneType'
type ZoneType string

// List of ZoneType
const (
	QUERY ZoneType = "QUERY"
	CONTENT ZoneType = "CONTENT"
	RICH_CONTENT ZoneType = "RICH_CONTENT"
	EXPECT_QUERY ZoneType = "EXPECT_QUERY"
	GENERATED_CONTENT ZoneType = "GENERATED_CONTENT"
)

// All allowed values of ZoneType enum
var AllowedZoneTypeEnumValues = []ZoneType{
	"QUERY",
	"CONTENT",
	"RICH_CONTENT",
	"EXPECT_QUERY",
	"GENERATED_CONTENT",
}

func (v *ZoneType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ZoneType(value)
	for _, existing := range AllowedZoneTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ZoneType", value)
}

// NewZoneTypeFromValue returns a pointer to a valid ZoneType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewZoneTypeFromValue(v string) (*ZoneType, error) {
	ev := ZoneType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ZoneType: valid values are %v", v, AllowedZoneTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ZoneType) IsValid() bool {
	for _, existing := range AllowedZoneTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ZoneType value
func (v ZoneType) Ptr() *ZoneType {
	return &v
}

type NullableZoneType struct {
	value *ZoneType
	isSet bool
}

func (v NullableZoneType) Get() *ZoneType {
	return v.value
}

func (v *NullableZoneType) Set(val *ZoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneType(val *ZoneType) *NullableZoneType {
	return &NullableZoneType{value: val, isSet: true}
}

func (v NullableZoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
