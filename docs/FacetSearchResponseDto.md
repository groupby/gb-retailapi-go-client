# FacetSearchResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalRequest** | [**SearchRequestDto**](SearchRequestDto.md) |  | 
**AvailableNavigation** | [**NavigationDto**](NavigationDto.md) |  | 

## Methods

### NewFacetSearchResponseDto

`func NewFacetSearchResponseDto(originalRequest SearchRequestDto, availableNavigation NavigationDto, ) *FacetSearchResponseDto`

NewFacetSearchResponseDto instantiates a new FacetSearchResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetSearchResponseDtoWithDefaults

`func NewFacetSearchResponseDtoWithDefaults() *FacetSearchResponseDto`

NewFacetSearchResponseDtoWithDefaults instantiates a new FacetSearchResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalRequest

`func (o *FacetSearchResponseDto) GetOriginalRequest() SearchRequestDto`

GetOriginalRequest returns the OriginalRequest field if non-nil, zero value otherwise.

### GetOriginalRequestOk

`func (o *FacetSearchResponseDto) GetOriginalRequestOk() (*SearchRequestDto, bool)`

GetOriginalRequestOk returns a tuple with the OriginalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequest

`func (o *FacetSearchResponseDto) SetOriginalRequest(v SearchRequestDto)`

SetOriginalRequest sets OriginalRequest field to given value.


### GetAvailableNavigation

`func (o *FacetSearchResponseDto) GetAvailableNavigation() NavigationDto`

GetAvailableNavigation returns the AvailableNavigation field if non-nil, zero value otherwise.

### GetAvailableNavigationOk

`func (o *FacetSearchResponseDto) GetAvailableNavigationOk() (*NavigationDto, bool)`

GetAvailableNavigationOk returns a tuple with the AvailableNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNavigation

`func (o *FacetSearchResponseDto) SetAvailableNavigation(v NavigationDto)`

SetAvailableNavigation sets AvailableNavigation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


