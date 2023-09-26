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

// checks if the RuleTemplateSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RuleTemplateSection{}

// RuleTemplateSection struct for RuleTemplateSection
type RuleTemplateSection struct {
	ZoneId int32 `json:"zoneId"`
	Name string `json:"name"`
	ZoneContent string `json:"zoneContent"`
	ZoneType ZoneType `json:"zoneType"`
}

// NewRuleTemplateSection instantiates a new RuleTemplateSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleTemplateSection(zoneId int32, name string, zoneContent string, zoneType ZoneType) *RuleTemplateSection {
	this := RuleTemplateSection{}
	this.ZoneId = zoneId
	this.Name = name
	this.ZoneContent = zoneContent
	this.ZoneType = zoneType
	return &this
}

// NewRuleTemplateSectionWithDefaults instantiates a new RuleTemplateSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleTemplateSectionWithDefaults() *RuleTemplateSection {
	this := RuleTemplateSection{}
	return &this
}

// GetZoneId returns the ZoneId field value
func (o *RuleTemplateSection) GetZoneId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value
// and a boolean to check if the value has been set.
func (o *RuleTemplateSection) GetZoneIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneId, true
}

// SetZoneId sets field value
func (o *RuleTemplateSection) SetZoneId(v int32) {
	o.ZoneId = v
}

// GetName returns the Name field value
func (o *RuleTemplateSection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RuleTemplateSection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RuleTemplateSection) SetName(v string) {
	o.Name = v
}

// GetZoneContent returns the ZoneContent field value
func (o *RuleTemplateSection) GetZoneContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ZoneContent
}

// GetZoneContentOk returns a tuple with the ZoneContent field value
// and a boolean to check if the value has been set.
func (o *RuleTemplateSection) GetZoneContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneContent, true
}

// SetZoneContent sets field value
func (o *RuleTemplateSection) SetZoneContent(v string) {
	o.ZoneContent = v
}

// GetZoneType returns the ZoneType field value
func (o *RuleTemplateSection) GetZoneType() ZoneType {
	if o == nil {
		var ret ZoneType
		return ret
	}

	return o.ZoneType
}

// GetZoneTypeOk returns a tuple with the ZoneType field value
// and a boolean to check if the value has been set.
func (o *RuleTemplateSection) GetZoneTypeOk() (*ZoneType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZoneType, true
}

// SetZoneType sets field value
func (o *RuleTemplateSection) SetZoneType(v ZoneType) {
	o.ZoneType = v
}

func (o RuleTemplateSection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleTemplateSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["zoneId"] = o.ZoneId
	toSerialize["name"] = o.Name
	toSerialize["zoneContent"] = o.ZoneContent
	toSerialize["zoneType"] = o.ZoneType
	return toSerialize, nil
}

type NullableRuleTemplateSection struct {
	value *RuleTemplateSection
	isSet bool
}

func (v NullableRuleTemplateSection) Get() *RuleTemplateSection {
	return v.value
}

func (v *NullableRuleTemplateSection) Set(val *RuleTemplateSection) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleTemplateSection) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleTemplateSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleTemplateSection(val *RuleTemplateSection) *NullableRuleTemplateSection {
	return &NullableRuleTemplateSection{value: val, isSet: true}
}

func (v NullableRuleTemplateSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleTemplateSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

