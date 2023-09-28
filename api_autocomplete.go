/*
GroupBy Retail

GroupBy Retail API

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gbretailapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// AutocompleteAPIService AutocompleteAPI service
type AutocompleteAPIService service

type ApiAutocompletesearchRequest struct {
	ctx context.Context
	ApiService *AutocompleteAPIService
	xGroupbyCustomerId *string
	identity *Identity
	merchandiser *Merchandiser
	request *Request
}

// Header on incoming HTTP requests that is populated by the API gateway and indicates the customer ID.
func (r ApiAutocompletesearchRequest) XGroupbyCustomerId(xGroupbyCustomerId string) ApiAutocompletesearchRequest {
	r.xGroupbyCustomerId = &xGroupbyCustomerId
	return r
}

func (r ApiAutocompletesearchRequest) Identity(identity Identity) ApiAutocompletesearchRequest {
	r.identity = &identity
	return r
}

func (r ApiAutocompletesearchRequest) Merchandiser(merchandiser Merchandiser) ApiAutocompletesearchRequest {
	r.merchandiser = &merchandiser
	return r
}

// Object which is represent autocomplete request and encapsulate all passed parameters. 
func (r ApiAutocompletesearchRequest) Request(request Request) ApiAutocompletesearchRequest {
	r.request = &request
	return r
}

func (r ApiAutocompletesearchRequest) Execute() (*SearchResults, *http.Response, error) {
	return r.ApiService.AutocompletesearchExecute(r)
}

/*
Autocompletesearch Method for Autocompletesearch

A simple request used to get completes the specified prefix with keyword suggestions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAutocompletesearchRequest
*/
func (a *AutocompleteAPIService) Autocompletesearch(ctx context.Context) ApiAutocompletesearchRequest {
	return ApiAutocompletesearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchResults
func (a *AutocompleteAPIService) AutocompletesearchExecute(r ApiAutocompletesearchRequest) (*SearchResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AutocompleteAPIService.Autocompletesearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/request"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xGroupbyCustomerId == nil {
		return localVarReturnValue, nil, reportError("xGroupbyCustomerId is required and must be specified")
	}
	if r.identity == nil {
		return localVarReturnValue, nil, reportError("identity is required and must be specified")
	}
	if r.merchandiser == nil {
		return localVarReturnValue, nil, reportError("merchandiser is required and must be specified")
	}

	if r.request != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "request", r.request, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "identity", r.identity, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "merchandiser", r.merchandiser, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Groupby-Customer-Id", r.xGroupbyCustomerId, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ClientKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
