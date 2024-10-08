# UpdateDeploymentConfigRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the deployment config | [optional] 
**InfraRollConfiguration** | Pointer to [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the deployment config | [optional] 
**RolloutPriorityList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateDeploymentConfigRequestBody

`func NewUpdateDeploymentConfigRequestBody() *UpdateDeploymentConfigRequestBody`

NewUpdateDeploymentConfigRequestBody instantiates a new UpdateDeploymentConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentConfigRequestBodyWithDefaults

`func NewUpdateDeploymentConfigRequestBodyWithDefaults() *UpdateDeploymentConfigRequestBody`

NewUpdateDeploymentConfigRequestBodyWithDefaults instantiates a new UpdateDeploymentConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateDeploymentConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeploymentConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeploymentConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeploymentConfigRequestBody) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequestBody) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *UpdateDeploymentConfigRequestBody) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequestBody) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.

### HasInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequestBody) HasInfraRollConfiguration() bool`

HasInfraRollConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeploymentConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeploymentConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeploymentConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeploymentConfigRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequestBody) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *UpdateDeploymentConfigRequestBody) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequestBody) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.

### HasRolloutPriorityList

`func (o *UpdateDeploymentConfigRequestBody) HasRolloutPriorityList() bool`

HasRolloutPriorityList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


