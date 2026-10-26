# DescribeServiceProviderOrganizationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDeploymentCellConfigurations** | Pointer to [**DeploymentCellConfigurations**](DeploymentCellConfigurations.md) |  | [optional] 
**DeploymentCellConfigurationsPerEnv** | Pointer to [**map[string]DeploymentCellConfigurations**](DeploymentCellConfigurations.md) | The default deployment cell configurations for the organization per environment. | [optional] 
**AwsAPNTagValue** | Pointer to **string** | Configured AWS APN tag override, omitted when unset. | [optional] 
**Id** | Pointer to **string** | ID of an Org | [optional] 
**IsPerEnvClusterEnabled** | Pointer to **bool** | Whether per-environment clusters are enabled for the organization | [optional] 
**JustInTimeAccessConfigurations** | Pointer to [**map[string]JustInTimeAccessConfiguration**](JustInTimeAccessConfiguration.md) | Stored just-in-time access settings per environment, excluding GLOBAL, including compatibility AUTO defaults seeded during provider sign-in. Omission remains compatible AUTO access; backend provenance distinguishes seeded defaults from explicit configuration. Viewing or backfilling defaults does not enable manual approval, request accounting, expiry cleanup, or session cutoff. | [optional] 

## Methods

### NewDescribeServiceProviderOrganizationResult

`func NewDescribeServiceProviderOrganizationResult() *DescribeServiceProviderOrganizationResult`

NewDescribeServiceProviderOrganizationResult instantiates a new DescribeServiceProviderOrganizationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceProviderOrganizationResultWithDefaults

`func NewDescribeServiceProviderOrganizationResultWithDefaults() *DescribeServiceProviderOrganizationResult`

NewDescribeServiceProviderOrganizationResultWithDefaults instantiates a new DescribeServiceProviderOrganizationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDeploymentCellConfigurations

`func (o *DescribeServiceProviderOrganizationResult) GetDefaultDeploymentCellConfigurations() DeploymentCellConfigurations`

GetDefaultDeploymentCellConfigurations returns the DefaultDeploymentCellConfigurations field if non-nil, zero value otherwise.

### GetDefaultDeploymentCellConfigurationsOk

`func (o *DescribeServiceProviderOrganizationResult) GetDefaultDeploymentCellConfigurationsOk() (*DeploymentCellConfigurations, bool)`

GetDefaultDeploymentCellConfigurationsOk returns a tuple with the DefaultDeploymentCellConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentCellConfigurations

`func (o *DescribeServiceProviderOrganizationResult) SetDefaultDeploymentCellConfigurations(v DeploymentCellConfigurations)`

SetDefaultDeploymentCellConfigurations sets DefaultDeploymentCellConfigurations field to given value.

### HasDefaultDeploymentCellConfigurations

`func (o *DescribeServiceProviderOrganizationResult) HasDefaultDeploymentCellConfigurations() bool`

HasDefaultDeploymentCellConfigurations returns a boolean if a field has been set.

### GetDeploymentCellConfigurationsPerEnv

`func (o *DescribeServiceProviderOrganizationResult) GetDeploymentCellConfigurationsPerEnv() map[string]DeploymentCellConfigurations`

GetDeploymentCellConfigurationsPerEnv returns the DeploymentCellConfigurationsPerEnv field if non-nil, zero value otherwise.

### GetDeploymentCellConfigurationsPerEnvOk

`func (o *DescribeServiceProviderOrganizationResult) GetDeploymentCellConfigurationsPerEnvOk() (*map[string]DeploymentCellConfigurations, bool)`

GetDeploymentCellConfigurationsPerEnvOk returns a tuple with the DeploymentCellConfigurationsPerEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentCellConfigurationsPerEnv

`func (o *DescribeServiceProviderOrganizationResult) SetDeploymentCellConfigurationsPerEnv(v map[string]DeploymentCellConfigurations)`

SetDeploymentCellConfigurationsPerEnv sets DeploymentCellConfigurationsPerEnv field to given value.

### HasDeploymentCellConfigurationsPerEnv

`func (o *DescribeServiceProviderOrganizationResult) HasDeploymentCellConfigurationsPerEnv() bool`

HasDeploymentCellConfigurationsPerEnv returns a boolean if a field has been set.

### GetAwsAPNTagValue

`func (o *DescribeServiceProviderOrganizationResult) GetAwsAPNTagValue() string`

GetAwsAPNTagValue returns the AwsAPNTagValue field if non-nil, zero value otherwise.

### GetAwsAPNTagValueOk

`func (o *DescribeServiceProviderOrganizationResult) GetAwsAPNTagValueOk() (*string, bool)`

GetAwsAPNTagValueOk returns a tuple with the AwsAPNTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAPNTagValue

`func (o *DescribeServiceProviderOrganizationResult) SetAwsAPNTagValue(v string)`

SetAwsAPNTagValue sets AwsAPNTagValue field to given value.

### HasAwsAPNTagValue

`func (o *DescribeServiceProviderOrganizationResult) HasAwsAPNTagValue() bool`

HasAwsAPNTagValue returns a boolean if a field has been set.

### GetId

`func (o *DescribeServiceProviderOrganizationResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServiceProviderOrganizationResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServiceProviderOrganizationResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DescribeServiceProviderOrganizationResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPerEnvClusterEnabled

`func (o *DescribeServiceProviderOrganizationResult) GetIsPerEnvClusterEnabled() bool`

GetIsPerEnvClusterEnabled returns the IsPerEnvClusterEnabled field if non-nil, zero value otherwise.

### GetIsPerEnvClusterEnabledOk

`func (o *DescribeServiceProviderOrganizationResult) GetIsPerEnvClusterEnabledOk() (*bool, bool)`

GetIsPerEnvClusterEnabledOk returns a tuple with the IsPerEnvClusterEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPerEnvClusterEnabled

`func (o *DescribeServiceProviderOrganizationResult) SetIsPerEnvClusterEnabled(v bool)`

SetIsPerEnvClusterEnabled sets IsPerEnvClusterEnabled field to given value.

### HasIsPerEnvClusterEnabled

`func (o *DescribeServiceProviderOrganizationResult) HasIsPerEnvClusterEnabled() bool`

HasIsPerEnvClusterEnabled returns a boolean if a field has been set.

### GetJustInTimeAccessConfigurations

`func (o *DescribeServiceProviderOrganizationResult) GetJustInTimeAccessConfigurations() map[string]JustInTimeAccessConfiguration`

GetJustInTimeAccessConfigurations returns the JustInTimeAccessConfigurations field if non-nil, zero value otherwise.

### GetJustInTimeAccessConfigurationsOk

`func (o *DescribeServiceProviderOrganizationResult) GetJustInTimeAccessConfigurationsOk() (*map[string]JustInTimeAccessConfiguration, bool)`

GetJustInTimeAccessConfigurationsOk returns a tuple with the JustInTimeAccessConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustInTimeAccessConfigurations

`func (o *DescribeServiceProviderOrganizationResult) SetJustInTimeAccessConfigurations(v map[string]JustInTimeAccessConfiguration)`

SetJustInTimeAccessConfigurations sets JustInTimeAccessConfigurations field to given value.

### HasJustInTimeAccessConfigurations

`func (o *DescribeServiceProviderOrganizationResult) HasJustInTimeAccessConfigurations() bool`

HasJustInTimeAccessConfigurations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


