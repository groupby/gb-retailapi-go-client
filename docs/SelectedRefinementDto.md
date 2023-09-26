# SelectedRefinementDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NavigationName** | **string** | The name of the navigation the refinement is for. | 
**Type** | [**NavigationTypeDto**](NavigationTypeDto.md) |  | 
**Value** | Pointer to **string** | Value of selected refinement, if type is value. | [optional] 
**Low** | Pointer to **float64** | The lowest end or value of the range, if applicable. | [optional] 
**High** | Pointer to **float64** | The highest end or value of the range, if applicable. | [optional] 
**Source** | Pointer to **string** | Field which is indicated that it is dynamic navigation. | [optional] 
**Or** | Pointer to **bool** | Navigation multiselect. Indicate that it is possibly to select more than one navigation value due to search request. | [optional] 

## Methods

### NewSelectedRefinementDto

`func NewSelectedRefinementDto(navigationName string, type_ NavigationTypeDto, ) *SelectedRefinementDto`

NewSelectedRefinementDto instantiates a new SelectedRefinementDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedRefinementDtoWithDefaults

`func NewSelectedRefinementDtoWithDefaults() *SelectedRefinementDto`

NewSelectedRefinementDtoWithDefaults instantiates a new SelectedRefinementDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNavigationName

`func (o *SelectedRefinementDto) GetNavigationName() string`

GetNavigationName returns the NavigationName field if non-nil, zero value otherwise.

### GetNavigationNameOk

`func (o *SelectedRefinementDto) GetNavigationNameOk() (*string, bool)`

GetNavigationNameOk returns a tuple with the NavigationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationName

`func (o *SelectedRefinementDto) SetNavigationName(v string)`

SetNavigationName sets NavigationName field to given value.


### GetType

`func (o *SelectedRefinementDto) GetType() NavigationTypeDto`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectedRefinementDto) GetTypeOk() (*NavigationTypeDto, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectedRefinementDto) SetType(v NavigationTypeDto)`

SetType sets Type field to given value.


### GetValue

`func (o *SelectedRefinementDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SelectedRefinementDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SelectedRefinementDto) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SelectedRefinementDto) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLow

`func (o *SelectedRefinementDto) GetLow() float64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *SelectedRefinementDto) GetLowOk() (*float64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *SelectedRefinementDto) SetLow(v float64)`

SetLow sets Low field to given value.

### HasLow

`func (o *SelectedRefinementDto) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetHigh

`func (o *SelectedRefinementDto) GetHigh() float64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *SelectedRefinementDto) GetHighOk() (*float64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *SelectedRefinementDto) SetHigh(v float64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *SelectedRefinementDto) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetSource

`func (o *SelectedRefinementDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SelectedRefinementDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SelectedRefinementDto) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SelectedRefinementDto) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetOr

`func (o *SelectedRefinementDto) GetOr() bool`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *SelectedRefinementDto) GetOrOk() (*bool, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *SelectedRefinementDto) SetOr(v bool)`

SetOr sets Or field to given value.

### HasOr

`func (o *SelectedRefinementDto) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


