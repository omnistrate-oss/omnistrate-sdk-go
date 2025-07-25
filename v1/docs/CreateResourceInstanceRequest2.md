# CreateResourceInstanceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | Pointer to **string** | The cloud provider name | [optional] 
**CustomNetworkId** | Pointer to **string** | Custom network for resource | [optional] 
**ExternalBillingId** | Pointer to **string** | This externalBillingId is deprecated and will be removed in the future | [optional] 
**NetworkType** | Pointer to **string** | The network type | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**Region** | Pointer to **string** | The region code | [optional] 
**RequestParams** | Pointer to **interface{}** | The request parameters | [optional] 

## Methods

### NewCreateResourceInstanceRequest2

`func NewCreateResourceInstanceRequest2() *CreateResourceInstanceRequest2`

NewCreateResourceInstanceRequest2 instantiates a new CreateResourceInstanceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResourceInstanceRequest2WithDefaults

`func NewCreateResourceInstanceRequest2WithDefaults() *CreateResourceInstanceRequest2`

NewCreateResourceInstanceRequest2WithDefaults instantiates a new CreateResourceInstanceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProvider

`func (o *CreateResourceInstanceRequest2) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CreateResourceInstanceRequest2) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CreateResourceInstanceRequest2) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *CreateResourceInstanceRequest2) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCustomNetworkId

`func (o *CreateResourceInstanceRequest2) GetCustomNetworkId() string`

GetCustomNetworkId returns the CustomNetworkId field if non-nil, zero value otherwise.

### GetCustomNetworkIdOk

`func (o *CreateResourceInstanceRequest2) GetCustomNetworkIdOk() (*string, bool)`

GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworkId

`func (o *CreateResourceInstanceRequest2) SetCustomNetworkId(v string)`

SetCustomNetworkId sets CustomNetworkId field to given value.

### HasCustomNetworkId

`func (o *CreateResourceInstanceRequest2) HasCustomNetworkId() bool`

HasCustomNetworkId returns a boolean if a field has been set.

### GetExternalBillingId

`func (o *CreateResourceInstanceRequest2) GetExternalBillingId() string`

GetExternalBillingId returns the ExternalBillingId field if non-nil, zero value otherwise.

### GetExternalBillingIdOk

`func (o *CreateResourceInstanceRequest2) GetExternalBillingIdOk() (*string, bool)`

GetExternalBillingIdOk returns a tuple with the ExternalBillingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalBillingId

`func (o *CreateResourceInstanceRequest2) SetExternalBillingId(v string)`

SetExternalBillingId sets ExternalBillingId field to given value.

### HasExternalBillingId

`func (o *CreateResourceInstanceRequest2) HasExternalBillingId() bool`

HasExternalBillingId returns a boolean if a field has been set.

### GetNetworkType

`func (o *CreateResourceInstanceRequest2) GetNetworkType() string`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *CreateResourceInstanceRequest2) GetNetworkTypeOk() (*string, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *CreateResourceInstanceRequest2) SetNetworkType(v string)`

SetNetworkType sets NetworkType field to given value.

### HasNetworkType

`func (o *CreateResourceInstanceRequest2) HasNetworkType() bool`

HasNetworkType returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *CreateResourceInstanceRequest2) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *CreateResourceInstanceRequest2) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *CreateResourceInstanceRequest2) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *CreateResourceInstanceRequest2) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetRegion

`func (o *CreateResourceInstanceRequest2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CreateResourceInstanceRequest2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CreateResourceInstanceRequest2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CreateResourceInstanceRequest2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRequestParams

`func (o *CreateResourceInstanceRequest2) GetRequestParams() interface{}`

GetRequestParams returns the RequestParams field if non-nil, zero value otherwise.

### GetRequestParamsOk

`func (o *CreateResourceInstanceRequest2) GetRequestParamsOk() (*interface{}, bool)`

GetRequestParamsOk returns a tuple with the RequestParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestParams

`func (o *CreateResourceInstanceRequest2) SetRequestParams(v interface{})`

SetRequestParams sets RequestParams field to given value.

### HasRequestParams

`func (o *CreateResourceInstanceRequest2) HasRequestParams() bool`

HasRequestParams returns a boolean if a field has been set.

### SetRequestParamsNil

`func (o *CreateResourceInstanceRequest2) SetRequestParamsNil(b bool)`

 SetRequestParamsNil sets the value for RequestParams to be an explicit nil

### UnsetRequestParams
`func (o *CreateResourceInstanceRequest2) UnsetRequestParams()`

UnsetRequestParams ensures that no value is present for RequestParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


