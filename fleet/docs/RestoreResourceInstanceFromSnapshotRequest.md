# RestoreResourceInstanceFromSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierKey** | **string** | The product tier name | 
**ResourceKey** | **string** | The resource key | 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentKey** | **string** | The service environment name | 
**ServiceKey** | **string** | The service name | 
**ServiceModelKey** | **string** | The service model name | 
**ServiceProviderId** | **string** | ID of a Service Provider | 
**SnapshotId** | **string** | ID of a Resource Instance Snapshot | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRestoreResourceInstanceFromSnapshotRequest

`func NewRestoreResourceInstanceFromSnapshotRequest(productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, snapshotId string, token string, ) *RestoreResourceInstanceFromSnapshotRequest`

NewRestoreResourceInstanceFromSnapshotRequest instantiates a new RestoreResourceInstanceFromSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreResourceInstanceFromSnapshotRequestWithDefaults

`func NewRestoreResourceInstanceFromSnapshotRequestWithDefaults() *RestoreResourceInstanceFromSnapshotRequest`

NewRestoreResourceInstanceFromSnapshotRequestWithDefaults instantiates a new RestoreResourceInstanceFromSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomNetworkId

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *RestoreResourceInstanceFromSnapshotRequest) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetNetworkType

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *RestoreResourceInstanceFromSnapshotRequest) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetResourceKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceAPIVersion

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetSnapshotId

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.


### GetSubscriptionId

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RestoreResourceInstanceFromSnapshotRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RestoreResourceInstanceFromSnapshotRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RestoreResourceInstanceFromSnapshotRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


