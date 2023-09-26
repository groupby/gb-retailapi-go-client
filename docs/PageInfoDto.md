# PageInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordStart** | Pointer to **int64** | Where in the list of products the page begins. | [optional] 
**RecordEnd** | Pointer to **int64** | Where in the list of products the page ends. | [optional] 

## Methods

### NewPageInfoDto

`func NewPageInfoDto() *PageInfoDto`

NewPageInfoDto instantiates a new PageInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoDtoWithDefaults

`func NewPageInfoDtoWithDefaults() *PageInfoDto`

NewPageInfoDtoWithDefaults instantiates a new PageInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordStart

`func (o *PageInfoDto) GetRecordStart() int64`

GetRecordStart returns the RecordStart field if non-nil, zero value otherwise.

### GetRecordStartOk

`func (o *PageInfoDto) GetRecordStartOk() (*int64, bool)`

GetRecordStartOk returns a tuple with the RecordStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStart

`func (o *PageInfoDto) SetRecordStart(v int64)`

SetRecordStart sets RecordStart field to given value.

### HasRecordStart

`func (o *PageInfoDto) HasRecordStart() bool`

HasRecordStart returns a boolean if a field has been set.

### GetRecordEnd

`func (o *PageInfoDto) GetRecordEnd() int64`

GetRecordEnd returns the RecordEnd field if non-nil, zero value otherwise.

### GetRecordEndOk

`func (o *PageInfoDto) GetRecordEndOk() (*int64, bool)`

GetRecordEndOk returns a tuple with the RecordEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordEnd

`func (o *PageInfoDto) SetRecordEnd(v int64)`

SetRecordEnd sets RecordEnd field to given value.

### HasRecordEnd

`func (o *PageInfoDto) HasRecordEnd() bool`

HasRecordEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


