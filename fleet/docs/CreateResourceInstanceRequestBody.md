# CreateResourceInstanceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 

## Methods

### NewCreateResourceInstanceRequestBody

`func NewCreateResourceInstanceRequestBody() *CreateResourceInstanceRequestBody`

NewCreateResourceInstanceRequestBody instantiates a new CreateResourceInstanceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceInstanceRequestBodyWithDefaults

`func NewCreateResourceInstanceRequestBodyWithDefaults() *CreateResourceInstanceRequestBody`

NewCreateResourceInstanceRequestBodyWithDefaults instantiates a new CreateResourceInstanceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateResourceInstanceRequestBody) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateResourceInstanceRequestBody) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateResourceInstanceRequestBody) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateResourceInstanceRequestBody) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCustomNetworkId

`func (o *CreateResourceInstanceRequestBody) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *CreateResourceInstanceRequestBody) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *CreateResourceInstanceRequestBody) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *CreateResourceInstanceRequestBody) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetNetworkType

`func (o *CreateResourceInstanceRequestBody) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *CreateResourceInstanceRequestBody) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *CreateResourceInstanceRequestBody) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *CreateResourceInstanceRequestBody) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *CreateResourceInstanceRequestBody) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *CreateResourceInstanceRequestBody) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *CreateResourceInstanceRequestBody) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *CreateResourceInstanceRequestBody) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetRegion

`func (o *CreateResourceInstanceRequestBody) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateResourceInstanceRequestBody) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateResourceInstanceRequestBody) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateResourceInstanceRequestBody) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestParams

`func (o *CreateResourceInstanceRequestBody) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *CreateResourceInstanceRequestBody) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *CreateResourceInstanceRequestBody) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *CreateResourceInstanceRequestBody) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *CreateResourceInstanceRequestBody) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *CreateResourceInstanceRequestBody) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil
### GetSubscriptionId

`func (o *CreateResourceInstanceRequestBody) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *CreateResourceInstanceRequestBody) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *CreateResourceInstanceRequestBody) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *CreateResourceInstanceRequestBody) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


