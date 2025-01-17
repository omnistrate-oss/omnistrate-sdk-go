# FleetListSubscriptionRequestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**Status** | Pointer to **string** | The status of the subscription request | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetListSubscriptionRequestsRequest

`func NewFleetListSubscriptionRequestsRequest(environmentId string, serviceId string, token string, ) *FleetListSubscriptionRequestsRequest`

NewFleetListSubscriptionRequestsRequest instantiates a new FleetListSubscriptionRequestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListSubscriptionRequestsRequestWithDefaults

`func NewFleetListSubscriptionRequestsRequestWithDefaults() *FleetListSubscriptionRequestsRequest`

NewFleetListSubscriptionRequestsRequestWithDefaults instantiates a new FleetListSubscriptionRequestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetListSubscriptionRequestsRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetListSubscriptionRequestsRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetListSubscriptionRequestsRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetServiceId

`func (o *FleetListSubscriptionRequestsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetListSubscriptionRequestsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetListSubscriptionRequestsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetStatus

`func (o *FleetListSubscriptionRequestsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetListSubscriptionRequestsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetListSubscriptionRequestsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetListSubscriptionRequestsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *FleetListSubscriptionRequestsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetListSubscriptionRequestsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetListSubscriptionRequestsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


