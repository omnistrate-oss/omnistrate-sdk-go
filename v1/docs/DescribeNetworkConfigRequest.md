# DescribeNetworkConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version of the infra config to describe. If not specified, the latest version is described. | [optional] 
**Id** | **string** | ID of a Network Config | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeNetworkConfigRequest

`func NewDescribeNetworkConfigRequest(id string, serviceId string, token string, ) *DescribeNetworkConfigRequest`

NewDescribeNetworkConfigRequest instantiates a new DescribeNetworkConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNetworkConfigRequestWithDefaults

`func NewDescribeNetworkConfigRequestWithDefaults() *DescribeNetworkConfigRequest`

NewDescribeNetworkConfigRequestWithDefaults instantiates a new DescribeNetworkConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *DescribeNetworkConfigRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeNetworkConfigRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeNetworkConfigRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *DescribeNetworkConfigRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *DescribeNetworkConfigRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *DescribeNetworkConfigRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *DescribeNetworkConfigRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *DescribeNetworkConfigRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetId

`func (o *DescribeNetworkConfigRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeNetworkConfigRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeNetworkConfigRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *DescribeNetworkConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeNetworkConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeNetworkConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DescribeNetworkConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeNetworkConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeNetworkConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


