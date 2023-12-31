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

// checks if the ProductDtoAudience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductDtoAudience{}

// ProductDtoAudience struct for ProductDtoAudience
type ProductDtoAudience struct {
	// The genders of the audience. Strongly encouraged to use the standard values: 'male', 'female', 'unisex'. At most 5 values are allowed. Length limit of 128 characters.
	Genders []string `json:"genders,omitempty"`
	// The age groups of the audience. Strongly encouraged to use the standard values: 'newborn' (up to 3 months old), 'infant' (3-12 months old), 'toddler' (1-5 years old), 'kids' (5-13 years old), 'adult' (typically teens or older). At most 5 values are allowed. Length limit of 128 characters.
	AgeGroups []string `json:"ageGroups,omitempty"`
}

// NewProductDtoAudience instantiates a new ProductDtoAudience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductDtoAudience() *ProductDtoAudience {
	this := ProductDtoAudience{}
	return &this
}

// NewProductDtoAudienceWithDefaults instantiates a new ProductDtoAudience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductDtoAudienceWithDefaults() *ProductDtoAudience {
	this := ProductDtoAudience{}
	return &this
}

// GetGenders returns the Genders field value if set, zero value otherwise.
func (o *ProductDtoAudience) GetGenders() []string {
	if o == nil || IsNil(o.Genders) {
		var ret []string
		return ret
	}
	return o.Genders
}

// GetGendersOk returns a tuple with the Genders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDtoAudience) GetGendersOk() ([]string, bool) {
	if o == nil || IsNil(o.Genders) {
		return nil, false
	}
	return o.Genders, true
}

// HasGenders returns a boolean if a field has been set.
func (o *ProductDtoAudience) HasGenders() bool {
	if o != nil && !IsNil(o.Genders) {
		return true
	}

	return false
}

// SetGenders gets a reference to the given []string and assigns it to the Genders field.
func (o *ProductDtoAudience) SetGenders(v []string) {
	o.Genders = v
}

// GetAgeGroups returns the AgeGroups field value if set, zero value otherwise.
func (o *ProductDtoAudience) GetAgeGroups() []string {
	if o == nil || IsNil(o.AgeGroups) {
		var ret []string
		return ret
	}
	return o.AgeGroups
}

// GetAgeGroupsOk returns a tuple with the AgeGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductDtoAudience) GetAgeGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.AgeGroups) {
		return nil, false
	}
	return o.AgeGroups, true
}

// HasAgeGroups returns a boolean if a field has been set.
func (o *ProductDtoAudience) HasAgeGroups() bool {
	if o != nil && !IsNil(o.AgeGroups) {
		return true
	}

	return false
}

// SetAgeGroups gets a reference to the given []string and assigns it to the AgeGroups field.
func (o *ProductDtoAudience) SetAgeGroups(v []string) {
	o.AgeGroups = v
}

func (o ProductDtoAudience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductDtoAudience) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Genders) {
		toSerialize["genders"] = o.Genders
	}
	if !IsNil(o.AgeGroups) {
		toSerialize["ageGroups"] = o.AgeGroups
	}
	return toSerialize, nil
}

type NullableProductDtoAudience struct {
	value *ProductDtoAudience
	isSet bool
}

func (v NullableProductDtoAudience) Get() *ProductDtoAudience {
	return v.value
}

func (v *NullableProductDtoAudience) Set(val *ProductDtoAudience) {
	v.value = val
	v.isSet = true
}

func (v NullableProductDtoAudience) IsSet() bool {
	return v.isSet
}

func (v *NullableProductDtoAudience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductDtoAudience(val *ProductDtoAudience) *NullableProductDtoAudience {
	return &NullableProductDtoAudience{value: val, isSet: true}
}

func (v NullableProductDtoAudience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductDtoAudience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


