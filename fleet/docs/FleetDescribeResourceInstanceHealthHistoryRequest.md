# FleetDescribeResourceInstanceHealthHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** | Exclusive end of the requested period, in RFC3339 format | 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**Limit** | Pointer to **int64** | Maximum number of changes to return in one page | [optional] [default to 100]
**NextPageToken** | Pointer to **string** | Continuation token bound to the provider, instance, and requested period | [optional] 
**ServiceId** | **string** | ID of a Service | 
**StartTime** | **time.Time** | Inclusive start of the requested period, in RFC3339 format | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetDescribeResourceInstanceHealthHistoryRequest

`func NewFleetDescribeResourceInstanceHealthHistoryRequest(endTime time.Time, environmentId string, instanceId string, serviceId string, startTime time.Time, token string, ) *FleetDescribeResourceInstanceHealthHistoryRequest`

NewFleetDescribeResourceInstanceHealthHistoryRequest instantiates a new FleetDescribeResourceInstanceHealthHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeResourceInstanceHealthHistoryRequestWithDefaults

`func NewFleetDescribeResourceInstanceHealthHistoryRequestWithDefaults() *FleetDescribeResourceInstanceHealthHistoryRequest`

NewFleetDescribeResourceInstanceHealthHistoryRequestWithDefaults instantiates a new FleetDescribeResourceInstanceHealthHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetEnvironmentId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetLimit

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextPageToken

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStartTime

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetToken

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetDescribeResourceInstanceHealthHistoryRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


