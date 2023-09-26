# BiasDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The field the bias refers to. | 
**Content** | Pointer to **NullableString** | The content the field must match for the bias to be applied. | [optional] 
**Strength** | [**BiasDtoStrengthDto**](BiasDtoStrengthDto.md) |  | 

## Methods

### NewBiasDto

`func NewBiasDto(field string, strength BiasDtoStrengthDto, ) *BiasDto`

NewBiasDto instantiates a new BiasDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiasDtoWithDefaults

`func NewBiasDtoWithDefaults() *BiasDto`

NewBiasDtoWithDefaults instantiates a new BiasDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *BiasDto) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *BiasDto) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *BiasDto) SetField(v string)`

SetField sets Field field to given value.


### GetContent

`func (o *BiasDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BiasDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BiasDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BiasDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *BiasDto) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *BiasDto) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetStrength

`func (o *BiasDto) GetStrength() BiasDtoStrengthDto`

GetStrength returns the Strength field if non-nil, zero value otherwise.

### GetStrengthOk

`func (o *BiasDto) GetStrengthOk() (*BiasDtoStrengthDto, bool)`

GetStrengthOk returns a tuple with the Strength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrength

`func (o *BiasDto) SetStrength(v BiasDtoStrengthDto)`

SetStrength sets Strength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


