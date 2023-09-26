# ValueFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field the value applies to. | 
**Value** | **string** | Value to filter on. | 
**NumberValue** | **float64** | Numeric value to filter on. | 
**Exclude** | **bool** | Describing whether the filter is negated or not: color that is NOT red. | 
**Type** | **interface{}** | Determine which field we need to use - value if &#39;TEXTUAL&#39; type or numberValue if &#39;NUMERIC&#39; type. | 

## Methods

### NewValueFilter

`func NewValueFilter(field string, value string, numberValue float64, exclude bool, type_ interface{}, ) *ValueFilter`

NewValueFilter instantiates a new ValueFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueFilterWithDefaults

`func NewValueFilterWithDefaults() *ValueFilter`

NewValueFilterWithDefaults instantiates a new ValueFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ValueFilter) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValueFilter) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValueFilter) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *ValueFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ValueFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ValueFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetNumberValue

`func (o *ValueFilter) GetNumberValue() float64`

GetNumberValue returns the NumberValue field if non-nil, zero value otherwise.

### GetNumberValueOk

`func (o *ValueFilter) GetNumberValueOk() (*float64, bool)`

GetNumberValueOk returns a tuple with the NumberValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValue

`func (o *ValueFilter) SetNumberValue(v float64)`

SetNumberValue sets NumberValue field to given value.


### GetExclude

`func (o *ValueFilter) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ValueFilter) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ValueFilter) SetExclude(v bool)`

SetExclude sets Exclude field to given value.


### GetType

`func (o *ValueFilter) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValueFilter) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValueFilter) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ValueFilter) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ValueFilter) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


