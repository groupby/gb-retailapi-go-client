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

// checks if the Filter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Filter{}

// Filter struct for Filter
type Filter struct {
	Field string `json:"field"`
	Value NullableString `json:"value,omitempty"`
	Exclude bool `json:"exclude"`
	DerivedFromProduct bool `json:"derivedFromProduct"`
}

// NewFilter instantiates a new Filter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilter(field string, exclude bool, derivedFromProduct bool) *Filter {
	this := Filter{}
	this.Field = field
	this.Exclude = exclude
	this.DerivedFromProduct = derivedFromProduct
	return &this
}

// NewFilterWithDefaults instantiates a new Filter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterWithDefaults() *Filter {
	this := Filter{}
	return &this
}

// GetField returns the Field field value
func (o *Filter) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *Filter) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *Filter) SetField(v string) {
	o.Field = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Filter) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Filter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *Filter) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *Filter) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *Filter) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *Filter) UnsetValue() {
	o.Value.Unset()
}

// GetExclude returns the Exclude field value
func (o *Filter) GetExclude() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value
// and a boolean to check if the value has been set.
func (o *Filter) GetExcludeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exclude, true
}

// SetExclude sets field value
func (o *Filter) SetExclude(v bool) {
	o.Exclude = v
}

// GetDerivedFromProduct returns the DerivedFromProduct field value
func (o *Filter) GetDerivedFromProduct() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DerivedFromProduct
}

// GetDerivedFromProductOk returns a tuple with the DerivedFromProduct field value
// and a boolean to check if the value has been set.
func (o *Filter) GetDerivedFromProductOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DerivedFromProduct, true
}

// SetDerivedFromProduct sets field value
func (o *Filter) SetDerivedFromProduct(v bool) {
	o.DerivedFromProduct = v
}

func (o Filter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	if o.Value.IsSet() {
		toSerialize["value"] = o.Value.Get()
	}
	toSerialize["exclude"] = o.Exclude
	toSerialize["derivedFromProduct"] = o.DerivedFromProduct
	return toSerialize, nil
}

type NullableFilter struct {
	value *Filter
	isSet bool
}

func (v NullableFilter) Get() *Filter {
	return v.value
}

func (v *NullableFilter) Set(val *Filter) {
	v.value = val
	v.isSet = true
}

func (v NullableFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilter(val *Filter) *NullableFilter {
	return &NullableFilter{value: val, isSet: true}
}

func (v NullableFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


