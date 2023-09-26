# ZoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the zone, ideally human-readable. | [optional] 
**Type** | Pointer to [**ZoneDtoType**](ZoneDtoType.md) |  | [optional] 
**Content** | Pointer to **string** | Zone content - it is can be any data, HTML - code, usual text or etc | [optional] 
**RichContent** | Pointer to **string** | Zone content - it is can be any data, HTML - code, usual text or etc | [optional] 

## Methods

### NewZoneDto

`func NewZoneDto() *ZoneDto`

NewZoneDto instantiates a new ZoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneDtoWithDefaults

`func NewZoneDtoWithDefaults() *ZoneDto`

NewZoneDtoWithDefaults instantiates a new ZoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ZoneDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ZoneDto) GetType() ZoneDtoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneDto) GetTypeOk() (*ZoneDtoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneDto) SetType(v ZoneDtoType)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContent

`func (o *ZoneDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ZoneDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ZoneDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ZoneDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRichContent

`func (o *ZoneDto) GetRichContent() string`

GetRichContent returns the RichContent field if non-nil, zero value otherwise.

### GetRichContentOk

`func (o *ZoneDto) GetRichContentOk() (*string, bool)`

GetRichContentOk returns a tuple with the RichContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichContent

`func (o *ZoneDto) SetRichContent(v string)`

SetRichContent sets RichContent field to given value.

### HasRichContent

`func (o *ZoneDto) HasRichContent() bool`

HasRichContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


