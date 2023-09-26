# FacetSearchRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Facet** | [**Facet**](Facet.md) |  | 
**OriginalRequest** | [**SearchRequestDto**](SearchRequestDto.md) |  | 

## Methods

### NewFacetSearchRequestDto

`func NewFacetSearchRequestDto(facet Facet, originalRequest SearchRequestDto, ) *FacetSearchRequestDto`

NewFacetSearchRequestDto instantiates a new FacetSearchRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetSearchRequestDtoWithDefaults

`func NewFacetSearchRequestDtoWithDefaults() *FacetSearchRequestDto`

NewFacetSearchRequestDtoWithDefaults instantiates a new FacetSearchRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFacet

`func (o *FacetSearchRequestDto) GetFacet() Facet`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *FacetSearchRequestDto) GetFacetOk() (*Facet, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *FacetSearchRequestDto) SetFacet(v Facet)`

SetFacet sets Facet field to given value.


### GetOriginalRequest

`func (o *FacetSearchRequestDto) GetOriginalRequest() SearchRequestDto`

GetOriginalRequest returns the OriginalRequest field if non-nil, zero value otherwise.

### GetOriginalRequestOk

`func (o *FacetSearchRequestDto) GetOriginalRequestOk() (*SearchRequestDto, bool)`

GetOriginalRequestOk returns a tuple with the OriginalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalRequest

`func (o *FacetSearchRequestDto) SetOriginalRequest(v SearchRequestDto)`

SetOriginalRequest sets OriginalRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


