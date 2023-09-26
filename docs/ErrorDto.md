# ErrorDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackingId** | Pointer to **string** | Identifier used for tracking purposes. | [optional] 
**Method** | Pointer to **string** | HTTP method of the API call which produced the error. | [optional] 
**Path** | Pointer to **string** | HTTP path of the API call which produced the error. | [optional] 
**Message** | Pointer to **string** | Error message. | [optional] 

## Methods

### NewErrorDto

`func NewErrorDto() *ErrorDto`

NewErrorDto instantiates a new ErrorDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDtoWithDefaults

`func NewErrorDtoWithDefaults() *ErrorDto`

NewErrorDtoWithDefaults instantiates a new ErrorDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackingId

`func (o *ErrorDto) GetTrackingId() string`

GetTrackingId returns the TrackingId field if non-nil, zero value otherwise.

### GetTrackingIdOk

`func (o *ErrorDto) GetTrackingIdOk() (*string, bool)`

GetTrackingIdOk returns a tuple with the TrackingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingId

`func (o *ErrorDto) SetTrackingId(v string)`

SetTrackingId sets TrackingId field to given value.

### HasTrackingId

`func (o *ErrorDto) HasTrackingId() bool`

HasTrackingId returns a boolean if a field has been set.

### GetMethod

`func (o *ErrorDto) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ErrorDto) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ErrorDto) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ErrorDto) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPath

`func (o *ErrorDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ErrorDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ErrorDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ErrorDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


