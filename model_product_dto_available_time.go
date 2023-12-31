/*
GroupBy Retail

GroupBy Retail API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
)

// checks if the ProductDtoAvailableTime type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDtoAvailableTime{}

// ProductDtoAvailableTime struct for ProductDtoAvailableTime
type ProductDtoAvailableTime struct {
	// Timestamp seconds.
	Seconds *int64 `json:"seconds,omitempty"`
	// Timestamp nanos.
	Nanos *int32 `json:"nanos,omitempty"`
}

// NewProductDtoAvailableTime instantiates a new ProductDtoAvailableTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDtoAvailableTime() *ProductDtoAvailableTime {
	this := ProductDtoAvailableTime{}
	return &this
}

// NewProductDtoAvailableTimeWithDefaults instantiates a new ProductDtoAvailableTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDtoAvailableTimeWithDefaults() *ProductDtoAvailableTime {
	this := ProductDtoAvailableTime{}
	return &this
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *ProductDtoAvailableTime) GetSeconds() int64 {
	if o == nil || IsNil(o.Seconds) {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDtoAvailableTime) GetSecondsOk() (*int64, bool) {
	if o == nil || IsNil(o.Seconds) {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *ProductDtoAvailableTime) HasSeconds() bool {
	if o != nil && !IsNil(o.Seconds) {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *ProductDtoAvailableTime) SetSeconds(v int64) {
	o.Seconds = &v
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *ProductDtoAvailableTime) GetNanos() int32 {
	if o == nil || IsNil(o.Nanos) {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDtoAvailableTime) GetNanosOk() (*int32, bool) {
	if o == nil || IsNil(o.Nanos) {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *ProductDtoAvailableTime) HasNanos() bool {
	if o != nil && !IsNil(o.Nanos) {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *ProductDtoAvailableTime) SetNanos(v int32) {
	o.Nanos = &v
}

func (o ProductDtoAvailableTime) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDtoAvailableTime) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Seconds) {
		toSerialize["seconds"] = o.Seconds
	}
	if !IsNil(o.Nanos) {
		toSerialize["nanos"] = o.Nanos
	}
	return toSerialize, nil
}

type NullableProductDtoAvailableTime struct {
	value *ProductDtoAvailableTime
	isSet bool
}

func (v NullableProductDtoAvailableTime) Get() *ProductDtoAvailableTime {
	return v.value
}

func (v *NullableProductDtoAvailableTime) Set(val *ProductDtoAvailableTime) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDtoAvailableTime) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDtoAvailableTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDtoAvailableTime(val *ProductDtoAvailableTime) *NullableProductDtoAvailableTime {
	return &NullableProductDtoAvailableTime{value: val, isSet: true}
}

func (v NullableProductDtoAvailableTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDtoAvailableTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


