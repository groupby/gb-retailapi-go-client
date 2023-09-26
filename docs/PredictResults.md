# PredictResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to **[]map[string]interface{}** | Warnings collected with validation and Recommendations AI API issues. | [optional] 
**Products** | Pointer to **[]map[string]interface{}** | Recommendations built by Recommendations AI model. | [optional] 
**Records** | Pointer to **[]map[string]interface{}** | Recommendations built by Recommendations AI model. | [optional] 
**ModelId** | Pointer to **string** | Model Id used for predictions | [optional] 
**ModelName** | Pointer to **string** | Model Name used for predictions | [optional] 
**ModelType** | Pointer to **string** |   Currently supported values:   &#x60;recommended-for-you&#x60;   &#x60;others-you-may-like&#x60;,   &#x60;frequently-bought-together&#x60;   &#x60;page-optimization&#x60;   &#x60;similar-items&#x60;,   &#x60;buy-it-again&#x60;   &#x60;on-sale-items&#x60;   &#x60;recently-viewed&#x60;    This field together with optimization_objective describe model metadata to use to control model training and   serving. See https://cloud.google.com/retail/docs/models for more details.  | [optional] 
**OptimizationObjective** | Pointer to **string** |   Currently supported values: &#x60;ctr&#x60;, &#x60;cvr&#x60;, &#x60;revenue-per-order&#x60;.     If not specified, we choose default based on model type. Default depends on type of recommendation:   &#x60;recommended-for-you&#x60; &#x3D;&gt; &#x60;ctr&#x60;   &#x60;others-you-may-like&#x60; &#x3D;&gt; &#x60;ctr&#x60;   &#x60;frequently-bought-together&#x60; &#x3D;&gt; &#x60;revenue_per_order&#x60;    This field together with modelType describe model metadata to use to control model training and serving.   See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which   combination of parameters are valid.  | [optional] 
**FilterSet** | Pointer to **string** | Filter set applied to the recommendation | [optional] 
**RawFilter** | Pointer to **string** | RawFilter applied to the recommendation | [optional] 
**Filters** | Pointer to [**[]FilterParameter**](FilterParameter.md) | Filters applied to the recommendation | [optional] 
**Metadata** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewPredictResults

`func NewPredictResults() *PredictResults`

NewPredictResults instantiates a new PredictResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictResultsWithDefaults

`func NewPredictResultsWithDefaults() *PredictResults`

NewPredictResultsWithDefaults instantiates a new PredictResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *PredictResults) GetWarnings() []map[string]interface{}`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *PredictResults) GetWarningsOk() (*[]map[string]interface{}, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *PredictResults) SetWarnings(v []map[string]interface{})`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *PredictResults) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetProducts

`func (o *PredictResults) GetProducts() []map[string]interface{}`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *PredictResults) GetProductsOk() (*[]map[string]interface{}, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *PredictResults) SetProducts(v []map[string]interface{})`

SetProducts sets Products field to given value.

### HasProducts

`func (o *PredictResults) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRecords

`func (o *PredictResults) GetRecords() []map[string]interface{}`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *PredictResults) GetRecordsOk() (*[]map[string]interface{}, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *PredictResults) SetRecords(v []map[string]interface{})`

SetRecords sets Records field to given value.

### HasRecords

`func (o *PredictResults) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetModelId

`func (o *PredictResults) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *PredictResults) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *PredictResults) SetModelId(v string)`

SetModelId sets ModelId field to given value.

### HasModelId

`func (o *PredictResults) HasModelId() bool`

HasModelId returns a boolean if a field has been set.

### GetModelName

`func (o *PredictResults) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *PredictResults) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *PredictResults) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *PredictResults) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelType

`func (o *PredictResults) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *PredictResults) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *PredictResults) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *PredictResults) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetOptimizationObjective

`func (o *PredictResults) GetOptimizationObjective() string`

GetOptimizationObjective returns the OptimizationObjective field if non-nil, zero value otherwise.

### GetOptimizationObjectiveOk

`func (o *PredictResults) GetOptimizationObjectiveOk() (*string, bool)`

GetOptimizationObjectiveOk returns a tuple with the OptimizationObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationObjective

`func (o *PredictResults) SetOptimizationObjective(v string)`

SetOptimizationObjective sets OptimizationObjective field to given value.

### HasOptimizationObjective

`func (o *PredictResults) HasOptimizationObjective() bool`

HasOptimizationObjective returns a boolean if a field has been set.

### GetFilterSet

`func (o *PredictResults) GetFilterSet() string`

GetFilterSet returns the FilterSet field if non-nil, zero value otherwise.

### GetFilterSetOk

`func (o *PredictResults) GetFilterSetOk() (*string, bool)`

GetFilterSetOk returns a tuple with the FilterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSet

`func (o *PredictResults) SetFilterSet(v string)`

SetFilterSet sets FilterSet field to given value.

### HasFilterSet

`func (o *PredictResults) HasFilterSet() bool`

HasFilterSet returns a boolean if a field has been set.

### GetRawFilter

`func (o *PredictResults) GetRawFilter() string`

GetRawFilter returns the RawFilter field if non-nil, zero value otherwise.

### GetRawFilterOk

`func (o *PredictResults) GetRawFilterOk() (*string, bool)`

GetRawFilterOk returns a tuple with the RawFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawFilter

`func (o *PredictResults) SetRawFilter(v string)`

SetRawFilter sets RawFilter field to given value.

### HasRawFilter

`func (o *PredictResults) HasRawFilter() bool`

HasRawFilter returns a boolean if a field has been set.

### GetFilters

`func (o *PredictResults) GetFilters() []FilterParameter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PredictResults) GetFiltersOk() (*[]FilterParameter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PredictResults) SetFilters(v []FilterParameter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PredictResults) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetadata

`func (o *PredictResults) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PredictResults) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PredictResults) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PredictResults) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


