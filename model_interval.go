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

// checks if the Interval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Interval{}

// Interval A floating point interval.
type Interval struct {
	// Inclusive lower bound. The lower bound of the interval. If neither of the min fields (minimum or exclusiveMinimum) are set, then the lower bound is negative infinity. This field must be not larger than maximum or exclusiveMaximum.
	Minimum *float64 `json:"minimum,omitempty"`
	// Exclusive lower bound. The lower bound of the interval. If neither of the min fields (minimum or exclusiveMinimum) are set, then the lower bound is negative infinity. This field must be not larger than maximum or exclusiveMaximum.
	ExclusiveMinimum *float64 `json:"exclusiveMinimum,omitempty"`
	// Inclusive upper bound. The upper bound of the interval. If neither of the max fields are set, then the upper bound is positive infinity. This field must be not smaller than minimum or exclusiveMinimum.
	Maximum *float64 `json:"maximum,omitempty"`
	// Exclusive upper bound. The upper bound of the interval. If neither of the max fields are set, then the upper bound is positive infinity. This field must be not smaller than minimum or exclusiveMinimum.
	ExclusiveMaximum *float64 `json:"exclusiveMaximum,omitempty"`
}

// NewInterval instantiates a new Interval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterval() *Interval {
	this := Interval{}
	return &this
}

// NewIntervalWithDefaults instantiates a new Interval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntervalWithDefaults() *Interval {
	this := Interval{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *Interval) GetMinimum() float64 {
	if o == nil || IsNil(o.Minimum) {
		var ret float64
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interval) GetMinimumOk() (*float64, bool) {
	if o == nil || IsNil(o.Minimum) {
		return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *Interval) HasMinimum() bool {
	if o != nil && !IsNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given float64 and assigns it to the Minimum field.
func (o *Interval) SetMinimum(v float64) {
	o.Minimum = &v
}

// GetExclusiveMinimum returns the ExclusiveMinimum field value if set, zero value otherwise.
func (o *Interval) GetExclusiveMinimum() float64 {
	if o == nil || IsNil(o.ExclusiveMinimum) {
		var ret float64
		return ret
	}
	return *o.ExclusiveMinimum
}

// GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interval) GetExclusiveMinimumOk() (*float64, bool) {
	if o == nil || IsNil(o.ExclusiveMinimum) {
		return nil, false
	}
	return o.ExclusiveMinimum, true
}

// HasExclusiveMinimum returns a boolean if a field has been set.
func (o *Interval) HasExclusiveMinimum() bool {
	if o != nil && !IsNil(o.ExclusiveMinimum) {
		return true
	}

	return false
}

// SetExclusiveMinimum gets a reference to the given float64 and assigns it to the ExclusiveMinimum field.
func (o *Interval) SetExclusiveMinimum(v float64) {
	o.ExclusiveMinimum = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *Interval) GetMaximum() float64 {
	if o == nil || IsNil(o.Maximum) {
		var ret float64
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interval) GetMaximumOk() (*float64, bool) {
	if o == nil || IsNil(o.Maximum) {
		return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *Interval) HasMaximum() bool {
	if o != nil && !IsNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given float64 and assigns it to the Maximum field.
func (o *Interval) SetMaximum(v float64) {
	o.Maximum = &v
}

// GetExclusiveMaximum returns the ExclusiveMaximum field value if set, zero value otherwise.
func (o *Interval) GetExclusiveMaximum() float64 {
	if o == nil || IsNil(o.ExclusiveMaximum) {
		var ret float64
		return ret
	}
	return *o.ExclusiveMaximum
}

// GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interval) GetExclusiveMaximumOk() (*float64, bool) {
	if o == nil || IsNil(o.ExclusiveMaximum) {
		return nil, false
	}
	return o.ExclusiveMaximum, true
}

// HasExclusiveMaximum returns a boolean if a field has been set.
func (o *Interval) HasExclusiveMaximum() bool {
	if o != nil && !IsNil(o.ExclusiveMaximum) {
		return true
	}

	return false
}

// SetExclusiveMaximum gets a reference to the given float64 and assigns it to the ExclusiveMaximum field.
func (o *Interval) SetExclusiveMaximum(v float64) {
	o.ExclusiveMaximum = &v
}

func (o Interval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Interval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !IsNil(o.ExclusiveMinimum) {
		toSerialize["exclusiveMinimum"] = o.ExclusiveMinimum
	}
	if !IsNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	if !IsNil(o.ExclusiveMaximum) {
		toSerialize["exclusiveMaximum"] = o.ExclusiveMaximum
	}
	return toSerialize, nil
}

type NullableInterval struct {
	value *Interval
	isSet bool
}

func (v NullableInterval) Get() *Interval {
	return v.value
}

func (v *NullableInterval) Set(val *Interval) {
	v.value = val
	v.isSet = true
}

func (v NullableInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterval(val *Interval) *NullableInterval {
	return &NullableInterval{value: val, isSet: true}
}

func (v NullableInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


