# HostCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** | ID of an Account Config | 
**AccountID** | **string** | The actual account ID (AWS Account ID, GCP Project ID, Azure Subscription ID) or &#39;Omnistrate Hosted&#39; for managed accounts | 
**CloudProvider** | **string** | Name of the Infra Provider | 
**CurrentNumberOfDeployments** | **int64** | The current number of deployments on the host cluster | 
**CustomNetworkDetail** | Pointer to [**CustomNetworkFleetDetail**](CustomNetworkFleetDetail.md) |  | [optional] 
**Description** | **string** |  | 
**HealthStatus** | Pointer to [**HostClusterHealthStatus**](HostClusterHealthStatus.md) |  | [optional] 
**HelmPackages** | Pointer to [**[]HelmPackage**](HelmPackage.md) | Helm packages installed on the host cluster | [optional] 
**Id** | **string** | ID of a Host Cluster | 
**IsCustomDeployment** | **bool** | Indicates if the host cluster is a custom deployment | 
**KubernetesDashboardEndpoint** | Pointer to **string** | Endpoint of the Kubernetes dashboard | [optional] 
**ModelType** | Pointer to **string** | The model type encapsulating this service | [optional] 
**Region** | **string** | The actual region name of the host cluster | 
**RegionId** | **string** | ID of a Region | 
**Role** | Pointer to **string** |  | [optional] 
**Status** | **string** | The status of an operation | 
**Type** | **string** | Type of the host cluster | 

## Methods

### NewHostCluster

`func NewHostCluster(accountConfigId string, accountID string, cloudProvider string, currentNumberOfDeployments int64, description string, id string, isCustomDeployment bool, region string, regionId string, status string, type_ string, ) *HostCluster`

NewHostCluster instantiates a new HostCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostClusterWithDefaults

`func NewHostClusterWithDefaults() *HostCluster`

NewHostClusterWithDefaults instantiates a new HostCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *HostCluster) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *HostCluster) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *HostCluster) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.


### GetAccountID

`func (o *HostCluster) GetAccountID() string`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *HostCluster) GetAccountIDOk() (*string, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountID

`func (o *HostCluster) SetAccountID(v string)`

SetAccountID sets AccountID field to given value.


### GetCloudProvider

`func (o *HostCluster) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *HostCluster) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *HostCluster) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCurrentNumberOfDeployments

`func (o *HostCluster) GetCurrentNumberOfDeployments() int64`

GetCurrentNumberOfDeployments returns the CurrentNumberOfDeployments field if non-nil, zero value otherwise.

### GetCurrentNumberOfDeploymentsOk

`func (o *HostCluster) GetCurrentNumberOfDeploymentsOk() (*int64, bool)`

GetCurrentNumberOfDeploymentsOk returns a tuple with the CurrentNumberOfDeployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberOfDeployments

`func (o *HostCluster) SetCurrentNumberOfDeployments(v int64)`

SetCurrentNumberOfDeployments sets CurrentNumberOfDeployments field to given value.


### GetCustomNetworkDetail

`func (o *HostCluster) GetCustomNetworkDetail() CustomNetworkFleetDetail`

GetCustomNetworkDetail returns the CustomNetworkDetail field if non-nil, zero value otherwise.

### GetCustomNetworkDetailOk

`func (o *HostCluster) GetCustomNetworkDetailOk() (*CustomNetworkFleetDetail, bool)`

GetCustomNetworkDetailOk returns a tuple with the CustomNetworkDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkDetail

`func (o *HostCluster) SetCustomNetworkDetail(v CustomNetworkFleetDetail)`

SetCustomNetworkDetail sets CustomNetworkDetail field to given value.

### HasCustomNetworkDetail

`func (o *HostCluster) HasCustomNetworkDetail() bool`

HasCustomNetworkDetail returns a boolean if a field has been set.

### GetDescription

`func (o *HostCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostCluster) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHealthStatus

`func (o *HostCluster) GetHealthStatus() HostClusterHealthStatus`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *HostCluster) GetHealthStatusOk() (*HostClusterHealthStatus, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *HostCluster) SetHealthStatus(v HostClusterHealthStatus)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *HostCluster) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetHelmPackages

`func (o *HostCluster) GetHelmPackages() []HelmPackage`

GetHelmPackages returns the HelmPackages field if non-nil, zero value otherwise.

### GetHelmPackagesOk

`func (o *HostCluster) GetHelmPackagesOk() (*[]HelmPackage, bool)`

GetHelmPackagesOk returns a tuple with the HelmPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmPackages

`func (o *HostCluster) SetHelmPackages(v []HelmPackage)`

SetHelmPackages sets HelmPackages field to given value.

### HasHelmPackages

`func (o *HostCluster) HasHelmPackages() bool`

HasHelmPackages returns a boolean if a field has been set.

### GetId

`func (o *HostCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostCluster) SetId(v string)`

SetId sets Id field to given value.


### GetIsCustomDeployment

`func (o *HostCluster) GetIsCustomDeployment() bool`

GetIsCustomDeployment returns the IsCustomDeployment field if non-nil, zero value otherwise.

### GetIsCustomDeploymentOk

`func (o *HostCluster) GetIsCustomDeploymentOk() (*bool, bool)`

GetIsCustomDeploymentOk returns a tuple with the IsCustomDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomDeployment

`func (o *HostCluster) SetIsCustomDeployment(v bool)`

SetIsCustomDeployment sets IsCustomDeployment field to given value.


### GetKubernetesDashboardEndpoint

`func (o *HostCluster) GetKubernetesDashboardEndpoint() string`

GetKubernetesDashboardEndpoint returns the KubernetesDashboardEndpoint field if non-nil, zero value otherwise.

### GetKubernetesDashboardEndpointOk

`func (o *HostCluster) GetKubernetesDashboardEndpointOk() (*string, bool)`

GetKubernetesDashboardEndpointOk returns a tuple with the KubernetesDashboardEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesDashboardEndpoint

`func (o *HostCluster) SetKubernetesDashboardEndpoint(v string)`

SetKubernetesDashboardEndpoint sets KubernetesDashboardEndpoint field to given value.

### HasKubernetesDashboardEndpoint

`func (o *HostCluster) HasKubernetesDashboardEndpoint() bool`

HasKubernetesDashboardEndpoint returns a boolean if a field has been set.

### GetModelType

`func (o *HostCluster) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *HostCluster) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *HostCluster) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *HostCluster) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetRegion

`func (o *HostCluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *HostCluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *HostCluster) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetRegionId

`func (o *HostCluster) GetRegionId() string`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *HostCluster) GetRegionIdOk() (*string, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *HostCluster) SetRegionId(v string)`

SetRegionId sets RegionId field to given value.


### GetRole

`func (o *HostCluster) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *HostCluster) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *HostCluster) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *HostCluster) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *HostCluster) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostCluster) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostCluster) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *HostCluster) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostCluster) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostCluster) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


