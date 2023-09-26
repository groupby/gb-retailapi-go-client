# TriggerSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryPatternTriggers** | [**[]QueryPatternTrigger**](QueryPatternTrigger.md) | Query pattern triggers. | 
**SelectedRefinementTriggers** | [**[]SelectedRefinementTrigger**](SelectedRefinementTrigger.md) | Selected refinement triggers. | 
**CustomParameterTriggers** | [**[]CustomParameterTrigger**](CustomParameterTrigger.md) | Custom parameter triggers. | 

## Methods

### NewTriggerSet

`func NewTriggerSet(queryPatternTriggers []QueryPatternTrigger, selectedRefinementTriggers []SelectedRefinementTrigger, customParameterTriggers []CustomParameterTrigger, ) *TriggerSet`

NewTriggerSet instantiates a new TriggerSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerSetWithDefaults

`func NewTriggerSetWithDefaults() *TriggerSet`

NewTriggerSetWithDefaults instantiates a new TriggerSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryPatternTriggers

`func (o *TriggerSet) GetQueryPatternTriggers() []QueryPatternTrigger`

GetQueryPatternTriggers returns the QueryPatternTriggers field if non-nil, zero value otherwise.

### GetQueryPatternTriggersOk

`func (o *TriggerSet) GetQueryPatternTriggersOk() (*[]QueryPatternTrigger, bool)`

GetQueryPatternTriggersOk returns a tuple with the QueryPatternTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryPatternTriggers

`func (o *TriggerSet) SetQueryPatternTriggers(v []QueryPatternTrigger)`

SetQueryPatternTriggers sets QueryPatternTriggers field to given value.


### GetSelectedRefinementTriggers

`func (o *TriggerSet) GetSelectedRefinementTriggers() []SelectedRefinementTrigger`

GetSelectedRefinementTriggers returns the SelectedRefinementTriggers field if non-nil, zero value otherwise.

### GetSelectedRefinementTriggersOk

`func (o *TriggerSet) GetSelectedRefinementTriggersOk() (*[]SelectedRefinementTrigger, bool)`

GetSelectedRefinementTriggersOk returns a tuple with the SelectedRefinementTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRefinementTriggers

`func (o *TriggerSet) SetSelectedRefinementTriggers(v []SelectedRefinementTrigger)`

SetSelectedRefinementTriggers sets SelectedRefinementTriggers field to given value.


### GetCustomParameterTriggers

`func (o *TriggerSet) GetCustomParameterTriggers() []CustomParameterTrigger`

GetCustomParameterTriggers returns the CustomParameterTriggers field if non-nil, zero value otherwise.

### GetCustomParameterTriggersOk

`func (o *TriggerSet) GetCustomParameterTriggersOk() (*[]CustomParameterTrigger, bool)`

GetCustomParameterTriggersOk returns a tuple with the CustomParameterTriggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameterTriggers

`func (o *TriggerSet) SetCustomParameterTriggers(v []CustomParameterTrigger)`

SetCustomParameterTriggers sets CustomParameterTriggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


