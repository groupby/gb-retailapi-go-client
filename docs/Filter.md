# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Exclude** | **bool** |  | 
**DerivedFromProduct** | **bool** |  | 

## Methods

### NewFilter

`func NewFilter(field string, exclude bool, derivedFromProduct bool, ) *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *Filter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Filter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Filter) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *Filter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Filter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Filter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Filter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *Filter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Filter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetExclude

`func (o *Filter) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *Filter) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *Filter) SetExclude(v bool)`

SetExclude sets Exclude field to given value.


### GetDerivedFromProduct

`func (o *Filter) GetDerivedFromProduct() bool`

GetDerivedFromProduct returns the DerivedFromProduct field if non-nil, zero value otherwise.

### GetDerivedFromProductOk

`func (o *Filter) GetDerivedFromProductOk() (*bool, bool)`

GetDerivedFromProductOk returns a tuple with the DerivedFromProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromProduct

`func (o *Filter) SetDerivedFromProduct(v bool)`

SetDerivedFromProduct sets DerivedFromProduct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


