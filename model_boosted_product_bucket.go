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

// checks if the BoostedProductBucket type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoostedProductBucket{}

// BoostedProductBucket struct for BoostedProductBucket
type BoostedProductBucket struct {
	Products []string `json:"products"`
}

// NewBoostedProductBucket instantiates a new BoostedProductBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoostedProductBucket(products []string) *BoostedProductBucket {
	this := BoostedProductBucket{}
	this.Products = products
	return &this
}

// NewBoostedProductBucketWithDefaults instantiates a new BoostedProductBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoostedProductBucketWithDefaults() *BoostedProductBucket {
	this := BoostedProductBucket{}
	return &this
}

// GetProducts returns the Products field value
func (o *BoostedProductBucket) GetProducts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *BoostedProductBucket) GetProductsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *BoostedProductBucket) SetProducts(v []string) {
	o.Products = v
}

func (o BoostedProductBucket) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoostedProductBucket) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["products"] = o.Products
	return toSerialize, nil
}

type NullableBoostedProductBucket struct {
	value *BoostedProductBucket
	isSet bool
}

func (v NullableBoostedProductBucket) Get() *BoostedProductBucket {
	return v.value
}

func (v *NullableBoostedProductBucket) Set(val *BoostedProductBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableBoostedProductBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableBoostedProductBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoostedProductBucket(val *BoostedProductBucket) *NullableBoostedProductBucket {
	return &NullableBoostedProductBucket{value: val, isSet: true}
}

func (v NullableBoostedProductBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoostedProductBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


