# RecommendationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** |  | 
**VisitorId** | Pointer to **NullableString** |  | [optional] 
**Limit** | Pointer to **NullableString** |  | [optional] 
**PageSize** | Pointer to **NullableString** |  | [optional] 
**EventType** | Pointer to **NullableString** |  | [optional] 
**LoginId** | Pointer to **NullableString** |  | [optional] 
**ProductID** | Pointer to **[]string** |  | [optional] 
**Fields** | Pointer to **[]string** |  | [optional] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**RawFilter** | Pointer to **NullableString** |  | [optional] 
**Placement** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StrictFiltering** | Pointer to **NullableBool** | The default is true. If strictFiltering true only products that are within the scope of the filter specified. If false, relax the filtering so that the response may contain other products that are outside the scope of the filtering. | [optional] 

## Methods

### NewRecommendationsRequest

`func NewRecommendationsRequest(collection string, ) *RecommendationsRequest`

NewRecommendationsRequest instantiates a new RecommendationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationsRequestWithDefaults

`func NewRecommendationsRequestWithDefaults() *RecommendationsRequest`

NewRecommendationsRequestWithDefaults instantiates a new RecommendationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *RecommendationsRequest) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *RecommendationsRequest) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *RecommendationsRequest) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetVisitorId

`func (o *RecommendationsRequest) GetVisitorId() string`

GetVisitorId returns the VisitorId field if non-nil, zero value otherwise.

### GetVisitorIdOk

`func (o *RecommendationsRequest) GetVisitorIdOk() (*string, bool)`

GetVisitorIdOk returns a tuple with the VisitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorId

`func (o *RecommendationsRequest) SetVisitorId(v string)`

SetVisitorId sets VisitorId field to given value.

### HasVisitorId

`func (o *RecommendationsRequest) HasVisitorId() bool`

HasVisitorId returns a boolean if a field has been set.

### SetVisitorIdNil

`func (o *RecommendationsRequest) SetVisitorIdNil(b bool)`

 SetVisitorIdNil sets the value for VisitorId to be an explicit nil

### UnsetVisitorId
`func (o *RecommendationsRequest) UnsetVisitorId()`

UnsetVisitorId ensures that no value is present for VisitorId, not even an explicit nil
### GetLimit

`func (o *RecommendationsRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecommendationsRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecommendationsRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RecommendationsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *RecommendationsRequest) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *RecommendationsRequest) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetPageSize

`func (o *RecommendationsRequest) GetPageSize() string`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RecommendationsRequest) GetPageSizeOk() (*string, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RecommendationsRequest) SetPageSize(v string)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *RecommendationsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *RecommendationsRequest) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *RecommendationsRequest) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetEventType

`func (o *RecommendationsRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *RecommendationsRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *RecommendationsRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *RecommendationsRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventTypeNil

`func (o *RecommendationsRequest) SetEventTypeNil(b bool)`

 SetEventTypeNil sets the value for EventType to be an explicit nil

### UnsetEventType
`func (o *RecommendationsRequest) UnsetEventType()`

UnsetEventType ensures that no value is present for EventType, not even an explicit nil
### GetLoginId

`func (o *RecommendationsRequest) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *RecommendationsRequest) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *RecommendationsRequest) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.

### HasLoginId

`func (o *RecommendationsRequest) HasLoginId() bool`

HasLoginId returns a boolean if a field has been set.

### SetLoginIdNil

`func (o *RecommendationsRequest) SetLoginIdNil(b bool)`

 SetLoginIdNil sets the value for LoginId to be an explicit nil

### UnsetLoginId
`func (o *RecommendationsRequest) UnsetLoginId()`

UnsetLoginId ensures that no value is present for LoginId, not even an explicit nil
### GetProductID

`func (o *RecommendationsRequest) GetProductID() []string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *RecommendationsRequest) GetProductIDOk() (*[]string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *RecommendationsRequest) SetProductID(v []string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *RecommendationsRequest) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### SetProductIDNil

`func (o *RecommendationsRequest) SetProductIDNil(b bool)`

 SetProductIDNil sets the value for ProductID to be an explicit nil

### UnsetProductID
`func (o *RecommendationsRequest) UnsetProductID()`

UnsetProductID ensures that no value is present for ProductID, not even an explicit nil
### GetFields

`func (o *RecommendationsRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RecommendationsRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RecommendationsRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *RecommendationsRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *RecommendationsRequest) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *RecommendationsRequest) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetFilters

`func (o *RecommendationsRequest) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RecommendationsRequest) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RecommendationsRequest) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RecommendationsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *RecommendationsRequest) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *RecommendationsRequest) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetRawFilter

`func (o *RecommendationsRequest) GetRawFilter() string`

GetRawFilter returns the RawFilter field if non-nil, zero value otherwise.

### GetRawFilterOk

`func (o *RecommendationsRequest) GetRawFilterOk() (*string, bool)`

GetRawFilterOk returns a tuple with the RawFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawFilter

`func (o *RecommendationsRequest) SetRawFilter(v string)`

SetRawFilter sets RawFilter field to given value.

### HasRawFilter

`func (o *RecommendationsRequest) HasRawFilter() bool`

HasRawFilter returns a boolean if a field has been set.

### SetRawFilterNil

`func (o *RecommendationsRequest) SetRawFilterNil(b bool)`

 SetRawFilterNil sets the value for RawFilter to be an explicit nil

### UnsetRawFilter
`func (o *RecommendationsRequest) UnsetRawFilter()`

UnsetRawFilter ensures that no value is present for RawFilter, not even an explicit nil
### GetPlacement

`func (o *RecommendationsRequest) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *RecommendationsRequest) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *RecommendationsRequest) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *RecommendationsRequest) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *RecommendationsRequest) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *RecommendationsRequest) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetName

`func (o *RecommendationsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecommendationsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecommendationsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecommendationsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RecommendationsRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RecommendationsRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStrictFiltering

`func (o *RecommendationsRequest) GetStrictFiltering() bool`

GetStrictFiltering returns the StrictFiltering field if non-nil, zero value otherwise.

### GetStrictFilteringOk

`func (o *RecommendationsRequest) GetStrictFilteringOk() (*bool, bool)`

GetStrictFilteringOk returns a tuple with the StrictFiltering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictFiltering

`func (o *RecommendationsRequest) SetStrictFiltering(v bool)`

SetStrictFiltering sets StrictFiltering field to given value.

### HasStrictFiltering

`func (o *RecommendationsRequest) HasStrictFiltering() bool`

HasStrictFiltering returns a boolean if a field has been set.

### SetStrictFilteringNil

`func (o *RecommendationsRequest) SetStrictFilteringNil(b bool)`

 SetStrictFilteringNil sets the value for StrictFiltering to be an explicit nil

### UnsetStrictFiltering
`func (o *RecommendationsRequest) UnsetStrictFiltering()`

UnsetStrictFiltering ensures that no value is present for StrictFiltering, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


