# FleetCustomNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR block for the network | [default to "10.0.0.0/16"]
**CloudProviderName** | **string** | The name of the cloud provider that the instance is running on. | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**Id** | **string** | ID of a custom network | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**NetworkDefinitionType** | Pointer to **string** | Type of the network definition | [optional] 
**NetworkFeaturesConfiguration** | Pointer to [**FleetNetworkFeaturesConfiguration**](FleetNetworkFeaturesConfiguration.md) |  | [optional] 
**NetworkInstances** | Pointer to [**[]FleetCustomNetworkInstance**](FleetCustomNetworkInstance.md) | List of network instances created within this custom network | [optional] 
**OwningOrgID** | **string** | ID of the owning organization | 
**OwningOrgName** | **string** | Name of the owning organization | 
**OwningUserID** | Pointer to **string** | ID of the owning user | [optional] 
**OwningUserName** | Pointer to **string** | Name of the owning user | [optional] 

## Methods

### NewFleetCustomNetwork

`func NewFleetCustomNetwork(cidr string, cloudProviderName string, cloudProviderRegion string, id string, owningOrgID string, owningOrgName string, ) *FleetCustomNetwork`

NewFleetCustomNetwork instantiates a new FleetCustomNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCustomNetworkWithDefaults

`func NewFleetCustomNetworkWithDefaults() *FleetCustomNetwork`

NewFleetCustomNetworkWithDefaults instantiates a new FleetCustomNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *FleetCustomNetwork) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FleetCustomNetwork) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FleetCustomNetwork) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCloudProviderName

`func (o *FleetCustomNetwork) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *FleetCustomNetwork) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *FleetCustomNetwork) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *FleetCustomNetwork) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *FleetCustomNetwork) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *FleetCustomNetwork) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetId

`func (o *FleetCustomNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetCustomNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetCustomNetwork) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FleetCustomNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetCustomNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetCustomNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FleetCustomNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkDefinitionType

`func (o *FleetCustomNetwork) GetNetworkDefinitionType() string`

GetNetworkDefinitionType returns the NetworkDefinitionType field if non-nil, zero value otherwise.

### GetNetworkDefinitionTypeOk

`func (o *FleetCustomNetwork) GetNetworkDefinitionTypeOk() (*string, bool)`

GetNetworkDefinitionTypeOk returns a tuple with the NetworkDefinitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDefinitionType

`func (o *FleetCustomNetwork) SetNetworkDefinitionType(v string)`

SetNetworkDefinitionType sets NetworkDefinitionType field to given value.

### HasNetworkDefinitionType

`func (o *FleetCustomNetwork) HasNetworkDefinitionType() bool`

HasNetworkDefinitionType returns a boolean if a field has been set.

### GetNetworkFeaturesConfiguration

`func (o *FleetCustomNetwork) GetNetworkFeaturesConfiguration() FleetNetworkFeaturesConfiguration`

GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field if non-nil, zero value otherwise.

### GetNetworkFeaturesConfigurationOk

`func (o *FleetCustomNetwork) GetNetworkFeaturesConfigurationOk() (*FleetNetworkFeaturesConfiguration, bool)`

GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFeaturesConfiguration

`func (o *FleetCustomNetwork) SetNetworkFeaturesConfiguration(v FleetNetworkFeaturesConfiguration)`

SetNetworkFeaturesConfiguration sets NetworkFeaturesConfiguration field to given value.

### HasNetworkFeaturesConfiguration

`func (o *FleetCustomNetwork) HasNetworkFeaturesConfiguration() bool`

HasNetworkFeaturesConfiguration returns a boolean if a field has been set.

### GetNetworkInstances

`func (o *FleetCustomNetwork) GetNetworkInstances() []FleetCustomNetworkInstance`

GetNetworkInstances returns the NetworkInstances field if non-nil, zero value otherwise.

### GetNetworkInstancesOk

`func (o *FleetCustomNetwork) GetNetworkInstancesOk() (*[]FleetCustomNetworkInstance, bool)`

GetNetworkInstancesOk returns a tuple with the NetworkInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInstances

`func (o *FleetCustomNetwork) SetNetworkInstances(v []FleetCustomNetworkInstance)`

SetNetworkInstances sets NetworkInstances field to given value.

### HasNetworkInstances

`func (o *FleetCustomNetwork) HasNetworkInstances() bool`

HasNetworkInstances returns a boolean if a field has been set.

### GetOwningOrgID

`func (o *FleetCustomNetwork) GetOwningOrgID() string`

GetOwningOrgID returns the OwningOrgID field if non-nil, zero value otherwise.

### GetOwningOrgIDOk

`func (o *FleetCustomNetwork) GetOwningOrgIDOk() (*string, bool)`

GetOwningOrgIDOk returns a tuple with the OwningOrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningOrgID

`func (o *FleetCustomNetwork) SetOwningOrgID(v string)`

SetOwningOrgID sets OwningOrgID field to given value.


### GetOwningOrgName

`func (o *FleetCustomNetwork) GetOwningOrgName() string`

GetOwningOrgName returns the OwningOrgName field if non-nil, zero value otherwise.

### GetOwningOrgNameOk

`func (o *FleetCustomNetwork) GetOwningOrgNameOk() (*string, bool)`

GetOwningOrgNameOk returns a tuple with the OwningOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningOrgName

`func (o *FleetCustomNetwork) SetOwningOrgName(v string)`

SetOwningOrgName sets OwningOrgName field to given value.


### GetOwningUserID

`func (o *FleetCustomNetwork) GetOwningUserID() string`

GetOwningUserID returns the OwningUserID field if non-nil, zero value otherwise.

### GetOwningUserIDOk

`func (o *FleetCustomNetwork) GetOwningUserIDOk() (*string, bool)`

GetOwningUserIDOk returns a tuple with the OwningUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningUserID

`func (o *FleetCustomNetwork) SetOwningUserID(v string)`

SetOwningUserID sets OwningUserID field to given value.

### HasOwningUserID

`func (o *FleetCustomNetwork) HasOwningUserID() bool`

HasOwningUserID returns a boolean if a field has been set.

### GetOwningUserName

`func (o *FleetCustomNetwork) GetOwningUserName() string`

GetOwningUserName returns the OwningUserName field if non-nil, zero value otherwise.

### GetOwningUserNameOk

`func (o *FleetCustomNetwork) GetOwningUserNameOk() (*string, bool)`

GetOwningUserNameOk returns a tuple with the OwningUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwningUserName

`func (o *FleetCustomNetwork) SetOwningUserName(v string)`

SetOwningUserName sets OwningUserName field to given value.

### HasOwningUserName

`func (o *FleetCustomNetwork) HasOwningUserName() bool`

HasOwningUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


