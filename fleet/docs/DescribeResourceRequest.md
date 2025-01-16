# DescribeResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | ID of a Service Environment | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**ServiceModelId** | **string** | ID of a Service Model | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeResourceRequest

`func NewDescribeResourceRequest(environmentId string, productTierId string, resourceId string, serviceId string, serviceModelId string, token string, ) *DescribeResourceRequest`

NewDescribeResourceRequest instantiates a new DescribeResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceRequestWithDefaults

`func NewDescribeResourceRequestWithDefaults() *DescribeResourceRequest`

NewDescribeResourceRequestWithDefaults instantiates a new DescribeResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *DescribeResourceRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeResourceRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeResourceRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetProductTierId

`func (o *DescribeResourceRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeResourceRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeResourceRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetResourceId

`func (o *DescribeResourceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DescribeResourceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DescribeResourceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *DescribeResourceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeResourceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeResourceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceModelId

`func (o *DescribeResourceRequest) GetServiceModelId() string`

GetServiceModelId returns the ServiceModelId field if non-nil, zero value otherwise.

### GetServiceModelIdOk

`func (o *DescribeResourceRequest) GetServiceModelIdOk() (*string, bool)`

GetServiceModelIdOk returns a tuple with the ServiceModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceModelId

`func (o *DescribeResourceRequest) SetServiceModelId(v string)`

SetServiceModelId sets ServiceModelId field to given value.


### GetToken

`func (o *DescribeResourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeResourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeResourceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


