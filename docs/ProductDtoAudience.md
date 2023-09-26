# ProductDtoAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genders** | Pointer to **[]string** | The genders of the audience. Strongly encouraged to use the standard values: &#39;male&#39;, &#39;female&#39;, &#39;unisex&#39;. At most 5 values are allowed. Length limit of 128 characters. | [optional] 
**AgeGroups** | Pointer to **[]string** | The age groups of the audience. Strongly encouraged to use the standard values: &#39;newborn&#39; (up to 3 months old), &#39;infant&#39; (3-12 months old), &#39;toddler&#39; (1-5 years old), &#39;kids&#39; (5-13 years old), &#39;adult&#39; (typically teens or older). At most 5 values are allowed. Length limit of 128 characters. | [optional] 

## Methods

### NewProductDtoAudience

`func NewProductDtoAudience() *ProductDtoAudience`

NewProductDtoAudience instantiates a new ProductDtoAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDtoAudienceWithDefaults

`func NewProductDtoAudienceWithDefaults() *ProductDtoAudience`

NewProductDtoAudienceWithDefaults instantiates a new ProductDtoAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenders

`func (o *ProductDtoAudience) GetGenders() []string`

GetGenders returns the Genders field if non-nil, zero value otherwise.

### GetGendersOk

`func (o *ProductDtoAudience) GetGendersOk() (*[]string, bool)`

GetGendersOk returns a tuple with the Genders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenders

`func (o *ProductDtoAudience) SetGenders(v []string)`

SetGenders sets Genders field to given value.

### HasGenders

`func (o *ProductDtoAudience) HasGenders() bool`

HasGenders returns a boolean if a field has been set.

### GetAgeGroups

`func (o *ProductDtoAudience) GetAgeGroups() []string`

GetAgeGroups returns the AgeGroups field if non-nil, zero value otherwise.

### GetAgeGroupsOk

`func (o *ProductDtoAudience) GetAgeGroupsOk() (*[]string, bool)`

GetAgeGroupsOk returns a tuple with the AgeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeGroups

`func (o *ProductDtoAudience) SetAgeGroups(v []string)`

SetAgeGroups sets AgeGroups field to given value.

### HasAgeGroups

`func (o *ProductDtoAudience) HasAgeGroups() bool`

HasAgeGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


