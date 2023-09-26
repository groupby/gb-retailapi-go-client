# SearchTerms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Suggestion itself. | [optional] 
**AdditionalInfo** | [**AdditionalInfo**](AdditionalInfo.md) |  | 

## Methods

### NewSearchTerms

`func NewSearchTerms(additionalInfo AdditionalInfo, ) *SearchTerms`

NewSearchTerms instantiates a new SearchTerms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchTermsWithDefaults

`func NewSearchTermsWithDefaults() *SearchTerms`

NewSearchTermsWithDefaults instantiates a new SearchTerms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SearchTerms) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchTerms) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchTerms) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchTerms) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *SearchTerms) GetAdditionalInfo() AdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *SearchTerms) GetAdditionalInfoOk() (*AdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *SearchTerms) SetAdditionalInfo(v AdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


