# TemplateDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the template. | [optional] 
**RuleName** | Pointer to **string** | Name of the rule which may have triggered. | [optional] 
**TriggerSet** | Pointer to [**TemplateDtoTriggerSet**](TemplateDtoTriggerSet.md) |  | [optional] 
**Zones** | Pointer to [**[]ZoneDto**](ZoneDto.md) | Zones for linked template. | [optional] 

## Methods

### NewTemplateDto

`func NewTemplateDto() *TemplateDto`

NewTemplateDto instantiates a new TemplateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateDtoWithDefaults

`func NewTemplateDtoWithDefaults() *TemplateDto`

NewTemplateDtoWithDefaults instantiates a new TemplateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TemplateDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TemplateDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TemplateDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TemplateDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleName

`func (o *TemplateDto) GetRuleName() string`

GetRuleName returns the RuleName field if non-nil, zero value otherwise.

### GetRuleNameOk

`func (o *TemplateDto) GetRuleNameOk() (*string, bool)`

GetRuleNameOk returns a tuple with the RuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleName

`func (o *TemplateDto) SetRuleName(v string)`

SetRuleName sets RuleName field to given value.

### HasRuleName

`func (o *TemplateDto) HasRuleName() bool`

HasRuleName returns a boolean if a field has been set.

### GetTriggerSet

`func (o *TemplateDto) GetTriggerSet() TemplateDtoTriggerSet`

GetTriggerSet returns the TriggerSet field if non-nil, zero value otherwise.

### GetTriggerSetOk

`func (o *TemplateDto) GetTriggerSetOk() (*TemplateDtoTriggerSet, bool)`

GetTriggerSetOk returns a tuple with the TriggerSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSet

`func (o *TemplateDto) SetTriggerSet(v TemplateDtoTriggerSet)`

SetTriggerSet sets TriggerSet field to given value.

### HasTriggerSet

`func (o *TemplateDto) HasTriggerSet() bool`

HasTriggerSet returns a boolean if a field has been set.

### GetZones

`func (o *TemplateDto) GetZones() []ZoneDto`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *TemplateDto) GetZonesOk() (*[]ZoneDto, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *TemplateDto) SetZones(v []ZoneDto)`

SetZones sets Zones field to given value.

### HasZones

`func (o *TemplateDto) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


