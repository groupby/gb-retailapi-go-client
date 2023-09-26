# RuleVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BiasingProfileName** | **string** |  | 
**IncludedNavigations** | **[]string** |  | 
**Template** | [**RuleTemplate**](RuleTemplate.md) |  | 
**BoostedProductBuckets** | [**[]BoostedProductBucket**](BoostedProductBucket.md) |  | 
**PinnedRefinements** | [**[]PinnedRefinement**](PinnedRefinement.md) |  | 
**Sort** | [**Sort**](Sort.md) |  | 
**ValueFilters** | [**[]ValueFilter**](ValueFilter.md) |  | 
**SearchFilters** | [**[]SearchFilter**](SearchFilter.md) |  | 
**RangeFilters** | [**[]RangeFilter**](RangeFilter.md) |  | 

## Methods

### NewRuleVariant

`func NewRuleVariant(biasingProfileName string, includedNavigations []string, template RuleTemplate, boostedProductBuckets []BoostedProductBucket, pinnedRefinements []PinnedRefinement, sort Sort, valueFilters []ValueFilter, searchFilters []SearchFilter, rangeFilters []RangeFilter, ) *RuleVariant`

NewRuleVariant instantiates a new RuleVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleVariantWithDefaults

`func NewRuleVariantWithDefaults() *RuleVariant`

NewRuleVariantWithDefaults instantiates a new RuleVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBiasingProfileName

`func (o *RuleVariant) GetBiasingProfileName() string`

GetBiasingProfileName returns the BiasingProfileName field if non-nil, zero value otherwise.

### GetBiasingProfileNameOk

`func (o *RuleVariant) GetBiasingProfileNameOk() (*string, bool)`

GetBiasingProfileNameOk returns a tuple with the BiasingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasingProfileName

`func (o *RuleVariant) SetBiasingProfileName(v string)`

SetBiasingProfileName sets BiasingProfileName field to given value.


### GetIncludedNavigations

`func (o *RuleVariant) GetIncludedNavigations() []string`

GetIncludedNavigations returns the IncludedNavigations field if non-nil, zero value otherwise.

### GetIncludedNavigationsOk

`func (o *RuleVariant) GetIncludedNavigationsOk() (*[]string, bool)`

GetIncludedNavigationsOk returns a tuple with the IncludedNavigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedNavigations

`func (o *RuleVariant) SetIncludedNavigations(v []string)`

SetIncludedNavigations sets IncludedNavigations field to given value.


### GetTemplate

`func (o *RuleVariant) GetTemplate() RuleTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *RuleVariant) GetTemplateOk() (*RuleTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *RuleVariant) SetTemplate(v RuleTemplate)`

SetTemplate sets Template field to given value.


### GetBoostedProductBuckets

`func (o *RuleVariant) GetBoostedProductBuckets() []BoostedProductBucket`

GetBoostedProductBuckets returns the BoostedProductBuckets field if non-nil, zero value otherwise.

### GetBoostedProductBucketsOk

`func (o *RuleVariant) GetBoostedProductBucketsOk() (*[]BoostedProductBucket, bool)`

GetBoostedProductBucketsOk returns a tuple with the BoostedProductBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostedProductBuckets

`func (o *RuleVariant) SetBoostedProductBuckets(v []BoostedProductBucket)`

SetBoostedProductBuckets sets BoostedProductBuckets field to given value.


### GetPinnedRefinements

`func (o *RuleVariant) GetPinnedRefinements() []PinnedRefinement`

GetPinnedRefinements returns the PinnedRefinements field if non-nil, zero value otherwise.

### GetPinnedRefinementsOk

`func (o *RuleVariant) GetPinnedRefinementsOk() (*[]PinnedRefinement, bool)`

GetPinnedRefinementsOk returns a tuple with the PinnedRefinements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedRefinements

`func (o *RuleVariant) SetPinnedRefinements(v []PinnedRefinement)`

SetPinnedRefinements sets PinnedRefinements field to given value.


### GetSort

`func (o *RuleVariant) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RuleVariant) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RuleVariant) SetSort(v Sort)`

SetSort sets Sort field to given value.


### GetValueFilters

`func (o *RuleVariant) GetValueFilters() []ValueFilter`

GetValueFilters returns the ValueFilters field if non-nil, zero value otherwise.

### GetValueFiltersOk

`func (o *RuleVariant) GetValueFiltersOk() (*[]ValueFilter, bool)`

GetValueFiltersOk returns a tuple with the ValueFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFilters

`func (o *RuleVariant) SetValueFilters(v []ValueFilter)`

SetValueFilters sets ValueFilters field to given value.


### GetSearchFilters

`func (o *RuleVariant) GetSearchFilters() []SearchFilter`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *RuleVariant) GetSearchFiltersOk() (*[]SearchFilter, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *RuleVariant) SetSearchFilters(v []SearchFilter)`

SetSearchFilters sets SearchFilters field to given value.


### GetRangeFilters

`func (o *RuleVariant) GetRangeFilters() []RangeFilter`

GetRangeFilters returns the RangeFilters field if non-nil, zero value otherwise.

### GetRangeFiltersOk

`func (o *RuleVariant) GetRangeFiltersOk() (*[]RangeFilter, bool)`

GetRangeFiltersOk returns a tuple with the RangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeFilters

`func (o *RuleVariant) SetRangeFilters(v []RangeFilter)`

SetRangeFilters sets RangeFilters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


