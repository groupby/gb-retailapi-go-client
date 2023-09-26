# SearchResultsStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SearchCount** | Pointer to **int32** | Count of suggested sentences in response. | [optional] 
**AutocompleteResponse** | Pointer to **int64** | Time in milliseconds taken by autocomplete service to handle request and send response. | [optional] 
**ExtendedAttributesCount** | Pointer to **int32** | Count of extended attributes in autocomplete response. | [optional] 
**ExtendedAttributesResponse** | Pointer to **int64** | Time in milliseconds taken by application to enrich response with extended attributes. | [optional] 

## Methods

### NewSearchResultsStats

`func NewSearchResultsStats() *SearchResultsStats`

NewSearchResultsStats instantiates a new SearchResultsStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsStatsWithDefaults

`func NewSearchResultsStatsWithDefaults() *SearchResultsStats`

NewSearchResultsStatsWithDefaults instantiates a new SearchResultsStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSearchCount

`func (o *SearchResultsStats) GetSearchCount() int32`

GetSearchCount returns the SearchCount field if non-nil, zero value otherwise.

### GetSearchCountOk

`func (o *SearchResultsStats) GetSearchCountOk() (*int32, bool)`

GetSearchCountOk returns a tuple with the SearchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchCount

`func (o *SearchResultsStats) SetSearchCount(v int32)`

SetSearchCount sets SearchCount field to given value.

### HasSearchCount

`func (o *SearchResultsStats) HasSearchCount() bool`

HasSearchCount returns a boolean if a field has been set.

### GetAutocompleteResponse

`func (o *SearchResultsStats) GetAutocompleteResponse() int64`

GetAutocompleteResponse returns the AutocompleteResponse field if non-nil, zero value otherwise.

### GetAutocompleteResponseOk

`func (o *SearchResultsStats) GetAutocompleteResponseOk() (*int64, bool)`

GetAutocompleteResponseOk returns a tuple with the AutocompleteResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutocompleteResponse

`func (o *SearchResultsStats) SetAutocompleteResponse(v int64)`

SetAutocompleteResponse sets AutocompleteResponse field to given value.

### HasAutocompleteResponse

`func (o *SearchResultsStats) HasAutocompleteResponse() bool`

HasAutocompleteResponse returns a boolean if a field has been set.

### GetExtendedAttributesCount

`func (o *SearchResultsStats) GetExtendedAttributesCount() int32`

GetExtendedAttributesCount returns the ExtendedAttributesCount field if non-nil, zero value otherwise.

### GetExtendedAttributesCountOk

`func (o *SearchResultsStats) GetExtendedAttributesCountOk() (*int32, bool)`

GetExtendedAttributesCountOk returns a tuple with the ExtendedAttributesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributesCount

`func (o *SearchResultsStats) SetExtendedAttributesCount(v int32)`

SetExtendedAttributesCount sets ExtendedAttributesCount field to given value.

### HasExtendedAttributesCount

`func (o *SearchResultsStats) HasExtendedAttributesCount() bool`

HasExtendedAttributesCount returns a boolean if a field has been set.

### GetExtendedAttributesResponse

`func (o *SearchResultsStats) GetExtendedAttributesResponse() int64`

GetExtendedAttributesResponse returns the ExtendedAttributesResponse field if non-nil, zero value otherwise.

### GetExtendedAttributesResponseOk

`func (o *SearchResultsStats) GetExtendedAttributesResponseOk() (*int64, bool)`

GetExtendedAttributesResponseOk returns a tuple with the ExtendedAttributesResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributesResponse

`func (o *SearchResultsStats) SetExtendedAttributesResponse(v int64)`

SetExtendedAttributesResponse sets ExtendedAttributesResponse field to given value.

### HasExtendedAttributesResponse

`func (o *SearchResultsStats) HasExtendedAttributesResponse() bool`

HasExtendedAttributesResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


