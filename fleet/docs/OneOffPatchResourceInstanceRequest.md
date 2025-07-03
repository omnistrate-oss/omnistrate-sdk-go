# OneOffPatchResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InstanceId** | **string** | ID of a Resource Instance | 
**ResourceId** | **string** | ID of a resource | 
**ResourceOverrideConfiguration** | Pointer to [**map[string]ResourceOneOffPatchConfigurationOverride**](ResourceOneOffPatchConfigurationOverride.md) | The resource override configuration for one-off patching. | [optional] 
**ServiceId** | **string** | ID of a Service | 
**TargetTierVersion** | Pointer to **string** | The target resource version. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewOneOffPatchResourceInstanceRequest

`func NewOneOffPatchResourceInstanceRequest(environmentId string, instanceId string, resourceId string, serviceId string, token string, ) *OneOffPatchResourceInstanceRequest`

NewOneOffPatchResourceInstanceRequest instantiates a new OneOffPatchResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneOffPatchResourceInstanceRequestWithDefaults

`func NewOneOffPatchResourceInstanceRequestWithDefaults() *OneOffPatchResourceInstanceRequest`

NewOneOffPatchResourceInstanceRequestWithDefaults instantiates a new OneOffPatchResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *OneOffPatchResourceInstanceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *OneOffPatchResourceInstanceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *OneOffPatchResourceInstanceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInstanceId

`func (o *OneOffPatchResourceInstanceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *OneOffPatchResourceInstanceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *OneOffPatchResourceInstanceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetResourceId

`func (o *OneOffPatchResourceInstanceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OneOffPatchResourceInstanceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OneOffPatchResourceInstanceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest) GetResourceOverrideConfiguration() map[string]ResourceOneOffPatchConfigurationOverride`

GetResourceOverrideConfiguration returns the ResourceOverrideConfiguration field if non-nil, zero value otherwise.

### GetResourceOverrideConfigurationOk

`func (o *OneOffPatchResourceInstanceRequest) GetResourceOverrideConfigurationOk() (*map[string]ResourceOneOffPatchConfigurationOverride, bool)`

GetResourceOverrideConfigurationOk returns a tuple with the ResourceOverrideConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest) SetResourceOverrideConfiguration(v map[string]ResourceOneOffPatchConfigurationOverride)`

SetResourceOverrideConfiguration sets ResourceOverrideConfiguration field to given value.

### HasResourceOverrideConfiguration

`func (o *OneOffPatchResourceInstanceRequest) HasResourceOverrideConfiguration() bool`

HasResourceOverrideConfiguration returns a boolean if a field has been set.

### GetServiceId

`func (o *OneOffPatchResourceInstanceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OneOffPatchResourceInstanceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OneOffPatchResourceInstanceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest) GetTargetTierVersion() string`

GetTargetTierVersion returns the TargetTierVersion field if non-nil, zero value otherwise.

### GetTargetTierVersionOk

`func (o *OneOffPatchResourceInstanceRequest) GetTargetTierVersionOk() (*string, bool)`

GetTargetTierVersionOk returns a tuple with the TargetTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest) SetTargetTierVersion(v string)`

SetTargetTierVersion sets TargetTierVersion field to given value.

### HasTargetTierVersion

`func (o *OneOffPatchResourceInstanceRequest) HasTargetTierVersion() bool`

HasTargetTierVersion returns a boolean if a field has been set.

### GetToken

`func (o *OneOffPatchResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OneOffPatchResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OneOffPatchResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


