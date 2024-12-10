# FleetAddCustomDNSToResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDNS** | **string** | The custom DNS to add | 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ResourceKey** | **string** | The resource key | 
**ServiceId** | **string** | ID of a Service | 
**TargetPort** | Pointer to **int64** | The target port | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetAddCustomDNSToResourceInstanceRequest

`func NewFleetAddCustomDNSToResourceInstanceRequest(customDNS string, environmentId string, instanceId string, resourceKey string, serviceId string, token string, ) *FleetAddCustomDNSToResourceInstanceRequest`

NewFleetAddCustomDNSToResourceInstanceRequest instantiates a new FleetAddCustomDNSToResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAddCustomDNSToResourceInstanceRequestWithDefaults

`func NewFleetAddCustomDNSToResourceInstanceRequestWithDefaults() *FleetAddCustomDNSToResourceInstanceRequest`

NewFleetAddCustomDNSToResourceInstanceRequestWithDefaults instantiates a new FleetAddCustomDNSToResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDNS

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetCustomDNS() string`

GetCustomDNS returns the CustomDNS field if non-nil, zero value otherwise.

### GetCustomDNSOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetCustomDNSOk() (*string, bool)`

GetCustomDNSOk returns a tuple with the CustomDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDNS

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetCustomDNS(v string)`

SetCustomDNS sets CustomDNS field to given value.


### GetEnvironmentId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetResourceKey

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetPort

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetTargetPort() int64`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetTargetPortOk() (*int64, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetTargetPort(v int64)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *FleetAddCustomDNSToResourceInstanceRequest) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetToken

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetAddCustomDNSToResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetAddCustomDNSToResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


