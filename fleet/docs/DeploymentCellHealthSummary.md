# DeploymentCellHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployingInstances** | **int64** | The number of instances currently deploying | 
**HealthyInstances** | **int64** | The number of healthy instances in the cell | 
**HostClusterID** | **string** | The ID of the host cluster | 
**Message** | **string** | The status message | 
**Status** | **string** | The status of the cell | 
**TotalInstances** | **int64** | The total number of instances in the cell | 
**UnhealthyInstances** | **int64** | The number of unhealthy instances in the cell | 

## Methods

### NewDeploymentCellHealthSummary

`func NewDeploymentCellHealthSummary(deployingInstances int64, healthyInstances int64, hostClusterID string, message string, status string, totalInstances int64, unhealthyInstances int64, ) *DeploymentCellHealthSummary`

NewDeploymentCellHealthSummary instantiates a new DeploymentCellHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellHealthSummaryWithDefaults

`func NewDeploymentCellHealthSummaryWithDefaults() *DeploymentCellHealthSummary`

NewDeploymentCellHealthSummaryWithDefaults instantiates a new DeploymentCellHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployingInstances

`func (o *DeploymentCellHealthSummary) GetDeployingInstances() int64`

GetDeployingInstances returns the DeployingInstances field if non-nil, zero value otherwise.

### GetDeployingInstancesOk

`func (o *DeploymentCellHealthSummary) GetDeployingInstancesOk() (*int64, bool)`

GetDeployingInstancesOk returns a tuple with the DeployingInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployingInstances

`func (o *DeploymentCellHealthSummary) SetDeployingInstances(v int64)`

SetDeployingInstances sets DeployingInstances field to given value.


### GetHealthyInstances

`func (o *DeploymentCellHealthSummary) GetHealthyInstances() int64`

GetHealthyInstances returns the HealthyInstances field if non-nil, zero value otherwise.

### GetHealthyInstancesOk

`func (o *DeploymentCellHealthSummary) GetHealthyInstancesOk() (*int64, bool)`

GetHealthyInstancesOk returns a tuple with the HealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyInstances

`func (o *DeploymentCellHealthSummary) SetHealthyInstances(v int64)`

SetHealthyInstances sets HealthyInstances field to given value.


### GetHostClusterID

`func (o *DeploymentCellHealthSummary) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DeploymentCellHealthSummary) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DeploymentCellHealthSummary) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetMessage

`func (o *DeploymentCellHealthSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DeploymentCellHealthSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DeploymentCellHealthSummary) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatus

`func (o *DeploymentCellHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentCellHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentCellHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalInstances

`func (o *DeploymentCellHealthSummary) GetTotalInstances() int64`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *DeploymentCellHealthSummary) GetTotalInstancesOk() (*int64, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *DeploymentCellHealthSummary) SetTotalInstances(v int64)`

SetTotalInstances sets TotalInstances field to given value.


### GetUnhealthyInstances

`func (o *DeploymentCellHealthSummary) GetUnhealthyInstances() int64`

GetUnhealthyInstances returns the UnhealthyInstances field if non-nil, zero value otherwise.

### GetUnhealthyInstancesOk

`func (o *DeploymentCellHealthSummary) GetUnhealthyInstancesOk() (*int64, bool)`

GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyInstances

`func (o *DeploymentCellHealthSummary) SetUnhealthyInstances(v int64)`

SetUnhealthyInstances sets UnhealthyInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


