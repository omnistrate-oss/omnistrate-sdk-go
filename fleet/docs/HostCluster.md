# HostCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** |  | 
**CustomNetworkDetail** | Pointer to [**CustomNetworkFleetDetail**](CustomNetworkFleetDetail.md) |  | [optional] 
**Description** | **string** |  | 
**HelmPackages** | Pointer to [**[]HelmPackage**](HelmPackage.md) | Helm packages installed on the host cluster | [optional] 
**Id** | **string** |  | 
**KubernetesDashboardEndpoint** | Pointer to **string** | Endpoint of the Kubernetes dashboard | [optional] 
**RegionId** | Pointer to **string** |  | [optional] 
**Status** | **string** | Status of the host cluster | 
**Type** | **string** | Type of the host cluster | 

## Methods

### NewHostCluster

`func NewHostCluster(accountConfigId string, description string, id string, status string, type_ string, ) *HostCluster`

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

### HasRegionId

`func (o *HostCluster) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.

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


