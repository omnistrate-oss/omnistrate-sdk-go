# AdoptResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**ExternalPayerId** | Pointer to **string** | The external payer id to record which customer should pay for this resource instance. This will override the subscription level external payer id if set. | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**ResourceConfiguration** | Pointer to **map[string]interface{}** | The resource configuration | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewAdoptResourceInstanceRequest2

`func NewAdoptResourceInstanceRequest2() *AdoptResourceInstanceRequest2`

NewAdoptResourceInstanceRequest2 instantiates a new AdoptResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptResourceInstanceRequest2WithDefaults

`func NewAdoptResourceInstanceRequest2WithDefaults() *AdoptResourceInstanceRequest2`

NewAdoptResourceInstanceRequest2WithDefaults instantiates a new AdoptResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *AdoptResourceInstanceRequest2) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *AdoptResourceInstanceRequest2) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *AdoptResourceInstanceRequest2) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *AdoptResourceInstanceRequest2) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCustomNetworkId

`func (o *AdoptResourceInstanceRequest2) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *AdoptResourceInstanceRequest2) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *AdoptResourceInstanceRequest2) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *AdoptResourceInstanceRequest2) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *AdoptResourceInstanceRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *AdoptResourceInstanceRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *AdoptResourceInstanceRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *AdoptResourceInstanceRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *AdoptResourceInstanceRequest2) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *AdoptResourceInstanceRequest2) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *AdoptResourceInstanceRequest2) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *AdoptResourceInstanceRequest2) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetRegion

`func (o *AdoptResourceInstanceRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AdoptResourceInstanceRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AdoptResourceInstanceRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AdoptResourceInstanceRequest2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourceConfiguration

`func (o *AdoptResourceInstanceRequest2) GetResourceConfiguration() map[string]interface{}`

GetResourceConfiguration returns the ResourceConfiguration field if non-nil, zero value otherwise.

### GetResourceConfigurationOk

`func (o *AdoptResourceInstanceRequest2) GetResourceConfigurationOk() (*map[string]interface{}, bool)`

GetResourceConfigurationOk returns a tuple with the ResourceConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceConfiguration

`func (o *AdoptResourceInstanceRequest2) SetResourceConfiguration(v map[string]interface{})`

SetResourceConfiguration sets ResourceConfiguration field to given value.

### HasResourceConfiguration

`func (o *AdoptResourceInstanceRequest2) HasResourceConfiguration() bool`

HasResourceConfiguration returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *AdoptResourceInstanceRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AdoptResourceInstanceRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AdoptResourceInstanceRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AdoptResourceInstanceRequest2) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


