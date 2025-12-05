# CreateCustomNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | CIDR block for the network | [optional] [default to "10.0.0.0/16"]
**CloudProviderName** | **string** | Name of the Infra Provider | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**NetworkFeaturesConfiguration** | Pointer to [**NetworkFeaturesConfiguration**](NetworkFeaturesConfiguration.md) |  | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateCustomNetworkRequest

`func NewCreateCustomNetworkRequest(cloudProviderName string, cloudProviderRegion string, token string, ) *CreateCustomNetworkRequest`

NewCreateCustomNetworkRequest instantiates a new CreateCustomNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomNetworkRequestWithDefaults

`func NewCreateCustomNetworkRequestWithDefaults() *CreateCustomNetworkRequest`

NewCreateCustomNetworkRequestWithDefaults instantiates a new CreateCustomNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CreateCustomNetworkRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateCustomNetworkRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateCustomNetworkRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CreateCustomNetworkRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudProviderName

`func (o *CreateCustomNetworkRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *CreateCustomNetworkRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *CreateCustomNetworkRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *CreateCustomNetworkRequest) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *CreateCustomNetworkRequest) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *CreateCustomNetworkRequest) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetName

`func (o *CreateCustomNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCustomNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest) GetNetworkFeaturesConfiguration() NetworkFeaturesConfiguration`

GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field if non-nil, zero value otherwise.

### GetNetworkFeaturesConfigurationOk

`func (o *CreateCustomNetworkRequest) GetNetworkFeaturesConfigurationOk() (*NetworkFeaturesConfiguration, bool)`

GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest) SetNetworkFeaturesConfiguration(v NetworkFeaturesConfiguration)`

SetNetworkFeaturesConfiguration sets NetworkFeaturesConfiguration field to given value.

### HasNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest) HasNetworkFeaturesConfiguration() bool`

HasNetworkFeaturesConfiguration returns a boolean if a field has been set.

### GetOrgId

`func (o *CreateCustomNetworkRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CreateCustomNetworkRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CreateCustomNetworkRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *CreateCustomNetworkRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetToken

`func (o *CreateCustomNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateCustomNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateCustomNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


