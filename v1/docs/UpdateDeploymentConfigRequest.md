# UpdateDeploymentConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A brief description of the deployment config | [optional] 
**Id** | **string** | ID of a Deployment Config | 
**InfraRollConfiguration** | Pointer to [**InfraRollConfiguration**](InfraRollConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the deployment config | [optional] 
**RolloutPriorityList** | Pointer to **[]string** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateDeploymentConfigRequest

`func NewUpdateDeploymentConfigRequest(id string, token string, ) *UpdateDeploymentConfigRequest`

NewUpdateDeploymentConfigRequest instantiates a new UpdateDeploymentConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentConfigRequestWithDefaults

`func NewUpdateDeploymentConfigRequestWithDefaults() *UpdateDeploymentConfigRequest`

NewUpdateDeploymentConfigRequestWithDefaults instantiates a new UpdateDeploymentConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateDeploymentConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeploymentConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeploymentConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeploymentConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *UpdateDeploymentConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDeploymentConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDeploymentConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest) GetInfraRollConfiguration() InfraRollConfiguration`

GetInfraRollConfiguration returns the InfraRollConfiguration field if non-nil, zero value otherwise.

### GetInfraRollConfigurationOk

`func (o *UpdateDeploymentConfigRequest) GetInfraRollConfigurationOk() (*InfraRollConfiguration, bool)`

GetInfraRollConfigurationOk returns a tuple with the InfraRollConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest) SetInfraRollConfiguration(v InfraRollConfiguration)`

SetInfraRollConfiguration sets InfraRollConfiguration field to given value.

### HasInfraRollConfiguration

`func (o *UpdateDeploymentConfigRequest) HasInfraRollConfiguration() bool`

HasInfraRollConfiguration returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeploymentConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeploymentConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeploymentConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeploymentConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest) GetRolloutPriorityList() []string`

GetRolloutPriorityList returns the RolloutPriorityList field if non-nil, zero value otherwise.

### GetRolloutPriorityListOk

`func (o *UpdateDeploymentConfigRequest) GetRolloutPriorityListOk() (*[]string, bool)`

GetRolloutPriorityListOk returns a tuple with the RolloutPriorityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest) SetRolloutPriorityList(v []string)`

SetRolloutPriorityList sets RolloutPriorityList field to given value.

### HasRolloutPriorityList

`func (o *UpdateDeploymentConfigRequest) HasRolloutPriorityList() bool`

HasRolloutPriorityList returns a boolean if a field has been set.

### GetToken

`func (o *UpdateDeploymentConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateDeploymentConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateDeploymentConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


