# HostCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**AccountConfigId** | **string** | ID of an Account Config | 
**AccountID** | **string** | The actual account ID (AWS Account ID, GCP Project ID, Azure Subscription ID) or &#39;Omnistrate Hosted&#39; for managed accounts | 
**Amenities** | Pointer to [**[]Amenity**](Amenity.md) | The amenities available in the host cluster | [optional] 
**CloudProvider** | **string** | Name of the Infra Provider | 
**CreatedAt** | Pointer to **string** | The timestamp when the host cluster was created | [optional] 
**CurrentNumberOfDeployments** | **int64** | The current number of deployments on the host cluster | 
**CustomNetworkDetail** | Pointer to [**CustomNetworkFleetDetail**](CustomNetworkFleetDetail.md) |  | [optional] 
**CustomerEmail** | Pointer to **string** | Email of the customer who owns the host cluster in case this is a BYOC / Adopted host cluster | [optional] 
**CustomerOrganizationName** | Pointer to **string** | Name of the customer organization that owns the host cluster in case this is a BYOC / Adopted host cluster | [optional] 
**Description** | **string** |  | 
**HasPendingChanges** | Pointer to **bool** | Whether the host cluster has pending changes | [optional] 
**HealthStatus** | Pointer to [**HostClusterHealthStatus**](HostClusterHealthStatus.md) |  | [optional] 
**HelmPackages** | Pointer to [**[]HelmPackage**](HelmPackage.md) | Helm packages installed on the host cluster | [optional] 
**Id** | **string** | ID of a Host Cluster | 
**IntermediaryAccountDetail** | Pointer to [**IntermediaryAccountDetail**](IntermediaryAccountDetail.md) |  | [optional] 
**IsCustomDeployment** | **bool** | Indicates if the host cluster is a custom deployment | 
**IsInSyncWithOrgTemplate** | Pointer to **bool** | Whether the host cluster is in sync with the org template | [optional] 
**Key** | **string** | Unique key for the host cluster, used for identification | 
**KubernetesDashboardEndpoint** | Pointer to **string** | Endpoint of the Kubernetes dashboard | [optional] 
**ModelType** | Pointer to **string** | The model type encapsulating this service | [optional] 
**PendingAmenities** | Pointer to [**[]Amenity**](Amenity.md) | The pending amenities for the host cluster | [optional] 
**Region** | **string** | The actual region name of the host cluster | 
**RegionId** | **string** | ID of a Region | 
**Role** | Pointer to **string** |  | [optional] 
**Status** | **string** | The status of an operation | 
**Type** | **string** | Type of the host cluster | 

## Methods

### NewHostCluster

`func NewHostCluster(accountConfigId string, accountID string, cloudProvider string, currentNumberOfDeployments int64, description string, id string, isCustomDeployment bool, key string, region string, regionId string, status string, type_ string, ) *HostCluster`

NewHostCluster instantiates a new HostCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostClusterWithDefaults

`func NewHostClusterWithDefaults() *HostCluster`

NewHostClusterWithDefaults instantiates a new HostCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentType

`func (o *HostCluster) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *HostCluster) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *HostCluster) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *HostCluster) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

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


### GetAmenities

`func (o *HostCluster) GetAmenities() []Amenity`

GetAmenities returns the Amenities field if non-nil, zero value otherwise.

### GetAmenitiesOk

`func (o *HostCluster) GetAmenitiesOk() (*[]Amenity, bool)`

GetAmenitiesOk returns a tuple with the Amenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmenities

`func (o *HostCluster) SetAmenities(v []Amenity)`

SetAmenities sets Amenities field to given value.

### HasAmenities

`func (o *HostCluster) HasAmenities() bool`

HasAmenities returns a boolean if a field has been set.

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


### GetCreatedAt

`func (o *HostCluster) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostCluster) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostCluster) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HostCluster) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

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

### GetCustomerEmail

`func (o *HostCluster) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *HostCluster) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *HostCluster) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *HostCluster) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetCustomerOrganizationName

`func (o *HostCluster) GetCustomerOrganizationName() string`

GetCustomerOrganizationName returns the CustomerOrganizationName field if non-nil, zero value otherwise.

### GetCustomerOrganizationNameOk

`func (o *HostCluster) GetCustomerOrganizationNameOk() (*string, bool)`

GetCustomerOrganizationNameOk returns a tuple with the CustomerOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerOrganizationName

`func (o *HostCluster) SetCustomerOrganizationName(v string)`

SetCustomerOrganizationName sets CustomerOrganizationName field to given value.

### HasCustomerOrganizationName

`func (o *HostCluster) HasCustomerOrganizationName() bool`

HasCustomerOrganizationName returns a boolean if a field has been set.

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


### GetHasPendingChanges

`func (o *HostCluster) GetHasPendingChanges() bool`

GetHasPendingChanges returns the HasPendingChanges field if non-nil, zero value otherwise.

### GetHasPendingChangesOk

`func (o *HostCluster) GetHasPendingChangesOk() (*bool, bool)`

GetHasPendingChangesOk returns a tuple with the HasPendingChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPendingChanges

`func (o *HostCluster) SetHasPendingChanges(v bool)`

SetHasPendingChanges sets HasPendingChanges field to given value.

### HasHasPendingChanges

`func (o *HostCluster) HasHasPendingChanges() bool`

HasHasPendingChanges returns a boolean if a field has been set.

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


### GetIntermediaryAccountDetail

`func (o *HostCluster) GetIntermediaryAccountDetail() IntermediaryAccountDetail`

GetIntermediaryAccountDetail returns the IntermediaryAccountDetail field if non-nil, zero value otherwise.

### GetIntermediaryAccountDetailOk

`func (o *HostCluster) GetIntermediaryAccountDetailOk() (*IntermediaryAccountDetail, bool)`

GetIntermediaryAccountDetailOk returns a tuple with the IntermediaryAccountDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediaryAccountDetail

`func (o *HostCluster) SetIntermediaryAccountDetail(v IntermediaryAccountDetail)`

SetIntermediaryAccountDetail sets IntermediaryAccountDetail field to given value.

### HasIntermediaryAccountDetail

`func (o *HostCluster) HasIntermediaryAccountDetail() bool`

HasIntermediaryAccountDetail returns a boolean if a field has been set.

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


### GetIsInSyncWithOrgTemplate

`func (o *HostCluster) GetIsInSyncWithOrgTemplate() bool`

GetIsInSyncWithOrgTemplate returns the IsInSyncWithOrgTemplate field if non-nil, zero value otherwise.

### GetIsInSyncWithOrgTemplateOk

`func (o *HostCluster) GetIsInSyncWithOrgTemplateOk() (*bool, bool)`

GetIsInSyncWithOrgTemplateOk returns a tuple with the IsInSyncWithOrgTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSyncWithOrgTemplate

`func (o *HostCluster) SetIsInSyncWithOrgTemplate(v bool)`

SetIsInSyncWithOrgTemplate sets IsInSyncWithOrgTemplate field to given value.

### HasIsInSyncWithOrgTemplate

`func (o *HostCluster) HasIsInSyncWithOrgTemplate() bool`

HasIsInSyncWithOrgTemplate returns a boolean if a field has been set.

### GetKey

`func (o *HostCluster) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *HostCluster) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *HostCluster) SetKey(v string)`

SetKey sets Key field to given value.


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

### GetPendingAmenities

`func (o *HostCluster) GetPendingAmenities() []Amenity`

GetPendingAmenities returns the PendingAmenities field if non-nil, zero value otherwise.

### GetPendingAmenitiesOk

`func (o *HostCluster) GetPendingAmenitiesOk() (*[]Amenity, bool)`

GetPendingAmenitiesOk returns a tuple with the PendingAmenities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingAmenities

`func (o *HostCluster) SetPendingAmenities(v []Amenity)`

SetPendingAmenities sets PendingAmenities field to given value.

### HasPendingAmenities

`func (o *HostCluster) HasPendingAmenities() bool`

HasPendingAmenities returns a boolean if a field has been set.

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


