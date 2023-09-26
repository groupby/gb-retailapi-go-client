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

// checks if the DebugDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebugDto{}

// DebugDto Contains debug info associated to response.
type DebugDto struct {
	// Search request in raw format executed internally against search engine.
	RawSearchRequest []map[string]interface{} `json:"rawSearchRequest,omitempty"`
	// Search response in raw format received internally from search engine.
	RawSearchResponse []map[string]interface{} `json:"rawSearchResponse,omitempty"`
}

// NewDebugDto instantiates a new DebugDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebugDto() *DebugDto {
	this := DebugDto{}
	return &this
}

// NewDebugDtoWithDefaults instantiates a new DebugDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebugDtoWithDefaults() *DebugDto {
	this := DebugDto{}
	return &this
}

// GetRawSearchRequest returns the RawSearchRequest field value if set, zero value otherwise.
func (o *DebugDto) GetRawSearchRequest() []map[string]interface{} {
	if o == nil || IsNil(o.RawSearchRequest) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RawSearchRequest
}

// GetRawSearchRequestOk returns a tuple with the RawSearchRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebugDto) GetRawSearchRequestOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.RawSearchRequest) {
		return nil, false
	}
	return o.RawSearchRequest, true
}

// HasRawSearchRequest returns a boolean if a field has been set.
func (o *DebugDto) HasRawSearchRequest() bool {
	if o != nil && !IsNil(o.RawSearchRequest) {
		return true
	}

	return false
}

// SetRawSearchRequest gets a reference to the given []map[string]interface{} and assigns it to the RawSearchRequest field.
func (o *DebugDto) SetRawSearchRequest(v []map[string]interface{}) {
	o.RawSearchRequest = v
}

// GetRawSearchResponse returns the RawSearchResponse field value if set, zero value otherwise.
func (o *DebugDto) GetRawSearchResponse() []map[string]interface{} {
	if o == nil || IsNil(o.RawSearchResponse) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RawSearchResponse
}

// GetRawSearchResponseOk returns a tuple with the RawSearchResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebugDto) GetRawSearchResponseOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.RawSearchResponse) {
		return nil, false
	}
	return o.RawSearchResponse, true
}

// HasRawSearchResponse returns a boolean if a field has been set.
func (o *DebugDto) HasRawSearchResponse() bool {
	if o != nil && !IsNil(o.RawSearchResponse) {
		return true
	}

	return false
}

// SetRawSearchResponse gets a reference to the given []map[string]interface{} and assigns it to the RawSearchResponse field.
func (o *DebugDto) SetRawSearchResponse(v []map[string]interface{}) {
	o.RawSearchResponse = v
}

func (o DebugDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebugDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RawSearchRequest) {
		toSerialize["rawSearchRequest"] = o.RawSearchRequest
	}
	if !IsNil(o.RawSearchResponse) {
		toSerialize["rawSearchResponse"] = o.RawSearchResponse
	}
	return toSerialize, nil
}

type NullableDebugDto struct {
	value *DebugDto
	isSet bool
}

func (v NullableDebugDto) Get() *DebugDto {
	return v.value
}

func (v *NullableDebugDto) Set(val *DebugDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDebugDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDebugDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebugDto(val *DebugDto) *NullableDebugDto {
	return &NullableDebugDto{value: val, isSet: true}
}

func (v NullableDebugDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebugDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


