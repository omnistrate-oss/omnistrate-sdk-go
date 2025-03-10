# PerDeploymentCellCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Cost** | [**[]CostDataPerDate**](CostDataPerDate.md) |  | 
**InstancesCost** | Pointer to [**[]PerInstanceCost**](PerInstanceCost.md) |  | [optional] 
**RegionName** | **string** | The name of the region | 
**TotalCost** | **float64** | The total cost of the fleet in the deployment cell | 

## Methods

### NewPerDeploymentCellCost

`func NewPerDeploymentCellCost(cloudProviderName string, cost []CostDataPerDate, regionName string, totalCost float64, ) *PerDeploymentCellCost`

NewPerDeploymentCellCost instantiates a new PerDeploymentCellCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerDeploymentCellCostWithDefaults

`func NewPerDeploymentCellCostWithDefaults() *PerDeploymentCellCost`

NewPerDeploymentCellCostWithDefaults instantiates a new PerDeploymentCellCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *PerDeploymentCellCost) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *PerDeploymentCellCost) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *PerDeploymentCellCost) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCost

`func (o *PerDeploymentCellCost) GetCost() []CostDataPerDate`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PerDeploymentCellCost) GetCostOk() (*[]CostDataPerDate, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PerDeploymentCellCost) SetCost(v []CostDataPerDate)`

SetCost sets Cost field to given value.


### GetInstancesCost

`func (o *PerDeploymentCellCost) GetInstancesCost() []PerInstanceCost`

GetInstancesCost returns the InstancesCost field if non-nil, zero value otherwise.

### GetInstancesCostOk

`func (o *PerDeploymentCellCost) GetInstancesCostOk() (*[]PerInstanceCost, bool)`

GetInstancesCostOk returns a tuple with the InstancesCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCost

`func (o *PerDeploymentCellCost) SetInstancesCost(v []PerInstanceCost)`

SetInstancesCost sets InstancesCost field to given value.

### HasInstancesCost

`func (o *PerDeploymentCellCost) HasInstancesCost() bool`

HasInstancesCost returns a boolean if a field has been set.

### GetRegionName

`func (o *PerDeploymentCellCost) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *PerDeploymentCellCost) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *PerDeploymentCellCost) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetTotalCost

`func (o *PerDeploymentCellCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PerDeploymentCellCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PerDeploymentCellCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


