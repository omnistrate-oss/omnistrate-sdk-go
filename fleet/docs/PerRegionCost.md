# PerRegionCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Cost** | [**[]CostDataPerDate**](CostDataPerDate.md) |  | 
**RegionName** | **string** | The name of the region | 
**TotalCost** | **float64** | The total cost of the fleet in the region | 

## Methods

### NewPerRegionCost

`func NewPerRegionCost(cloudProviderName string, cost []CostDataPerDate, regionName string, totalCost float64, ) *PerRegionCost`

NewPerRegionCost instantiates a new PerRegionCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerRegionCostWithDefaults

`func NewPerRegionCostWithDefaults() *PerRegionCost`

NewPerRegionCostWithDefaults instantiates a new PerRegionCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *PerRegionCost) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *PerRegionCost) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *PerRegionCost) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCost

`func (o *PerRegionCost) GetCost() []CostDataPerDate`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PerRegionCost) GetCostOk() (*[]CostDataPerDate, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PerRegionCost) SetCost(v []CostDataPerDate)`

SetCost sets Cost field to given value.


### GetRegionName

`func (o *PerRegionCost) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *PerRegionCost) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *PerRegionCost) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetTotalCost

`func (o *PerRegionCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PerRegionCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PerRegionCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


