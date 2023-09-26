# FieldMask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to **[]string** | Paths for retrievable fields (array). | [optional] 

## Methods

### NewFieldMask

`func NewFieldMask() *FieldMask`

NewFieldMask instantiates a new FieldMask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMaskWithDefaults

`func NewFieldMaskWithDefaults() *FieldMask`

NewFieldMaskWithDefaults instantiates a new FieldMask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *FieldMask) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *FieldMask) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *FieldMask) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *FieldMask) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


