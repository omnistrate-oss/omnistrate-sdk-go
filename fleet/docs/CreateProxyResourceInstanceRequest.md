# CreateProxyResourceInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**ProductTierKey** | **string** | The product tier name | 
**Region** | Pointer to **string** | The region code | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 
**ServiceAPIVersion** | **string** | The service API version | 
**ServiceEnvironmentKey** | **string** | The service environment name | 
**ServiceKey** | **string** | The service name | 
**ServiceModelKey** | **string** | The service model name | 
**ServiceProviderId** | **string** | ID of a Service Provider | 
**SourceResourceKey** | **string** | The serverless resource key | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateProxyResourceInstanceRequest

`func NewCreateProxyResourceInstanceRequest(productTierKey string, serviceAPIVersion string, serviceEnvironmentKey string, serviceKey string, serviceModelKey string, serviceProviderId string, sourceResourceKey string, token string, ) *CreateProxyResourceInstanceRequest`

NewCreateProxyResourceInstanceRequest instantiates a new CreateProxyResourceInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProxyResourceInstanceRequestWithDefaults

`func NewCreateProxyResourceInstanceRequestWithDefaults() *CreateProxyResourceInstanceRequest`

NewCreateProxyResourceInstanceRequestWithDefaults instantiates a new CreateProxyResourceInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateProxyResourceInstanceRequest) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateProxyResourceInstanceRequest) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateProxyResourceInstanceRequest) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateProxyResourceInstanceRequest) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetProductTierKey

`func (o *CreateProxyResourceInstanceRequest) GetProductTierKey() string`

GetProductTierKey returns the ProductTierKey field if non-nil, zero value otherwise.

### GetProductTierKeyOk

`func (o *CreateProxyResourceInstanceRequest) GetProductTierKeyOk() (*string, bool)`

GetProductTierKeyOk returns a tuple with the ProductTierKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierKey

`func (o *CreateProxyResourceInstanceRequest) SetProductTierKey(v string)`

SetProductTierKey sets ProductTierKey field to given value.


### GetRegion

`func (o *CreateProxyResourceInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateProxyResourceInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateProxyResourceInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateProxyResourceInstanceRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestParams

`func (o *CreateProxyResourceInstanceRequest) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *CreateProxyResourceInstanceRequest) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *CreateProxyResourceInstanceRequest) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *CreateProxyResourceInstanceRequest) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *CreateProxyResourceInstanceRequest) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *CreateProxyResourceInstanceRequest) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetServiceAPIVersion

`func (o *CreateProxyResourceInstanceRequest) GetServiceAPIVersion() string`

GetServiceAPIVersion returns the ServiceAPIVersion field if non-nil, zero value otherwise.

### GetServiceAPIVersionOk

`func (o *CreateProxyResourceInstanceRequest) GetServiceAPIVersionOk() (*string, bool)`

GetServiceAPIVersionOk returns a tuple with the ServiceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIVersion

`func (o *CreateProxyResourceInstanceRequest) SetServiceAPIVersion(v string)`

SetServiceAPIVersion sets ServiceAPIVersion field to given value.


### GetServiceEnvironmentKey

`func (o *CreateProxyResourceInstanceRequest) GetServiceEnvironmentKey() string`

GetServiceEnvironmentKey returns the ServiceEnvironmentKey field if non-nil, zero value otherwise.

### GetServiceEnvironmentKeyOk

`func (o *CreateProxyResourceInstanceRequest) GetServiceEnvironmentKeyOk() (*string, bool)`

GetServiceEnvironmentKeyOk returns a tuple with the ServiceEnvironmentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentKey

`func (o *CreateProxyResourceInstanceRequest) SetServiceEnvironmentKey(v string)`

SetServiceEnvironmentKey sets ServiceEnvironmentKey field to given value.


### GetServiceKey

`func (o *CreateProxyResourceInstanceRequest) GetServiceKey() string`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *CreateProxyResourceInstanceRequest) GetServiceKeyOk() (*string, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *CreateProxyResourceInstanceRequest) SetServiceKey(v string)`

SetServiceKey sets ServiceKey field to given value.


### GetServiceModelKey

`func (o *CreateProxyResourceInstanceRequest) GetServiceModelKey() string`

GetServiceModelKey returns the ServiceModelKey field if non-nil, zero value otherwise.

### GetServiceModelKeyOk

`func (o *CreateProxyResourceInstanceRequest) GetServiceModelKeyOk() (*string, bool)`

GetServiceModelKeyOk returns a tuple with the ServiceModelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelKey

`func (o *CreateProxyResourceInstanceRequest) SetServiceModelKey(v string)`

SetServiceModelKey sets ServiceModelKey field to given value.


### GetServiceProviderId

`func (o *CreateProxyResourceInstanceRequest) GetServiceProviderId() string`

GetServiceProviderId returns the ServiceProviderId field if non-nil, zero value otherwise.

### GetServiceProviderIdOk

`func (o *CreateProxyResourceInstanceRequest) GetServiceProviderIdOk() (*string, bool)`

GetServiceProviderIdOk returns a tuple with the ServiceProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviderId

`func (o *CreateProxyResourceInstanceRequest) SetServiceProviderId(v string)`

SetServiceProviderId sets ServiceProviderId field to given value.


### GetSourceResourceKey

`func (o *CreateProxyResourceInstanceRequest) GetSourceResourceKey() string`

GetSourceResourceKey returns the SourceResourceKey field if non-nil, zero value otherwise.

### GetSourceResourceKeyOk

`func (o *CreateProxyResourceInstanceRequest) GetSourceResourceKeyOk() (*string, bool)`

GetSourceResourceKeyOk returns a tuple with the SourceResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceResourceKey

`func (o *CreateProxyResourceInstanceRequest) SetSourceResourceKey(v string)`

SetSourceResourceKey sets SourceResourceKey field to given value.


### GetToken

`func (o *CreateProxyResourceInstanceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateProxyResourceInstanceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateProxyResourceInstanceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


