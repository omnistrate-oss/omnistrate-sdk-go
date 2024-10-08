# ResourceDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionHookDeploymentStatus** | **string** | The status of the resource action hook deployment. | 
**ComputeDeploymentStatus** | **string** | The status of the resource compute deployment. | 
**ConfigurationStatus** | **string** | The status of the resource deployment configuration. | 
**InfraDeploymentStatus** | **string** | The status of the resource infrastructure deployment. | 
**MonitoringStatus** | **string** | The status of the resource monitoring configuring. | 
**NetworkDeploymentStatus** | **string** | The status of the resource network deployment. | 
**ResourceName** | **string** | Name of the resource | 
**Status** | **string** | The status of the resource deployment. | 
**StorageDeploymentStatus** | **string** | The status of the resource storage deployment. | 

## Methods

### NewResourceDeploymentStatus

`func NewResourceDeploymentStatus(actionHookDeploymentStatus string, computeDeploymentStatus string, configurationStatus string, infraDeploymentStatus string, monitoringStatus string, networkDeploymentStatus string, resourceName string, status string, storageDeploymentStatus string, ) *ResourceDeploymentStatus`

NewResourceDeploymentStatus instantiates a new ResourceDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceDeploymentStatusWithDefaults

`func NewResourceDeploymentStatusWithDefaults() *ResourceDeploymentStatus`

NewResourceDeploymentStatusWithDefaults instantiates a new ResourceDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionHookDeploymentStatus

`func (o *ResourceDeploymentStatus) GetActionHookDeploymentStatus() string`

GetActionHookDeploymentStatus returns the ActionHookDeploymentStatus field if non-nil, zero value otherwise.

### GetActionHookDeploymentStatusOk

`func (o *ResourceDeploymentStatus) GetActionHookDeploymentStatusOk() (*string, bool)`

GetActionHookDeploymentStatusOk returns a tuple with the ActionHookDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionHookDeploymentStatus

`func (o *ResourceDeploymentStatus) SetActionHookDeploymentStatus(v string)`

SetActionHookDeploymentStatus sets ActionHookDeploymentStatus field to given value.


### GetComputeDeploymentStatus

`func (o *ResourceDeploymentStatus) GetComputeDeploymentStatus() string`

GetComputeDeploymentStatus returns the ComputeDeploymentStatus field if non-nil, zero value otherwise.

### GetComputeDeploymentStatusOk

`func (o *ResourceDeploymentStatus) GetComputeDeploymentStatusOk() (*string, bool)`

GetComputeDeploymentStatusOk returns a tuple with the ComputeDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeDeploymentStatus

`func (o *ResourceDeploymentStatus) SetComputeDeploymentStatus(v string)`

SetComputeDeploymentStatus sets ComputeDeploymentStatus field to given value.


### GetConfigurationStatus

`func (o *ResourceDeploymentStatus) GetConfigurationStatus() string`

GetConfigurationStatus returns the ConfigurationStatus field if non-nil, zero value otherwise.

### GetConfigurationStatusOk

`func (o *ResourceDeploymentStatus) GetConfigurationStatusOk() (*string, bool)`

GetConfigurationStatusOk returns a tuple with the ConfigurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationStatus

`func (o *ResourceDeploymentStatus) SetConfigurationStatus(v string)`

SetConfigurationStatus sets ConfigurationStatus field to given value.


### GetInfraDeploymentStatus

`func (o *ResourceDeploymentStatus) GetInfraDeploymentStatus() string`

GetInfraDeploymentStatus returns the InfraDeploymentStatus field if non-nil, zero value otherwise.

### GetInfraDeploymentStatusOk

`func (o *ResourceDeploymentStatus) GetInfraDeploymentStatusOk() (*string, bool)`

GetInfraDeploymentStatusOk returns a tuple with the InfraDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraDeploymentStatus

`func (o *ResourceDeploymentStatus) SetInfraDeploymentStatus(v string)`

SetInfraDeploymentStatus sets InfraDeploymentStatus field to given value.


### GetMonitoringStatus

`func (o *ResourceDeploymentStatus) GetMonitoringStatus() string`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *ResourceDeploymentStatus) GetMonitoringStatusOk() (*string, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *ResourceDeploymentStatus) SetMonitoringStatus(v string)`

SetMonitoringStatus sets MonitoringStatus field to given value.


### GetNetworkDeploymentStatus

`func (o *ResourceDeploymentStatus) GetNetworkDeploymentStatus() string`

GetNetworkDeploymentStatus returns the NetworkDeploymentStatus field if non-nil, zero value otherwise.

### GetNetworkDeploymentStatusOk

`func (o *ResourceDeploymentStatus) GetNetworkDeploymentStatusOk() (*string, bool)`

GetNetworkDeploymentStatusOk returns a tuple with the NetworkDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDeploymentStatus

`func (o *ResourceDeploymentStatus) SetNetworkDeploymentStatus(v string)`

SetNetworkDeploymentStatus sets NetworkDeploymentStatus field to given value.


### GetResourceName

`func (o *ResourceDeploymentStatus) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceDeploymentStatus) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceDeploymentStatus) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetStatus

`func (o *ResourceDeploymentStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceDeploymentStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceDeploymentStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStorageDeploymentStatus

`func (o *ResourceDeploymentStatus) GetStorageDeploymentStatus() string`

GetStorageDeploymentStatus returns the StorageDeploymentStatus field if non-nil, zero value otherwise.

### GetStorageDeploymentStatusOk

`func (o *ResourceDeploymentStatus) GetStorageDeploymentStatusOk() (*string, bool)`

GetStorageDeploymentStatusOk returns a tuple with the StorageDeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDeploymentStatus

`func (o *ResourceDeploymentStatus) SetStorageDeploymentStatus(v string)`

SetStorageDeploymentStatus sets StorageDeploymentStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


