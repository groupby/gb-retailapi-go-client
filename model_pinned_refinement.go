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

// checks if the PinnedRefinement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinnedRefinement{}

// PinnedRefinement struct for PinnedRefinement
type PinnedRefinement struct {
	Navigation string `json:"navigation"`
	Refinements []Refinement `json:"refinements"`
}

// NewPinnedRefinement instantiates a new PinnedRefinement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinnedRefinement(navigation string, refinements []Refinement) *PinnedRefinement {
	this := PinnedRefinement{}
	this.Navigation = navigation
	this.Refinements = refinements
	return &this
}

// NewPinnedRefinementWithDefaults instantiates a new PinnedRefinement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinnedRefinementWithDefaults() *PinnedRefinement {
	this := PinnedRefinement{}
	return &this
}

// GetNavigation returns the Navigation field value
func (o *PinnedRefinement) GetNavigation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Navigation
}

// GetNavigationOk returns a tuple with the Navigation field value
// and a boolean to check if the value has been set.
func (o *PinnedRefinement) GetNavigationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Navigation, true
}

// SetNavigation sets field value
func (o *PinnedRefinement) SetNavigation(v string) {
	o.Navigation = v
}

// GetRefinements returns the Refinements field value
func (o *PinnedRefinement) GetRefinements() []Refinement {
	if o == nil {
		var ret []Refinement
		return ret
	}

	return o.Refinements
}

// GetRefinementsOk returns a tuple with the Refinements field value
// and a boolean to check if the value has been set.
func (o *PinnedRefinement) GetRefinementsOk() ([]Refinement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Refinements, true
}

// SetRefinements sets field value
func (o *PinnedRefinement) SetRefinements(v []Refinement) {
	o.Refinements = v
}

func (o PinnedRefinement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinnedRefinement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["navigation"] = o.Navigation
	toSerialize["refinements"] = o.Refinements
	return toSerialize, nil
}

type NullablePinnedRefinement struct {
	value *PinnedRefinement
	isSet bool
}

func (v NullablePinnedRefinement) Get() *PinnedRefinement {
	return v.value
}

func (v *NullablePinnedRefinement) Set(val *PinnedRefinement) {
	v.value = val
	v.isSet = true
}

func (v NullablePinnedRefinement) IsSet() bool {
	return v.isSet
}

func (v *NullablePinnedRefinement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinnedRefinement(val *PinnedRefinement) *NullablePinnedRefinement {
	return &NullablePinnedRefinement{value: val, isSet: true}
}

func (v NullablePinnedRefinement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinnedRefinement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


