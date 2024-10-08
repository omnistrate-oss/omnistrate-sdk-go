# DescribeDeploymentConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the deployment config | 
**Id** | **string** | The deployment configuration ID | 
**InfraRollConfiguration** | [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | 
**Name** | **string** | Name of the deployment config | 
**RolloutPriorityList** | **[]string** |  | 

## Methods

### NewDescribeDeploymentConfigResult

`func NewDescribeDeploymentConfigResult(description string, id string, infraRollConfiguration InfraRollConfiguration, name string, rolloutPriorityList []string, ) *DescribeDeploymentConfigResult`

NewDescribeDeploymentConfigResult instantiates a new DescribeDeploymentConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDeploymentConfigResultWithDefaults

`func NewDescribeDeploymentConfigResultWithDefaults() *DescribeDeploymentConfigResult`

NewDescribeDeploymentConfigResultWithDefaults instantiates a new DescribeDeploymentConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DescribeDeploymentConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeDeploymentConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeDeploymentConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeDeploymentConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeDeploymentConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeDeploymentConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetInfraRollConfiguration

`func (o *DescribeDeploymentConfigResult) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *DescribeDeploymentConfigResult) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *DescribeDeploymentConfigResult) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.


### GetName

`func (o *DescribeDeploymentConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeDeploymentConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeDeploymentConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetRolloutPriorityList

`func (o *DescribeDeploymentConfigResult) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *DescribeDeploymentConfigResult) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *DescribeDeploymentConfigResult) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


