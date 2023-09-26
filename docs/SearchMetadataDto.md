# SearchMetadataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchAttributionToken** | Pointer to **NullableString** | Token to assist beacon collectors in correlating searches to user events. | [optional] 
**Cached** | Pointer to **bool** | Were the search results from a previous call. | [optional] 
**TotalTime** | Pointer to **int64** | Total time spent performing the Google search in milliseconds. | [optional] 

## Methods

### NewSearchMetadataDto

`func NewSearchMetadataDto() *SearchMetadataDto`

NewSearchMetadataDto instantiates a new SearchMetadataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMetadataDtoWithDefaults

`func NewSearchMetadataDtoWithDefaults() *SearchMetadataDto`

NewSearchMetadataDtoWithDefaults instantiates a new SearchMetadataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchAttributionToken

`func (o *SearchMetadataDto) GetSearchAttributionToken() string`

GetSearchAttributionToken returns the SearchAttributionToken field if non-nil, zero value otherwise.

### GetSearchAttributionTokenOk

`func (o *SearchMetadataDto) GetSearchAttributionTokenOk() (*string, bool)`

GetSearchAttributionTokenOk returns a tuple with the SearchAttributionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchAttributionToken

`func (o *SearchMetadataDto) SetSearchAttributionToken(v string)`

SetSearchAttributionToken sets SearchAttributionToken field to given value.

### HasSearchAttributionToken

`func (o *SearchMetadataDto) HasSearchAttributionToken() bool`

HasSearchAttributionToken returns a boolean if a field has been set.

### SetSearchAttributionTokenNil

`func (o *SearchMetadataDto) SetSearchAttributionTokenNil(b bool)`

 SetSearchAttributionTokenNil sets the value for SearchAttributionToken to be an explicit nil

### UnsetSearchAttributionToken
`func (o *SearchMetadataDto) UnsetSearchAttributionToken()`

UnsetSearchAttributionToken ensures that no value is present for SearchAttributionToken, not even an explicit nil
### GetCached

`func (o *SearchMetadataDto) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *SearchMetadataDto) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *SearchMetadataDto) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *SearchMetadataDto) HasCached() bool`

HasCached returns a boolean if a field has been set.

### GetTotalTime

`func (o *SearchMetadataDto) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *SearchMetadataDto) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *SearchMetadataDto) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *SearchMetadataDto) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


