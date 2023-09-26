# QueryPatternTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**QueryPatternTriggerType**](QueryPatternTriggerType.md) |  | 
**Values** | Pointer to **[]string** |  | [optional] 
**Patterns** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewQueryPatternTrigger

`func NewQueryPatternTrigger(type_ QueryPatternTriggerType, ) *QueryPatternTrigger`

NewQueryPatternTrigger instantiates a new QueryPatternTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPatternTriggerWithDefaults

`func NewQueryPatternTriggerWithDefaults() *QueryPatternTrigger`

NewQueryPatternTriggerWithDefaults instantiates a new QueryPatternTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *QueryPatternTrigger) GetType() QueryPatternTriggerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueryPatternTrigger) GetTypeOk() (*QueryPatternTriggerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueryPatternTrigger) SetType(v QueryPatternTriggerType)`

SetType sets Type field to given value.


### GetValues

`func (o *QueryPatternTrigger) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *QueryPatternTrigger) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *QueryPatternTrigger) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *QueryPatternTrigger) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetPatterns

`func (o *QueryPatternTrigger) GetPatterns() []map[string]interface{}`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *QueryPatternTrigger) GetPatternsOk() (*[]map[string]interface{}, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *QueryPatternTrigger) SetPatterns(v []map[string]interface{})`

SetPatterns sets Patterns field to given value.

### HasPatterns

`func (o *QueryPatternTrigger) HasPatterns() bool`

HasPatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


