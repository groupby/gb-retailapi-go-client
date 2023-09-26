# PriceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | Currency code. | [optional] 
**Price** | Pointer to **float32** | Price value. | [optional] 
**OriginalPrice** | Pointer to **float32** | Original price value. | [optional] 
**Cost** | Pointer to **float32** | Cost | [optional] 
**PriceEffectiveTime** | Pointer to [**PriceInfoPriceEffectiveTime**](PriceInfoPriceEffectiveTime.md) |  | [optional] 
**PriceExpireTime** | Pointer to [**PriceInfoPriceExpireTime**](PriceInfoPriceExpireTime.md) |  | [optional] 
**PriceRange** | Pointer to [**PriceInfoPriceRange**](PriceInfoPriceRange.md) |  | [optional] 

## Methods

### NewPriceInfo

`func NewPriceInfo() *PriceInfo`

NewPriceInfo instantiates a new PriceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceInfoWithDefaults

`func NewPriceInfoWithDefaults() *PriceInfo`

NewPriceInfoWithDefaults instantiates a new PriceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *PriceInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *PriceInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *PriceInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *PriceInfo) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPrice

`func (o *PriceInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *PriceInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *PriceInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *PriceInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *PriceInfo) GetOriginalPrice() float32`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *PriceInfo) GetOriginalPriceOk() (*float32, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *PriceInfo) SetOriginalPrice(v float32)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *PriceInfo) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetCost

`func (o *PriceInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PriceInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PriceInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *PriceInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPriceEffectiveTime

`func (o *PriceInfo) GetPriceEffectiveTime() PriceInfoPriceEffectiveTime`

GetPriceEffectiveTime returns the PriceEffectiveTime field if non-nil, zero value otherwise.

### GetPriceEffectiveTimeOk

`func (o *PriceInfo) GetPriceEffectiveTimeOk() (*PriceInfoPriceEffectiveTime, bool)`

GetPriceEffectiveTimeOk returns a tuple with the PriceEffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveTime

`func (o *PriceInfo) SetPriceEffectiveTime(v PriceInfoPriceEffectiveTime)`

SetPriceEffectiveTime sets PriceEffectiveTime field to given value.

### HasPriceEffectiveTime

`func (o *PriceInfo) HasPriceEffectiveTime() bool`

HasPriceEffectiveTime returns a boolean if a field has been set.

### GetPriceExpireTime

`func (o *PriceInfo) GetPriceExpireTime() PriceInfoPriceExpireTime`

GetPriceExpireTime returns the PriceExpireTime field if non-nil, zero value otherwise.

### GetPriceExpireTimeOk

`func (o *PriceInfo) GetPriceExpireTimeOk() (*PriceInfoPriceExpireTime, bool)`

GetPriceExpireTimeOk returns a tuple with the PriceExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExpireTime

`func (o *PriceInfo) SetPriceExpireTime(v PriceInfoPriceExpireTime)`

SetPriceExpireTime sets PriceExpireTime field to given value.

### HasPriceExpireTime

`func (o *PriceInfo) HasPriceExpireTime() bool`

HasPriceExpireTime returns a boolean if a field has been set.

### GetPriceRange

`func (o *PriceInfo) GetPriceRange() PriceInfoPriceRange`

GetPriceRange returns the PriceRange field if non-nil, zero value otherwise.

### GetPriceRangeOk

`func (o *PriceInfo) GetPriceRangeOk() (*PriceInfoPriceRange, bool)`

GetPriceRangeOk returns a tuple with the PriceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRange

`func (o *PriceInfo) SetPriceRange(v PriceInfoPriceRange)`

SetPriceRange sets PriceRange field to given value.

### HasPriceRange

`func (o *PriceInfo) HasPriceRange() bool`

HasPriceRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


