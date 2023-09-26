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

// checks if the FacetSearchRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FacetSearchRequestDto{}

// FacetSearchRequestDto Request that should be populated to configure a search API call, made by the client on behalf of a shopper. Contains original request and information about facet for which extra keys requested.
type FacetSearchRequestDto struct {
	Facet Facet `json:"facet"`
	OriginalRequest SearchRequestDto `json:"originalRequest"`
}

// NewFacetSearchRequestDto instantiates a new FacetSearchRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetSearchRequestDto(facet Facet, originalRequest SearchRequestDto) *FacetSearchRequestDto {
	this := FacetSearchRequestDto{}
	this.Facet = facet
	this.OriginalRequest = originalRequest
	return &this
}

// NewFacetSearchRequestDtoWithDefaults instantiates a new FacetSearchRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetSearchRequestDtoWithDefaults() *FacetSearchRequestDto {
	this := FacetSearchRequestDto{}
	return &this
}

// GetFacet returns the Facet field value
func (o *FacetSearchRequestDto) GetFacet() Facet {
	if o == nil {
		var ret Facet
		return ret
	}

	return o.Facet
}

// GetFacetOk returns a tuple with the Facet field value
// and a boolean to check if the value has been set.
func (o *FacetSearchRequestDto) GetFacetOk() (*Facet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Facet, true
}

// SetFacet sets field value
func (o *FacetSearchRequestDto) SetFacet(v Facet) {
	o.Facet = v
}

// GetOriginalRequest returns the OriginalRequest field value
func (o *FacetSearchRequestDto) GetOriginalRequest() SearchRequestDto {
	if o == nil {
		var ret SearchRequestDto
		return ret
	}

	return o.OriginalRequest
}

// GetOriginalRequestOk returns a tuple with the OriginalRequest field value
// and a boolean to check if the value has been set.
func (o *FacetSearchRequestDto) GetOriginalRequestOk() (*SearchRequestDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalRequest, true
}

// SetOriginalRequest sets field value
func (o *FacetSearchRequestDto) SetOriginalRequest(v SearchRequestDto) {
	o.OriginalRequest = v
}

func (o FacetSearchRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FacetSearchRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["facet"] = o.Facet
	toSerialize["originalRequest"] = o.OriginalRequest
	return toSerialize, nil
}

type NullableFacetSearchRequestDto struct {
	value *FacetSearchRequestDto
	isSet bool
}

func (v NullableFacetSearchRequestDto) Get() *FacetSearchRequestDto {
	return v.value
}

func (v *NullableFacetSearchRequestDto) Set(val *FacetSearchRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetSearchRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetSearchRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetSearchRequestDto(val *FacetSearchRequestDto) *NullableFacetSearchRequestDto {
	return &NullableFacetSearchRequestDto{value: val, isSet: true}
}

func (v NullableFacetSearchRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetSearchRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

