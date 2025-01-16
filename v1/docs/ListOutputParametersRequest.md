# ListOutputParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version of the infra config to describe. If not specified, the latest version is described. | [optional] 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListOutputParametersRequest

`func NewListOutputParametersRequest(resourceId string, serviceId string, token string, ) *ListOutputParametersRequest`

NewListOutputParametersRequest instantiates a new ListOutputParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOutputParametersRequestWithDefaults

`func NewListOutputParametersRequestWithDefaults() *ListOutputParametersRequest`

NewListOutputParametersRequestWithDefaults instantiates a new ListOutputParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *ListOutputParametersRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListOutputParametersRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListOutputParametersRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ListOutputParametersRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *ListOutputParametersRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ListOutputParametersRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ListOutputParametersRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ListOutputParametersRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetResourceId

`func (o *ListOutputParametersRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ListOutputParametersRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ListOutputParametersRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *ListOutputParametersRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListOutputParametersRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListOutputParametersRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListOutputParametersRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListOutputParametersRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListOutputParametersRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


