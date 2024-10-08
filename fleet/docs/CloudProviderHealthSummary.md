# CloudProviderHealthSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | The cloud provider | 
**DeployingInstances** | **int64** | The number of instances currently deploying | 
**HealthyInstances** | **int64** | The number of healthy instances in the cloud provider | 
**Message** | **string** | The status message | 
**RegionalHealthSummary** | [**map[string]RegionalHealthSummary**](RegionalHealthSummary.md) | The summary of health by region | 
**Status** | **string** | The status of the cloud provider | 
**TotalInstances** | **int64** | The total number of instances in the cloud provider | 
**UnhealthyInstances** | **int64** | The number of unhealthy instances in the cloud provider | 

## Methods

### NewCloudProviderHealthSummary

`func NewCloudProviderHealthSummary(cloudProvider string, deployingInstances int64, healthyInstances int64, message string, regionalHealthSummary map[string]RegionalHealthSummary, status string, totalInstances int64, unhealthyInstances int64, ) *CloudProviderHealthSummary`

NewCloudProviderHealthSummary instantiates a new CloudProviderHealthSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderHealthSummaryWithDefaults

`func NewCloudProviderHealthSummaryWithDefaults() *CloudProviderHealthSummary`

NewCloudProviderHealthSummaryWithDefaults instantiates a new CloudProviderHealthSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CloudProviderHealthSummary) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CloudProviderHealthSummary) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CloudProviderHealthSummary) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetDeployingInstances

`func (o *CloudProviderHealthSummary) GetDeployingInstances() int64`

GetDeployingInstances returns the DeployingInstances field if non-nil, zero value otherwise.

### GetDeployingInstancesOk

`func (o *CloudProviderHealthSummary) GetDeployingInstancesOk() (*int64, bool)`

GetDeployingInstancesOk returns a tuple with the DeployingInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployingInstances

`func (o *CloudProviderHealthSummary) SetDeployingInstances(v int64)`

SetDeployingInstances sets DeployingInstances field to given value.


### GetHealthyInstances

`func (o *CloudProviderHealthSummary) GetHealthyInstances() int64`

GetHealthyInstances returns the HealthyInstances field if non-nil, zero value otherwise.

### GetHealthyInstancesOk

`func (o *CloudProviderHealthSummary) GetHealthyInstancesOk() (*int64, bool)`

GetHealthyInstancesOk returns a tuple with the HealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyInstances

`func (o *CloudProviderHealthSummary) SetHealthyInstances(v int64)`

SetHealthyInstances sets HealthyInstances field to given value.


### GetMessage

`func (o *CloudProviderHealthSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudProviderHealthSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudProviderHealthSummary) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRegionalHealthSummary

`func (o *CloudProviderHealthSummary) GetRegionalHealthSummary() map[string]RegionalHealthSummary`

GetRegionalHealthSummary returns the RegionalHealthSummary field if non-nil, zero value otherwise.

### GetRegionalHealthSummaryOk

`func (o *CloudProviderHealthSummary) GetRegionalHealthSummaryOk() (*map[string]RegionalHealthSummary, bool)`

GetRegionalHealthSummaryOk returns a tuple with the RegionalHealthSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionalHealthSummary

`func (o *CloudProviderHealthSummary) SetRegionalHealthSummary(v map[string]RegionalHealthSummary)`

SetRegionalHealthSummary sets RegionalHealthSummary field to given value.


### GetStatus

`func (o *CloudProviderHealthSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudProviderHealthSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudProviderHealthSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTotalInstances

`func (o *CloudProviderHealthSummary) GetTotalInstances() int64`

GetTotalInstances returns the TotalInstances field if non-nil, zero value otherwise.

### GetTotalInstancesOk

`func (o *CloudProviderHealthSummary) GetTotalInstancesOk() (*int64, bool)`

GetTotalInstancesOk returns a tuple with the TotalInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstances

`func (o *CloudProviderHealthSummary) SetTotalInstances(v int64)`

SetTotalInstances sets TotalInstances field to given value.


### GetUnhealthyInstances

`func (o *CloudProviderHealthSummary) GetUnhealthyInstances() int64`

GetUnhealthyInstances returns the UnhealthyInstances field if non-nil, zero value otherwise.

### GetUnhealthyInstancesOk

`func (o *CloudProviderHealthSummary) GetUnhealthyInstancesOk() (*int64, bool)`

GetUnhealthyInstancesOk returns a tuple with the UnhealthyInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthyInstances

`func (o *CloudProviderHealthSummary) SetUnhealthyInstances(v int64)`

SetUnhealthyInstances sets UnhealthyInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


