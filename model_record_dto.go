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

// checks if the RecordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordDto{}

// RecordDto Information regarding a product in the proprietary Group By API format.
type RecordDto struct {
	// Identifier of the record as an MD5 hash.
	Id *string `json:"_id,omitempty"`
	// Logical URL of the record.
	U *string `json:"_u,omitempty"`
	// Title of the record.
	T *string `json:"_t,omitempty"`
	// Collection the record was queried from or resides.
	Collection *string `json:"collection,omitempty"`
	// All other metadata, read fields, the record has.
	AllMeta map[string]interface{} `json:"allMeta,omitempty"`
}

// NewRecordDto instantiates a new RecordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordDto() *RecordDto {
	this := RecordDto{}
	return &this
}

// NewRecordDtoWithDefaults instantiates a new RecordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordDtoWithDefaults() *RecordDto {
	this := RecordDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecordDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecordDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecordDto) SetId(v string) {
	o.Id = &v
}

// GetU returns the U field value if set, zero value otherwise.
func (o *RecordDto) GetU() string {
	if o == nil || IsNil(o.U) {
		var ret string
		return ret
	}
	return *o.U
}

// GetUOk returns a tuple with the U field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordDto) GetUOk() (*string, bool) {
	if o == nil || IsNil(o.U) {
		return nil, false
	}
	return o.U, true
}

// HasU returns a boolean if a field has been set.
func (o *RecordDto) HasU() bool {
	if o != nil && !IsNil(o.U) {
		return true
	}

	return false
}

// SetU gets a reference to the given string and assigns it to the U field.
func (o *RecordDto) SetU(v string) {
	o.U = &v
}

// GetT returns the T field value if set, zero value otherwise.
func (o *RecordDto) GetT() string {
	if o == nil || IsNil(o.T) {
		var ret string
		return ret
	}
	return *o.T
}

// GetTOk returns a tuple with the T field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordDto) GetTOk() (*string, bool) {
	if o == nil || IsNil(o.T) {
		return nil, false
	}
	return o.T, true
}

// HasT returns a boolean if a field has been set.
func (o *RecordDto) HasT() bool {
	if o != nil && !IsNil(o.T) {
		return true
	}

	return false
}

// SetT gets a reference to the given string and assigns it to the T field.
func (o *RecordDto) SetT(v string) {
	o.T = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *RecordDto) GetCollection() string {
	if o == nil || IsNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordDto) GetCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *RecordDto) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *RecordDto) SetCollection(v string) {
	o.Collection = &v
}

// GetAllMeta returns the AllMeta field value if set, zero value otherwise.
func (o *RecordDto) GetAllMeta() map[string]interface{} {
	if o == nil || IsNil(o.AllMeta) {
		var ret map[string]interface{}
		return ret
	}
	return o.AllMeta
}

// GetAllMetaOk returns a tuple with the AllMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordDto) GetAllMetaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AllMeta) {
		return map[string]interface{}{}, false
	}
	return o.AllMeta, true
}

// HasAllMeta returns a boolean if a field has been set.
func (o *RecordDto) HasAllMeta() bool {
	if o != nil && !IsNil(o.AllMeta) {
		return true
	}

	return false
}

// SetAllMeta gets a reference to the given map[string]interface{} and assigns it to the AllMeta field.
func (o *RecordDto) SetAllMeta(v map[string]interface{}) {
	o.AllMeta = v
}

func (o RecordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.U) {
		toSerialize["_u"] = o.U
	}
	if !IsNil(o.T) {
		toSerialize["_t"] = o.T
	}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.AllMeta) {
		toSerialize["allMeta"] = o.AllMeta
	}
	return toSerialize, nil
}

type NullableRecordDto struct {
	value *RecordDto
	isSet bool
}

func (v NullableRecordDto) Get() *RecordDto {
	return v.value
}

func (v *NullableRecordDto) Set(val *RecordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordDto(val *RecordDto) *NullableRecordDto {
	return &NullableRecordDto{value: val, isSet: true}
}

func (v NullableRecordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


