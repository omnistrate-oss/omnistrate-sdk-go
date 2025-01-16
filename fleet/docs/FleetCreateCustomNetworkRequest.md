# FleetCreateCustomNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | CIDR block for the network | [default to "10.0.0.0/16"]
**CloudProviderName** | **string** | Name of the Infra Provider | 
**CloudProviderRegion** | **string** | The region of the cloud provider that the instance is running in. | 
**Name** | Pointer to **string** | User friendly network name to help distinguish networks with same CIDRs | [optional] 
**NetworkFeaturesConfiguration** | Pointer to [**FleetNetworkFeaturesConfiguration**](FleetNetworkFeaturesConfiguration.md) |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateCustomNetworkRequest

`func NewFleetCreateCustomNetworkRequest(cidr string, cloudProviderName string, cloudProviderRegion string, token string, ) *FleetCreateCustomNetworkRequest`

NewFleetCreateCustomNetworkRequest instantiates a new FleetCreateCustomNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateCustomNetworkRequestWithDefaults

`func NewFleetCreateCustomNetworkRequestWithDefaults() *FleetCreateCustomNetworkRequest`

NewFleetCreateCustomNetworkRequestWithDefaults instantiates a new FleetCreateCustomNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *FleetCreateCustomNetworkRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FleetCreateCustomNetworkRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FleetCreateCustomNetworkRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetCloudProviderName

`func (o *FleetCreateCustomNetworkRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *FleetCreateCustomNetworkRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *FleetCreateCustomNetworkRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCloudProviderRegion

`func (o *FleetCreateCustomNetworkRequest) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *FleetCreateCustomNetworkRequest) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *FleetCreateCustomNetworkRequest) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetName

`func (o *FleetCreateCustomNetworkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetCreateCustomNetworkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetCreateCustomNetworkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FleetCreateCustomNetworkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFeaturesConfiguration

`func (o *FleetCreateCustomNetworkRequest) GetNetworkFeaturesConfiguration() FleetNetworkFeaturesConfiguration`

GetNetworkFeaturesConfiguration returns the NetworkFeaturesConfiguration field if non-nil, zero value otherwise.

### GetNetworkFeaturesConfigurationOk

`func (o *FleetCreateCustomNetworkRequest) GetNetworkFeaturesConfigurationOk() (*FleetNetworkFeaturesConfiguration, bool)`

GetNetworkFeaturesConfigurationOk returns a tuple with the NetworkFeaturesConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFeaturesConfiguration

`func (o *FleetCreateCustomNetworkRequest) SetNetworkFeaturesConfiguration(v FleetNetworkFeaturesConfiguration)`

SetNetworkFeaturesConfiguration sets NetworkFeaturesConfiguration field to given value.

### HasNetworkFeaturesConfiguration

`func (o *FleetCreateCustomNetworkRequest) HasNetworkFeaturesConfiguration() bool`

HasNetworkFeaturesConfiguration returns a boolean if a field has been set.

### GetToken

`func (o *FleetCreateCustomNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateCustomNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateCustomNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


