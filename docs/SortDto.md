# SortDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field the order will be applied to. | 
**Order** | [**NullableSortDtoOrderDto**](SortDtoOrderDto.md) |  | 

## Methods

### NewSortDto

`func NewSortDto(field string, order NullableSortDtoOrderDto, ) *SortDto`

NewSortDto instantiates a new SortDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortDtoWithDefaults

`func NewSortDtoWithDefaults() *SortDto`

NewSortDtoWithDefaults instantiates a new SortDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *SortDto) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *SortDto) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *SortDto) SetField(v string)`

SetField sets Field field to given value.


### GetOrder

`func (o *SortDto) GetOrder() SortDtoOrderDto`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *SortDto) GetOrderOk() (*SortDtoOrderDto, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *SortDto) SetOrder(v SortDtoOrderDto)`

SetOrder sets Order field to given value.


### SetOrderNil

`func (o *SortDto) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *SortDto) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


