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

// checks if the ExperimentVariant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExperimentVariant{}

// ExperimentVariant struct for ExperimentVariant
type ExperimentVariant struct {
	Name string `json:"name"`
	RuleVariant RuleVariant `json:"ruleVariant"`
	VariantTriggerPercentage int32 `json:"variantTriggerPercentage"`
}

// NewExperimentVariant instantiates a new ExperimentVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperimentVariant(name string, ruleVariant RuleVariant, variantTriggerPercentage int32) *ExperimentVariant {
	this := ExperimentVariant{}
	this.Name = name
	this.RuleVariant = ruleVariant
	this.VariantTriggerPercentage = variantTriggerPercentage
	return &this
}

// NewExperimentVariantWithDefaults instantiates a new ExperimentVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperimentVariantWithDefaults() *ExperimentVariant {
	this := ExperimentVariant{}
	return &this
}

// GetName returns the Name field value
func (o *ExperimentVariant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExperimentVariant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExperimentVariant) SetName(v string) {
	o.Name = v
}

// GetRuleVariant returns the RuleVariant field value
func (o *ExperimentVariant) GetRuleVariant() RuleVariant {
	if o == nil {
		var ret RuleVariant
		return ret
	}

	return o.RuleVariant
}

// GetRuleVariantOk returns a tuple with the RuleVariant field value
// and a boolean to check if the value has been set.
func (o *ExperimentVariant) GetRuleVariantOk() (*RuleVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleVariant, true
}

// SetRuleVariant sets field value
func (o *ExperimentVariant) SetRuleVariant(v RuleVariant) {
	o.RuleVariant = v
}

// GetVariantTriggerPercentage returns the VariantTriggerPercentage field value
func (o *ExperimentVariant) GetVariantTriggerPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VariantTriggerPercentage
}

// GetVariantTriggerPercentageOk returns a tuple with the VariantTriggerPercentage field value
// and a boolean to check if the value has been set.
func (o *ExperimentVariant) GetVariantTriggerPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariantTriggerPercentage, true
}

// SetVariantTriggerPercentage sets field value
func (o *ExperimentVariant) SetVariantTriggerPercentage(v int32) {
	o.VariantTriggerPercentage = v
}

func (o ExperimentVariant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExperimentVariant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["ruleVariant"] = o.RuleVariant
	toSerialize["variantTriggerPercentage"] = o.VariantTriggerPercentage
	return toSerialize, nil
}

type NullableExperimentVariant struct {
	value *ExperimentVariant
	isSet bool
}

func (v NullableExperimentVariant) Get() *ExperimentVariant {
	return v.value
}

func (v *NullableExperimentVariant) Set(val *ExperimentVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableExperimentVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableExperimentVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperimentVariant(val *ExperimentVariant) *NullableExperimentVariant {
	return &NullableExperimentVariant{value: val, isSet: true}
}

func (v NullableExperimentVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperimentVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


