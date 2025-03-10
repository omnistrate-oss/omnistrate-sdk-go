# FleetUpdateResourceInstanceDebugModeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Enable debug mode | 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateResourceInstanceDebugModeRequest

`func NewFleetUpdateResourceInstanceDebugModeRequest(enable bool, environmentId string, instanceId string, serviceId string, token string, ) *FleetUpdateResourceInstanceDebugModeRequest`

NewFleetUpdateResourceInstanceDebugModeRequest instantiates a new FleetUpdateResourceInstanceDebugModeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateResourceInstanceDebugModeRequestWithDefaults

`func NewFleetUpdateResourceInstanceDebugModeRequestWithDefaults() *FleetUpdateResourceInstanceDebugModeRequest`

NewFleetUpdateResourceInstanceDebugModeRequestWithDefaults instantiates a new FleetUpdateResourceInstanceDebugModeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *FleetUpdateResourceInstanceDebugModeRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetEnvironmentId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateResourceInstanceDebugModeRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateResourceInstanceDebugModeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateResourceInstanceDebugModeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


