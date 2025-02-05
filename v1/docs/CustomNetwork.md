# CustomNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **string** | CIDR block for the network | [optional] [default to "10.0.0.0/16"]
**CloudProviderName** | **string** | Name of the Infra Provider | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**Id** | **string** | ID of a custom network | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**NetworkDefinitionType** | Pointer to **string** | Type of the network definition | [optional] 
**NetworkFeaturesConfiguration** | Pointer to [**NetworkFeaturesConfiguration**](NetworkFeaturesConfiguration.md) |  | [optional] 
**NetworkInstances** | Pointer to [**[]CustomNetworkInstance**](CustomNetworkInstance.md) | List of network instances created within this custom network | [optional] 
**OwningUserId** | Pointer to **string** | ID of a User | [optional] 
**OwningUserName** | Pointer to **string** | The name of the user that created and owns the custom network | [optional] 

## Methods

### NewCustomNetwork

`func NewCustomNetwork(cloudProviderName string, cloudProviderRegion string, id string, ) *CustomNetwork`

NewCustomNetwork instantiates a new CustomNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomNetworkWithDefaults

`func NewCustomNetworkWithDefaults() *CustomNetwork`

NewCustomNetworkWithDefaults instantiates a new CustomNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CustomNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CustomNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CustomNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *CustomNetwork) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCloudProviderName

`func (o *CustomNetwork) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *CustomNetwork) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *CustomNetwork) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *CustomNetwork) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *CustomNetwork) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *CustomNetwork) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetId

`func (o *CustomNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CustomNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkDefinitionType

`func (o *CustomNetwork) GetNetworkDefinitionType() string`

GetNetworkDefinitionType returns the NetworkDefinitionType field if non-nil, zero value otherwise.

### GetNetworkDefinitionTypeOk

`func (o *CustomNetwork) GetNetworkDefinitionTypeOk() (*string, bool)`

GetNetworkDefinitionTypeOk returns a tuple with the NetworkDefinitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDefinitionType

`func (o *CustomNetwork) SetNetworkDefinitionType(v string)`

SetNetworkDefinitionType sets NetworkDefinitionType field to given value.

### HasNetworkDefinitionType

`func (o *CustomNetwork) HasNetworkDefinitionType() bool`

HasNetworkDefinitionType returns a boolean if a field has been set.

### GetNetworkFeaturesConfiguration

`func (o *CustomNetwork) GetNetworkFeaturesConfiguration() NetworkFeaturesConfiguration`

GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field if non-nil, zero value otherwise.

### GetNetworkFeaturesConfigurationOk

`func (o *CustomNetwork) GetNetworkFeaturesConfigurationOk() (*NetworkFeaturesConfiguration, bool)`

GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFeaturesConfiguration

`func (o *CustomNetwork) SetNetworkFeaturesConfiguration(v NetworkFeaturesConfiguration)`

SetNetworkFeaturesConfiguration sets NetworkFeaturesConfiguration field to given value.

### HasNetworkFeaturesConfiguration

`func (o *CustomNetwork) HasNetworkFeaturesConfiguration() bool`

HasNetworkFeaturesConfiguration returns a boolean if a field has been set.

### GetNetworkInstances

`func (o *CustomNetwork) GetNetworkInstances() []CustomNetworkInstance`

GetNetworkInstances returns the NetworkInstances field if non-nil, zero value otherwise.

### GetNetworkInstancesOk

`func (o *CustomNetwork) GetNetworkInstancesOk() (*[]CustomNetworkInstance, bool)`

GetNetworkInstancesOk returns a tuple with the NetworkInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstances

`func (o *CustomNetwork) SetNetworkInstances(v []CustomNetworkInstance)`

SetNetworkInstances sets NetworkInstances field to given value.

### HasNetworkInstances

`func (o *CustomNetwork) HasNetworkInstances() bool`

HasNetworkInstances returns a boolean if a field has been set.

### GetOwningUserId

`func (o *CustomNetwork) GetOwningUserId() string`

GetOwningUserId returns the OwningUserId field if non-nil, zero value otherwise.

### GetOwningUserIdOk

`func (o *CustomNetwork) GetOwningUserIdOk() (*string, bool)`

GetOwningUserIdOk returns a tuple with the OwningUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningUserId

`func (o *CustomNetwork) SetOwningUserId(v string)`

SetOwningUserId sets OwningUserId field to given value.

### HasOwningUserId

`func (o *CustomNetwork) HasOwningUserId() bool`

HasOwningUserId returns a boolean if a field has been set.

### GetOwningUserName

`func (o *CustomNetwork) GetOwningUserName() string`

GetOwningUserName returns the OwningUserName field if non-nil, zero value otherwise.

### GetOwningUserNameOk

`func (o *CustomNetwork) GetOwningUserNameOk() (*string, bool)`

GetOwningUserNameOk returns a tuple with the OwningUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningUserName

`func (o *CustomNetwork) SetOwningUserName(v string)`

SetOwningUserName sets OwningUserName field to given value.

### HasOwningUserName

`func (o *CustomNetwork) HasOwningUserName() bool`

HasOwningUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


