# DescribeResourceInstanceHealthHistoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** | Exclusive end of the requested period, in RFC3339 format | 
**Id** | **string** | The instance ID | 
**Limit** | Pointer to **int64** | Maximum number of changes to return in one page | [optional] [default to 100]
**NextPageToken** | Pointer to **string** | Continuation token bound to the provider, instance, and requested period | [optional] 
**ProductTierKey** | **string** | The product tier name | 
**ResourceKey** | **string** | The resource key | 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentKey** | **string** | The service environment name | 
**ServiceKey** | **string** | The service name | 
**ServiceModelKey** | **string** | The service model name | 
**ServiceProviderId** | **string** | ID of a Service Provider | 
**StartTime** | **time.Time** | Inclusive start of the requested period, in RFC3339 format | 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeResourceInstanceHealthHistoryRequest

`func NewDescribeResourceInstanceHealthHistoryRequest(endTime time.Time, id string, productTierKey string, resourceKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, startTime time.Time, token string, ) *DescribeResourceInstanceHealthHistoryRequest`

NewDescribeResourceInstanceHealthHistoryRequest instantiates a new DescribeResourceInstanceHealthHistoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceInstanceHealthHistoryRequestWithDefaults

`func NewDescribeResourceInstanceHealthHistoryRequestWithDefaults() *DescribeResourceInstanceHealthHistoryRequest`

NewDescribeResourceInstanceHealthHistoryRequestWithDefaults instantiates a new DescribeResourceInstanceHealthHistoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetId

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLimit

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DescribeResourceInstanceHealthHistoryRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextPageToken

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *DescribeResourceInstanceHealthHistoryRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetProductTierKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetResourceKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetServiceAPIVersion

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetStartTime

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetSubscriptionId

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *DescribeResourceInstanceHealthHistoryRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeResourceInstanceHealthHistoryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeResourceInstanceHealthHistoryRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


