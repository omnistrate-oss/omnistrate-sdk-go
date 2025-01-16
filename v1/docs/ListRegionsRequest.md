# ListRegionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**ModelType** | Pointer to **string** | The model type encapsulating this service | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceModelId** | Pointer to **string** | ID of a Service Model | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListRegionsRequest

`func NewListRegionsRequest(cloudProviderName string, token string, ) *ListRegionsRequest`

NewListRegionsRequest instantiates a new ListRegionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRegionsRequestWithDefaults

`func NewListRegionsRequestWithDefaults() *ListRegionsRequest`

NewListRegionsRequestWithDefaults instantiates a new ListRegionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *ListRegionsRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ListRegionsRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ListRegionsRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetModelType

`func (o *ListRegionsRequest) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *ListRegionsRequest) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *ListRegionsRequest) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *ListRegionsRequest) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListRegionsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListRegionsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListRegionsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ListRegionsRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *ListRegionsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListRegionsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListRegionsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListRegionsRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceModelId

`func (o *ListRegionsRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ListRegionsRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ListRegionsRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.

### HasServiceModelId

`func (o *ListRegionsRequest) HasServiceModelId() bool`

HasServiceModelId returns a boolean if a field has been set.

### GetToken

`func (o *ListRegionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListRegionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListRegionsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


