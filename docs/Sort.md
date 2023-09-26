# Sort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field the order will be applied to. | 
**Order** | **interface{}** | OrderDto the products will appear in. | 

## Methods

### NewSort

`func NewSort(field string, order interface{}, ) *Sort`

NewSort instantiates a new Sort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortWithDefaults

`func NewSortWithDefaults() *Sort`

NewSortWithDefaults instantiates a new Sort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *Sort) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Sort) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Sort) SetField(v string)`

SetField sets Field field to given value.


### GetOrder

`func (o *Sort) GetOrder() interface{}`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Sort) GetOrderOk() (*interface{}, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Sort) SetOrder(v interface{})`

SetOrder sets Order field to given value.


### SetOrderNil

`func (o *Sort) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *Sort) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


