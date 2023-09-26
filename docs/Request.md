# Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | **string** | Area i.e. &#39;Production&#39; Will not be used immediately. This will be useful when we eventually need to support a US area vs a Canada area. But this would require using the custom dataset instead of user-generated. | 
**Collection** | **string** | Name of the collection used to determine the retail backend to use. Usually it is name which is associated with retail project in command center (project configuration). | 
**SearchItems** | **int32** | Completion max suggestions. The maximum allowed max suggestions is 20. | 
**Query** | **string** | String. Required. The query used to generate suggestions. The maximum number of allowed characters is 255. | 
**EnableAttributeSuggestion** | Pointer to **bool** | Enable attribute suggestions, by setting the field enableAttributeSuggestion&#x3D;true in the API request. Then in response, there will be an additional field attributeResults to show direct match results on brands/categories  Note that attribute results directly come from the products we import and Google does not apply any cleaning on them.  | [optional] 
**ExtendedSuggestions** | Pointer to **NullableBool** | Optional param which is define if extended suggestions will be returned in autocomplete response or not. Possibly values: true, false, not specified (If not specified default collection setting will be used).  | [optional] 
**ExtendedAttributes** | Pointer to **[]string** |     List with extended attributes which are would be returned in autocomplete response.     Requires extendedSuggestions to be enabled using search param or on collection layer.  | [optional] 
**VisitorId** | Pointer to **NullableString** | String. Not required field. A unique identifier for tracking visitors. For example, this could be implemented with an HTTP cookie, which should be able to uniquely identify a visitor on a single device. This unique identifier should not change if the visitor logs in or out of the website. The field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.  | [optional] 
**Site** | Pointer to **NullableString** | Name of site filter. If not specified, the specified area&#39;s default site will be applied if configured in Command Center. To not use default specify empty value i.e.\&quot;\&quot;.  If the site doesn&#39;t exist then the search will execute without the site filter. | [optional] 

## Methods

### NewRequest

`func NewRequest(area string, collection string, searchItems int32, query string, ) *Request`

NewRequest instantiates a new Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestWithDefaults

`func NewRequestWithDefaults() *Request`

NewRequestWithDefaults instantiates a new Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *Request) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *Request) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *Request) SetArea(v string)`

SetArea sets Area field to given value.


### GetCollection

`func (o *Request) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *Request) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *Request) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetSearchItems

`func (o *Request) GetSearchItems() int32`

GetSearchItems returns the SearchItems field if non-nil, zero value otherwise.

### GetSearchItemsOk

`func (o *Request) GetSearchItemsOk() (*int32, bool)`

GetSearchItemsOk returns a tuple with the SearchItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchItems

`func (o *Request) SetSearchItems(v int32)`

SetSearchItems sets SearchItems field to given value.


### GetQuery

`func (o *Request) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Request) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Request) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetEnableAttributeSuggestion

`func (o *Request) GetEnableAttributeSuggestion() bool`

GetEnableAttributeSuggestion returns the EnableAttributeSuggestion field if non-nil, zero value otherwise.

### GetEnableAttributeSuggestionOk

`func (o *Request) GetEnableAttributeSuggestionOk() (*bool, bool)`

GetEnableAttributeSuggestionOk returns a tuple with the EnableAttributeSuggestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAttributeSuggestion

`func (o *Request) SetEnableAttributeSuggestion(v bool)`

SetEnableAttributeSuggestion sets EnableAttributeSuggestion field to given value.

### HasEnableAttributeSuggestion

`func (o *Request) HasEnableAttributeSuggestion() bool`

HasEnableAttributeSuggestion returns a boolean if a field has been set.

### GetExtendedSuggestions

`func (o *Request) GetExtendedSuggestions() bool`

GetExtendedSuggestions returns the ExtendedSuggestions field if non-nil, zero value otherwise.

### GetExtendedSuggestionsOk

`func (o *Request) GetExtendedSuggestionsOk() (*bool, bool)`

GetExtendedSuggestionsOk returns a tuple with the ExtendedSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSuggestions

`func (o *Request) SetExtendedSuggestions(v bool)`

SetExtendedSuggestions sets ExtendedSuggestions field to given value.

### HasExtendedSuggestions

`func (o *Request) HasExtendedSuggestions() bool`

HasExtendedSuggestions returns a boolean if a field has been set.

### SetExtendedSuggestionsNil

`func (o *Request) SetExtendedSuggestionsNil(b bool)`

 SetExtendedSuggestionsNil sets the value for ExtendedSuggestions to be an explicit nil

### UnsetExtendedSuggestions
`func (o *Request) UnsetExtendedSuggestions()`

UnsetExtendedSuggestions ensures that no value is present for ExtendedSuggestions, not even an explicit nil
### GetExtendedAttributes

`func (o *Request) GetExtendedAttributes() []string`

GetExtendedAttributes returns the ExtendedAttributes field if non-nil, zero value otherwise.

### GetExtendedAttributesOk

`func (o *Request) GetExtendedAttributesOk() (*[]string, bool)`

GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedAttributes

`func (o *Request) SetExtendedAttributes(v []string)`

SetExtendedAttributes sets ExtendedAttributes field to given value.

### HasExtendedAttributes

`func (o *Request) HasExtendedAttributes() bool`

HasExtendedAttributes returns a boolean if a field has been set.

### SetExtendedAttributesNil

`func (o *Request) SetExtendedAttributesNil(b bool)`

 SetExtendedAttributesNil sets the value for ExtendedAttributes to be an explicit nil

### UnsetExtendedAttributes
`func (o *Request) UnsetExtendedAttributes()`

UnsetExtendedAttributes ensures that no value is present for ExtendedAttributes, not even an explicit nil
### GetVisitorId

`func (o *Request) GetVisitorId() string`

GetVisitorId returns the VisitorId field if non-nil, zero value otherwise.

### GetVisitorIdOk

`func (o *Request) GetVisitorIdOk() (*string, bool)`

GetVisitorIdOk returns a tuple with the VisitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitorId

`func (o *Request) SetVisitorId(v string)`

SetVisitorId sets VisitorId field to given value.

### HasVisitorId

`func (o *Request) HasVisitorId() bool`

HasVisitorId returns a boolean if a field has been set.

### SetVisitorIdNil

`func (o *Request) SetVisitorIdNil(b bool)`

 SetVisitorIdNil sets the value for VisitorId to be an explicit nil

### UnsetVisitorId
`func (o *Request) UnsetVisitorId()`

UnsetVisitorId ensures that no value is present for VisitorId, not even an explicit nil
### GetSite

`func (o *Request) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Request) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Request) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *Request) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *Request) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *Request) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


