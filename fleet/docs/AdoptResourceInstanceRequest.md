# AdoptResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | **string** | The host cluster ID or key | 
**PrimaryResourceKey** | **string** | The primary resource key to adopt. This is the top-level resource that will be managed by Omnistrate. | 
**ResourceAdoptionConfiguration** | Pointer to [**map[string]ResourceAdoptionConfiguration**](ResourceAdoptionConfiguration.md) | The resource adoption configuration | [optional] 
**ServiceID** | **string** | ID of a Service | 
**ServicePlanID** | **string** | ID of a Product Tier | 
**ServicePlanVersion** | Pointer to **string** | The service plan version | [optional] 
**SubscriptionID** | Pointer to **string** | ID of a Subscription | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewAdoptResourceInstanceRequest

`func NewAdoptResourceInstanceRequest(hostClusterID string, primaryResourceKey string, serviceID string, servicePlanID string, token string, ) *AdoptResourceInstanceRequest`

NewAdoptResourceInstanceRequest instantiates a new AdoptResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptResourceInstanceRequestWithDefaults

`func NewAdoptResourceInstanceRequestWithDefaults() *AdoptResourceInstanceRequest`

NewAdoptResourceInstanceRequestWithDefaults instantiates a new AdoptResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *AdoptResourceInstanceRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *AdoptResourceInstanceRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *AdoptResourceInstanceRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.


### GetPrimaryResourceKey

`func (o *AdoptResourceInstanceRequest) GetPrimaryResourceKey() string`

GetPrimaryResourceKey returns the PrimaryResourceKey field if non-nil, zero value otherwise.

### GetPrimaryResourceKeyOk

`func (o *AdoptResourceInstanceRequest) GetPrimaryResourceKeyOk() (*string, bool)`

GetPrimaryResourceKeyOk returns a tuple with the PrimaryResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryResourceKey

`func (o *AdoptResourceInstanceRequest) SetPrimaryResourceKey(v string)`

SetPrimaryResourceKey sets PrimaryResourceKey field to given value.


### GetResourceAdoptionConfiguration

`func (o *AdoptResourceInstanceRequest) GetResourceAdoptionConfiguration() map[string]ResourceAdoptionConfiguration`

GetResourceAdoptionConfiguration returns the ResourceAdoptionConfiguration field if non-nil, zero value otherwise.

### GetResourceAdoptionConfigurationOk

`func (o *AdoptResourceInstanceRequest) GetResourceAdoptionConfigurationOk() (*map[string]ResourceAdoptionConfiguration, bool)`

GetResourceAdoptionConfigurationOk returns a tuple with the ResourceAdoptionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAdoptionConfiguration

`func (o *AdoptResourceInstanceRequest) SetResourceAdoptionConfiguration(v map[string]ResourceAdoptionConfiguration)`

SetResourceAdoptionConfiguration sets ResourceAdoptionConfiguration field to given value.

### HasResourceAdoptionConfiguration

`func (o *AdoptResourceInstanceRequest) HasResourceAdoptionConfiguration() bool`

HasResourceAdoptionConfiguration returns a boolean if a field has been set.

### GetServiceID

`func (o *AdoptResourceInstanceRequest) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *AdoptResourceInstanceRequest) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *AdoptResourceInstanceRequest) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetServicePlanID

`func (o *AdoptResourceInstanceRequest) GetServicePlanID() string`

GetServicePlanID returns the ServicePlanID field if non-nil, zero value otherwise.

### GetServicePlanIDOk

`func (o *AdoptResourceInstanceRequest) GetServicePlanIDOk() (*string, bool)`

GetServicePlanIDOk returns a tuple with the ServicePlanID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanID

`func (o *AdoptResourceInstanceRequest) SetServicePlanID(v string)`

SetServicePlanID sets ServicePlanID field to given value.


### GetServicePlanVersion

`func (o *AdoptResourceInstanceRequest) GetServicePlanVersion() string`

GetServicePlanVersion returns the ServicePlanVersion field if non-nil, zero value otherwise.

### GetServicePlanVersionOk

`func (o *AdoptResourceInstanceRequest) GetServicePlanVersionOk() (*string, bool)`

GetServicePlanVersionOk returns a tuple with the ServicePlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanVersion

`func (o *AdoptResourceInstanceRequest) SetServicePlanVersion(v string)`

SetServicePlanVersion sets ServicePlanVersion field to given value.

### HasServicePlanVersion

`func (o *AdoptResourceInstanceRequest) HasServicePlanVersion() bool`

HasServicePlanVersion returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *AdoptResourceInstanceRequest) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *AdoptResourceInstanceRequest) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *AdoptResourceInstanceRequest) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *AdoptResourceInstanceRequest) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetToken

`func (o *AdoptResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AdoptResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AdoptResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


