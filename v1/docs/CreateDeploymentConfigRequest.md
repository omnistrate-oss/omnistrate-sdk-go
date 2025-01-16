# CreateDeploymentConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the deployment config | 
**InfraRollConfiguration** | Pointer to [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the deployment config | 
**RolloutPriorityList** | Pointer to **[]string** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateDeploymentConfigRequest

`func NewCreateDeploymentConfigRequest(description string, name string, token string, ) *CreateDeploymentConfigRequest`

NewCreateDeploymentConfigRequest instantiates a new CreateDeploymentConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeploymentConfigRequestWithDefaults

`func NewCreateDeploymentConfigRequestWithDefaults() *CreateDeploymentConfigRequest`

NewCreateDeploymentConfigRequestWithDefaults instantiates a new CreateDeploymentConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDeploymentConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDeploymentConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDeploymentConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetInfraRollConfiguration

`func (o *CreateDeploymentConfigRequest) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *CreateDeploymentConfigRequest) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *CreateDeploymentConfigRequest) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.

### HasInfraRollConfiguration

`func (o *CreateDeploymentConfigRequest) HasInfraRollConfiguration() bool`

HasInfraRollConfiguration returns a boolean if a field has been set.

### GetName

`func (o *CreateDeploymentConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeploymentConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeploymentConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRolloutPriorityList

`func (o *CreateDeploymentConfigRequest) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *CreateDeploymentConfigRequest) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *CreateDeploymentConfigRequest) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.

### HasRolloutPriorityList

`func (o *CreateDeploymentConfigRequest) HasRolloutPriorityList() bool`

HasRolloutPriorityList returns a boolean if a field has been set.

### GetToken

`func (o *CreateDeploymentConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDeploymentConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDeploymentConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


