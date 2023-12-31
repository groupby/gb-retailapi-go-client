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

// checks if the Merchandiser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Merchandiser{}

// Merchandiser struct for Merchandiser
type Merchandiser struct {
	MerchandiserId string `json:"merchandiserId"`
}

// NewMerchandiser instantiates a new Merchandiser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchandiser(merchandiserId string) *Merchandiser {
	this := Merchandiser{}
	this.MerchandiserId = merchandiserId
	return &this
}

// NewMerchandiserWithDefaults instantiates a new Merchandiser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchandiserWithDefaults() *Merchandiser {
	this := Merchandiser{}
	return &this
}

// GetMerchandiserId returns the MerchandiserId field value
func (o *Merchandiser) GetMerchandiserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchandiserId
}

// GetMerchandiserIdOk returns a tuple with the MerchandiserId field value
// and a boolean to check if the value has been set.
func (o *Merchandiser) GetMerchandiserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchandiserId, true
}

// SetMerchandiserId sets field value
func (o *Merchandiser) SetMerchandiserId(v string) {
	o.MerchandiserId = v
}

func (o Merchandiser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Merchandiser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merchandiserId"] = o.MerchandiserId
	return toSerialize, nil
}

type NullableMerchandiser struct {
	value *Merchandiser
	isSet bool
}

func (v NullableMerchandiser) Get() *Merchandiser {
	return v.value
}

func (v *NullableMerchandiser) Set(val *Merchandiser) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchandiser) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchandiser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchandiser(val *Merchandiser) *NullableMerchandiser {
	return &NullableMerchandiser{value: val, isSet: true}
}

func (v NullableMerchandiser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchandiser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


