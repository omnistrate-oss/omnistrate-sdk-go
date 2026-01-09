# FleetUpdateResourceInstanceMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionProtectionOverride** | Pointer to **bool** | Set to true to enable deletion protection, false to disable it, or omit to use plan defaults | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetUpdateResourceInstanceMetadataRequest

`func NewFleetUpdateResourceInstanceMetadataRequest(environmentId string, instanceId string, serviceId string, token string, ) *FleetUpdateResourceInstanceMetadataRequest`

NewFleetUpdateResourceInstanceMetadataRequest instantiates a new FleetUpdateResourceInstanceMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetUpdateResourceInstanceMetadataRequestWithDefaults

`func NewFleetUpdateResourceInstanceMetadataRequestWithDefaults() *FleetUpdateResourceInstanceMetadataRequest`

NewFleetUpdateResourceInstanceMetadataRequestWithDefaults instantiates a new FleetUpdateResourceInstanceMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionProtectionOverride

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetDeletionProtectionOverride() bool`

GetDeletionProtectionOverride returns the DeletionProtectionOverride field if non-nil, zero value otherwise.

### GetDeletionProtectionOverrideOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetDeletionProtectionOverrideOk() (*bool, bool)`

GetDeletionProtectionOverrideOk returns a tuple with the DeletionProtectionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtectionOverride

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetDeletionProtectionOverride(v bool)`

SetDeletionProtectionOverride sets DeletionProtectionOverride field to given value.

### HasDeletionProtectionOverride

`func (o *FleetUpdateResourceInstanceMetadataRequest) HasDeletionProtectionOverride() bool`

HasDeletionProtectionOverride returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetServiceId

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


