# RuleTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**EnableExactMatching** | **bool** |  | 
**Sections** | [**[]RuleTemplateSection**](RuleTemplateSection.md) |  | 

## Methods

### NewRuleTemplate

`func NewRuleTemplate(name string, enableExactMatching bool, sections []RuleTemplateSection, ) *RuleTemplate`

NewRuleTemplate instantiates a new RuleTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleTemplateWithDefaults

`func NewRuleTemplateWithDefaults() *RuleTemplate`

NewRuleTemplateWithDefaults instantiates a new RuleTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetEnableExactMatching

`func (o *RuleTemplate) GetEnableExactMatching() bool`

GetEnableExactMatching returns the EnableExactMatching field if non-nil, zero value otherwise.

### GetEnableExactMatchingOk

`func (o *RuleTemplate) GetEnableExactMatchingOk() (*bool, bool)`

GetEnableExactMatchingOk returns a tuple with the EnableExactMatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExactMatching

`func (o *RuleTemplate) SetEnableExactMatching(v bool)`

SetEnableExactMatching sets EnableExactMatching field to given value.


### GetSections

`func (o *RuleTemplate) GetSections() []RuleTemplateSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *RuleTemplate) GetSectionsOk() (*[]RuleTemplateSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *RuleTemplate) SetSections(v []RuleTemplateSection)`

SetSections sets Sections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


