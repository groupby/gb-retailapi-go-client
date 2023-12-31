/*
GroupBy Retail

GroupBy Retail API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"encoding/json"
)

// checks if the SearchResponseDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResponseDto{}

// SearchResponseDto Response of calling the search API, including various elements of the original request context, matching records and general metadata relating to the results.
type SearchResponseDto struct {
	// Unique identifier for the search.
	Id *string `json:"id,omitempty"`
	// Area Id the search was performed in.
	Area *string `json:"area,omitempty"`
	// Original search query.
	Query *string `json:"query,omitempty"`
	// Search query after any changes/corrections are done by the engine.
	CorrectedQuery *string `json:"correctedQuery,omitempty"`
	// Name of the biasing profile which was used to bias products in the search results.
	BiasingProfile *string `json:"biasingProfile,omitempty"`
	// Id of the biasing profile which was used to bias products in the search results.
	BiasingProfileAppliedId *int32 `json:"biasingProfileAppliedId,omitempty"`
	Filter string `json:"filter"`
	OriginalRequest SearchRequestDto `json:"originalRequest"`
	// The list of records that match the search.
	Records []RecordDto `json:"records,omitempty"`
	// The total number of products that match the search. If all products were filtered out on S4R site equals to 0.
	TotalRecordCount *int64 `json:"totalRecordCount,omitempty"`
	Metadata SearchMetadataDto `json:"metadata"`
	PageInfo PageInfoDto `json:"pageInfo"`
	AvailableNavigation []NavigationDto `json:"availableNavigation"`
	SelectedNavigation []NavigationDto `json:"selectedNavigation"`
	// URL to which the merchandiser should redirect the shopper to.
	Redirect *string `json:"redirect,omitempty"`
	Experiments []Experiment `json:"experiments"`
	Template TemplateDto `json:"template"`
	// True if the search yielded no results, otherwise false.
	Empty *bool `json:"empty,omitempty"`
	SiteParams []Metadata `json:"siteParams"`
	Debug DebugDto `json:"debug"`
	// Warning messages containing information about invalid products, etc.
	Warnings []string `json:"warnings,omitempty"`
	// When a shopper uses an ambiguous or a multi-word search phrase, they can get an empty response. After turning on include expanded results, Retail Search analyzes the request and returns the expanded list of products based on the parsed search query. For example, if you search \"Google Pixel 5\" without query expansion, you might only get \"google_pixel_5\" in the result. With query expansion, you might get \"google_pixel_4a_with_5g\", \"google_pixel_4a\" and \"google_pixel_5_case\" as well.The default value is configured in the tenant settings or true if there is no such setting
	IncludeExpandedResults *bool `json:"includeExpandedResults,omitempty"`
	// Maximum of facet values that should be returned for this facet. If not specified, defaults to 20. The maximum allowed value is 300. Values above 300 will be coerced to 300.  If this field is negative, an INVALID_ARGUMENT is returned.  This limit (300) is configured on Google side, but Google have an ability to change it for specific project. 
	FacetLimit *int32 `json:"facetLimit,omitempty"`
	RedirectMetadata []Metadata `json:"redirectMetadata"`
}

// NewSearchResponseDto instantiates a new SearchResponseDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResponseDto(filter string, originalRequest SearchRequestDto, metadata SearchMetadataDto, pageInfo PageInfoDto, availableNavigation []NavigationDto, selectedNavigation []NavigationDto, experiments []Experiment, template TemplateDto, siteParams []Metadata, debug DebugDto, redirectMetadata []Metadata) *SearchResponseDto {
	this := SearchResponseDto{}
	this.Filter = filter
	this.OriginalRequest = originalRequest
	this.Metadata = metadata
	this.PageInfo = pageInfo
	this.AvailableNavigation = availableNavigation
	this.SelectedNavigation = selectedNavigation
	this.Experiments = experiments
	this.Template = template
	this.SiteParams = siteParams
	this.Debug = debug
	this.RedirectMetadata = redirectMetadata
	return &this
}

// NewSearchResponseDtoWithDefaults instantiates a new SearchResponseDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResponseDtoWithDefaults() *SearchResponseDto {
	this := SearchResponseDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SearchResponseDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SearchResponseDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SearchResponseDto) SetId(v string) {
	o.Id = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *SearchResponseDto) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *SearchResponseDto) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *SearchResponseDto) SetArea(v string) {
	o.Area = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SearchResponseDto) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SearchResponseDto) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *SearchResponseDto) SetQuery(v string) {
	o.Query = &v
}

// GetCorrectedQuery returns the CorrectedQuery field value if set, zero value otherwise.
func (o *SearchResponseDto) GetCorrectedQuery() string {
	if o == nil || IsNil(o.CorrectedQuery) {
		var ret string
		return ret
	}
	return *o.CorrectedQuery
}

// GetCorrectedQueryOk returns a tuple with the CorrectedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetCorrectedQueryOk() (*string, bool) {
	if o == nil || IsNil(o.CorrectedQuery) {
		return nil, false
	}
	return o.CorrectedQuery, true
}

// HasCorrectedQuery returns a boolean if a field has been set.
func (o *SearchResponseDto) HasCorrectedQuery() bool {
	if o != nil && !IsNil(o.CorrectedQuery) {
		return true
	}

	return false
}

// SetCorrectedQuery gets a reference to the given string and assigns it to the CorrectedQuery field.
func (o *SearchResponseDto) SetCorrectedQuery(v string) {
	o.CorrectedQuery = &v
}

// GetBiasingProfile returns the BiasingProfile field value if set, zero value otherwise.
func (o *SearchResponseDto) GetBiasingProfile() string {
	if o == nil || IsNil(o.BiasingProfile) {
		var ret string
		return ret
	}
	return *o.BiasingProfile
}

// GetBiasingProfileOk returns a tuple with the BiasingProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetBiasingProfileOk() (*string, bool) {
	if o == nil || IsNil(o.BiasingProfile) {
		return nil, false
	}
	return o.BiasingProfile, true
}

// HasBiasingProfile returns a boolean if a field has been set.
func (o *SearchResponseDto) HasBiasingProfile() bool {
	if o != nil && !IsNil(o.BiasingProfile) {
		return true
	}

	return false
}

// SetBiasingProfile gets a reference to the given string and assigns it to the BiasingProfile field.
func (o *SearchResponseDto) SetBiasingProfile(v string) {
	o.BiasingProfile = &v
}

// GetBiasingProfileAppliedId returns the BiasingProfileAppliedId field value if set, zero value otherwise.
func (o *SearchResponseDto) GetBiasingProfileAppliedId() int32 {
	if o == nil || IsNil(o.BiasingProfileAppliedId) {
		var ret int32
		return ret
	}
	return *o.BiasingProfileAppliedId
}

// GetBiasingProfileAppliedIdOk returns a tuple with the BiasingProfileAppliedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetBiasingProfileAppliedIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BiasingProfileAppliedId) {
		return nil, false
	}
	return o.BiasingProfileAppliedId, true
}

// HasBiasingProfileAppliedId returns a boolean if a field has been set.
func (o *SearchResponseDto) HasBiasingProfileAppliedId() bool {
	if o != nil && !IsNil(o.BiasingProfileAppliedId) {
		return true
	}

	return false
}

// SetBiasingProfileAppliedId gets a reference to the given int32 and assigns it to the BiasingProfileAppliedId field.
func (o *SearchResponseDto) SetBiasingProfileAppliedId(v int32) {
	o.BiasingProfileAppliedId = &v
}

// GetFilter returns the Filter field value
func (o *SearchResponseDto) GetFilter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetFilterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *SearchResponseDto) SetFilter(v string) {
	o.Filter = v
}

// GetOriginalRequest returns the OriginalRequest field value
func (o *SearchResponseDto) GetOriginalRequest() SearchRequestDto {
	if o == nil {
		var ret SearchRequestDto
		return ret
	}

	return o.OriginalRequest
}

// GetOriginalRequestOk returns a tuple with the OriginalRequest field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetOriginalRequestOk() (*SearchRequestDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalRequest, true
}

// SetOriginalRequest sets field value
func (o *SearchResponseDto) SetOriginalRequest(v SearchRequestDto) {
	o.OriginalRequest = v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SearchResponseDto) GetRecords() []RecordDto {
	if o == nil || IsNil(o.Records) {
		var ret []RecordDto
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetRecordsOk() ([]RecordDto, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SearchResponseDto) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []RecordDto and assigns it to the Records field.
func (o *SearchResponseDto) SetRecords(v []RecordDto) {
	o.Records = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *SearchResponseDto) GetTotalRecordCount() int64 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int64
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetTotalRecordCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *SearchResponseDto) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int64 and assigns it to the TotalRecordCount field.
func (o *SearchResponseDto) SetTotalRecordCount(v int64) {
	o.TotalRecordCount = &v
}

// GetMetadata returns the Metadata field value
func (o *SearchResponseDto) GetMetadata() SearchMetadataDto {
	if o == nil {
		var ret SearchMetadataDto
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetMetadataOk() (*SearchMetadataDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SearchResponseDto) SetMetadata(v SearchMetadataDto) {
	o.Metadata = v
}

// GetPageInfo returns the PageInfo field value
func (o *SearchResponseDto) GetPageInfo() PageInfoDto {
	if o == nil {
		var ret PageInfoDto
		return ret
	}

	return o.PageInfo
}

// GetPageInfoOk returns a tuple with the PageInfo field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetPageInfoOk() (*PageInfoDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageInfo, true
}

// SetPageInfo sets field value
func (o *SearchResponseDto) SetPageInfo(v PageInfoDto) {
	o.PageInfo = v
}

// GetAvailableNavigation returns the AvailableNavigation field value
func (o *SearchResponseDto) GetAvailableNavigation() []NavigationDto {
	if o == nil {
		var ret []NavigationDto
		return ret
	}

	return o.AvailableNavigation
}

// GetAvailableNavigationOk returns a tuple with the AvailableNavigation field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetAvailableNavigationOk() ([]NavigationDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvailableNavigation, true
}

// SetAvailableNavigation sets field value
func (o *SearchResponseDto) SetAvailableNavigation(v []NavigationDto) {
	o.AvailableNavigation = v
}

// GetSelectedNavigation returns the SelectedNavigation field value
func (o *SearchResponseDto) GetSelectedNavigation() []NavigationDto {
	if o == nil {
		var ret []NavigationDto
		return ret
	}

	return o.SelectedNavigation
}

// GetSelectedNavigationOk returns a tuple with the SelectedNavigation field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetSelectedNavigationOk() ([]NavigationDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.SelectedNavigation, true
}

// SetSelectedNavigation sets field value
func (o *SearchResponseDto) SetSelectedNavigation(v []NavigationDto) {
	o.SelectedNavigation = v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *SearchResponseDto) GetRedirect() string {
	if o == nil || IsNil(o.Redirect) {
		var ret string
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetRedirectOk() (*string, bool) {
	if o == nil || IsNil(o.Redirect) {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *SearchResponseDto) HasRedirect() bool {
	if o != nil && !IsNil(o.Redirect) {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given string and assigns it to the Redirect field.
func (o *SearchResponseDto) SetRedirect(v string) {
	o.Redirect = &v
}

// GetExperiments returns the Experiments field value
func (o *SearchResponseDto) GetExperiments() []Experiment {
	if o == nil {
		var ret []Experiment
		return ret
	}

	return o.Experiments
}

// GetExperimentsOk returns a tuple with the Experiments field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetExperimentsOk() ([]Experiment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Experiments, true
}

// SetExperiments sets field value
func (o *SearchResponseDto) SetExperiments(v []Experiment) {
	o.Experiments = v
}

// GetTemplate returns the Template field value
func (o *SearchResponseDto) GetTemplate() TemplateDto {
	if o == nil {
		var ret TemplateDto
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetTemplateOk() (*TemplateDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *SearchResponseDto) SetTemplate(v TemplateDto) {
	o.Template = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *SearchResponseDto) GetEmpty() bool {
	if o == nil || IsNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *SearchResponseDto) HasEmpty() bool {
	if o != nil && !IsNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *SearchResponseDto) SetEmpty(v bool) {
	o.Empty = &v
}

// GetSiteParams returns the SiteParams field value
func (o *SearchResponseDto) GetSiteParams() []Metadata {
	if o == nil {
		var ret []Metadata
		return ret
	}

	return o.SiteParams
}

// GetSiteParamsOk returns a tuple with the SiteParams field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetSiteParamsOk() ([]Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.SiteParams, true
}

// SetSiteParams sets field value
func (o *SearchResponseDto) SetSiteParams(v []Metadata) {
	o.SiteParams = v
}

// GetDebug returns the Debug field value
func (o *SearchResponseDto) GetDebug() DebugDto {
	if o == nil {
		var ret DebugDto
		return ret
	}

	return o.Debug
}

// GetDebugOk returns a tuple with the Debug field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetDebugOk() (*DebugDto, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Debug, true
}

// SetDebug sets field value
func (o *SearchResponseDto) SetDebug(v DebugDto) {
	o.Debug = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SearchResponseDto) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SearchResponseDto) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *SearchResponseDto) SetWarnings(v []string) {
	o.Warnings = v
}

// GetIncludeExpandedResults returns the IncludeExpandedResults field value if set, zero value otherwise.
func (o *SearchResponseDto) GetIncludeExpandedResults() bool {
	if o == nil || IsNil(o.IncludeExpandedResults) {
		var ret bool
		return ret
	}
	return *o.IncludeExpandedResults
}

// GetIncludeExpandedResultsOk returns a tuple with the IncludeExpandedResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetIncludeExpandedResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeExpandedResults) {
		return nil, false
	}
	return o.IncludeExpandedResults, true
}

// HasIncludeExpandedResults returns a boolean if a field has been set.
func (o *SearchResponseDto) HasIncludeExpandedResults() bool {
	if o != nil && !IsNil(o.IncludeExpandedResults) {
		return true
	}

	return false
}

// SetIncludeExpandedResults gets a reference to the given bool and assigns it to the IncludeExpandedResults field.
func (o *SearchResponseDto) SetIncludeExpandedResults(v bool) {
	o.IncludeExpandedResults = &v
}

// GetFacetLimit returns the FacetLimit field value if set, zero value otherwise.
func (o *SearchResponseDto) GetFacetLimit() int32 {
	if o == nil || IsNil(o.FacetLimit) {
		var ret int32
		return ret
	}
	return *o.FacetLimit
}

// GetFacetLimitOk returns a tuple with the FacetLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetFacetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.FacetLimit) {
		return nil, false
	}
	return o.FacetLimit, true
}

// HasFacetLimit returns a boolean if a field has been set.
func (o *SearchResponseDto) HasFacetLimit() bool {
	if o != nil && !IsNil(o.FacetLimit) {
		return true
	}

	return false
}

// SetFacetLimit gets a reference to the given int32 and assigns it to the FacetLimit field.
func (o *SearchResponseDto) SetFacetLimit(v int32) {
	o.FacetLimit = &v
}

// GetRedirectMetadata returns the RedirectMetadata field value
func (o *SearchResponseDto) GetRedirectMetadata() []Metadata {
	if o == nil {
		var ret []Metadata
		return ret
	}

	return o.RedirectMetadata
}

// GetRedirectMetadataOk returns a tuple with the RedirectMetadata field value
// and a boolean to check if the value has been set.
func (o *SearchResponseDto) GetRedirectMetadataOk() ([]Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.RedirectMetadata, true
}

// SetRedirectMetadata sets field value
func (o *SearchResponseDto) SetRedirectMetadata(v []Metadata) {
	o.RedirectMetadata = v
}

func (o SearchResponseDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResponseDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !IsNil(o.CorrectedQuery) {
		toSerialize["correctedQuery"] = o.CorrectedQuery
	}
	if !IsNil(o.BiasingProfile) {
		toSerialize["biasingProfile"] = o.BiasingProfile
	}
	if !IsNil(o.BiasingProfileAppliedId) {
		toSerialize["biasingProfileAppliedId"] = o.BiasingProfileAppliedId
	}
	toSerialize["filter"] = o.Filter
	toSerialize["originalRequest"] = o.OriginalRequest
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["totalRecordCount"] = o.TotalRecordCount
	}
	toSerialize["metadata"] = o.Metadata
	toSerialize["pageInfo"] = o.PageInfo
	toSerialize["availableNavigation"] = o.AvailableNavigation
	toSerialize["selectedNavigation"] = o.SelectedNavigation
	if !IsNil(o.Redirect) {
		toSerialize["redirect"] = o.Redirect
	}
	toSerialize["experiments"] = o.Experiments
	toSerialize["template"] = o.Template
	if !IsNil(o.Empty) {
		toSerialize["empty"] = o.Empty
	}
	toSerialize["siteParams"] = o.SiteParams
	toSerialize["debug"] = o.Debug
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.IncludeExpandedResults) {
		toSerialize["includeExpandedResults"] = o.IncludeExpandedResults
	}
	if !IsNil(o.FacetLimit) {
		toSerialize["facetLimit"] = o.FacetLimit
	}
	toSerialize["redirectMetadata"] = o.RedirectMetadata
	return toSerialize, nil
}

type NullableSearchResponseDto struct {
	value *SearchResponseDto
	isSet bool
}

func (v NullableSearchResponseDto) Get() *SearchResponseDto {
	return v.value
}

func (v *NullableSearchResponseDto) Set(val *SearchResponseDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResponseDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResponseDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResponseDto(val *SearchResponseDto) *NullableSearchResponseDto {
	return &NullableSearchResponseDto{value: val, isSet: true}
}

func (v NullableSearchResponseDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResponseDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


