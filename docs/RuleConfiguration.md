# RuleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleInExperiment** | Pointer to **bool** |  | [optional] 
**Id** | **int32** |  | 
**Name** | **string** |  | 
**AreaId** | **int32** |  | 
**Priority** | **int32** |  | 
**Enabled** | **bool** |  | 
**ActiveHoursEnabled** | **bool** |  | 
**ActiveFrom** | **int64** |  | 
**ActiveTo** | **int64** |  | 
**TriggerSets** | [**[]TriggerSet**](TriggerSet.md) |  | 
**BiasingProfileName** | **string** |  | 
**Sort** | [**Sort**](Sort.md) |  | 
**IncludedNavigations** | **[]string** |  | 
**ValueFilters** | [**[]ValueFilter**](ValueFilter.md) |  | 
**SearchFilters** | [**[]SearchFilter**](SearchFilter.md) |  | 
**RangeFilters** | [**[]RangeFilter**](RangeFilter.md) |  | 
**Template** | [**RuleTemplate**](RuleTemplate.md) |  | 
**BoostedProductBuckets** | [**[]BoostedProductBucket**](BoostedProductBucket.md) |  | 
**PinnedRefinements** | [**[]PinnedRefinement**](PinnedRefinement.md) |  | 
**MessageType** | [**MessageType**](MessageType.md) |  | 
**Type** | [**RuleType**](RuleType.md) |  | 
**Variants** | [**[]ExperimentVariant**](ExperimentVariant.md) |  | 

## Methods

### NewRuleConfiguration

`func NewRuleConfiguration(id int32, name string, areaId int32, priority int32, enabled bool, activeHoursEnabled bool, activeFrom int64, activeTo int64, triggerSets []TriggerSet, biasingProfileName string, sort Sort, includedNavigations []string, valueFilters []ValueFilter, searchFilters []SearchFilter, rangeFilters []RangeFilter, template RuleTemplate, boostedProductBuckets []BoostedProductBucket, pinnedRefinements []PinnedRefinement, messageType MessageType, type_ RuleType, variants []ExperimentVariant, ) *RuleConfiguration`

NewRuleConfiguration instantiates a new RuleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConfigurationWithDefaults

`func NewRuleConfigurationWithDefaults() *RuleConfiguration`

NewRuleConfigurationWithDefaults instantiates a new RuleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleInExperiment

`func (o *RuleConfiguration) GetRuleInExperiment() bool`

GetRuleInExperiment returns the RuleInExperiment field if non-nil, zero value otherwise.

### GetRuleInExperimentOk

`func (o *RuleConfiguration) GetRuleInExperimentOk() (*bool, bool)`

GetRuleInExperimentOk returns a tuple with the RuleInExperiment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleInExperiment

`func (o *RuleConfiguration) SetRuleInExperiment(v bool)`

SetRuleInExperiment sets RuleInExperiment field to given value.

### HasRuleInExperiment

`func (o *RuleConfiguration) HasRuleInExperiment() bool`

HasRuleInExperiment returns a boolean if a field has been set.

### GetId

`func (o *RuleConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RuleConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RuleConfiguration) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *RuleConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetAreaId

`func (o *RuleConfiguration) GetAreaId() int32`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *RuleConfiguration) GetAreaIdOk() (*int32, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *RuleConfiguration) SetAreaId(v int32)`

SetAreaId sets AreaId field to given value.


### GetPriority

`func (o *RuleConfiguration) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RuleConfiguration) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RuleConfiguration) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetEnabled

`func (o *RuleConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RuleConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RuleConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetActiveHoursEnabled

`func (o *RuleConfiguration) GetActiveHoursEnabled() bool`

GetActiveHoursEnabled returns the ActiveHoursEnabled field if non-nil, zero value otherwise.

### GetActiveHoursEnabledOk

`func (o *RuleConfiguration) GetActiveHoursEnabledOk() (*bool, bool)`

GetActiveHoursEnabledOk returns a tuple with the ActiveHoursEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveHoursEnabled

`func (o *RuleConfiguration) SetActiveHoursEnabled(v bool)`

SetActiveHoursEnabled sets ActiveHoursEnabled field to given value.


### GetActiveFrom

`func (o *RuleConfiguration) GetActiveFrom() int64`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *RuleConfiguration) GetActiveFromOk() (*int64, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *RuleConfiguration) SetActiveFrom(v int64)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveTo

`func (o *RuleConfiguration) GetActiveTo() int64`

GetActiveTo returns the ActiveTo field if non-nil, zero value otherwise.

### GetActiveToOk

`func (o *RuleConfiguration) GetActiveToOk() (*int64, bool)`

GetActiveToOk returns a tuple with the ActiveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveTo

`func (o *RuleConfiguration) SetActiveTo(v int64)`

SetActiveTo sets ActiveTo field to given value.


### GetTriggerSets

`func (o *RuleConfiguration) GetTriggerSets() []TriggerSet`

GetTriggerSets returns the TriggerSets field if non-nil, zero value otherwise.

### GetTriggerSetsOk

`func (o *RuleConfiguration) GetTriggerSetsOk() (*[]TriggerSet, bool)`

GetTriggerSetsOk returns a tuple with the TriggerSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSets

`func (o *RuleConfiguration) SetTriggerSets(v []TriggerSet)`

SetTriggerSets sets TriggerSets field to given value.


### GetBiasingProfileName

`func (o *RuleConfiguration) GetBiasingProfileName() string`

GetBiasingProfileName returns the BiasingProfileName field if non-nil, zero value otherwise.

### GetBiasingProfileNameOk

`func (o *RuleConfiguration) GetBiasingProfileNameOk() (*string, bool)`

GetBiasingProfileNameOk returns a tuple with the BiasingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiasingProfileName

`func (o *RuleConfiguration) SetBiasingProfileName(v string)`

SetBiasingProfileName sets BiasingProfileName field to given value.


### GetSort

`func (o *RuleConfiguration) GetSort() Sort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RuleConfiguration) GetSortOk() (*Sort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RuleConfiguration) SetSort(v Sort)`

SetSort sets Sort field to given value.


### GetIncludedNavigations

`func (o *RuleConfiguration) GetIncludedNavigations() []string`

GetIncludedNavigations returns the IncludedNavigations field if non-nil, zero value otherwise.

### GetIncludedNavigationsOk

`func (o *RuleConfiguration) GetIncludedNavigationsOk() (*[]string, bool)`

GetIncludedNavigationsOk returns a tuple with the IncludedNavigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedNavigations

`func (o *RuleConfiguration) SetIncludedNavigations(v []string)`

SetIncludedNavigations sets IncludedNavigations field to given value.


### GetValueFilters

`func (o *RuleConfiguration) GetValueFilters() []ValueFilter`

GetValueFilters returns the ValueFilters field if non-nil, zero value otherwise.

### GetValueFiltersOk

`func (o *RuleConfiguration) GetValueFiltersOk() (*[]ValueFilter, bool)`

GetValueFiltersOk returns a tuple with the ValueFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFilters

`func (o *RuleConfiguration) SetValueFilters(v []ValueFilter)`

SetValueFilters sets ValueFilters field to given value.


### GetSearchFilters

`func (o *RuleConfiguration) GetSearchFilters() []SearchFilter`

GetSearchFilters returns the SearchFilters field if non-nil, zero value otherwise.

### GetSearchFiltersOk

`func (o *RuleConfiguration) GetSearchFiltersOk() (*[]SearchFilter, bool)`

GetSearchFiltersOk returns a tuple with the SearchFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilters

`func (o *RuleConfiguration) SetSearchFilters(v []SearchFilter)`

SetSearchFilters sets SearchFilters field to given value.


### GetRangeFilters

`func (o *RuleConfiguration) GetRangeFilters() []RangeFilter`

GetRangeFilters returns the RangeFilters field if non-nil, zero value otherwise.

### GetRangeFiltersOk

`func (o *RuleConfiguration) GetRangeFiltersOk() (*[]RangeFilter, bool)`

GetRangeFiltersOk returns a tuple with the RangeFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeFilters

`func (o *RuleConfiguration) SetRangeFilters(v []RangeFilter)`

SetRangeFilters sets RangeFilters field to given value.


### GetTemplate

`func (o *RuleConfiguration) GetTemplate() RuleTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *RuleConfiguration) GetTemplateOk() (*RuleTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *RuleConfiguration) SetTemplate(v RuleTemplate)`

SetTemplate sets Template field to given value.


### GetBoostedProductBuckets

`func (o *RuleConfiguration) GetBoostedProductBuckets() []BoostedProductBucket`

GetBoostedProductBuckets returns the BoostedProductBuckets field if non-nil, zero value otherwise.

### GetBoostedProductBucketsOk

`func (o *RuleConfiguration) GetBoostedProductBucketsOk() (*[]BoostedProductBucket, bool)`

GetBoostedProductBucketsOk returns a tuple with the BoostedProductBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostedProductBuckets

`func (o *RuleConfiguration) SetBoostedProductBuckets(v []BoostedProductBucket)`

SetBoostedProductBuckets sets BoostedProductBuckets field to given value.


### SetBoostedProductBucketsNil

`func (o *RuleConfiguration) SetBoostedProductBucketsNil(b bool)`

 SetBoostedProductBucketsNil sets the value for BoostedProductBuckets to be an explicit nil

### UnsetBoostedProductBuckets
`func (o *RuleConfiguration) UnsetBoostedProductBuckets()`

UnsetBoostedProductBuckets ensures that no value is present for BoostedProductBuckets, not even an explicit nil
### GetPinnedRefinements

`func (o *RuleConfiguration) GetPinnedRefinements() []PinnedRefinement`

GetPinnedRefinements returns the PinnedRefinements field if non-nil, zero value otherwise.

### GetPinnedRefinementsOk

`func (o *RuleConfiguration) GetPinnedRefinementsOk() (*[]PinnedRefinement, bool)`

GetPinnedRefinementsOk returns a tuple with the PinnedRefinements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedRefinements

`func (o *RuleConfiguration) SetPinnedRefinements(v []PinnedRefinement)`

SetPinnedRefinements sets PinnedRefinements field to given value.


### GetMessageType

`func (o *RuleConfiguration) GetMessageType() MessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *RuleConfiguration) GetMessageTypeOk() (*MessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *RuleConfiguration) SetMessageType(v MessageType)`

SetMessageType sets MessageType field to given value.


### GetType

`func (o *RuleConfiguration) GetType() RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleConfiguration) GetTypeOk() (*RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleConfiguration) SetType(v RuleType)`

SetType sets Type field to given value.


### GetVariants

`func (o *RuleConfiguration) GetVariants() []ExperimentVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *RuleConfiguration) GetVariantsOk() (*[]ExperimentVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *RuleConfiguration) SetVariants(v []ExperimentVariant)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


