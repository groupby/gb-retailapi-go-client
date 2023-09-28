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

// checks if the PageInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageInfoDto{}

// PageInfoDto Information regarding where the page of results starts and ends.
type PageInfoDto struct {
	// Where in the list of products the page begins.
	RecordStart *int64 `json:"recordStart,omitempty"`
	// Where in the list of products the page ends.
	RecordEnd *int64 `json:"recordEnd,omitempty"`
}

// NewPageInfoDto instantiates a new PageInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageInfoDto() *PageInfoDto {
	this := PageInfoDto{}
	return &this
}

// NewPageInfoDtoWithDefaults instantiates a new PageInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageInfoDtoWithDefaults() *PageInfoDto {
	this := PageInfoDto{}
	return &this
}

// GetRecordStart returns the RecordStart field value if set, zero value otherwise.
func (o *PageInfoDto) GetRecordStart() int64 {
	if o == nil || IsNil(o.RecordStart) {
		var ret int64
		return ret
	}
	return *o.RecordStart
}

// GetRecordStartOk returns a tuple with the RecordStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageInfoDto) GetRecordStartOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordStart) {
		return nil, false
	}
	return o.RecordStart, true
}

// HasRecordStart returns a boolean if a field has been set.
func (o *PageInfoDto) HasRecordStart() bool {
	if o != nil && !IsNil(o.RecordStart) {
		return true
	}

	return false
}

// SetRecordStart gets a reference to the given int64 and assigns it to the RecordStart field.
func (o *PageInfoDto) SetRecordStart(v int64) {
	o.RecordStart = &v
}

// GetRecordEnd returns the RecordEnd field value if set, zero value otherwise.
func (o *PageInfoDto) GetRecordEnd() int64 {
	if o == nil || IsNil(o.RecordEnd) {
		var ret int64
		return ret
	}
	return *o.RecordEnd
}

// GetRecordEndOk returns a tuple with the RecordEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageInfoDto) GetRecordEndOk() (*int64, bool) {
	if o == nil || IsNil(o.RecordEnd) {
		return nil, false
	}
	return o.RecordEnd, true
}

// HasRecordEnd returns a boolean if a field has been set.
func (o *PageInfoDto) HasRecordEnd() bool {
	if o != nil && !IsNil(o.RecordEnd) {
		return true
	}

	return false
}

// SetRecordEnd gets a reference to the given int64 and assigns it to the RecordEnd field.
func (o *PageInfoDto) SetRecordEnd(v int64) {
	o.RecordEnd = &v
}

func (o PageInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordStart) {
		toSerialize["recordStart"] = o.RecordStart
	}
	if !IsNil(o.RecordEnd) {
		toSerialize["recordEnd"] = o.RecordEnd
	}
	return toSerialize, nil
}

type NullablePageInfoDto struct {
	value *PageInfoDto
	isSet bool
}

func (v NullablePageInfoDto) Get() *PageInfoDto {
	return v.value
}

func (v *NullablePageInfoDto) Set(val *PageInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePageInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePageInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageInfoDto(val *PageInfoDto) *NullablePageInfoDto {
	return &NullablePageInfoDto{value: val, isSet: true}
}

func (v NullablePageInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


