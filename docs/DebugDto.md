# DebugDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RawSearchRequest** | Pointer to **[]map[string]interface{}** | Search request in raw format executed internally against search engine. | [optional] 
**RawSearchResponse** | Pointer to **[]map[string]interface{}** | Search response in raw format received internally from search engine. | [optional] 

## Methods

### NewDebugDto

`func NewDebugDto() *DebugDto`

NewDebugDto instantiates a new DebugDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDebugDtoWithDefaults

`func NewDebugDtoWithDefaults() *DebugDto`

NewDebugDtoWithDefaults instantiates a new DebugDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRawSearchRequest

`func (o *DebugDto) GetRawSearchRequest() []map[string]interface{}`

GetRawSearchRequest returns the RawSearchRequest field if non-nil, zero value otherwise.

### GetRawSearchRequestOk

`func (o *DebugDto) GetRawSearchRequestOk() (*[]map[string]interface{}, bool)`

GetRawSearchRequestOk returns a tuple with the RawSearchRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSearchRequest

`func (o *DebugDto) SetRawSearchRequest(v []map[string]interface{})`

SetRawSearchRequest sets RawSearchRequest field to given value.

### HasRawSearchRequest

`func (o *DebugDto) HasRawSearchRequest() bool`

HasRawSearchRequest returns a boolean if a field has been set.

### GetRawSearchResponse

`func (o *DebugDto) GetRawSearchResponse() []map[string]interface{}`

GetRawSearchResponse returns the RawSearchResponse field if non-nil, zero value otherwise.

### GetRawSearchResponseOk

`func (o *DebugDto) GetRawSearchResponseOk() (*[]map[string]interface{}, bool)`

GetRawSearchResponseOk returns a tuple with the RawSearchResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSearchResponse

`func (o *DebugDto) SetRawSearchResponse(v []map[string]interface{})`

SetRawSearchResponse sets RawSearchResponse field to given value.

### HasRawSearchResponse

`func (o *DebugDto) HasRawSearchResponse() bool`

HasRawSearchResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


