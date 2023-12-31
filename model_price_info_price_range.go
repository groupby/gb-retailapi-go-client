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

// checks if the PriceInfoPriceRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceInfoPriceRange{}

// PriceInfoPriceRange struct for PriceInfoPriceRange
type PriceInfoPriceRange struct {
	Price *PriceInfoPriceRangePrice `json:"price,omitempty"`
	OriginalPrice *PriceInfoPriceRangeOriginalPrice `json:"originalPrice,omitempty"`
}

// NewPriceInfoPriceRange instantiates a new PriceInfoPriceRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceInfoPriceRange() *PriceInfoPriceRange {
	this := PriceInfoPriceRange{}
	return &this
}

// NewPriceInfoPriceRangeWithDefaults instantiates a new PriceInfoPriceRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceInfoPriceRangeWithDefaults() *PriceInfoPriceRange {
	this := PriceInfoPriceRange{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *PriceInfoPriceRange) GetPrice() PriceInfoPriceRangePrice {
	if o == nil || IsNil(o.Price) {
		var ret PriceInfoPriceRangePrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceInfoPriceRange) GetPriceOk() (*PriceInfoPriceRangePrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *PriceInfoPriceRange) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given PriceInfoPriceRangePrice and assigns it to the Price field.
func (o *PriceInfoPriceRange) SetPrice(v PriceInfoPriceRangePrice) {
	o.Price = &v
}

// GetOriginalPrice returns the OriginalPrice field value if set, zero value otherwise.
func (o *PriceInfoPriceRange) GetOriginalPrice() PriceInfoPriceRangeOriginalPrice {
	if o == nil || IsNil(o.OriginalPrice) {
		var ret PriceInfoPriceRangeOriginalPrice
		return ret
	}
	return *o.OriginalPrice
}

// GetOriginalPriceOk returns a tuple with the OriginalPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceInfoPriceRange) GetOriginalPriceOk() (*PriceInfoPriceRangeOriginalPrice, bool) {
	if o == nil || IsNil(o.OriginalPrice) {
		return nil, false
	}
	return o.OriginalPrice, true
}

// HasOriginalPrice returns a boolean if a field has been set.
func (o *PriceInfoPriceRange) HasOriginalPrice() bool {
	if o != nil && !IsNil(o.OriginalPrice) {
		return true
	}

	return false
}

// SetOriginalPrice gets a reference to the given PriceInfoPriceRangeOriginalPrice and assigns it to the OriginalPrice field.
func (o *PriceInfoPriceRange) SetOriginalPrice(v PriceInfoPriceRangeOriginalPrice) {
	o.OriginalPrice = &v
}

func (o PriceInfoPriceRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceInfoPriceRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.OriginalPrice) {
		toSerialize["originalPrice"] = o.OriginalPrice
	}
	return toSerialize, nil
}

type NullablePriceInfoPriceRange struct {
	value *PriceInfoPriceRange
	isSet bool
}

func (v NullablePriceInfoPriceRange) Get() *PriceInfoPriceRange {
	return v.value
}

func (v *NullablePriceInfoPriceRange) Set(val *PriceInfoPriceRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceInfoPriceRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceInfoPriceRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceInfoPriceRange(val *PriceInfoPriceRange) *NullablePriceInfoPriceRange {
	return &NullablePriceInfoPriceRange{value: val, isSet: true}
}

func (v NullablePriceInfoPriceRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceInfoPriceRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


