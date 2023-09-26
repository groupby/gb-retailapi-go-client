# RefinementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NavigationTypeDto**](NavigationTypeDto.md) |  | 
**Count** | Pointer to **int64** | Number of products which match this refinement value or range. | [optional] 
**Value** | Pointer to **string** | Value of the refinement. | [optional] 
**Low** | Pointer to **float64** | Lower bound of the refinement range. | [optional] 
**High** | Pointer to **float64** | Upper bound  of the refinement range. | [optional] 

## Methods

### NewRefinementDto

`func NewRefinementDto(type_ NavigationTypeDto, ) *RefinementDto`

NewRefinementDto instantiates a new RefinementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefinementDtoWithDefaults

`func NewRefinementDtoWithDefaults() *RefinementDto`

NewRefinementDtoWithDefaults instantiates a new RefinementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RefinementDto) GetType() NavigationTypeDto`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RefinementDto) GetTypeOk() (*NavigationTypeDto, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RefinementDto) SetType(v NavigationTypeDto)`

SetType sets Type field to given value.


### GetCount

`func (o *RefinementDto) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RefinementDto) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RefinementDto) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RefinementDto) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetValue

`func (o *RefinementDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RefinementDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RefinementDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RefinementDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLow

`func (o *RefinementDto) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *RefinementDto) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *RefinementDto) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *RefinementDto) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetHigh

`func (o *RefinementDto) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *RefinementDto) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *RefinementDto) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *RefinementDto) HasHigh() bool`

HasHigh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


