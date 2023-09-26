/*
GroupBy Retail

GroupBy Retail API

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
)

// checks if the RefinementDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefinementDto{}

// RefinementDto Refinement value or range in the navigation.
type RefinementDto struct {
	Type NavigationTypeDto `json:"type"`
	// Number of products which match this refinement value or range.
	Count *int64 `json:"count,omitempty"`
	// Value of the refinement.
	Value *string `json:"value,omitempty"`
	// Lower bound of the refinement range.
	Low *float64 `json:"low,omitempty"`
	// Upper bound  of the refinement range.
	High *float64 `json:"high,omitempty"`
}

// NewRefinementDto instantiates a new RefinementDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefinementDto(type_ NavigationTypeDto) *RefinementDto {
	this := RefinementDto{}
	this.Type = type_
	return &this
}

// NewRefinementDtoWithDefaults instantiates a new RefinementDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefinementDtoWithDefaults() *RefinementDto {
	this := RefinementDto{}
	return &this
}

// GetType returns the Type field value
func (o *RefinementDto) GetType() NavigationTypeDto {
	if o == nil {
		var ret NavigationTypeDto
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RefinementDto) GetTypeOk() (*NavigationTypeDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RefinementDto) SetType(v NavigationTypeDto) {
	o.Type = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RefinementDto) GetCount() int64 {
	if o == nil || IsNil(o.Count) {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementDto) GetCountOk() (*int64, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RefinementDto) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RefinementDto) SetCount(v int64) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *RefinementDto) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementDto) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *RefinementDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *RefinementDto) SetValue(v string) {
	o.Value = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *RefinementDto) GetLow() float64 {
	if o == nil || IsNil(o.Low) {
		var ret float64
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementDto) GetLowOk() (*float64, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *RefinementDto) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given float64 and assigns it to the Low field.
func (o *RefinementDto) SetLow(v float64) {
	o.Low = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *RefinementDto) GetHigh() float64 {
	if o == nil || IsNil(o.High) {
		var ret float64
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefinementDto) GetHighOk() (*float64, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *RefinementDto) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given float64 and assigns it to the High field.
func (o *RefinementDto) SetHigh(v float64) {
	o.High = &v
}

func (o RefinementDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefinementDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	return toSerialize, nil
}

type NullableRefinementDto struct {
	value *RefinementDto
	isSet bool
}

func (v NullableRefinementDto) Get() *RefinementDto {
	return v.value
}

func (v *NullableRefinementDto) Set(val *RefinementDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRefinementDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRefinementDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefinementDto(val *RefinementDto) *NullableRefinementDto {
	return &NullableRefinementDto{value: val, isSet: true}
}

func (v NullableRefinementDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefinementDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

