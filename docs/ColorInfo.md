# ColorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColorFamilies** | Pointer to **[]string** | Product color families (array). | [optional] 
**Colors** | Pointer to **[]string** | Product color (array). | [optional] 

## Methods

### NewColorInfo

`func NewColorInfo() *ColorInfo`

NewColorInfo instantiates a new ColorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorInfoWithDefaults

`func NewColorInfoWithDefaults() *ColorInfo`

NewColorInfoWithDefaults instantiates a new ColorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColorFamilies

`func (o *ColorInfo) GetColorFamilies() []string`

GetColorFamilies returns the ColorFamilies field if non-nil, zero value otherwise.

### GetColorFamiliesOk

`func (o *ColorInfo) GetColorFamiliesOk() (*[]string, bool)`

GetColorFamiliesOk returns a tuple with the ColorFamilies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorFamilies

`func (o *ColorInfo) SetColorFamilies(v []string)`

SetColorFamilies sets ColorFamilies field to given value.

### HasColorFamilies

`func (o *ColorInfo) HasColorFamilies() bool`

HasColorFamilies returns a boolean if a field has been set.

### GetColors

`func (o *ColorInfo) GetColors() []string`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *ColorInfo) GetColorsOk() (*[]string, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *ColorInfo) SetColors(v []string)`

SetColors sets Colors field to given value.

### HasColors

`func (o *ColorInfo) HasColors() bool`

HasColors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


