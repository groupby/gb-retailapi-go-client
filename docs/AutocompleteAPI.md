# \AutocompleteAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Autocompletesearch**](AutocompleteAPI.md#Autocompletesearch) | **Get** /api/request | 



## Autocompletesearch

> SearchResults Autocompletesearch(ctx).XGroupbyCustomerId(xGroupbyCustomerId).Identity(identity).Merchandiser(merchandiser).Request(request).Execute()





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
    xGroupbyCustomerId := "xGroupbyCustomerId_example" // string | Header on incoming HTTP requests that is populated by the API gateway and indicates the customer ID.
    identity := *openapiclient.NewIdentity("Subject_example", "Company_example", []openapiclient.Role{*openapiclient.NewRole("Name_example")}) // Identity | 
    merchandiser := *openapiclient.NewMerchandiser("MerchandiserId_example") // Merchandiser | 
    request := *openapiclient.NewRequest("Area_example", "Collection_example", int32(123), "Query_example") // Request | Object which is represent autocomplete request and encapsulate all passed parameters.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutocompleteAPI.Autocompletesearch(context.Background()).XGroupbyCustomerId(xGroupbyCustomerId).Identity(identity).Merchandiser(merchandiser).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutocompleteAPI.Autocompletesearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Autocompletesearch`: SearchResults
    fmt.Fprintf(os.Stdout, "Response from `AutocompleteAPI.Autocompletesearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAutocompletesearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xGroupbyCustomerId** | **string** | Header on incoming HTTP requests that is populated by the API gateway and indicates the customer ID. | 
 **identity** | [**Identity**](Identity.md) |  | 
 **merchandiser** | [**Merchandiser**](Merchandiser.md) |  | 
 **request** | [**Request**](Request.md) | Object which is represent autocomplete request and encapsulate all passed parameters.  | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[GroupByIncEmployee](../README.md#GroupByIncEmployee), [ClientKey](../README.md#ClientKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

