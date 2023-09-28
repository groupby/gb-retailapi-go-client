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

// checks if the TriggerSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TriggerSet{}

// TriggerSet struct for TriggerSet
type TriggerSet struct {
	// Query pattern triggers.
	QueryPatternTriggers []QueryPatternTrigger `json:"queryPatternTriggers"`
	// Selected refinement triggers.
	SelectedRefinementTriggers []SelectedRefinementTrigger `json:"selectedRefinementTriggers"`
	// Custom parameter triggers.
	CustomParameterTriggers []CustomParameterTrigger `json:"customParameterTriggers"`
}

// NewTriggerSet instantiates a new TriggerSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerSet(queryPatternTriggers []QueryPatternTrigger, selectedRefinementTriggers []SelectedRefinementTrigger, customParameterTriggers []CustomParameterTrigger) *TriggerSet {
	this := TriggerSet{}
	this.QueryPatternTriggers = queryPatternTriggers
	this.SelectedRefinementTriggers = selectedRefinementTriggers
	this.CustomParameterTriggers = customParameterTriggers
	return &this
}

// NewTriggerSetWithDefaults instantiates a new TriggerSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerSetWithDefaults() *TriggerSet {
	this := TriggerSet{}
	return &this
}

// GetQueryPatternTriggers returns the QueryPatternTriggers field value
func (o *TriggerSet) GetQueryPatternTriggers() []QueryPatternTrigger {
	if o == nil {
		var ret []QueryPatternTrigger
		return ret
	}

	return o.QueryPatternTriggers
}

// GetQueryPatternTriggersOk returns a tuple with the QueryPatternTriggers field value
// and a boolean to check if the value has been set.
func (o *TriggerSet) GetQueryPatternTriggersOk() ([]QueryPatternTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueryPatternTriggers, true
}

// SetQueryPatternTriggers sets field value
func (o *TriggerSet) SetQueryPatternTriggers(v []QueryPatternTrigger) {
	o.QueryPatternTriggers = v
}

// GetSelectedRefinementTriggers returns the SelectedRefinementTriggers field value
func (o *TriggerSet) GetSelectedRefinementTriggers() []SelectedRefinementTrigger {
	if o == nil {
		var ret []SelectedRefinementTrigger
		return ret
	}

	return o.SelectedRefinementTriggers
}

// GetSelectedRefinementTriggersOk returns a tuple with the SelectedRefinementTriggers field value
// and a boolean to check if the value has been set.
func (o *TriggerSet) GetSelectedRefinementTriggersOk() ([]SelectedRefinementTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedRefinementTriggers, true
}

// SetSelectedRefinementTriggers sets field value
func (o *TriggerSet) SetSelectedRefinementTriggers(v []SelectedRefinementTrigger) {
	o.SelectedRefinementTriggers = v
}

// GetCustomParameterTriggers returns the CustomParameterTriggers field value
func (o *TriggerSet) GetCustomParameterTriggers() []CustomParameterTrigger {
	if o == nil {
		var ret []CustomParameterTrigger
		return ret
	}

	return o.CustomParameterTriggers
}

// GetCustomParameterTriggersOk returns a tuple with the CustomParameterTriggers field value
// and a boolean to check if the value has been set.
func (o *TriggerSet) GetCustomParameterTriggersOk() ([]CustomParameterTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomParameterTriggers, true
}

// SetCustomParameterTriggers sets field value
func (o *TriggerSet) SetCustomParameterTriggers(v []CustomParameterTrigger) {
	o.CustomParameterTriggers = v
}

func (o TriggerSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["queryPatternTriggers"] = o.QueryPatternTriggers
	toSerialize["selectedRefinementTriggers"] = o.SelectedRefinementTriggers
	toSerialize["customParameterTriggers"] = o.CustomParameterTriggers
	return toSerialize, nil
}

type NullableTriggerSet struct {
	value *TriggerSet
	isSet bool
}

func (v NullableTriggerSet) Get() *TriggerSet {
	return v.value
}

func (v *NullableTriggerSet) Set(val *TriggerSet) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerSet) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerSet(val *TriggerSet) *NullableTriggerSet {
	return &NullableTriggerSet{value: val, isSet: true}
}

func (v NullableTriggerSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


