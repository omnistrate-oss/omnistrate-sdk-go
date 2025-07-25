# FleetApplyPendingChangesToHostClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**Id** | **string** | ID of a Host Cluster | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetApplyPendingChangesToHostClusterRequest

`func NewFleetApplyPendingChangesToHostClusterRequest(environmentId string, id string, serviceId string, token string, ) *FleetApplyPendingChangesToHostClusterRequest`

NewFleetApplyPendingChangesToHostClusterRequest instantiates a new FleetApplyPendingChangesToHostClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetApplyPendingChangesToHostClusterRequestWithDefaults

`func NewFleetApplyPendingChangesToHostClusterRequestWithDefaults() *FleetApplyPendingChangesToHostClusterRequest`

NewFleetApplyPendingChangesToHostClusterRequestWithDefaults instantiates a new FleetApplyPendingChangesToHostClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetApplyPendingChangesToHostClusterRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetId

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetApplyPendingChangesToHostClusterRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetApplyPendingChangesToHostClusterRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetApplyPendingChangesToHostClusterRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetApplyPendingChangesToHostClusterRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


