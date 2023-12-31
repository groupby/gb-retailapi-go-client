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

// checks if the SearchMetadataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMetadataDto{}

// SearchMetadataDto Metadata relating to the search results, or how they were generated.
type SearchMetadataDto struct {
	// Token to assist beacon collectors in correlating searches to user events.
	SearchAttributionToken NullableString `json:"searchAttributionToken,omitempty"`
	// Were the search results from a previous call.
	Cached *bool `json:"cached,omitempty"`
	// Total time spent performing the Google search in milliseconds.
	TotalTime *int64 `json:"totalTime,omitempty"`
}

// NewSearchMetadataDto instantiates a new SearchMetadataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMetadataDto() *SearchMetadataDto {
	this := SearchMetadataDto{}
	return &this
}

// NewSearchMetadataDtoWithDefaults instantiates a new SearchMetadataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMetadataDtoWithDefaults() *SearchMetadataDto {
	this := SearchMetadataDto{}
	return &this
}

// GetSearchAttributionToken returns the SearchAttributionToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchMetadataDto) GetSearchAttributionToken() string {
	if o == nil || IsNil(o.SearchAttributionToken.Get()) {
		var ret string
		return ret
	}
	return *o.SearchAttributionToken.Get()
}

// GetSearchAttributionTokenOk returns a tuple with the SearchAttributionToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchMetadataDto) GetSearchAttributionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchAttributionToken.Get(), o.SearchAttributionToken.IsSet()
}

// HasSearchAttributionToken returns a boolean if a field has been set.
func (o *SearchMetadataDto) HasSearchAttributionToken() bool {
	if o != nil && o.SearchAttributionToken.IsSet() {
		return true
	}

	return false
}

// SetSearchAttributionToken gets a reference to the given NullableString and assigns it to the SearchAttributionToken field.
func (o *SearchMetadataDto) SetSearchAttributionToken(v string) {
	o.SearchAttributionToken.Set(&v)
}
// SetSearchAttributionTokenNil sets the value for SearchAttributionToken to be an explicit nil
func (o *SearchMetadataDto) SetSearchAttributionTokenNil() {
	o.SearchAttributionToken.Set(nil)
}

// UnsetSearchAttributionToken ensures that no value is present for SearchAttributionToken, not even an explicit nil
func (o *SearchMetadataDto) UnsetSearchAttributionToken() {
	o.SearchAttributionToken.Unset()
}

// GetCached returns the Cached field value if set, zero value otherwise.
func (o *SearchMetadataDto) GetCached() bool {
	if o == nil || IsNil(o.Cached) {
		var ret bool
		return ret
	}
	return *o.Cached
}

// GetCachedOk returns a tuple with the Cached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMetadataDto) GetCachedOk() (*bool, bool) {
	if o == nil || IsNil(o.Cached) {
		return nil, false
	}
	return o.Cached, true
}

// HasCached returns a boolean if a field has been set.
func (o *SearchMetadataDto) HasCached() bool {
	if o != nil && !IsNil(o.Cached) {
		return true
	}

	return false
}

// SetCached gets a reference to the given bool and assigns it to the Cached field.
func (o *SearchMetadataDto) SetCached(v bool) {
	o.Cached = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *SearchMetadataDto) GetTotalTime() int64 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int64
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMetadataDto) GetTotalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *SearchMetadataDto) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int64 and assigns it to the TotalTime field.
func (o *SearchMetadataDto) SetTotalTime(v int64) {
	o.TotalTime = &v
}

func (o SearchMetadataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMetadataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchAttributionToken.IsSet() {
		toSerialize["searchAttributionToken"] = o.SearchAttributionToken.Get()
	}
	if !IsNil(o.Cached) {
		toSerialize["cached"] = o.Cached
	}
	if !IsNil(o.TotalTime) {
		toSerialize["totalTime"] = o.TotalTime
	}
	return toSerialize, nil
}

type NullableSearchMetadataDto struct {
	value *SearchMetadataDto
	isSet bool
}

func (v NullableSearchMetadataDto) Get() *SearchMetadataDto {
	return v.value
}

func (v *NullableSearchMetadataDto) Set(val *SearchMetadataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMetadataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMetadataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMetadataDto(val *SearchMetadataDto) *NullableSearchMetadataDto {
	return &NullableSearchMetadataDto{value: val, isSet: true}
}

func (v NullableSearchMetadataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMetadataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


