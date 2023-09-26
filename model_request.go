/*
GroupBy Retail

GroupBy Retail API

API version: 0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
)

// checks if the Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Request{}

// Request Object which is represent autocomplete request and encapsulate all passed parameters. 
type Request struct {
	// Area i.e. 'Production' Will not be used immediately. This will be useful when we eventually need to support a US area vs a Canada area. But this would require using the custom dataset instead of user-generated.
	Area string `json:"area"`
	// Name of the collection used to determine the retail backend to use. Usually it is name which is associated with retail project in command center (project configuration).
	Collection string `json:"collection"`
	// Completion max suggestions. The maximum allowed max suggestions is 20.
	SearchItems int32 `json:"searchItems"`
	// String. Required. The query used to generate suggestions. The maximum number of allowed characters is 255.
	Query string `json:"query"`
	// Enable attribute suggestions, by setting the field enableAttributeSuggestion=true in the API request. Then in response, there will be an additional field attributeResults to show direct match results on brands/categories  Note that attribute results directly come from the products we import and Google does not apply any cleaning on them. 
	EnableAttributeSuggestion *bool `json:"enableAttributeSuggestion,omitempty"`
	// Optional param which is define if extended suggestions will be returned in autocomplete response or not. Possibly values: true, false, not specified (If not specified default collection setting will be used). 
	ExtendedSuggestions NullableBool `json:"extendedSuggestions,omitempty"`
	//     List with extended attributes which are would be returned in autocomplete response.     Requires extendedSuggestions to be enabled using search param or on collection layer. 
	ExtendedAttributes []string `json:"extendedAttributes,omitempty"`
	// String. Not required field. A unique identifier for tracking visitors. For example, this could be implemented with an HTTP cookie, which should be able to uniquely identify a visitor on a single device. This unique identifier should not change if the visitor logs in or out of the website. The field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned. 
	VisitorId NullableString `json:"visitorId,omitempty"`
	// Name of site filter. If not specified, the specified area's default site will be applied if configured in Command Center. To not use default specify empty value i.e.\"\".  If the site doesn't exist then the search will execute without the site filter.
	Site NullableString `json:"site,omitempty"`
}

// NewRequest instantiates a new Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequest(area string, collection string, searchItems int32, query string) *Request {
	this := Request{}
	this.Area = area
	this.Collection = collection
	this.SearchItems = searchItems
	this.Query = query
	return &this
}

// NewRequestWithDefaults instantiates a new Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestWithDefaults() *Request {
	this := Request{}
	return &this
}

// GetArea returns the Area field value
func (o *Request) GetArea() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Area
}

// GetAreaOk returns a tuple with the Area field value
// and a boolean to check if the value has been set.
func (o *Request) GetAreaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Area, true
}

// SetArea sets field value
func (o *Request) SetArea(v string) {
	o.Area = v
}

// GetCollection returns the Collection field value
func (o *Request) GetCollection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *Request) GetCollectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *Request) SetCollection(v string) {
	o.Collection = v
}

// GetSearchItems returns the SearchItems field value
func (o *Request) GetSearchItems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SearchItems
}

// GetSearchItemsOk returns a tuple with the SearchItems field value
// and a boolean to check if the value has been set.
func (o *Request) GetSearchItemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchItems, true
}

// SetSearchItems sets field value
func (o *Request) SetSearchItems(v int32) {
	o.SearchItems = v
}

// GetQuery returns the Query field value
func (o *Request) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *Request) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *Request) SetQuery(v string) {
	o.Query = v
}

// GetEnableAttributeSuggestion returns the EnableAttributeSuggestion field value if set, zero value otherwise.
func (o *Request) GetEnableAttributeSuggestion() bool {
	if o == nil || IsNil(o.EnableAttributeSuggestion) {
		var ret bool
		return ret
	}
	return *o.EnableAttributeSuggestion
}

// GetEnableAttributeSuggestionOk returns a tuple with the EnableAttributeSuggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Request) GetEnableAttributeSuggestionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAttributeSuggestion) {
		return nil, false
	}
	return o.EnableAttributeSuggestion, true
}

// HasEnableAttributeSuggestion returns a boolean if a field has been set.
func (o *Request) HasEnableAttributeSuggestion() bool {
	if o != nil && !IsNil(o.EnableAttributeSuggestion) {
		return true
	}

	return false
}

// SetEnableAttributeSuggestion gets a reference to the given bool and assigns it to the EnableAttributeSuggestion field.
func (o *Request) SetEnableAttributeSuggestion(v bool) {
	o.EnableAttributeSuggestion = &v
}

// GetExtendedSuggestions returns the ExtendedSuggestions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetExtendedSuggestions() bool {
	if o == nil || IsNil(o.ExtendedSuggestions.Get()) {
		var ret bool
		return ret
	}
	return *o.ExtendedSuggestions.Get()
}

// GetExtendedSuggestionsOk returns a tuple with the ExtendedSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetExtendedSuggestionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExtendedSuggestions.Get(), o.ExtendedSuggestions.IsSet()
}

// HasExtendedSuggestions returns a boolean if a field has been set.
func (o *Request) HasExtendedSuggestions() bool {
	if o != nil && o.ExtendedSuggestions.IsSet() {
		return true
	}

	return false
}

// SetExtendedSuggestions gets a reference to the given NullableBool and assigns it to the ExtendedSuggestions field.
func (o *Request) SetExtendedSuggestions(v bool) {
	o.ExtendedSuggestions.Set(&v)
}
// SetExtendedSuggestionsNil sets the value for ExtendedSuggestions to be an explicit nil
func (o *Request) SetExtendedSuggestionsNil() {
	o.ExtendedSuggestions.Set(nil)
}

// UnsetExtendedSuggestions ensures that no value is present for ExtendedSuggestions, not even an explicit nil
func (o *Request) UnsetExtendedSuggestions() {
	o.ExtendedSuggestions.Unset()
}

// GetExtendedAttributes returns the ExtendedAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetExtendedAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtendedAttributes
}

// GetExtendedAttributesOk returns a tuple with the ExtendedAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetExtendedAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtendedAttributes) {
		return nil, false
	}
	return o.ExtendedAttributes, true
}

// HasExtendedAttributes returns a boolean if a field has been set.
func (o *Request) HasExtendedAttributes() bool {
	if o != nil && IsNil(o.ExtendedAttributes) {
		return true
	}

	return false
}

// SetExtendedAttributes gets a reference to the given []string and assigns it to the ExtendedAttributes field.
func (o *Request) SetExtendedAttributes(v []string) {
	o.ExtendedAttributes = v
}

// GetVisitorId returns the VisitorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetVisitorId() string {
	if o == nil || IsNil(o.VisitorId.Get()) {
		var ret string
		return ret
	}
	return *o.VisitorId.Get()
}

// GetVisitorIdOk returns a tuple with the VisitorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetVisitorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisitorId.Get(), o.VisitorId.IsSet()
}

// HasVisitorId returns a boolean if a field has been set.
func (o *Request) HasVisitorId() bool {
	if o != nil && o.VisitorId.IsSet() {
		return true
	}

	return false
}

// SetVisitorId gets a reference to the given NullableString and assigns it to the VisitorId field.
func (o *Request) SetVisitorId(v string) {
	o.VisitorId.Set(&v)
}
// SetVisitorIdNil sets the value for VisitorId to be an explicit nil
func (o *Request) SetVisitorIdNil() {
	o.VisitorId.Set(nil)
}

// UnsetVisitorId ensures that no value is present for VisitorId, not even an explicit nil
func (o *Request) UnsetVisitorId() {
	o.VisitorId.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Request) GetSite() string {
	if o == nil || IsNil(o.Site.Get()) {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Request) GetSiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *Request) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *Request) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *Request) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *Request) UnsetSite() {
	o.Site.Unset()
}

func (o Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["area"] = o.Area
	toSerialize["collection"] = o.Collection
	toSerialize["searchItems"] = o.SearchItems
	toSerialize["query"] = o.Query
	if !IsNil(o.EnableAttributeSuggestion) {
		toSerialize["enableAttributeSuggestion"] = o.EnableAttributeSuggestion
	}
	if o.ExtendedSuggestions.IsSet() {
		toSerialize["extendedSuggestions"] = o.ExtendedSuggestions.Get()
	}
	if o.ExtendedAttributes != nil {
		toSerialize["extendedAttributes"] = o.ExtendedAttributes
	}
	if o.VisitorId.IsSet() {
		toSerialize["visitorId"] = o.VisitorId.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	return toSerialize, nil
}

type NullableRequest struct {
	value *Request
	isSet bool
}

func (v NullableRequest) Get() *Request {
	return v.value
}

func (v *NullableRequest) Set(val *Request) {
	v.value = val
	v.isSet = true
}

func (v NullableRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequest(val *Request) *NullableRequest {
	return &NullableRequest{value: val, isSet: true}
}

func (v NullableRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


