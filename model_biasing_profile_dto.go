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

// checks if the BiasingProfileDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BiasingProfileDto{}

// BiasingProfileDto Inline biasing profile, which should be applied to the search. Takes priority over biasing profile.
type BiasingProfileDto struct {
	Biases []BiasDto `json:"biases"`
}

// NewBiasingProfileDto instantiates a new BiasingProfileDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBiasingProfileDto(biases []BiasDto) *BiasingProfileDto {
	this := BiasingProfileDto{}
	this.Biases = biases
	return &this
}

// NewBiasingProfileDtoWithDefaults instantiates a new BiasingProfileDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBiasingProfileDtoWithDefaults() *BiasingProfileDto {
	this := BiasingProfileDto{}
	return &this
}

// GetBiases returns the Biases field value
func (o *BiasingProfileDto) GetBiases() []BiasDto {
	if o == nil {
		var ret []BiasDto
		return ret
	}

	return o.Biases
}

// GetBiasesOk returns a tuple with the Biases field value
// and a boolean to check if the value has been set.
func (o *BiasingProfileDto) GetBiasesOk() ([]BiasDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Biases, true
}

// SetBiases sets field value
func (o *BiasingProfileDto) SetBiases(v []BiasDto) {
	o.Biases = v
}

func (o BiasingProfileDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BiasingProfileDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["biases"] = o.Biases
	return toSerialize, nil
}

type NullableBiasingProfileDto struct {
	value *BiasingProfileDto
	isSet bool
}

func (v NullableBiasingProfileDto) Get() *BiasingProfileDto {
	return v.value
}

func (v *NullableBiasingProfileDto) Set(val *BiasingProfileDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBiasingProfileDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBiasingProfileDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBiasingProfileDto(val *BiasingProfileDto) *NullableBiasingProfileDto {
	return &NullableBiasingProfileDto{value: val, isSet: true}
}

func (v NullableBiasingProfileDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBiasingProfileDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


