# SearchResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stats** | Pointer to [**SearchResultsStats**](SearchResultsStats.md) |  | [optional] 
**SearchTerms** | [**[]SearchTerms**](SearchTerms.md) |  | 
**ExtendedAttributes** | Pointer to **map[string][]string** | Map with extended attributes which are returned in autocomplete response.  | [optional] 
**AttributeResults** | Pointer to [**map[string]AttributeSuggestion**](AttributeSuggestion.md) | SAYT response attributes. Contains list of direct matching attributes. | [optional] 
**SiteFilter** | Pointer to **string** | SiteFilter object used with request. | [optional] 

## Methods

### NewSearchResults

`func NewSearchResults(searchTerms []SearchTerms, ) *SearchResults`

NewSearchResults instantiates a new SearchResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultsWithDefaults

`func NewSearchResultsWithDefaults() *SearchResults`

NewSearchResultsWithDefaults instantiates a new SearchResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStats

`func (o *SearchResults) GetStats() SearchResultsStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SearchResults) GetStatsOk() (*SearchResultsStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SearchResults) SetStats(v SearchResultsStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SearchResults) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetSearchTerms

`func (o *SearchResults) GetSearchTerms() []SearchTerms`

GetSearchTerms returns the SearchTerms field if non-nil, zero value otherwise.

### GetSearchTermsOk

`func (o *SearchResults) GetSearchTermsOk() (*[]SearchTerms, bool)`

GetSearchTermsOk returns a tuple with the SearchTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerms

`func (o *SearchResults) SetSearchTerms(v []SearchTerms)`

SetSearchTerms sets SearchTerms field to given value.


### GetExtendedAttributes

`func (o *SearchResults) GetExtendedAttributes() map[string][]string`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *SearchResults) GetExtendedAttributesOk() (*map[string][]string, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *SearchResults) SetExtendedAttributes(v map[string][]string)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *SearchResults) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### GetAttributeResults

`func (o *SearchResults) GetAttributeResults() map[string]AttributeSuggestion`

GetAttributeResults returns the AttributeResults field if non-nil, zero value otherwise.

### GetAttributeResultsOk

`func (o *SearchResults) GetAttributeResultsOk() (*map[string]AttributeSuggestion, bool)`

GetAttributeResultsOk returns a tuple with the AttributeResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeResults

`func (o *SearchResults) SetAttributeResults(v map[string]AttributeSuggestion)`

SetAttributeResults sets AttributeResults field to given value.

### HasAttributeResults

`func (o *SearchResults) HasAttributeResults() bool`

HasAttributeResults returns a boolean if a field has been set.

### GetSiteFilter

`func (o *SearchResults) GetSiteFilter() string`

GetSiteFilter returns the SiteFilter field if non-nil, zero value otherwise.

### GetSiteFilterOk

`func (o *SearchResults) GetSiteFilterOk() (*string, bool)`

GetSiteFilterOk returns a tuple with the SiteFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteFilter

`func (o *SearchResults) SetSiteFilter(v string)`

SetSiteFilter sets SiteFilter field to given value.

### HasSiteFilter

`func (o *SearchResults) HasSiteFilter() bool`

HasSiteFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


