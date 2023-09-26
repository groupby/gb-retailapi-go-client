# SearchResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the search. | [optional] 
**Area** | Pointer to **string** | Area Id the search was performed in. | [optional] 
**Query** | Pointer to **string** | Original search query. | [optional] 
**CorrectedQuery** | Pointer to **string** | Search query after any changes/corrections are done by the engine. | [optional] 
**BiasingProfile** | Pointer to **string** | Name of the biasing profile which was used to bias products in the search results. | [optional] 
**BiasingProfileAppliedId** | Pointer to **int32** | Id of the biasing profile which was used to bias products in the search results. | [optional] 
**Filter** | **string** |  | 
**OriginalRequest** | [**SearchRequestDto**](SearchRequestDto.md) |  | 
**Records** | Pointer to [**[]RecordDto**](RecordDto.md) | The list of records that match the search. | [optional] 
**TotalRecordCount** | Pointer to **int64** | The total number of products that match the search. If all products were filtered out on S4R site equals to 0. | [optional] 
**Metadata** | [**SearchMetadataDto**](SearchMetadataDto.md) |  | 
**PageInfo** | [**PageInfoDto**](PageInfoDto.md) |  | 
**AvailableNavigation** | [**[]NavigationDto**](NavigationDto.md) |  | 
**SelectedNavigation** | [**[]NavigationDto**](NavigationDto.md) |  | 
**Redirect** | Pointer to **string** | URL to which the merchandiser should redirect the shopper to. | [optional] 
**Experiments** | [**[]Experiment**](Experiment.md) |  | 
**Template** | [**TemplateDto**](TemplateDto.md) |  | 
**Empty** | Pointer to **bool** | True if the search yielded no results, otherwise false. | [optional] 
**SiteParams** | [**[]Metadata**](Metadata.md) |  | 
**Debug** | [**DebugDto**](DebugDto.md) |  | 
**Warnings** | Pointer to **[]string** | Warning messages containing information about invalid products, etc. | [optional] 
**IncludeExpandedResults** | Pointer to **bool** | When a shopper uses an ambiguous or a multi-word search phrase, they can get an empty response. After turning on include expanded results, Retail Search analyzes the request and returns the expanded list of products based on the parsed search query. For example, if you search \&quot;Google Pixel 5\&quot; without query expansion, you might only get \&quot;google_pixel_5\&quot; in the result. With query expansion, you might get \&quot;google_pixel_4a_with_5g\&quot;, \&quot;google_pixel_4a\&quot; and \&quot;google_pixel_5_case\&quot; as well.The default value is configured in the tenant settings or true if there is no such setting | [optional] 
**FacetLimit** | Pointer to **int32** | Maximum of facet values that should be returned for this facet. If not specified, defaults to 20. The maximum allowed value is 300. Values above 300 will be coerced to 300.  If this field is negative, an INVALID_ARGUMENT is returned.  This limit (300) is configured on Google side, but Google have an ability to change it for specific project.  | [optional] 
**RedirectMetadata** | [**[]Metadata**](Metadata.md) |  | 

## Methods

### NewSearchResponseDto

`func NewSearchResponseDto(filter string, originalRequest SearchRequestDto, metadata SearchMetadataDto, pageInfo PageInfoDto, availableNavigation []NavigationDto, selectedNavigation []NavigationDto, experiments []Experiment, template TemplateDto, siteParams []Metadata, debug DebugDto, redirectMetadata []Metadata, ) *SearchResponseDto`

NewSearchResponseDto instantiates a new SearchResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResponseDtoWithDefaults

`func NewSearchResponseDtoWithDefaults() *SearchResponseDto`

NewSearchResponseDtoWithDefaults instantiates a new SearchResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchResponseDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchResponseDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchResponseDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchResponseDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetArea

`func (o *SearchResponseDto) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *SearchResponseDto) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *SearchResponseDto) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *SearchResponseDto) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetQuery

`func (o *SearchResponseDto) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchResponseDto) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchResponseDto) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SearchResponseDto) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCorrectedQuery

`func (o *SearchResponseDto) GetCorrectedQuery() string`

GetCorrectedQuery returns the CorrectedQuery field if non-nil, zero value otherwise.

### GetCorrectedQueryOk

`func (o *SearchResponseDto) GetCorrectedQueryOk() (*string, bool)`

GetCorrectedQueryOk returns a tuple with the CorrectedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedQuery

`func (o *SearchResponseDto) SetCorrectedQuery(v string)`

SetCorrectedQuery sets CorrectedQuery field to given value.

### HasCorrectedQuery

`func (o *SearchResponseDto) HasCorrectedQuery() bool`

HasCorrectedQuery returns a boolean if a field has been set.

### GetBiasingProfile

`func (o *SearchResponseDto) GetBiasingProfile() string`

GetBiasingProfile returns the BiasingProfile field if non-nil, zero value otherwise.

### GetBiasingProfileOk

`func (o *SearchResponseDto) GetBiasingProfileOk() (*string, bool)`

GetBiasingProfileOk returns a tuple with the BiasingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasingProfile

`func (o *SearchResponseDto) SetBiasingProfile(v string)`

SetBiasingProfile sets BiasingProfile field to given value.

### HasBiasingProfile

`func (o *SearchResponseDto) HasBiasingProfile() bool`

HasBiasingProfile returns a boolean if a field has been set.

### GetBiasingProfileAppliedId

`func (o *SearchResponseDto) GetBiasingProfileAppliedId() int32`

GetBiasingProfileAppliedId returns the BiasingProfileAppliedId field if non-nil, zero value otherwise.

### GetBiasingProfileAppliedIdOk

`func (o *SearchResponseDto) GetBiasingProfileAppliedIdOk() (*int32, bool)`

GetBiasingProfileAppliedIdOk returns a tuple with the BiasingProfileAppliedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasingProfileAppliedId

`func (o *SearchResponseDto) SetBiasingProfileAppliedId(v int32)`

SetBiasingProfileAppliedId sets BiasingProfileAppliedId field to given value.

### HasBiasingProfileAppliedId

`func (o *SearchResponseDto) HasBiasingProfileAppliedId() bool`

HasBiasingProfileAppliedId returns a boolean if a field has been set.

### GetFilter

`func (o *SearchResponseDto) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SearchResponseDto) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SearchResponseDto) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetOriginalRequest

`func (o *SearchResponseDto) GetOriginalRequest() SearchRequestDto`

GetOriginalRequest returns the OriginalRequest field if non-nil, zero value otherwise.

### GetOriginalRequestOk

`func (o *SearchResponseDto) GetOriginalRequestOk() (*SearchRequestDto, bool)`

GetOriginalRequestOk returns a tuple with the OriginalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequest

`func (o *SearchResponseDto) SetOriginalRequest(v SearchRequestDto)`

SetOriginalRequest sets OriginalRequest field to given value.


### GetRecords

`func (o *SearchResponseDto) GetRecords() []RecordDto`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *SearchResponseDto) GetRecordsOk() (*[]RecordDto, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *SearchResponseDto) SetRecords(v []RecordDto)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *SearchResponseDto) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *SearchResponseDto) GetTotalRecordCount() int64`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *SearchResponseDto) GetTotalRecordCountOk() (*int64, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *SearchResponseDto) SetTotalRecordCount(v int64)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *SearchResponseDto) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetMetadata

`func (o *SearchResponseDto) GetMetadata() SearchMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchResponseDto) GetMetadataOk() (*SearchMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchResponseDto) SetMetadata(v SearchMetadataDto)`

SetMetadata sets Metadata field to given value.


### GetPageInfo

`func (o *SearchResponseDto) GetPageInfo() PageInfoDto`

GetPageInfo returns the PageInfo field if non-nil, zero value otherwise.

### GetPageInfoOk

`func (o *SearchResponseDto) GetPageInfoOk() (*PageInfoDto, bool)`

GetPageInfoOk returns a tuple with the PageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageInfo

`func (o *SearchResponseDto) SetPageInfo(v PageInfoDto)`

SetPageInfo sets PageInfo field to given value.


### GetAvailableNavigation

`func (o *SearchResponseDto) GetAvailableNavigation() []NavigationDto`

GetAvailableNavigation returns the AvailableNavigation field if non-nil, zero value otherwise.

### GetAvailableNavigationOk

`func (o *SearchResponseDto) GetAvailableNavigationOk() (*[]NavigationDto, bool)`

GetAvailableNavigationOk returns a tuple with the AvailableNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNavigation

`func (o *SearchResponseDto) SetAvailableNavigation(v []NavigationDto)`

SetAvailableNavigation sets AvailableNavigation field to given value.


### GetSelectedNavigation

`func (o *SearchResponseDto) GetSelectedNavigation() []NavigationDto`

GetSelectedNavigation returns the SelectedNavigation field if non-nil, zero value otherwise.

### GetSelectedNavigationOk

`func (o *SearchResponseDto) GetSelectedNavigationOk() (*[]NavigationDto, bool)`

GetSelectedNavigationOk returns a tuple with the SelectedNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedNavigation

`func (o *SearchResponseDto) SetSelectedNavigation(v []NavigationDto)`

SetSelectedNavigation sets SelectedNavigation field to given value.


### GetRedirect

`func (o *SearchResponseDto) GetRedirect() string`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SearchResponseDto) GetRedirectOk() (*string, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SearchResponseDto) SetRedirect(v string)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *SearchResponseDto) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetExperiments

`func (o *SearchResponseDto) GetExperiments() []Experiment`

GetExperiments returns the Experiments field if non-nil, zero value otherwise.

### GetExperimentsOk

`func (o *SearchResponseDto) GetExperimentsOk() (*[]Experiment, bool)`

GetExperimentsOk returns a tuple with the Experiments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiments

`func (o *SearchResponseDto) SetExperiments(v []Experiment)`

SetExperiments sets Experiments field to given value.


### GetTemplate

`func (o *SearchResponseDto) GetTemplate() TemplateDto`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SearchResponseDto) GetTemplateOk() (*TemplateDto, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SearchResponseDto) SetTemplate(v TemplateDto)`

SetTemplate sets Template field to given value.


### GetEmpty

`func (o *SearchResponseDto) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *SearchResponseDto) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *SearchResponseDto) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *SearchResponseDto) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetSiteParams

`func (o *SearchResponseDto) GetSiteParams() []Metadata`

GetSiteParams returns the SiteParams field if non-nil, zero value otherwise.

### GetSiteParamsOk

`func (o *SearchResponseDto) GetSiteParamsOk() (*[]Metadata, bool)`

GetSiteParamsOk returns a tuple with the SiteParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteParams

`func (o *SearchResponseDto) SetSiteParams(v []Metadata)`

SetSiteParams sets SiteParams field to given value.


### GetDebug

`func (o *SearchResponseDto) GetDebug() DebugDto`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *SearchResponseDto) GetDebugOk() (*DebugDto, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *SearchResponseDto) SetDebug(v DebugDto)`

SetDebug sets Debug field to given value.


### GetWarnings

`func (o *SearchResponseDto) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SearchResponseDto) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SearchResponseDto) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SearchResponseDto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetIncludeExpandedResults

`func (o *SearchResponseDto) GetIncludeExpandedResults() bool`

GetIncludeExpandedResults returns the IncludeExpandedResults field if non-nil, zero value otherwise.

### GetIncludeExpandedResultsOk

`func (o *SearchResponseDto) GetIncludeExpandedResultsOk() (*bool, bool)`

GetIncludeExpandedResultsOk returns a tuple with the IncludeExpandedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpandedResults

`func (o *SearchResponseDto) SetIncludeExpandedResults(v bool)`

SetIncludeExpandedResults sets IncludeExpandedResults field to given value.

### HasIncludeExpandedResults

`func (o *SearchResponseDto) HasIncludeExpandedResults() bool`

HasIncludeExpandedResults returns a boolean if a field has been set.

### GetFacetLimit

`func (o *SearchResponseDto) GetFacetLimit() int32`

GetFacetLimit returns the FacetLimit field if non-nil, zero value otherwise.

### GetFacetLimitOk

`func (o *SearchResponseDto) GetFacetLimitOk() (*int32, bool)`

GetFacetLimitOk returns a tuple with the FacetLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacetLimit

`func (o *SearchResponseDto) SetFacetLimit(v int32)`

SetFacetLimit sets FacetLimit field to given value.

### HasFacetLimit

`func (o *SearchResponseDto) HasFacetLimit() bool`

HasFacetLimit returns a boolean if a field has been set.

### GetRedirectMetadata

`func (o *SearchResponseDto) GetRedirectMetadata() []Metadata`

GetRedirectMetadata returns the RedirectMetadata field if non-nil, zero value otherwise.

### GetRedirectMetadataOk

`func (o *SearchResponseDto) GetRedirectMetadataOk() (*[]Metadata, bool)`

GetRedirectMetadataOk returns a tuple with the RedirectMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectMetadata

`func (o *SearchResponseDto) SetRedirectMetadata(v []Metadata)`

SetRedirectMetadata sets RedirectMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


