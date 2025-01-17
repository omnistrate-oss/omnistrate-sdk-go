# RemoveCapacityFromResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapacityToBeRemoved** | **int64** | Number of replicas to be removed | 
**Id** | **string** | The instance ID | 
**ProductTierKey** | **string** | The product tier name | 
**ResourceKey** | **string** | The resource key | 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentKey** | **string** | The service environment name | 
**ServiceKey** | **string** | The service name | 
**ServiceModelKey** | **string** | The service model name | 
**ServiceProviderId** | **string** | ID of a Service Provider | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRemoveCapacityFromResourceInstanceRequest

`func NewRemoveCapacityFromResourceInstanceRequest(capacityToBeRemoved int64, id string, productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, token string, ) *RemoveCapacityFromResourceInstanceRequest`

NewRemoveCapacityFromResourceInstanceRequest instantiates a new RemoveCapacityFromResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveCapacityFromResourceInstanceRequestWithDefaults

`func NewRemoveCapacityFromResourceInstanceRequestWithDefaults() *RemoveCapacityFromResourceInstanceRequest`

NewRemoveCapacityFromResourceInstanceRequestWithDefaults instantiates a new RemoveCapacityFromResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacityToBeRemoved

`func (o *RemoveCapacityFromResourceInstanceRequest) GetCapacityToBeRemoved() int64`

GetCapacityToBeRemoved returns the CapacityToBeRemoved field if non-nil, zero value otherwise.

### GetCapacityToBeRemovedOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetCapacityToBeRemovedOk() (*int64, bool)`

GetCapacityToBeRemovedOk returns a tuple with the CapacityToBeRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityToBeRemoved

`func (o *RemoveCapacityFromResourceInstanceRequest) SetCapacityToBeRemoved(v int64)`

SetCapacityToBeRemoved sets CapacityToBeRemoved field to given value.


### GetId

`func (o *RemoveCapacityFromResourceInstanceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoveCapacityFromResourceInstanceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierKey

`func (o *RemoveCapacityFromResourceInstanceRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *RemoveCapacityFromResourceInstanceRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetResourceKey

`func (o *RemoveCapacityFromResourceInstanceRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *RemoveCapacityFromResourceInstanceRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceAPIVersion

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *RemoveCapacityFromResourceInstanceRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *RemoveCapacityFromResourceInstanceRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *RemoveCapacityFromResourceInstanceRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *RemoveCapacityFromResourceInstanceRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *RemoveCapacityFromResourceInstanceRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetSubscriptionId

`func (o *RemoveCapacityFromResourceInstanceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *RemoveCapacityFromResourceInstanceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *RemoveCapacityFromResourceInstanceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *RemoveCapacityFromResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RemoveCapacityFromResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RemoveCapacityFromResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


