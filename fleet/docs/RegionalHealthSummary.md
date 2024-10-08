# RegionalHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployingInstances** | **int64** | The number of instances currently deploying | 
**DeploymentCellHealthSummary** | [**map[string]DeploymentCellHealthSummary**](DeploymentCellHealthSummary.md) | The summary of health by deployment cell | 
**HealthyInstances** | **int64** | The number of healthy instances in the region | 
**Message** | **string** | The status message | 
**Region** | **string** | The region | 
**Status** | **string** | The status of the region | 
**TotalInstances** | **int64** | The total number of instances in the region | 
**UnhealthyInstances** | **int64** | The number of unhealthy instances in the region | 

## Methods

### NewRegionalHealthSummary

`func NewRegionalHealthSummary(deployingInstances int64, deploymentCellHealthSummary map[string]DeploymentCellHealthSummary, healthyInstances int64, message string, region string, status string, totalInstances int64, unhealthyInstances int64, ) *RegionalHealthSummary`

NewRegionalHealthSummary instantiates a new RegionalHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionalHealthSummaryWithDefaults

`func NewRegionalHealthSummaryWithDefaults() *RegionalHealthSummary`

NewRegionalHealthSummaryWithDefaults instantiates a new RegionalHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployingInstances

`func (o *RegionalHealthSummary) GetDeployingInstances() int64`

GetDeployingInstances returns the DeployingInstances field if non-nil, zero value otherwise.

### GetDeployingInstancesOk

`func (o *RegionalHealthSummary) GetDeployingInstancesOk() (*int64, bool)`

GetDeployingInstancesOk returns a tuple with the DeployingInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployingInstances

`func (o *RegionalHealthSummary) SetDeployingInstances(v int64)`

SetDeployingInstances sets DeployingInstances field to given value.


### GetDeploymentCellHealthSummary

`func (o *RegionalHealthSummary) GetDeploymentCellHealthSummary() map[string]DeploymentCellHealthSummary`

GetDeploymentCellHealthSummary returns the DeploymentCellHealthSummary field if non-nil, zero value otherwise.

### GetDeploymentCellHealthSummaryOk

`func (o *RegionalHealthSummary) GetDeploymentCellHealthSummaryOk() (*map[string]DeploymentCellHealthSummary, bool)`

GetDeploymentCellHealthSummaryOk returns a tuple with the DeploymentCellHealthSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellHealthSummary

`func (o *RegionalHealthSummary) SetDeploymentCellHealthSummary(v map[string]DeploymentCellHealthSummary)`

SetDeploymentCellHealthSummary sets DeploymentCellHealthSummary field to given value.


### GetHealthyInstances

`func (o *RegionalHealthSummary) GetHealthyInstances() int64`

GetHealthyInstances returns the HealthyInstances field if non-nil, zero value otherwise.

### GetHealthyInstancesOk

`func (o *RegionalHealthSummary) GetHealthyInstancesOk() (*int64, bool)`

GetHealthyInstancesOk returns a tuple with the HealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyInstances

`func (o *RegionalHealthSummary) SetHealthyInstances(v int64)`

SetHealthyInstances sets HealthyInstances field to given value.


### GetMessage

`func (o *RegionalHealthSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegionalHealthSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegionalHealthSummary) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRegion

`func (o *RegionalHealthSummary) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionalHealthSummary) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionalHealthSummary) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetStatus

`func (o *RegionalHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionalHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionalHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalInstances

`func (o *RegionalHealthSummary) GetTotalInstances() int64`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *RegionalHealthSummary) GetTotalInstancesOk() (*int64, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *RegionalHealthSummary) SetTotalInstances(v int64)`

SetTotalInstances sets TotalInstances field to given value.


### GetUnhealthyInstances

`func (o *RegionalHealthSummary) GetUnhealthyInstances() int64`

GetUnhealthyInstances returns the UnhealthyInstances field if non-nil, zero value otherwise.

### GetUnhealthyInstancesOk

`func (o *RegionalHealthSummary) GetUnhealthyInstancesOk() (*int64, bool)`

GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyInstances

`func (o *RegionalHealthSummary) SetUnhealthyInstances(v int64)`

SetUnhealthyInstances sets UnhealthyInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


