# DescribeServiceProviderOrganizationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDeploymentCellConfigurations** | Pointer to [**map[string]DefaultDeploymentCellConfigurations**](DefaultDeploymentCellConfigurations.md) | The default deployment cell configurations for the organization per environment. | [optional] 
**Id** | Pointer to **string** | ID of an Org | [optional] 

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

`func (o *DescribeServiceProviderOrganizationResult) GetDefaultDeploymentCellConfigurations() map[string]DefaultDeploymentCellConfigurations`

GetDefaultDeploymentCellConfigurations returns the DefaultDeploymentCellConfigurations field if non-nil, zero value otherwise.

### GetDefaultDeploymentCellConfigurationsOk

`func (o *DescribeServiceProviderOrganizationResult) GetDefaultDeploymentCellConfigurationsOk() (*map[string]DefaultDeploymentCellConfigurations, bool)`

GetDefaultDeploymentCellConfigurationsOk returns a tuple with the DefaultDeploymentCellConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentCellConfigurations

`func (o *DescribeServiceProviderOrganizationResult) SetDefaultDeploymentCellConfigurations(v map[string]DefaultDeploymentCellConfigurations)`

SetDefaultDeploymentCellConfigurations sets DefaultDeploymentCellConfigurations field to given value.

### HasDefaultDeploymentCellConfigurations

`func (o *DescribeServiceProviderOrganizationResult) HasDefaultDeploymentCellConfigurations() bool`

HasDefaultDeploymentCellConfigurations returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


