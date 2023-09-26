# Facet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** | Only get facet values that start with the given string prefix. For example, suppose \&quot;categories\&quot; has three values \&quot;Women &gt; Shoe\&quot;, \&quot;Women &gt; Dress\&quot; and \&quot;Men &gt; Shoe\&quot;. If set \&quot;prefixes\&quot; to \&quot;Women\&quot;, the \&quot;categories\&quot; facet will give only \&quot;Women &gt; Shoe\&quot; and \&quot;Women &gt; Dress\&quot;. Only supported on textual fields. Maximum is 10. This field is case-sensitive | [optional] 
**Contains** | Pointer to **string** | Only get facet values that contains the given strings. For example, suppose \&quot;categories\&quot; has three values \&quot;Women &gt; Shoe\&quot;, \&quot;Women &gt; Dress\&quot; and \&quot;Men &gt; Shoe\&quot;. If set \&quot;contains\&quot; to \&quot;Shoe\&quot;, the \&quot;categories\&quot; facet will give only \&quot;Women &gt; Shoe\&quot; and \&quot;Men &gt; Shoe\&quot;. Only supported on textual fields. Maximum is 10. This field is case-sensitive | [optional] 
**DisplayName** | Pointer to **string** | Display name of facet | [optional] 
**Type** | Pointer to [**NavigationType**](NavigationType.md) |  | [optional] 
**NavigationName** | Pointer to **string** | Represents the name of navigation. | [optional] 

## Methods

### NewFacet

`func NewFacet() *Facet`

NewFacet instantiates a new Facet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacetWithDefaults

`func NewFacetWithDefaults() *Facet`

NewFacetWithDefaults instantiates a new Facet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *Facet) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Facet) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Facet) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *Facet) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetContains

`func (o *Facet) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *Facet) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *Facet) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *Facet) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetDisplayName

`func (o *Facet) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Facet) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Facet) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Facet) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *Facet) GetType() NavigationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Facet) GetTypeOk() (*NavigationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Facet) SetType(v NavigationType)`

SetType sets Type field to given value.

### HasType

`func (o *Facet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNavigationName

`func (o *Facet) GetNavigationName() string`

GetNavigationName returns the NavigationName field if non-nil, zero value otherwise.

### GetNavigationNameOk

`func (o *Facet) GetNavigationNameOk() (*string, bool)`

GetNavigationNameOk returns a tuple with the NavigationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationName

`func (o *Facet) SetNavigationName(v string)`

SetNavigationName sets NavigationName field to given value.

### HasNavigationName

`func (o *Facet) HasNavigationName() bool`

HasNavigationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


