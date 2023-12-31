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

// checks if the Facet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Facet{}

// Facet A facet specification to perform faceted search.
type Facet struct {
	// Only get facet values that start with the given string prefix. For example, suppose \"categories\" has three values \"Women > Shoe\", \"Women > Dress\" and \"Men > Shoe\". If set \"prefixes\" to \"Women\", the \"categories\" facet will give only \"Women > Shoe\" and \"Women > Dress\". Only supported on textual fields. Maximum is 10. This field is case-sensitive
	Prefix *string `json:"prefix,omitempty"`
	// Only get facet values that contains the given strings. For example, suppose \"categories\" has three values \"Women > Shoe\", \"Women > Dress\" and \"Men > Shoe\". If set \"contains\" to \"Shoe\", the \"categories\" facet will give only \"Women > Shoe\" and \"Men > Shoe\". Only supported on textual fields. Maximum is 10. This field is case-sensitive
	Contains *string `json:"contains,omitempty"`
	// Display name of facet
	DisplayName *string `json:"displayName,omitempty"`
	Type *NavigationType `json:"type,omitempty"`
	// Represents the name of navigation.
	NavigationName *string `json:"navigationName,omitempty"`
}

// NewFacet instantiates a new Facet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacet() *Facet {
	this := Facet{}
	return &this
}

// NewFacetWithDefaults instantiates a new Facet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetWithDefaults() *Facet {
	this := Facet{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *Facet) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facet) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *Facet) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *Facet) SetPrefix(v string) {
	o.Prefix = &v
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *Facet) GetContains() string {
	if o == nil || IsNil(o.Contains) {
		var ret string
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facet) GetContainsOk() (*string, bool) {
	if o == nil || IsNil(o.Contains) {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *Facet) HasContains() bool {
	if o != nil && !IsNil(o.Contains) {
		return true
	}

	return false
}

// SetContains gets a reference to the given string and assigns it to the Contains field.
func (o *Facet) SetContains(v string) {
	o.Contains = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Facet) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facet) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Facet) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Facet) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Facet) GetType() NavigationType {
	if o == nil || IsNil(o.Type) {
		var ret NavigationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facet) GetTypeOk() (*NavigationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Facet) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given NavigationType and assigns it to the Type field.
func (o *Facet) SetType(v NavigationType) {
	o.Type = &v
}

// GetNavigationName returns the NavigationName field value if set, zero value otherwise.
func (o *Facet) GetNavigationName() string {
	if o == nil || IsNil(o.NavigationName) {
		var ret string
		return ret
	}
	return *o.NavigationName
}

// GetNavigationNameOk returns a tuple with the NavigationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Facet) GetNavigationNameOk() (*string, bool) {
	if o == nil || IsNil(o.NavigationName) {
		return nil, false
	}
	return o.NavigationName, true
}

// HasNavigationName returns a boolean if a field has been set.
func (o *Facet) HasNavigationName() bool {
	if o != nil && !IsNil(o.NavigationName) {
		return true
	}

	return false
}

// SetNavigationName gets a reference to the given string and assigns it to the NavigationName field.
func (o *Facet) SetNavigationName(v string) {
	o.NavigationName = &v
}

func (o Facet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Facet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Contains) {
		toSerialize["contains"] = o.Contains
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NavigationName) {
		toSerialize["navigationName"] = o.NavigationName
	}
	return toSerialize, nil
}

type NullableFacet struct {
	value *Facet
	isSet bool
}

func (v NullableFacet) Get() *Facet {
	return v.value
}

func (v *NullableFacet) Set(val *Facet) {
	v.value = val
	v.isSet = true
}

func (v NullableFacet) IsSet() bool {
	return v.isSet
}

func (v *NullableFacet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacet(val *Facet) *NullableFacet {
	return &NullableFacet{value: val, isSet: true}
}

func (v NullableFacet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


