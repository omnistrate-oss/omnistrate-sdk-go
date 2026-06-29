# FleetRestoreResourceInstanceFromSnapshotRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**InputParametersOverride** | Pointer to **interface{}** | Custom input parameters override | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierVersionOverride** | Pointer to **string** | The product tier version | [optional] 
**RestoreToSourceInstance** | Pointer to **bool** | If true, restore to the same instance ID as the previously deleted source instance. This preserves the original instance ID and endpoint. | [optional] 
**SubscriptionId** | Pointer to **string** | The target subscription ID. If omitted, restores to the snapshot&#39;s original subscription. Cross-subscription restore is only supported for service provider hosted deployments on the same service, product tier, and host cluster. | [optional] 

## Methods

### NewFleetRestoreResourceInstanceFromSnapshotRequest2

`func NewFleetRestoreResourceInstanceFromSnapshotRequest2() *FleetRestoreResourceInstanceFromSnapshotRequest2`

NewFleetRestoreResourceInstanceFromSnapshotRequest2 instantiates a new FleetRestoreResourceInstanceFromSnapshotRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetRestoreResourceInstanceFromSnapshotRequest2WithDefaults

`func NewFleetRestoreResourceInstanceFromSnapshotRequest2WithDefaults() *FleetRestoreResourceInstanceFromSnapshotRequest2`

NewFleetRestoreResourceInstanceFromSnapshotRequest2WithDefaults instantiates a new FleetRestoreResourceInstanceFromSnapshotRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomNetworkId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetInputParametersOverride() interface{}`

GetInputParametersOverride returns the InputParametersOverride field if non-nil, zero value otherwise.

### GetInputParametersOverrideOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetInputParametersOverrideOk() (*interface{}, bool)`

GetInputParametersOverrideOk returns a tuple with the InputParametersOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetInputParametersOverride(v interface{})`

SetInputParametersOverride sets InputParametersOverride field to given value.

### HasInputParametersOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasInputParametersOverride() bool`

HasInputParametersOverride returns a boolean if a field has been set.

### SetInputParametersOverrideNil

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetInputParametersOverrideNil(b bool)`

 SetInputParametersOverrideNil sets the value for InputParametersOverride to be an explicit nil

### UnsetInputParametersOverride
`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) UnsetInputParametersOverride()`

UnsetInputParametersOverride ensures that no value is present for InputParametersOverride, not even an explicit nil
### GetNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetProductTierVersionOverride() string`

GetProductTierVersionOverride returns the ProductTierVersionOverride field if non-nil, zero value otherwise.

### GetProductTierVersionOverrideOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetProductTierVersionOverrideOk() (*string, bool)`

GetProductTierVersionOverrideOk returns a tuple with the ProductTierVersionOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetProductTierVersionOverride(v string)`

SetProductTierVersionOverride sets ProductTierVersionOverride field to given value.

### HasProductTierVersionOverride

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasProductTierVersionOverride() bool`

HasProductTierVersionOverride returns a boolean if a field has been set.

### GetRestoreToSourceInstance

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetRestoreToSourceInstance() bool`

GetRestoreToSourceInstance returns the RestoreToSourceInstance field if non-nil, zero value otherwise.

### GetRestoreToSourceInstanceOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetRestoreToSourceInstanceOk() (*bool, bool)`

GetRestoreToSourceInstanceOk returns a tuple with the RestoreToSourceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreToSourceInstance

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetRestoreToSourceInstance(v bool)`

SetRestoreToSourceInstance sets RestoreToSourceInstance field to given value.

### HasRestoreToSourceInstance

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasRestoreToSourceInstance() bool`

HasRestoreToSourceInstance returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetRestoreResourceInstanceFromSnapshotRequest2) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


