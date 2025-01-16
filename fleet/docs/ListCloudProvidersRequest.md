# ListCloudProvidersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceModelId** | Pointer to **string** | ID of a Service Model | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListCloudProvidersRequest

`func NewListCloudProvidersRequest(token string, ) *ListCloudProvidersRequest`

NewListCloudProvidersRequest instantiates a new ListCloudProvidersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudProvidersRequestWithDefaults

`func NewListCloudProvidersRequestWithDefaults() *ListCloudProvidersRequest`

NewListCloudProvidersRequestWithDefaults instantiates a new ListCloudProvidersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *ListCloudProvidersRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListCloudProvidersRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListCloudProvidersRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ListCloudProvidersRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetServiceId

`func (o *ListCloudProvidersRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListCloudProvidersRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListCloudProvidersRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ListCloudProvidersRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceModelId

`func (o *ListCloudProvidersRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *ListCloudProvidersRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *ListCloudProvidersRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.

### HasServiceModelId

`func (o *ListCloudProvidersRequest) HasServiceModelId() bool`

HasServiceModelId returns a boolean if a field has been set.

### GetToken

`func (o *ListCloudProvidersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListCloudProvidersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListCloudProvidersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


