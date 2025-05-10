# FleetRestoreResourceInstanceFromSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**InputParametersOverride** | Pointer to **interface{}** | Custom input parameters override | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierVersionOverride** | Pointer to **string** | The product tier version | [optional] 
**ServiceId** | **string** | ID of a Service | 
**SnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetRestoreResourceInstanceFromSnapshotRequest

`func NewFleetRestoreResourceInstanceFromSnapshotRequest(environmentId string, serviceId string, snapshotId string, token string, ) *FleetRestoreResourceInstanceFromSnapshotRequest`

NewFleetRestoreResourceInstanceFromSnapshotRequest instantiates a new FleetRestoreResourceInstanceFromSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetRestoreResourceInstanceFromSnapshotRequestWithDefaults

`func NewFleetRestoreResourceInstanceFromSnapshotRequestWithDefaults() *FleetRestoreResourceInstanceFromSnapshotRequest`

NewFleetRestoreResourceInstanceFromSnapshotRequestWithDefaults instantiates a new FleetRestoreResourceInstanceFromSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetInputParametersOverride() interface{}`

GetInputParametersOverride returns the InputParametersOverride field if non-nil, zero value otherwise.

### GetInputParametersOverrideOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetInputParametersOverrideOk() (*interface{}, bool)`

GetInputParametersOverrideOk returns a tuple with the InputParametersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetInputParametersOverride(v interface{})`

SetInputParametersOverride sets InputParametersOverride field to given value.

### HasInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) HasInputParametersOverride() bool`

HasInputParametersOverride returns a boolean if a field has been set.

### SetInputParametersOverrideNil

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetInputParametersOverrideNil(b bool)`

 SetInputParametersOverrideNil sets the value for InputParametersOverride to be an explicit nil

### UnsetInputParametersOverride
`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) UnsetInputParametersOverride()`

UnsetInputParametersOverride ensures that no value is present for InputParametersOverride, not even an explicit nil
### GetNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetProductTierVersionOverride() string`

GetProductTierVersionOverride returns the ProductTierVersionOverride field if non-nil, zero value otherwise.

### GetProductTierVersionOverrideOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetProductTierVersionOverrideOk() (*string, bool)`

GetProductTierVersionOverrideOk returns a tuple with the ProductTierVersionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetProductTierVersionOverride(v string)`

SetProductTierVersionOverride sets ProductTierVersionOverride field to given value.

### HasProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) HasProductTierVersionOverride() bool`

HasProductTierVersionOverride returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSnapshotId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetToken

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


