# ProductCustomAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **[]string** | The textual values of this custom attribute. At most 400 values are allowed. Empty values are not allowed. Length limit of 256 characters. Exactly one of text or numbers should be set. | [optional] 
**Numbers** | Pointer to **[]float64** | The numerical values of this custom attribute. At most 400 values are allowed. Exactly one of text or numbers should be set. | [optional] 
**Searchable** | Pointer to **bool** | If true, custom attribute values are searchable by text queries in. search. Only set if type text set. | [optional] 
**Indexable** | Pointer to **bool** | If true, custom attribute values are indexed, so that it can be filtered, faceted or boosted in search. | [optional] 

## Methods

### NewProductCustomAttribute

`func NewProductCustomAttribute() *ProductCustomAttribute`

NewProductCustomAttribute instantiates a new ProductCustomAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductCustomAttributeWithDefaults

`func NewProductCustomAttributeWithDefaults() *ProductCustomAttribute`

NewProductCustomAttributeWithDefaults instantiates a new ProductCustomAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ProductCustomAttribute) GetText() []string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ProductCustomAttribute) GetTextOk() (*[]string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ProductCustomAttribute) SetText(v []string)`

SetText sets Text field to given value.

### HasText

`func (o *ProductCustomAttribute) HasText() bool`

HasText returns a boolean if a field has been set.

### GetNumbers

`func (o *ProductCustomAttribute) GetNumbers() []float64`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ProductCustomAttribute) GetNumbersOk() (*[]float64, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ProductCustomAttribute) SetNumbers(v []float64)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *ProductCustomAttribute) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetSearchable

`func (o *ProductCustomAttribute) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *ProductCustomAttribute) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *ProductCustomAttribute) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *ProductCustomAttribute) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetIndexable

`func (o *ProductCustomAttribute) GetIndexable() bool`

GetIndexable returns the Indexable field if non-nil, zero value otherwise.

### GetIndexableOk

`func (o *ProductCustomAttribute) GetIndexableOk() (*bool, bool)`

GetIndexableOk returns a tuple with the Indexable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexable

`func (o *ProductCustomAttribute) SetIndexable(v bool)`

SetIndexable sets Indexable field to given value.

### HasIndexable

`func (o *ProductCustomAttribute) HasIndexable() bool`

HasIndexable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


