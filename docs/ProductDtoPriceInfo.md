# ProductDtoPriceInfo

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

### NewProductDtoPriceInfo

`func NewProductDtoPriceInfo() *ProductDtoPriceInfo`

NewProductDtoPriceInfo instantiates a new ProductDtoPriceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDtoPriceInfoWithDefaults

`func NewProductDtoPriceInfoWithDefaults() *ProductDtoPriceInfo`

NewProductDtoPriceInfoWithDefaults instantiates a new ProductDtoPriceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *ProductDtoPriceInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ProductDtoPriceInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ProductDtoPriceInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ProductDtoPriceInfo) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetPrice

`func (o *ProductDtoPriceInfo) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductDtoPriceInfo) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductDtoPriceInfo) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductDtoPriceInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOriginalPrice

`func (o *ProductDtoPriceInfo) GetOriginalPrice() float32`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *ProductDtoPriceInfo) GetOriginalPriceOk() (*float32, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *ProductDtoPriceInfo) SetOriginalPrice(v float32)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *ProductDtoPriceInfo) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### GetCost

`func (o *ProductDtoPriceInfo) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductDtoPriceInfo) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductDtoPriceInfo) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProductDtoPriceInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPriceEffectiveTime

`func (o *ProductDtoPriceInfo) GetPriceEffectiveTime() PriceInfoPriceEffectiveTime`

GetPriceEffectiveTime returns the PriceEffectiveTime field if non-nil, zero value otherwise.

### GetPriceEffectiveTimeOk

`func (o *ProductDtoPriceInfo) GetPriceEffectiveTimeOk() (*PriceInfoPriceEffectiveTime, bool)`

GetPriceEffectiveTimeOk returns a tuple with the PriceEffectiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceEffectiveTime

`func (o *ProductDtoPriceInfo) SetPriceEffectiveTime(v PriceInfoPriceEffectiveTime)`

SetPriceEffectiveTime sets PriceEffectiveTime field to given value.

### HasPriceEffectiveTime

`func (o *ProductDtoPriceInfo) HasPriceEffectiveTime() bool`

HasPriceEffectiveTime returns a boolean if a field has been set.

### GetPriceExpireTime

`func (o *ProductDtoPriceInfo) GetPriceExpireTime() PriceInfoPriceExpireTime`

GetPriceExpireTime returns the PriceExpireTime field if non-nil, zero value otherwise.

### GetPriceExpireTimeOk

`func (o *ProductDtoPriceInfo) GetPriceExpireTimeOk() (*PriceInfoPriceExpireTime, bool)`

GetPriceExpireTimeOk returns a tuple with the PriceExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceExpireTime

`func (o *ProductDtoPriceInfo) SetPriceExpireTime(v PriceInfoPriceExpireTime)`

SetPriceExpireTime sets PriceExpireTime field to given value.

### HasPriceExpireTime

`func (o *ProductDtoPriceInfo) HasPriceExpireTime() bool`

HasPriceExpireTime returns a boolean if a field has been set.

### GetPriceRange

`func (o *ProductDtoPriceInfo) GetPriceRange() PriceInfoPriceRange`

GetPriceRange returns the PriceRange field if non-nil, zero value otherwise.

### GetPriceRangeOk

`func (o *ProductDtoPriceInfo) GetPriceRangeOk() (*PriceInfoPriceRange, bool)`

GetPriceRangeOk returns a tuple with the PriceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRange

`func (o *ProductDtoPriceInfo) SetPriceRange(v PriceInfoPriceRange)`

SetPriceRange sets PriceRange field to given value.

### HasPriceRange

`func (o *ProductDtoPriceInfo) HasPriceRange() bool`

HasPriceRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


