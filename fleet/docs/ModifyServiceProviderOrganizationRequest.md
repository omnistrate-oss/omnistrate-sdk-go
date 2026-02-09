# ModifyServiceProviderOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellConfigurations** | Pointer to [**map[string]DeploymentCellConfigurations**](DeploymentCellConfigurations.md) | The default deployment cell configurations for the organization per environment. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewModifyServiceProviderOrganizationRequest

`func NewModifyServiceProviderOrganizationRequest(token string, ) *ModifyServiceProviderOrganizationRequest`

NewModifyServiceProviderOrganizationRequest instantiates a new ModifyServiceProviderOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyServiceProviderOrganizationRequestWithDefaults

`func NewModifyServiceProviderOrganizationRequestWithDefaults() *ModifyServiceProviderOrganizationRequest`

NewModifyServiceProviderOrganizationRequestWithDefaults instantiates a new ModifyServiceProviderOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest) GetDeploymentCellConfigurations() map[string]DeploymentCellConfigurations`

GetDeploymentCellConfigurations returns the DeploymentCellConfigurations field if non-nil, zero value otherwise.

### GetDeploymentCellConfigurationsOk

`func (o *ModifyServiceProviderOrganizationRequest) GetDeploymentCellConfigurationsOk() (*map[string]DeploymentCellConfigurations, bool)`

GetDeploymentCellConfigurationsOk returns a tuple with the DeploymentCellConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest) SetDeploymentCellConfigurations(v map[string]DeploymentCellConfigurations)`

SetDeploymentCellConfigurations sets DeploymentCellConfigurations field to given value.

### HasDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest) HasDeploymentCellConfigurations() bool`

HasDeploymentCellConfigurations returns a boolean if a field has been set.

### GetToken

`func (o *ModifyServiceProviderOrganizationRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ModifyServiceProviderOrganizationRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ModifyServiceProviderOrganizationRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


