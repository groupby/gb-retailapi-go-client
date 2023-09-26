# RecordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of the record as an MD5 hash. | [optional] 
**U** | Pointer to **string** | Logical URL of the record. | [optional] 
**T** | Pointer to **string** | Title of the record. | [optional] 
**Collection** | Pointer to **string** | Collection the record was queried from or resides. | [optional] 
**AllMeta** | Pointer to **map[string]interface{}** | All other metadata, read fields, the record has. | [optional] 

## Methods

### NewRecordDto

`func NewRecordDto() *RecordDto`

NewRecordDto instantiates a new RecordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordDtoWithDefaults

`func NewRecordDtoWithDefaults() *RecordDto`

NewRecordDtoWithDefaults instantiates a new RecordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecordDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecordDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetU

`func (o *RecordDto) GetU() string`

GetU returns the U field if non-nil, zero value otherwise.

### GetUOk

`func (o *RecordDto) GetUOk() (*string, bool)`

GetUOk returns a tuple with the U field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetU

`func (o *RecordDto) SetU(v string)`

SetU sets U field to given value.

### HasU

`func (o *RecordDto) HasU() bool`

HasU returns a boolean if a field has been set.

### GetT

`func (o *RecordDto) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *RecordDto) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *RecordDto) SetT(v string)`

SetT sets T field to given value.

### HasT

`func (o *RecordDto) HasT() bool`

HasT returns a boolean if a field has been set.

### GetCollection

`func (o *RecordDto) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *RecordDto) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *RecordDto) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *RecordDto) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetAllMeta

`func (o *RecordDto) GetAllMeta() map[string]interface{}`

GetAllMeta returns the AllMeta field if non-nil, zero value otherwise.

### GetAllMetaOk

`func (o *RecordDto) GetAllMetaOk() (*map[string]interface{}, bool)`

GetAllMetaOk returns a tuple with the AllMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllMeta

`func (o *RecordDto) SetAllMeta(v map[string]interface{})`

SetAllMeta sets AllMeta field to given value.

### HasAllMeta

`func (o *RecordDto) HasAllMeta() bool`

HasAllMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


