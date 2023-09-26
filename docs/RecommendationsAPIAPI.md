# \RecommendationsAPIAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Predict**](RecommendationsAPIAPI.md#Predict) | **Post** /api/predict | Provide Recommendations AI functionality.
[**PredictV2**](RecommendationsAPIAPI.md#PredictV2) | **Post** /api/recommendation | Provide Recommendations AI functionality.



## Predict

> PredictResults Predict(ctx).XGroupbyCustomerId(xGroupbyCustomerId).RecommendationsRequest(recommendationsRequest).Execute()

Provide Recommendations AI functionality.



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
    recommendationsRequest := *openapiclient.NewRecommendationsRequest("Collection_example") // RecommendationsRequest | Request that should be populated to configure a recommendations API call made by the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecommendationsAPIAPI.Predict(context.Background()).XGroupbyCustomerId(xGroupbyCustomerId).RecommendationsRequest(recommendationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecommendationsAPIAPI.Predict``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Predict`: PredictResults
    fmt.Fprintf(os.Stdout, "Response from `RecommendationsAPIAPI.Predict`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredictRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGroupbyCustomerId** | **string** | Custom HTTP header which may contain tenant name. Used to define a client by its name. | 
 **recommendationsRequest** | [**RecommendationsRequest**](RecommendationsRequest.md) | Request that should be populated to configure a recommendations API call made by the client. | 

### Return type

[**PredictResults**](PredictResults.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PredictV2

> PredictResults PredictV2(ctx).XGroupbyCustomerId(xGroupbyCustomerId).RecommendationsRequest(recommendationsRequest).Execute()

Provide Recommendations AI functionality.



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
    recommendationsRequest := *openapiclient.NewRecommendationsRequest("Collection_example") // RecommendationsRequest | Request that should be populated to configure a recommendations API call made by the client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecommendationsAPIAPI.PredictV2(context.Background()).XGroupbyCustomerId(xGroupbyCustomerId).RecommendationsRequest(recommendationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecommendationsAPIAPI.PredictV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PredictV2`: PredictResults
    fmt.Fprintf(os.Stdout, "Response from `RecommendationsAPIAPI.PredictV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPredictV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGroupbyCustomerId** | **string** | Custom HTTP header which may contain tenant name. Used to define a client by its name. | 
 **recommendationsRequest** | [**RecommendationsRequest**](RecommendationsRequest.md) | Request that should be populated to configure a recommendations API call made by the client. | 

### Return type

[**PredictResults**](PredictResults.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

