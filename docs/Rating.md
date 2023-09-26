# Rating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RatingCount** | Pointer to **int32** | The total number of ratings. This value is independent of the value of histogram.This value must be nonnegative. | [optional] 
**AverageRating** | Pointer to **float32** | The average rating of the product. The rating is scaled at 1-5. | [optional] 
**RatingHistogram** | Pointer to **[]int32** | List of rating counts per rating value (index &#x3D; rating - 1). The list is empty if there is no rating. If the list is non-empty, its size is always 5. For example, [41, 14, 13, 47, 303]. It means that the product got 41 ratings with 1 star, 14 ratings with 2 star, and so on. | [optional] 

## Methods

### NewRating

`func NewRating() *Rating`

NewRating instantiates a new Rating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingWithDefaults

`func NewRatingWithDefaults() *Rating`

NewRatingWithDefaults instantiates a new Rating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRatingCount

`func (o *Rating) GetRatingCount() int32`

GetRatingCount returns the RatingCount field if non-nil, zero value otherwise.

### GetRatingCountOk

`func (o *Rating) GetRatingCountOk() (*int32, bool)`

GetRatingCountOk returns a tuple with the RatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingCount

`func (o *Rating) SetRatingCount(v int32)`

SetRatingCount sets RatingCount field to given value.

### HasRatingCount

`func (o *Rating) HasRatingCount() bool`

HasRatingCount returns a boolean if a field has been set.

### GetAverageRating

`func (o *Rating) GetAverageRating() float32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *Rating) GetAverageRatingOk() (*float32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *Rating) SetAverageRating(v float32)`

SetAverageRating sets AverageRating field to given value.

### HasAverageRating

`func (o *Rating) HasAverageRating() bool`

HasAverageRating returns a boolean if a field has been set.

### GetRatingHistogram

`func (o *Rating) GetRatingHistogram() []int32`

GetRatingHistogram returns the RatingHistogram field if non-nil, zero value otherwise.

### GetRatingHistogramOk

`func (o *Rating) GetRatingHistogramOk() (*[]int32, bool)`

GetRatingHistogramOk returns a tuple with the RatingHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingHistogram

`func (o *Rating) SetRatingHistogram(v []int32)`

SetRatingHistogram sets RatingHistogram field to given value.

### HasRatingHistogram

`func (o *Rating) HasRatingHistogram() bool`

HasRatingHistogram returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


