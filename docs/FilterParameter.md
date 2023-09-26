# FilterParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Value** | Pointer to **NullableString** |  | [optional] 
**Exclude** | **bool** |  | 
**DerivedFromProduct** | **bool** |  | 

## Methods

### NewFilterParameter

`func NewFilterParameter(field string, exclude bool, derivedFromProduct bool, ) *FilterParameter`

NewFilterParameter instantiates a new FilterParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterParameterWithDefaults

`func NewFilterParameterWithDefaults() *FilterParameter`

NewFilterParameterWithDefaults instantiates a new FilterParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterParameter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterParameter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterParameter) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *FilterParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FilterParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *FilterParameter) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FilterParameter) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetExclude

`func (o *FilterParameter) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *FilterParameter) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *FilterParameter) SetExclude(v bool)`

SetExclude sets Exclude field to given value.


### GetDerivedFromProduct

`func (o *FilterParameter) GetDerivedFromProduct() bool`

GetDerivedFromProduct returns the DerivedFromProduct field if non-nil, zero value otherwise.

### GetDerivedFromProductOk

`func (o *FilterParameter) GetDerivedFromProductOk() (*bool, bool)`

GetDerivedFromProductOk returns a tuple with the DerivedFromProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromProduct

`func (o *FilterParameter) SetDerivedFromProduct(v bool)`

SetDerivedFromProduct sets DerivedFromProduct field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


