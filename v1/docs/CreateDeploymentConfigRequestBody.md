# CreateDeploymentConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the deployment config | 
**InfraRollConfiguration** | Pointer to [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the deployment config | 
**RolloutPriorityList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateDeploymentConfigRequestBody

`func NewCreateDeploymentConfigRequestBody(description string, name string, ) *CreateDeploymentConfigRequestBody`

NewCreateDeploymentConfigRequestBody instantiates a new CreateDeploymentConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeploymentConfigRequestBodyWithDefaults

`func NewCreateDeploymentConfigRequestBodyWithDefaults() *CreateDeploymentConfigRequestBody`

NewCreateDeploymentConfigRequestBodyWithDefaults instantiates a new CreateDeploymentConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDeploymentConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDeploymentConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDeploymentConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInfraRollConfiguration

`func (o *CreateDeploymentConfigRequestBody) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *CreateDeploymentConfigRequestBody) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *CreateDeploymentConfigRequestBody) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.

### HasInfraRollConfiguration

`func (o *CreateDeploymentConfigRequestBody) HasInfraRollConfiguration() bool`

HasInfraRollConfiguration returns a boolean if a field has been set.

### GetName

`func (o *CreateDeploymentConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeploymentConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeploymentConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetRolloutPriorityList

`func (o *CreateDeploymentConfigRequestBody) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *CreateDeploymentConfigRequestBody) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *CreateDeploymentConfigRequestBody) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.

### HasRolloutPriorityList

`func (o *CreateDeploymentConfigRequestBody) HasRolloutPriorityList() bool`

HasRolloutPriorityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


