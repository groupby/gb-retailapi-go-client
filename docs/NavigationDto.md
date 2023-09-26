# NavigationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the field the navigation in on. | [optional] 
**DisplayName** | Pointer to **string** | Name of the navigation for display purposes. | [optional] 
**Type** | [**NavigationTypeDto**](NavigationTypeDto.md) |  | 
**Refinements** | [**[]RefinementDto**](RefinementDto.md) |  | 
**Or** | Pointer to **bool** | Flag if this navigation supports or queries. | [optional] 
**Source** | **string** |  | 
**Metadata** | [**[]Metadata**](Metadata.md) |  | 
**PlaceId** | **string** | Place id for inventory navigations. | 

## Methods

### NewNavigationDto

`func NewNavigationDto(type_ NavigationTypeDto, refinements []RefinementDto, source string, metadata []Metadata, placeId string, ) *NavigationDto`

NewNavigationDto instantiates a new NavigationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationDtoWithDefaults

`func NewNavigationDtoWithDefaults() *NavigationDto`

NewNavigationDtoWithDefaults instantiates a new NavigationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NavigationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NavigationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NavigationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NavigationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *NavigationDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NavigationDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NavigationDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NavigationDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *NavigationDto) GetType() NavigationTypeDto`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NavigationDto) GetTypeOk() (*NavigationTypeDto, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NavigationDto) SetType(v NavigationTypeDto)`

SetType sets Type field to given value.


### GetRefinements

`func (o *NavigationDto) GetRefinements() []RefinementDto`

GetRefinements returns the Refinements field if non-nil, zero value otherwise.

### GetRefinementsOk

`func (o *NavigationDto) GetRefinementsOk() (*[]RefinementDto, bool)`

GetRefinementsOk returns a tuple with the Refinements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinements

`func (o *NavigationDto) SetRefinements(v []RefinementDto)`

SetRefinements sets Refinements field to given value.


### GetOr

`func (o *NavigationDto) GetOr() bool`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *NavigationDto) GetOrOk() (*bool, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *NavigationDto) SetOr(v bool)`

SetOr sets Or field to given value.

### HasOr

`func (o *NavigationDto) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetSource

`func (o *NavigationDto) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NavigationDto) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NavigationDto) SetSource(v string)`

SetSource sets Source field to given value.


### GetMetadata

`func (o *NavigationDto) GetMetadata() []Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NavigationDto) GetMetadataOk() (*[]Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NavigationDto) SetMetadata(v []Metadata)`

SetMetadata sets Metadata field to given value.


### GetPlaceId

`func (o *NavigationDto) GetPlaceId() string`

GetPlaceId returns the PlaceId field if non-nil, zero value otherwise.

### GetPlaceIdOk

`func (o *NavigationDto) GetPlaceIdOk() (*string, bool)`

GetPlaceIdOk returns a tuple with the PlaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceId

`func (o *NavigationDto) SetPlaceId(v string)`

SetPlaceId sets PlaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


