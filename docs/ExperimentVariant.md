# ExperimentVariant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RuleVariant** | [**RuleVariant**](RuleVariant.md) |  | 
**VariantTriggerPercentage** | **int32** |  | 

## Methods

### NewExperimentVariant

`func NewExperimentVariant(name string, ruleVariant RuleVariant, variantTriggerPercentage int32, ) *ExperimentVariant`

NewExperimentVariant instantiates a new ExperimentVariant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExperimentVariantWithDefaults

`func NewExperimentVariantWithDefaults() *ExperimentVariant`

NewExperimentVariantWithDefaults instantiates a new ExperimentVariant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExperimentVariant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExperimentVariant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExperimentVariant) SetName(v string)`

SetName sets Name field to given value.


### GetRuleVariant

`func (o *ExperimentVariant) GetRuleVariant() RuleVariant`

GetRuleVariant returns the RuleVariant field if non-nil, zero value otherwise.

### GetRuleVariantOk

`func (o *ExperimentVariant) GetRuleVariantOk() (*RuleVariant, bool)`

GetRuleVariantOk returns a tuple with the RuleVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleVariant

`func (o *ExperimentVariant) SetRuleVariant(v RuleVariant)`

SetRuleVariant sets RuleVariant field to given value.


### GetVariantTriggerPercentage

`func (o *ExperimentVariant) GetVariantTriggerPercentage() int32`

GetVariantTriggerPercentage returns the VariantTriggerPercentage field if non-nil, zero value otherwise.

### GetVariantTriggerPercentageOk

`func (o *ExperimentVariant) GetVariantTriggerPercentageOk() (*int32, bool)`

GetVariantTriggerPercentageOk returns a tuple with the VariantTriggerPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantTriggerPercentage

`func (o *ExperimentVariant) SetVariantTriggerPercentage(v int32)`

SetVariantTriggerPercentage sets VariantTriggerPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


