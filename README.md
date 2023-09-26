# Go API client for gbretailapi

GroupBy Retail API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import gbretailapi "github.com/groupby/gb-retailapi-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), gbretailapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), gbretailapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), gbretailapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), gbretailapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AutocompleteAPI* | [**Autocompletesearch**](docs/AutocompleteAPI.md#autocompletesearch) | **Get** /api/request | 
*ProductAPI* | [**GetByProductIds**](docs/ProductAPI.md#getbyproductids) | **Get** /api/search/product | Provided product search functionality
*RecommendationsAPIAPI* | [**Predict**](docs/RecommendationsAPIAPI.md#predict) | **Post** /api/predict | Provide Recommendations AI functionality.
*RecommendationsAPIAPI* | [**PredictV2**](docs/RecommendationsAPIAPI.md#predictv2) | **Post** /api/recommendation | Provide Recommendations AI functionality.
*SearchAPI* | [**FacetSearch**](docs/SearchAPI.md#facetsearch) | **Post** /api/search/facet | Provided search functionality
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Post** /api/search | Provided search functionality


## Documentation For Models

 - [AdditionalInfo](docs/AdditionalInfo.md)
 - [AttributeSuggestion](docs/AttributeSuggestion.md)
 - [Audience](docs/Audience.md)
 - [BiasDto](docs/BiasDto.md)
 - [BiasDtoStrengthDto](docs/BiasDtoStrengthDto.md)
 - [BiasingProfileDto](docs/BiasingProfileDto.md)
 - [BoostedProductBucket](docs/BoostedProductBucket.md)
 - [ColorInfo](docs/ColorInfo.md)
 - [CustomParameterDto](docs/CustomParameterDto.md)
 - [CustomParameterTrigger](docs/CustomParameterTrigger.md)
 - [DebugDto](docs/DebugDto.md)
 - [ErrorDto](docs/ErrorDto.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Experiment](docs/Experiment.md)
 - [ExperimentVariant](docs/ExperimentVariant.md)
 - [Facet](docs/Facet.md)
 - [FacetSearchRequestDto](docs/FacetSearchRequestDto.md)
 - [FacetSearchResponseDto](docs/FacetSearchResponseDto.md)
 - [FieldMask](docs/FieldMask.md)
 - [Filter](docs/Filter.md)
 - [FilterParameter](docs/FilterParameter.md)
 - [FulfillmentInfo](docs/FulfillmentInfo.md)
 - [Identity](docs/Identity.md)
 - [Image](docs/Image.md)
 - [Interval](docs/Interval.md)
 - [Merchandiser](docs/Merchandiser.md)
 - [MessageType](docs/MessageType.md)
 - [Metadata](docs/Metadata.md)
 - [NavigationDto](docs/NavigationDto.md)
 - [NavigationType](docs/NavigationType.md)
 - [NavigationTypeDto](docs/NavigationTypeDto.md)
 - [Order](docs/Order.md)
 - [Overwrites](docs/Overwrites.md)
 - [PageInfoDto](docs/PageInfoDto.md)
 - [PinnedRefinement](docs/PinnedRefinement.md)
 - [PredictResults](docs/PredictResults.md)
 - [PriceInfo](docs/PriceInfo.md)
 - [PriceInfoPriceEffectiveTime](docs/PriceInfoPriceEffectiveTime.md)
 - [PriceInfoPriceExpireTime](docs/PriceInfoPriceExpireTime.md)
 - [PriceInfoPriceRange](docs/PriceInfoPriceRange.md)
 - [PriceInfoPriceRangeOriginalPrice](docs/PriceInfoPriceRangeOriginalPrice.md)
 - [PriceInfoPriceRangePrice](docs/PriceInfoPriceRangePrice.md)
 - [ProductCustomAttribute](docs/ProductCustomAttribute.md)
 - [ProductDto](docs/ProductDto.md)
 - [ProductDtoAudience](docs/ProductDtoAudience.md)
 - [ProductDtoAvailableTime](docs/ProductDtoAvailableTime.md)
 - [ProductDtoColorInfo](docs/ProductDtoColorInfo.md)
 - [ProductDtoPriceInfo](docs/ProductDtoPriceInfo.md)
 - [ProductDtoPublishTime](docs/ProductDtoPublishTime.md)
 - [ProductDtoRating](docs/ProductDtoRating.md)
 - [ProductDtoRetrievableFields](docs/ProductDtoRetrievableFields.md)
 - [Promotion](docs/Promotion.md)
 - [QueryPatternTrigger](docs/QueryPatternTrigger.md)
 - [QueryPatternTriggerType](docs/QueryPatternTriggerType.md)
 - [Range](docs/Range.md)
 - [RangeFilter](docs/RangeFilter.md)
 - [Rating](docs/Rating.md)
 - [RecommendationsErrorResponse](docs/RecommendationsErrorResponse.md)
 - [RecommendationsRequest](docs/RecommendationsRequest.md)
 - [RecordDto](docs/RecordDto.md)
 - [Refinement](docs/Refinement.md)
 - [RefinementDto](docs/RefinementDto.md)
 - [Request](docs/Request.md)
 - [Role](docs/Role.md)
 - [RuleConfiguration](docs/RuleConfiguration.md)
 - [RuleTemplate](docs/RuleTemplate.md)
 - [RuleTemplateSection](docs/RuleTemplateSection.md)
 - [RuleType](docs/RuleType.md)
 - [RuleVariant](docs/RuleVariant.md)
 - [SearchFilter](docs/SearchFilter.md)
 - [SearchMetadataDto](docs/SearchMetadataDto.md)
 - [SearchRequestDto](docs/SearchRequestDto.md)
 - [SearchRequestDtoOverwrites](docs/SearchRequestDtoOverwrites.md)
 - [SearchResponseDto](docs/SearchResponseDto.md)
 - [SearchResults](docs/SearchResults.md)
 - [SearchResultsStats](docs/SearchResultsStats.md)
 - [SearchTerms](docs/SearchTerms.md)
 - [SelectedRefinementDto](docs/SelectedRefinementDto.md)
 - [SelectedRefinementTrigger](docs/SelectedRefinementTrigger.md)
 - [SelectedRefinementTriggerType](docs/SelectedRefinementTriggerType.md)
 - [Sort](docs/Sort.md)
 - [SortDto](docs/SortDto.md)
 - [SortDtoOrderDto](docs/SortDtoOrderDto.md)
 - [SpellCorrectionMode](docs/SpellCorrectionMode.md)
 - [Stats](docs/Stats.md)
 - [TemplateDto](docs/TemplateDto.md)
 - [TemplateDtoTriggerSet](docs/TemplateDtoTriggerSet.md)
 - [Timestamp](docs/Timestamp.md)
 - [TriggerSet](docs/TriggerSet.md)
 - [ValueFilter](docs/ValueFilter.md)
 - [ValueFilterValueFilterType](docs/ValueFilterValueFilterType.md)
 - [ZoneDto](docs/ZoneDto.md)
 - [ZoneDtoType](docs/ZoneDtoType.md)
 - [ZoneType](docs/ZoneType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ClientKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### GroupByIncEmployee

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



