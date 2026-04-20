# AccountConfigCloudNativeNetworkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigId** | **string** | ID of an Account Config | 
**Cidr** | Pointer to **string** | The primary CIDR block | [optional] 
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) | 
**CreatedAt** | **time.Time** | When this cloud native network was registered | 
**Id** | **string** | Internal cloud native network registration ID | 
**Name** | Pointer to **string** | The name of the cloud native network | [optional] 
**PrivateSubnets** | Pointer to [**[]SubnetDetail**](SubnetDetail.md) | Private subnets in this cloud native network | [optional] 
**PublicSubnets** | Pointer to [**[]SubnetDetail**](SubnetDetail.md) | Public subnets in this cloud native network | [optional] 
**Region** | **string** | The cloud region where the network resides | 
**Status** | **string** | The status of a cloud native network registered under an account configuration | 
**StatusMessage** | Pointer to **string** | Detailed status message | [optional] 
**SupportsPrivateDeployment** | Pointer to **bool** | Whether this network supports private deployments (has private subnets with egress) | [optional] 
**SupportsPublicDeployment** | Pointer to **bool** | Whether this network supports public deployments (has public subnets with IGW) | [optional] 
**UpdatedAt** | **time.Time** | When this cloud native network was last updated | 

## Methods

### NewAccountConfigCloudNativeNetworkResult

`func NewAccountConfigCloudNativeNetworkResult(accountConfigId string, cloudNativeNetworkId string, createdAt time.Time, id string, region string, status string, updatedAt time.Time, ) *AccountConfigCloudNativeNetworkResult`

NewAccountConfigCloudNativeNetworkResult instantiates a new AccountConfigCloudNativeNetworkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigCloudNativeNetworkResultWithDefaults

`func NewAccountConfigCloudNativeNetworkResultWithDefaults() *AccountConfigCloudNativeNetworkResult`

NewAccountConfigCloudNativeNetworkResultWithDefaults instantiates a new AccountConfigCloudNativeNetworkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigId

`func (o *AccountConfigCloudNativeNetworkResult) GetAccountConfigId() string`

GetAccountConfigId returns the AccountConfigId field if non-nil, zero value otherwise.

### GetAccountConfigIdOk

`func (o *AccountConfigCloudNativeNetworkResult) GetAccountConfigIdOk() (*string, bool)`

GetAccountConfigIdOk returns a tuple with the AccountConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigId

`func (o *AccountConfigCloudNativeNetworkResult) SetAccountConfigId(v string)`

SetAccountConfigId sets AccountConfigId field to given value.


### GetCidr

`func (o *AccountConfigCloudNativeNetworkResult) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *AccountConfigCloudNativeNetworkResult) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *AccountConfigCloudNativeNetworkResult) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *AccountConfigCloudNativeNetworkResult) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudNativeNetworkId

`func (o *AccountConfigCloudNativeNetworkResult) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *AccountConfigCloudNativeNetworkResult) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *AccountConfigCloudNativeNetworkResult) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetCreatedAt

`func (o *AccountConfigCloudNativeNetworkResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountConfigCloudNativeNetworkResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountConfigCloudNativeNetworkResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *AccountConfigCloudNativeNetworkResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountConfigCloudNativeNetworkResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountConfigCloudNativeNetworkResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountConfigCloudNativeNetworkResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountConfigCloudNativeNetworkResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountConfigCloudNativeNetworkResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountConfigCloudNativeNetworkResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateSubnets

`func (o *AccountConfigCloudNativeNetworkResult) GetPrivateSubnets() []SubnetDetail`

GetPrivateSubnets returns the PrivateSubnets field if non-nil, zero value otherwise.

### GetPrivateSubnetsOk

`func (o *AccountConfigCloudNativeNetworkResult) GetPrivateSubnetsOk() (*[]SubnetDetail, bool)`

GetPrivateSubnetsOk returns a tuple with the PrivateSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateSubnets

`func (o *AccountConfigCloudNativeNetworkResult) SetPrivateSubnets(v []SubnetDetail)`

SetPrivateSubnets sets PrivateSubnets field to given value.

### HasPrivateSubnets

`func (o *AccountConfigCloudNativeNetworkResult) HasPrivateSubnets() bool`

HasPrivateSubnets returns a boolean if a field has been set.

### GetPublicSubnets

`func (o *AccountConfigCloudNativeNetworkResult) GetPublicSubnets() []SubnetDetail`

GetPublicSubnets returns the PublicSubnets field if non-nil, zero value otherwise.

### GetPublicSubnetsOk

`func (o *AccountConfigCloudNativeNetworkResult) GetPublicSubnetsOk() (*[]SubnetDetail, bool)`

GetPublicSubnetsOk returns a tuple with the PublicSubnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSubnets

`func (o *AccountConfigCloudNativeNetworkResult) SetPublicSubnets(v []SubnetDetail)`

SetPublicSubnets sets PublicSubnets field to given value.

### HasPublicSubnets

`func (o *AccountConfigCloudNativeNetworkResult) HasPublicSubnets() bool`

HasPublicSubnets returns a boolean if a field has been set.

### GetRegion

`func (o *AccountConfigCloudNativeNetworkResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AccountConfigCloudNativeNetworkResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AccountConfigCloudNativeNetworkResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetStatus

`func (o *AccountConfigCloudNativeNetworkResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountConfigCloudNativeNetworkResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountConfigCloudNativeNetworkResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *AccountConfigCloudNativeNetworkResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AccountConfigCloudNativeNetworkResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AccountConfigCloudNativeNetworkResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AccountConfigCloudNativeNetworkResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetSupportsPrivateDeployment

`func (o *AccountConfigCloudNativeNetworkResult) GetSupportsPrivateDeployment() bool`

GetSupportsPrivateDeployment returns the SupportsPrivateDeployment field if non-nil, zero value otherwise.

### GetSupportsPrivateDeploymentOk

`func (o *AccountConfigCloudNativeNetworkResult) GetSupportsPrivateDeploymentOk() (*bool, bool)`

GetSupportsPrivateDeploymentOk returns a tuple with the SupportsPrivateDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPrivateDeployment

`func (o *AccountConfigCloudNativeNetworkResult) SetSupportsPrivateDeployment(v bool)`

SetSupportsPrivateDeployment sets SupportsPrivateDeployment field to given value.

### HasSupportsPrivateDeployment

`func (o *AccountConfigCloudNativeNetworkResult) HasSupportsPrivateDeployment() bool`

HasSupportsPrivateDeployment returns a boolean if a field has been set.

### GetSupportsPublicDeployment

`func (o *AccountConfigCloudNativeNetworkResult) GetSupportsPublicDeployment() bool`

GetSupportsPublicDeployment returns the SupportsPublicDeployment field if non-nil, zero value otherwise.

### GetSupportsPublicDeploymentOk

`func (o *AccountConfigCloudNativeNetworkResult) GetSupportsPublicDeploymentOk() (*bool, bool)`

GetSupportsPublicDeploymentOk returns a tuple with the SupportsPublicDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPublicDeployment

`func (o *AccountConfigCloudNativeNetworkResult) SetSupportsPublicDeployment(v bool)`

SetSupportsPublicDeployment sets SupportsPublicDeployment field to given value.

### HasSupportsPublicDeployment

`func (o *AccountConfigCloudNativeNetworkResult) HasSupportsPublicDeployment() bool`

HasSupportsPublicDeployment returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountConfigCloudNativeNetworkResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountConfigCloudNativeNetworkResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountConfigCloudNativeNetworkResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


