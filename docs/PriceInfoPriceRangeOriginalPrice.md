# PriceInfoPriceRangeOriginalPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minimum** | Pointer to **float64** | Inclusive lower bound. The lower bound of the interval. If neither of the min fields (minimum or exclusiveMinimum) are set, then the lower bound is negative infinity. This field must be not larger than maximum or exclusiveMaximum. | [optional] 
**ExclusiveMinimum** | Pointer to **float64** | Exclusive lower bound. The lower bound of the interval. If neither of the min fields (minimum or exclusiveMinimum) are set, then the lower bound is negative infinity. This field must be not larger than maximum or exclusiveMaximum. | [optional] 
**Maximum** | Pointer to **float64** | Inclusive upper bound. The upper bound of the interval. If neither of the max fields are set, then the upper bound is positive infinity. This field must be not smaller than minimum or exclusiveMinimum. | [optional] 
**ExclusiveMaximum** | Pointer to **float64** | Exclusive upper bound. The upper bound of the interval. If neither of the max fields are set, then the upper bound is positive infinity. This field must be not smaller than minimum or exclusiveMinimum. | [optional] 

## Methods

### NewPriceInfoPriceRangeOriginalPrice

`func NewPriceInfoPriceRangeOriginalPrice() *PriceInfoPriceRangeOriginalPrice`

NewPriceInfoPriceRangeOriginalPrice instantiates a new PriceInfoPriceRangeOriginalPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceInfoPriceRangeOriginalPriceWithDefaults

`func NewPriceInfoPriceRangeOriginalPriceWithDefaults() *PriceInfoPriceRangeOriginalPrice`

NewPriceInfoPriceRangeOriginalPriceWithDefaults instantiates a new PriceInfoPriceRangeOriginalPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) GetMinimum() float64`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *PriceInfoPriceRangeOriginalPrice) GetMinimumOk() (*float64, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) SetMinimum(v float64)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetExclusiveMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) GetExclusiveMinimum() float64`

GetExclusiveMinimum returns the ExclusiveMinimum field if non-nil, zero value otherwise.

### GetExclusiveMinimumOk

`func (o *PriceInfoPriceRangeOriginalPrice) GetExclusiveMinimumOk() (*float64, bool)`

GetExclusiveMinimumOk returns a tuple with the ExclusiveMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) SetExclusiveMinimum(v float64)`

SetExclusiveMinimum sets ExclusiveMinimum field to given value.

### HasExclusiveMinimum

`func (o *PriceInfoPriceRangeOriginalPrice) HasExclusiveMinimum() bool`

HasExclusiveMinimum returns a boolean if a field has been set.

### GetMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) GetMaximum() float64`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *PriceInfoPriceRangeOriginalPrice) GetMaximumOk() (*float64, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) SetMaximum(v float64)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetExclusiveMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) GetExclusiveMaximum() float64`

GetExclusiveMaximum returns the ExclusiveMaximum field if non-nil, zero value otherwise.

### GetExclusiveMaximumOk

`func (o *PriceInfoPriceRangeOriginalPrice) GetExclusiveMaximumOk() (*float64, bool)`

GetExclusiveMaximumOk returns a tuple with the ExclusiveMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusiveMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) SetExclusiveMaximum(v float64)`

SetExclusiveMaximum sets ExclusiveMaximum field to given value.

### HasExclusiveMaximum

`func (o *PriceInfoPriceRangeOriginalPrice) HasExclusiveMaximum() bool`

HasExclusiveMaximum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


