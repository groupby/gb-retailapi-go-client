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

// checks if the Experiment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Experiment{}

// Experiment Information about Rule based Experiment.
type Experiment struct {
	// Experiment id.
	ExperimentId *string `json:"experimentId,omitempty"`
	// Experiment variant.
	ExperimentVariant *string `json:"experimentVariant,omitempty"`
}

// NewExperiment instantiates a new Experiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperiment() *Experiment {
	this := Experiment{}
	return &this
}

// NewExperimentWithDefaults instantiates a new Experiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentWithDefaults() *Experiment {
	this := Experiment{}
	return &this
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise.
func (o *Experiment) GetExperimentId() string {
	if o == nil || IsNil(o.ExperimentId) {
		var ret string
		return ret
	}
	return *o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Experiment) GetExperimentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentId) {
		return nil, false
	}
	return o.ExperimentId, true
}

// HasExperimentId returns a boolean if a field has been set.
func (o *Experiment) HasExperimentId() bool {
	if o != nil && !IsNil(o.ExperimentId) {
		return true
	}

	return false
}

// SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.
func (o *Experiment) SetExperimentId(v string) {
	o.ExperimentId = &v
}

// GetExperimentVariant returns the ExperimentVariant field value if set, zero value otherwise.
func (o *Experiment) GetExperimentVariant() string {
	if o == nil || IsNil(o.ExperimentVariant) {
		var ret string
		return ret
	}
	return *o.ExperimentVariant
}

// GetExperimentVariantOk returns a tuple with the ExperimentVariant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Experiment) GetExperimentVariantOk() (*string, bool) {
	if o == nil || IsNil(o.ExperimentVariant) {
		return nil, false
	}
	return o.ExperimentVariant, true
}

// HasExperimentVariant returns a boolean if a field has been set.
func (o *Experiment) HasExperimentVariant() bool {
	if o != nil && !IsNil(o.ExperimentVariant) {
		return true
	}

	return false
}

// SetExperimentVariant gets a reference to the given string and assigns it to the ExperimentVariant field.
func (o *Experiment) SetExperimentVariant(v string) {
	o.ExperimentVariant = &v
}

func (o Experiment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Experiment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExperimentId) {
		toSerialize["experimentId"] = o.ExperimentId
	}
	if !IsNil(o.ExperimentVariant) {
		toSerialize["experimentVariant"] = o.ExperimentVariant
	}
	return toSerialize, nil
}

type NullableExperiment struct {
	value *Experiment
	isSet bool
}

func (v NullableExperiment) Get() *Experiment {
	return v.value
}

func (v *NullableExperiment) Set(val *Experiment) {
	v.value = val
	v.isSet = true
}

func (v NullableExperiment) IsSet() bool {
	return v.isSet
}

func (v *NullableExperiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperiment(val *Experiment) *NullableExperiment {
	return &NullableExperiment{value: val, isSet: true}
}

func (v NullableExperiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


