# FailoverResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedReplicaAction** | Pointer to **string** | The action to take for the replica that has failed | [optional] 
**FailedReplicaID** | **string** | The failed replica ID | 
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

### NewFailoverResourceInstanceRequest

`func NewFailoverResourceInstanceRequest(failedReplicaID string, id string, productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, token string, ) *FailoverResourceInstanceRequest`

NewFailoverResourceInstanceRequest instantiates a new FailoverResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailoverResourceInstanceRequestWithDefaults

`func NewFailoverResourceInstanceRequestWithDefaults() *FailoverResourceInstanceRequest`

NewFailoverResourceInstanceRequestWithDefaults instantiates a new FailoverResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedReplicaAction

`func (o *FailoverResourceInstanceRequest) GetFailedReplicaAction() string`

GetFailedReplicaAction returns the FailedReplicaAction field if non-nil, zero value otherwise.

### GetFailedReplicaActionOk

`func (o *FailoverResourceInstanceRequest) GetFailedReplicaActionOk() (*string, bool)`

GetFailedReplicaActionOk returns a tuple with the FailedReplicaAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaAction

`func (o *FailoverResourceInstanceRequest) SetFailedReplicaAction(v string)`

SetFailedReplicaAction sets FailedReplicaAction field to given value.

### HasFailedReplicaAction

`func (o *FailoverResourceInstanceRequest) HasFailedReplicaAction() bool`

HasFailedReplicaAction returns a boolean if a field has been set.

### GetFailedReplicaID

`func (o *FailoverResourceInstanceRequest) GetFailedReplicaID() string`

GetFailedReplicaID returns the FailedReplicaID field if non-nil, zero value otherwise.

### GetFailedReplicaIDOk

`func (o *FailoverResourceInstanceRequest) GetFailedReplicaIDOk() (*string, bool)`

GetFailedReplicaIDOk returns a tuple with the FailedReplicaID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReplicaID

`func (o *FailoverResourceInstanceRequest) SetFailedReplicaID(v string)`

SetFailedReplicaID sets FailedReplicaID field to given value.


### GetId

`func (o *FailoverResourceInstanceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailoverResourceInstanceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailoverResourceInstanceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierKey

`func (o *FailoverResourceInstanceRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *FailoverResourceInstanceRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *FailoverResourceInstanceRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetResourceKey

`func (o *FailoverResourceInstanceRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *FailoverResourceInstanceRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *FailoverResourceInstanceRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceAPIVersion

`func (o *FailoverResourceInstanceRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *FailoverResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *FailoverResourceInstanceRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *FailoverResourceInstanceRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *FailoverResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *FailoverResourceInstanceRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *FailoverResourceInstanceRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *FailoverResourceInstanceRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *FailoverResourceInstanceRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *FailoverResourceInstanceRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *FailoverResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *FailoverResourceInstanceRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *FailoverResourceInstanceRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *FailoverResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *FailoverResourceInstanceRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetSubscriptionId

`func (o *FailoverResourceInstanceRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FailoverResourceInstanceRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FailoverResourceInstanceRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FailoverResourceInstanceRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *FailoverResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FailoverResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FailoverResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


