# InstanceTypeCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumVMs** | **int64** | The number of VMs of this type | 
**TotalCost** | **float64** | The total cost across all VMs of this type | 
**TotalUptimeHours** | **float64** | The total uptime hours across all VMs of this type | 

## Methods

### NewInstanceTypeCost

`func NewInstanceTypeCost(numVMs int64, totalCost float64, totalUptimeHours float64, ) *InstanceTypeCost`

NewInstanceTypeCost instantiates a new InstanceTypeCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeCostWithDefaults

`func NewInstanceTypeCostWithDefaults() *InstanceTypeCost`

NewInstanceTypeCostWithDefaults instantiates a new InstanceTypeCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumVMs

`func (o *InstanceTypeCost) GetNumVMs() int64`

GetNumVMs returns the NumVMs field if non-nil, zero value otherwise.

### GetNumVMsOk

`func (o *InstanceTypeCost) GetNumVMsOk() (*int64, bool)`

GetNumVMsOk returns a tuple with the NumVMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVMs

`func (o *InstanceTypeCost) SetNumVMs(v int64)`

SetNumVMs sets NumVMs field to given value.


### GetTotalCost

`func (o *InstanceTypeCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *InstanceTypeCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *InstanceTypeCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetTotalUptimeHours

`func (o *InstanceTypeCost) GetTotalUptimeHours() float64`

GetTotalUptimeHours returns the TotalUptimeHours field if non-nil, zero value otherwise.

### GetTotalUptimeHoursOk

`func (o *InstanceTypeCost) GetTotalUptimeHoursOk() (*float64, bool)`

GetTotalUptimeHoursOk returns a tuple with the TotalUptimeHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUptimeHours

`func (o *InstanceTypeCost) SetTotalUptimeHours(v float64)`

SetTotalUptimeHours sets TotalUptimeHours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


