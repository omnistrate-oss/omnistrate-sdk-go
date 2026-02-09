# CountWithRecent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Last7DaysCount** | **int64** | Count of items created or executed in the last 7 days | 
**TotalCount** | **int64** | Total count across the organization for the specified environment type | 

## Methods

### NewCountWithRecent

`func NewCountWithRecent(last7DaysCount int64, totalCount int64, ) *CountWithRecent`

NewCountWithRecent instantiates a new CountWithRecent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountWithRecentWithDefaults

`func NewCountWithRecentWithDefaults() *CountWithRecent`

NewCountWithRecentWithDefaults instantiates a new CountWithRecent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLast7DaysCount

`func (o *CountWithRecent) GetLast7DaysCount() int64`

GetLast7DaysCount returns the Last7DaysCount field if non-nil, zero value otherwise.

### GetLast7DaysCountOk

`func (o *CountWithRecent) GetLast7DaysCountOk() (*int64, bool)`

GetLast7DaysCountOk returns a tuple with the Last7DaysCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast7DaysCount

`func (o *CountWithRecent) SetLast7DaysCount(v int64)`

SetLast7DaysCount sets Last7DaysCount field to given value.


### GetTotalCount

`func (o *CountWithRecent) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CountWithRecent) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CountWithRecent) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


