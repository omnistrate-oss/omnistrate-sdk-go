# ListEligibleInstancesPerUpgradeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]InstanceUpgrade**](InstanceUpgrade.md) | The list of instances that are in the upgrade path | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewListEligibleInstancesPerUpgradeResult

`func NewListEligibleInstancesPerUpgradeResult(instances []InstanceUpgrade, productTierId string, serviceId string, upgradePathId string, ) *ListEligibleInstancesPerUpgradeResult`

NewListEligibleInstancesPerUpgradeResult instantiates a new ListEligibleInstancesPerUpgradeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEligibleInstancesPerUpgradeResultWithDefaults

`func NewListEligibleInstancesPerUpgradeResultWithDefaults() *ListEligibleInstancesPerUpgradeResult`

NewListEligibleInstancesPerUpgradeResultWithDefaults instantiates a new ListEligibleInstancesPerUpgradeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListEligibleInstancesPerUpgradeResult) GetInstances() []InstanceUpgrade`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListEligibleInstancesPerUpgradeResult) GetInstancesOk() (*[]InstanceUpgrade, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListEligibleInstancesPerUpgradeResult) SetInstances(v []InstanceUpgrade)`

SetInstances sets Instances field to given value.


### GetProductTierId

`func (o *ListEligibleInstancesPerUpgradeResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListEligibleInstancesPerUpgradeResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListEligibleInstancesPerUpgradeResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ListEligibleInstancesPerUpgradeResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListEligibleInstancesPerUpgradeResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListEligibleInstancesPerUpgradeResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetUpgradePathId

`func (o *ListEligibleInstancesPerUpgradeResult) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *ListEligibleInstancesPerUpgradeResult) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *ListEligibleInstancesPerUpgradeResult) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


