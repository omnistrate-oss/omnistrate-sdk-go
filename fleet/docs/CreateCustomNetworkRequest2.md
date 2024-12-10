# CreateCustomNetworkRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR block for the network | [default to "10.0.0.0/16"]
**CloudProviderName** | **string** | The name of the cloud provider that the instance is running on. | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**NetworkFeaturesConfiguration** | Pointer to [**FleetNetworkFeaturesConfiguration**](FleetNetworkFeaturesConfiguration.md) |  | [optional] 

## Methods

### NewCreateCustomNetworkRequest2

`func NewCreateCustomNetworkRequest2(cidr string, cloudProviderName string, cloudProviderRegion string, ) *CreateCustomNetworkRequest2`

NewCreateCustomNetworkRequest2 instantiates a new CreateCustomNetworkRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomNetworkRequest2WithDefaults

`func NewCreateCustomNetworkRequest2WithDefaults() *CreateCustomNetworkRequest2`

NewCreateCustomNetworkRequest2WithDefaults instantiates a new CreateCustomNetworkRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *CreateCustomNetworkRequest2) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *CreateCustomNetworkRequest2) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *CreateCustomNetworkRequest2) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCloudProviderName

`func (o *CreateCustomNetworkRequest2) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *CreateCustomNetworkRequest2) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *CreateCustomNetworkRequest2) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *CreateCustomNetworkRequest2) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *CreateCustomNetworkRequest2) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *CreateCustomNetworkRequest2) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetName

`func (o *CreateCustomNetworkRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomNetworkRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomNetworkRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCustomNetworkRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest2) GetNetworkFeaturesConfiguration() FleetNetworkFeaturesConfiguration`

GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field if non-nil, zero value otherwise.

### GetNetworkFeaturesConfigurationOk

`func (o *CreateCustomNetworkRequest2) GetNetworkFeaturesConfigurationOk() (*FleetNetworkFeaturesConfiguration, bool)`

GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest2) SetNetworkFeaturesConfiguration(v FleetNetworkFeaturesConfiguration)`

SetNetworkFeaturesConfiguration sets NetworkFeaturesConfiguration field to given value.

### HasNetworkFeaturesConfiguration

`func (o *CreateCustomNetworkRequest2) HasNetworkFeaturesConfiguration() bool`

HasNetworkFeaturesConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


