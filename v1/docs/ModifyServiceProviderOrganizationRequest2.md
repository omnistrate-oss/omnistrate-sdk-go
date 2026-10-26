# ModifyServiceProviderOrganizationRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentCellConfigurations** | Pointer to [**map[string]DeploymentCellConfigurations**](DeploymentCellConfigurations.md) | The default deployment cell configurations for the organization per environment. | [optional] 
**AwsAPNTagValue** | Pointer to **string** | AWS APN tag override for this provider organization. Omit to leave unchanged or send an empty string to restore the default. Root and admin users only. | [optional] 
**JustInTimeAccessConfigurations** | Pointer to [**map[string]JustInTimeAccessConfiguration**](JustInTimeAccessConfiguration.md) | Root-only just-in-time access configuration patch for one environment of the authenticated service provider organization per call. GLOBAL is not supported by the backend for these settings. Omit the field or a nested property to preserve stored settings and sibling templates. Provider sign-in may seed missing AUTO defaults without enabling manual approval or request accounting. Explicit AUTO-to-MANUAL changes require backend cutoff/readiness; ordinary reads or unrelated updates do not interrupt access. | [optional] 

## Methods

### NewModifyServiceProviderOrganizationRequest2

`func NewModifyServiceProviderOrganizationRequest2() *ModifyServiceProviderOrganizationRequest2`

NewModifyServiceProviderOrganizationRequest2 instantiates a new ModifyServiceProviderOrganizationRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyServiceProviderOrganizationRequest2WithDefaults

`func NewModifyServiceProviderOrganizationRequest2WithDefaults() *ModifyServiceProviderOrganizationRequest2`

NewModifyServiceProviderOrganizationRequest2WithDefaults instantiates a new ModifyServiceProviderOrganizationRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) GetDeploymentCellConfigurations() map[string]DeploymentCellConfigurations`

GetDeploymentCellConfigurations returns the DeploymentCellConfigurations field if non-nil, zero value otherwise.

### GetDeploymentCellConfigurationsOk

`func (o *ModifyServiceProviderOrganizationRequest2) GetDeploymentCellConfigurationsOk() (*map[string]DeploymentCellConfigurations, bool)`

GetDeploymentCellConfigurationsOk returns a tuple with the DeploymentCellConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) SetDeploymentCellConfigurations(v map[string]DeploymentCellConfigurations)`

SetDeploymentCellConfigurations sets DeploymentCellConfigurations field to given value.

### HasDeploymentCellConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) HasDeploymentCellConfigurations() bool`

HasDeploymentCellConfigurations returns a boolean if a field has been set.

### GetAwsAPNTagValue

`func (o *ModifyServiceProviderOrganizationRequest2) GetAwsAPNTagValue() string`

GetAwsAPNTagValue returns the AwsAPNTagValue field if non-nil, zero value otherwise.

### GetAwsAPNTagValueOk

`func (o *ModifyServiceProviderOrganizationRequest2) GetAwsAPNTagValueOk() (*string, bool)`

GetAwsAPNTagValueOk returns a tuple with the AwsAPNTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAPNTagValue

`func (o *ModifyServiceProviderOrganizationRequest2) SetAwsAPNTagValue(v string)`

SetAwsAPNTagValue sets AwsAPNTagValue field to given value.

### HasAwsAPNTagValue

`func (o *ModifyServiceProviderOrganizationRequest2) HasAwsAPNTagValue() bool`

HasAwsAPNTagValue returns a boolean if a field has been set.

### GetJustInTimeAccessConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) GetJustInTimeAccessConfigurations() map[string]JustInTimeAccessConfiguration`

GetJustInTimeAccessConfigurations returns the JustInTimeAccessConfigurations field if non-nil, zero value otherwise.

### GetJustInTimeAccessConfigurationsOk

`func (o *ModifyServiceProviderOrganizationRequest2) GetJustInTimeAccessConfigurationsOk() (*map[string]JustInTimeAccessConfiguration, bool)`

GetJustInTimeAccessConfigurationsOk returns a tuple with the JustInTimeAccessConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustInTimeAccessConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) SetJustInTimeAccessConfigurations(v map[string]JustInTimeAccessConfiguration)`

SetJustInTimeAccessConfigurations sets JustInTimeAccessConfigurations field to given value.

### HasJustInTimeAccessConfigurations

`func (o *ModifyServiceProviderOrganizationRequest2) HasJustInTimeAccessConfigurations() bool`

HasJustInTimeAccessConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


