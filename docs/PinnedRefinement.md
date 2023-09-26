# PinnedRefinement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Navigation** | **string** |  | 
**Refinements** | [**[]Refinement**](Refinement.md) |  | 

## Methods

### NewPinnedRefinement

`func NewPinnedRefinement(navigation string, refinements []Refinement, ) *PinnedRefinement`

NewPinnedRefinement instantiates a new PinnedRefinement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedRefinementWithDefaults

`func NewPinnedRefinementWithDefaults() *PinnedRefinement`

NewPinnedRefinementWithDefaults instantiates a new PinnedRefinement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNavigation

`func (o *PinnedRefinement) GetNavigation() string`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *PinnedRefinement) GetNavigationOk() (*string, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *PinnedRefinement) SetNavigation(v string)`

SetNavigation sets Navigation field to given value.


### GetRefinements

`func (o *PinnedRefinement) GetRefinements() []Refinement`

GetRefinements returns the Refinements field if non-nil, zero value otherwise.

### GetRefinementsOk

`func (o *PinnedRefinement) GetRefinementsOk() (*[]Refinement, bool)`

GetRefinementsOk returns a tuple with the Refinements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefinements

`func (o *PinnedRefinement) SetRefinements(v []Refinement)`

SetRefinements sets Refinements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


