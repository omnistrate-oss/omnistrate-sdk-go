# FleetAddCapacityToResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityToBeAdded** | **int64** | Number of replicas to be added | 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetAddCapacityToResourceInstanceRequest

`func NewFleetAddCapacityToResourceInstanceRequest(capacityToBeAdded int64, environmentId string, instanceId string, resourceId string, serviceId string, token string, ) *FleetAddCapacityToResourceInstanceRequest`

NewFleetAddCapacityToResourceInstanceRequest instantiates a new FleetAddCapacityToResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAddCapacityToResourceInstanceRequestWithDefaults

`func NewFleetAddCapacityToResourceInstanceRequestWithDefaults() *FleetAddCapacityToResourceInstanceRequest`

NewFleetAddCapacityToResourceInstanceRequestWithDefaults instantiates a new FleetAddCapacityToResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityToBeAdded

`func (o *FleetAddCapacityToResourceInstanceRequest) GetCapacityToBeAdded() int64`

GetCapacityToBeAdded returns the CapacityToBeAdded field if non-nil, zero value otherwise.

### GetCapacityToBeAddedOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetCapacityToBeAddedOk() (*int64, bool)`

GetCapacityToBeAddedOk returns a tuple with the CapacityToBeAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityToBeAdded

`func (o *FleetAddCapacityToResourceInstanceRequest) SetCapacityToBeAdded(v int64)`

SetCapacityToBeAdded sets CapacityToBeAdded field to given value.


### GetEnvironmentId

`func (o *FleetAddCapacityToResourceInstanceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetAddCapacityToResourceInstanceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetAddCapacityToResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetAddCapacityToResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetResourceId

`func (o *FleetAddCapacityToResourceInstanceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *FleetAddCapacityToResourceInstanceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *FleetAddCapacityToResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetAddCapacityToResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetAddCapacityToResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetAddCapacityToResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetAddCapacityToResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


