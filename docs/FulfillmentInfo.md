# FulfillmentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Fulfillment type. Place where product fulfilled. | [optional] 
**PlaceIds** | Pointer to **[]string** | Place ids where product fulfilled (array). | [optional] 

## Methods

### NewFulfillmentInfo

`func NewFulfillmentInfo() *FulfillmentInfo`

NewFulfillmentInfo instantiates a new FulfillmentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFulfillmentInfoWithDefaults

`func NewFulfillmentInfoWithDefaults() *FulfillmentInfo`

NewFulfillmentInfoWithDefaults instantiates a new FulfillmentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FulfillmentInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FulfillmentInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FulfillmentInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FulfillmentInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPlaceIds

`func (o *FulfillmentInfo) GetPlaceIds() []string`

GetPlaceIds returns the PlaceIds field if non-nil, zero value otherwise.

### GetPlaceIdsOk

`func (o *FulfillmentInfo) GetPlaceIdsOk() (*[]string, bool)`

GetPlaceIdsOk returns a tuple with the PlaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceIds

`func (o *FulfillmentInfo) SetPlaceIds(v []string)`

SetPlaceIds sets PlaceIds field to given value.

### HasPlaceIds

`func (o *FulfillmentInfo) HasPlaceIds() bool`

HasPlaceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


