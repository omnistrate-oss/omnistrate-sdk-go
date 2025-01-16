# AddCustomDNSToResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDNS** | **string** | The custom DNS to add | 
**Id** | **string** | The instance ID | 
**ProductTierKey** | **string** | The product tier name | 
**ResourceKey** | **string** | The resource key | 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentKey** | **string** | The service environment name | 
**ServiceKey** | **string** | The service name | 
**ServiceModelKey** | **string** | The service model name | 
**ServiceProviderId** | **string** | ID of a Service Provider | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**TargetPort** | Pointer to **int64** | The target port | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAddCustomDNSToResourceInstanceRequest

`func NewAddCustomDNSToResourceInstanceRequest(customDNS string, id string, productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, token string, ) *AddCustomDNSToResourceInstanceRequest`

NewAddCustomDNSToResourceInstanceRequest instantiates a new AddCustomDNSToResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCustomDNSToResourceInstanceRequestWithDefaults

`func NewAddCustomDNSToResourceInstanceRequestWithDefaults() *AddCustomDNSToResourceInstanceRequest`

NewAddCustomDNSToResourceInstanceRequestWithDefaults instantiates a new AddCustomDNSToResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequest) GetCustomDNS() string`

GetCustomDNS returns the CustomDNS field if non-nil, zero value otherwise.

### GetCustomDNSOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetCustomDNSOk() (*string, bool)`

GetCustomDNSOk returns a tuple with the CustomDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDNS

`func (o *AddCustomDNSToResourceInstanceRequest) SetCustomDNS(v string)`

SetCustomDNS sets CustomDNS field to given value.


### GetId

`func (o *AddCustomDNSToResourceInstanceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCustomDNSToResourceInstanceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierKey

`func (o *AddCustomDNSToResourceInstanceRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *AddCustomDNSToResourceInstanceRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetResourceKey

`func (o *AddCustomDNSToResourceInstanceRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *AddCustomDNSToResourceInstanceRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceAPIVersion

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *AddCustomDNSToResourceInstanceRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *AddCustomDNSToResourceInstanceRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AddCustomDNSToResourceInstanceRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *AddCustomDNSToResourceInstanceRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *AddCustomDNSToResourceInstanceRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetSubscriptionId

`func (o *AddCustomDNSToResourceInstanceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *AddCustomDNSToResourceInstanceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *AddCustomDNSToResourceInstanceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest) GetTargetPort() int64`

GetTargetPort returns the TargetPort field if non-nil, zero value otherwise.

### GetTargetPortOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetTargetPortOk() (*int64, bool)`

GetTargetPortOk returns a tuple with the TargetPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest) SetTargetPort(v int64)`

SetTargetPort sets TargetPort field to given value.

### HasTargetPort

`func (o *AddCustomDNSToResourceInstanceRequest) HasTargetPort() bool`

HasTargetPort returns a boolean if a field has been set.

### GetToken

`func (o *AddCustomDNSToResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddCustomDNSToResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddCustomDNSToResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


