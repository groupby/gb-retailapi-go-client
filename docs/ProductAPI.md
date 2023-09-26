# \ProductAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByProductIds**](ProductAPI.md#GetByProductIds) | **Get** /api/search/product | Provided product search functionality



## GetByProductIds

> ProductDto GetByProductIds(ctx).Collection(collection).ProductId(productId).XGroupbyCustomerId(xGroupbyCustomerId).VariantIds(variantIds).Execute()

Provided product search functionality



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
    collection := "collection_example" // string | Collection name
    productId := "productId_example" // string | Product ID which need to be search
    xGroupbyCustomerId := "xGroupbyCustomerId_example" // string | Required. This parameter will extract from header X-Groupby-Customer-Id. May contain tenant name. Used to define a                           client by its name.
    variantIds := []string{"Inner_example"} // []string | Not required. If the product has variant list and the request specifies the variantIds, requested variants will be the                           first in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductAPI.GetByProductIds(context.Background()).Collection(collection).ProductId(productId).XGroupbyCustomerId(xGroupbyCustomerId).VariantIds(variantIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductAPI.GetByProductIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByProductIds`: ProductDto
    fmt.Fprintf(os.Stdout, "Response from `ProductAPI.GetByProductIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetByProductIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection** | **string** | Collection name | 
 **productId** | **string** | Product ID which need to be search | 
 **xGroupbyCustomerId** | **string** | Required. This parameter will extract from header X-Groupby-Customer-Id. May contain tenant name. Used to define a                           client by its name. | 
 **variantIds** | **[]string** | Not required. If the product has variant list and the request specifies the variantIds, requested variants will be the                           first in the response. | 

### Return type

[**ProductDto**](ProductDto.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

