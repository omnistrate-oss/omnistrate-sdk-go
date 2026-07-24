# FleetAccountConfigCloudNativeNetworkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** | ID of an Account Config | 
**Cidr** | Pointer to **string** | The primary CIDR block | [optional] 
**CloudNativeNetworkId** | **string** | The provider-native network ID (for example, AWS VPC ID, GCP VPC name, or Azure VNet resource ID) | 
**CreatedAt** | **time.Time** | When this cloud native network was registered | 
**HostClusters** | Pointer to [**[]FleetAccountConfigCloudNativeNetworkHostClusterResult**](FleetAccountConfigCloudNativeNetworkHostClusterResult.md) | Host clusters discovered within this cloud native network | [optional] 
**Id** | **string** | Internal cloud native network registration ID | 
**Imported** | Pointer to **bool** | Whether this network has been imported for deployments via the import API. Independent of validation status. | [optional] 
**InUse** | Pointer to **bool** | Whether this network is currently referenced by at least one host cluster. Unimport is rejected while true. | [optional] 
**Name** | Pointer to **string** | The name of the cloud native network | [optional] 
**PrivateSubnets** | Pointer to [**[]SubnetDetail**](SubnetDetail.md) | Private subnets in this cloud native network | [optional] 
**PublicSubnets** | Pointer to [**[]SubnetDetail**](SubnetDetail.md) | Public subnets in this cloud native network | [optional] 
**Region** | **string** | The cloud region where the network resides | 
**Status** | **string** | The validation status of a cloud native network registered under an account configuration | 
**StatusMessage** | Pointer to **string** | Detailed status message | [optional] 
**SupportsPrivateDeployment** | Pointer to **bool** | Whether this network supports private deployments (has private subnets with egress) | [optional] 
**SupportsPublicDeployment** | Pointer to **bool** | Whether this network supports public deployments | [optional] 
**UpdatedAt** | **time.Time** | When this cloud native network was last updated | 

## Methods

### NewFleetAccountConfigCloudNativeNetworkResult

`func NewFleetAccountConfigCloudNativeNetworkResult(accountConfigId string, cloudNativeNetworkId string, createdAt time.Time, id string, region string, status string, updatedAt time.Time, ) *FleetAccountConfigCloudNativeNetworkResult`

NewFleetAccountConfigCloudNativeNetworkResult instantiates a new FleetAccountConfigCloudNativeNetworkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAccountConfigCloudNativeNetworkResultWithDefaults

`func NewFleetAccountConfigCloudNativeNetworkResultWithDefaults() *FleetAccountConfigCloudNativeNetworkResult`

NewFleetAccountConfigCloudNativeNetworkResultWithDefaults instantiates a new FleetAccountConfigCloudNativeNetworkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.


### GetCidr

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudNativeNetworkId

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetCreatedAt

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetHostClusters

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetHostClusters() []FleetAccountConfigCloudNativeNetworkHostClusterResult`

GetHostClusters returns the HostClusters field if non-nil, zero value otherwise.

### GetHostClustersOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetHostClustersOk() (*[]FleetAccountConfigCloudNativeNetworkHostClusterResult, bool)`

GetHostClustersOk returns a tuple with the HostClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusters

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetHostClusters(v []FleetAccountConfigCloudNativeNetworkHostClusterResult)`

SetHostClusters sets HostClusters field to given value.

### HasHostClusters

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasHostClusters() bool`

HasHostClusters returns a boolean if a field has been set.

### GetId

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetId(v string)`

SetId sets Id field to given value.


### GetImported

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasImported() bool`

HasImported returns a boolean if a field has been set.

### GetInUse

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetInUse() bool`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetInUseOk() (*bool, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetInUse(v bool)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetName

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetPrivateSubnets() []SubnetDetail`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetPrivateSubnetsOk() (*[]SubnetDetail, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetPrivateSubnets(v []SubnetDetail)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetPublicSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetPublicSubnets() []SubnetDetail`

GetPublicSubnets returns the PublicSubnets field if non-nil, zero value otherwise.

### GetPublicSubnetsOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetPublicSubnetsOk() (*[]SubnetDetail, bool)`

GetPublicSubnetsOk returns a tuple with the PublicSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetPublicSubnets(v []SubnetDetail)`

SetPublicSubnets sets PublicSubnets field to given value.

### HasPublicSubnets

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasPublicSubnets() bool`

HasPublicSubnets returns a boolean if a field has been set.

### GetRegion

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetStatus

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSupportsPrivateDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetSupportsPrivateDeployment() bool`

GetSupportsPrivateDeployment returns the SupportsPrivateDeployment field if non-nil, zero value otherwise.

### GetSupportsPrivateDeploymentOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetSupportsPrivateDeploymentOk() (*bool, bool)`

GetSupportsPrivateDeploymentOk returns a tuple with the SupportsPrivateDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPrivateDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetSupportsPrivateDeployment(v bool)`

SetSupportsPrivateDeployment sets SupportsPrivateDeployment field to given value.

### HasSupportsPrivateDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasSupportsPrivateDeployment() bool`

HasSupportsPrivateDeployment returns a boolean if a field has been set.

### GetSupportsPublicDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetSupportsPublicDeployment() bool`

GetSupportsPublicDeployment returns the SupportsPublicDeployment field if non-nil, zero value otherwise.

### GetSupportsPublicDeploymentOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetSupportsPublicDeploymentOk() (*bool, bool)`

GetSupportsPublicDeploymentOk returns a tuple with the SupportsPublicDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPublicDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetSupportsPublicDeployment(v bool)`

SetSupportsPublicDeployment sets SupportsPublicDeployment field to given value.

### HasSupportsPublicDeployment

`func (o *FleetAccountConfigCloudNativeNetworkResult) HasSupportsPublicDeployment() bool`

HasSupportsPublicDeployment returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FleetAccountConfigCloudNativeNetworkResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FleetAccountConfigCloudNativeNetworkResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


