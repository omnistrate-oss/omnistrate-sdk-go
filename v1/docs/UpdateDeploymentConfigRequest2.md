# UpdateDeploymentConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the deployment config | [optional] 
**InfraRollConfiguration** | Pointer to [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the deployment config | [optional] 
**RolloutPriorityList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateDeploymentConfigRequest2

`func NewUpdateDeploymentConfigRequest2() *UpdateDeploymentConfigRequest2`

NewUpdateDeploymentConfigRequest2 instantiates a new UpdateDeploymentConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentConfigRequest2WithDefaults

`func NewUpdateDeploymentConfigRequest2WithDefaults() *UpdateDeploymentConfigRequest2`

NewUpdateDeploymentConfigRequest2WithDefaults instantiates a new UpdateDeploymentConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateDeploymentConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeploymentConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeploymentConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeploymentConfigRequest2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest2) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *UpdateDeploymentConfigRequest2) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest2) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.

### HasInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest2) HasInfraRollConfiguration() bool`

HasInfraRollConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeploymentConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeploymentConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeploymentConfigRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeploymentConfigRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest2) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *UpdateDeploymentConfigRequest2) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest2) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.

### HasRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest2) HasRolloutPriorityList() bool`

HasRolloutPriorityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


