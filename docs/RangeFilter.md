# RangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field the range applies to. | 
**Range** | **interface{}** | Range of values the field value can be. | 

## Methods

### NewRangeFilter

`func NewRangeFilter(field string, range_ interface{}, ) *RangeFilter`

NewRangeFilter instantiates a new RangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeFilterWithDefaults

`func NewRangeFilterWithDefaults() *RangeFilter`

NewRangeFilterWithDefaults instantiates a new RangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *RangeFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RangeFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RangeFilter) SetField(v string)`

SetField sets Field field to given value.


### GetRange

`func (o *RangeFilter) GetRange() interface{}`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *RangeFilter) GetRangeOk() (*interface{}, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *RangeFilter) SetRange(v interface{})`

SetRange sets Range field to given value.


### SetRangeNil

`func (o *RangeFilter) SetRangeNil(b bool)`

 SetRangeNil sets the value for Range to be an explicit nil

### UnsetRange
`func (o *RangeFilter) UnsetRange()`

UnsetRange ensures that no value is present for Range, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


