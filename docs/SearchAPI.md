# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FacetSearch**](SearchAPI.md#FacetSearch) | **Post** /api/search/facet | Provided search functionality
[**Search**](SearchAPI.md#Search) | **Post** /api/search | Provided search functionality



## FacetSearch

> FacetSearchResponseDto FacetSearch(ctx).XGroupbyCustomerId(xGroupbyCustomerId).FacetSearchRequestDto(facetSearchRequestDto).Execute()

Provided search functionality



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/groupby/gb-retailapi-go-client"
)

func main() {
    xGroupbyCustomerId := "xGroupbyCustomerId_example" // string | Custom HTTP header which may contain tenant name. Used to define a client by its name.
    facetSearchRequestDto := *openapiclient.NewFacetSearchRequestDto(*openapiclient.NewFacet(), *openapiclient.NewSearchRequestDto([]openapiclient.SelectedRefinementDto{*openapiclient.NewSelectedRefinementDto("brands", openapiclient.NavigationTypeDto("Value"))}, "TODO", []openapiclient.CustomParameterDto{*openapiclient.NewCustomParameterDto("landing-page", "easter-2021")}, []openapiclient.SortDto{*openapiclient.NewSortDto("rating", "TODO")})) // FacetSearchRequestDto | Request that should be populated to configure a search API call, made by the client on behalf of a shopper. Contains original request and information about facet for which extra keys requested.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.FacetSearch(context.Background()).XGroupbyCustomerId(xGroupbyCustomerId).FacetSearchRequestDto(facetSearchRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.FacetSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FacetSearch`: FacetSearchResponseDto
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.FacetSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacetSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGroupbyCustomerId** | **string** | Custom HTTP header which may contain tenant name. Used to define a client by its name. | 
 **facetSearchRequestDto** | [**FacetSearchRequestDto**](FacetSearchRequestDto.md) | Request that should be populated to configure a search API call, made by the client on behalf of a shopper. Contains original request and information about facet for which extra keys requested. | 

### Return type

[**FacetSearchResponseDto**](FacetSearchResponseDto.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResponseDto Search(ctx).XGroupbyCustomerId(xGroupbyCustomerId).SearchRequestDto(searchRequestDto).Execute()

Provided search functionality



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/groupby/gb-retailapi-go-client"
)

func main() {
    xGroupbyCustomerId := "xGroupbyCustomerId_example" // string | Custom HTTP header which may contain tenant name. Used to define a client by its name.
    searchRequestDto := *openapiclient.NewSearchRequestDto([]openapiclient.SelectedRefinementDto{*openapiclient.NewSelectedRefinementDto("brands", openapiclient.NavigationTypeDto("Value"))}, "TODO", []openapiclient.CustomParameterDto{*openapiclient.NewCustomParameterDto("landing-page", "easter-2021")}, []openapiclient.SortDto{*openapiclient.NewSortDto("rating", "TODO")}) // SearchRequestDto | Request that should be populated to configure a search API call, made by the client on behalf of a shopper.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchAPI.Search(context.Background()).XGroupbyCustomerId(xGroupbyCustomerId).SearchRequestDto(searchRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResponseDto
    fmt.Fprintf(os.Stdout, "Response from `SearchAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGroupbyCustomerId** | **string** | Custom HTTP header which may contain tenant name. Used to define a client by its name. | 
 **searchRequestDto** | [**SearchRequestDto**](SearchRequestDto.md) | Request that should be populated to configure a search API call, made by the client on behalf of a shopper. | 

### Return type

[**SearchResponseDto**](SearchResponseDto.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

