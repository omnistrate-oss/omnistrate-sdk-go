# FleetUpdateResourceInstanceMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomTags** | Pointer to [**[]CustomTag**](CustomTag.md) | The custom tags for the resource instance. | [optional] 
**DeletionProtection** | Pointer to **bool** | Set to true to enable deletion protection or false to disable it | [optional] 
**EnableDebugMode** | Pointer to **bool** | Enable debug mode | [optional] 
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

### GetCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetCustomTags() []CustomTag`

GetCustomTags returns the CustomTags field if non-nil, zero value otherwise.

### GetCustomTagsOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetCustomTagsOk() (*[]CustomTag, bool)`

GetCustomTagsOk returns a tuple with the CustomTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetCustomTags(v []CustomTag)`

SetCustomTags sets CustomTags field to given value.

### HasCustomTags

`func (o *FleetUpdateResourceInstanceMetadataRequest) HasCustomTags() bool`

HasCustomTags returns a boolean if a field has been set.

### GetDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetDeletionProtection() bool`

GetDeletionProtection returns the DeletionProtection field if non-nil, zero value otherwise.

### GetDeletionProtectionOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetDeletionProtectionOk() (*bool, bool)`

GetDeletionProtectionOk returns a tuple with the DeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetDeletionProtection(v bool)`

SetDeletionProtection sets DeletionProtection field to given value.

### HasDeletionProtection

`func (o *FleetUpdateResourceInstanceMetadataRequest) HasDeletionProtection() bool`

HasDeletionProtection returns a boolean if a field has been set.

### GetEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetEnableDebugMode() bool`

GetEnableDebugMode returns the EnableDebugMode field if non-nil, zero value otherwise.

### GetEnableDebugModeOk

`func (o *FleetUpdateResourceInstanceMetadataRequest) GetEnableDebugModeOk() (*bool, bool)`

GetEnableDebugModeOk returns a tuple with the EnableDebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest) SetEnableDebugMode(v bool)`

SetEnableDebugMode sets EnableDebugMode field to given value.

### HasEnableDebugMode

`func (o *FleetUpdateResourceInstanceMetadataRequest) HasEnableDebugMode() bool`

HasEnableDebugMode returns a boolean if a field has been set.

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


