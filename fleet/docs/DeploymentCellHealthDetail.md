# DeploymentCellHealthDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | **string** | The cloud provider of the host cluster | 
**HostClusterID** | **string** | The ID of the host cluster | 
**InstanceHealth** | [**map[string]InstanceHealthSummary**](InstanceHealthSummary.md) | The health summary of the instances by instance ID | 
**RegionCode** | **string** | The region code of the host cluster | 
**ServiceEnvironmentName** | **string** | The name of the service environment | 
**ServiceName** | **string** | The name of the service | 

## Methods

### NewDeploymentCellHealthDetail

`func NewDeploymentCellHealthDetail(cloudProvider string, hostClusterID string, instanceHealth map[string]InstanceHealthSummary, regionCode string, serviceEnvironmentName string, serviceName string, ) *DeploymentCellHealthDetail`

NewDeploymentCellHealthDetail instantiates a new DeploymentCellHealthDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentCellHealthDetailWithDefaults

`func NewDeploymentCellHealthDetailWithDefaults() *DeploymentCellHealthDetail`

NewDeploymentCellHealthDetailWithDefaults instantiates a new DeploymentCellHealthDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *DeploymentCellHealthDetail) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DeploymentCellHealthDetail) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DeploymentCellHealthDetail) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetHostClusterID

`func (o *DeploymentCellHealthDetail) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *DeploymentCellHealthDetail) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *DeploymentCellHealthDetail) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetInstanceHealth

`func (o *DeploymentCellHealthDetail) GetInstanceHealth() map[string]InstanceHealthSummary`

GetInstanceHealth returns the InstanceHealth field if non-nil, zero value otherwise.

### GetInstanceHealthOk

`func (o *DeploymentCellHealthDetail) GetInstanceHealthOk() (*map[string]InstanceHealthSummary, bool)`

GetInstanceHealthOk returns a tuple with the InstanceHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceHealth

`func (o *DeploymentCellHealthDetail) SetInstanceHealth(v map[string]InstanceHealthSummary)`

SetInstanceHealth sets InstanceHealth field to given value.


### GetRegionCode

`func (o *DeploymentCellHealthDetail) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *DeploymentCellHealthDetail) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *DeploymentCellHealthDetail) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetServiceEnvironmentName

`func (o *DeploymentCellHealthDetail) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *DeploymentCellHealthDetail) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *DeploymentCellHealthDetail) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceName

`func (o *DeploymentCellHealthDetail) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DeploymentCellHealthDetail) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DeploymentCellHealthDetail) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


