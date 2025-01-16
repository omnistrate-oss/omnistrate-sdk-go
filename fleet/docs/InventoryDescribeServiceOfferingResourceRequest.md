# InventoryDescribeServiceOfferingResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | Pointer to **string** | The instance ID | [optional] [default to "none"]
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**ResourceId** | **string** | ID of a resource | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewInventoryDescribeServiceOfferingResourceRequest

`func NewInventoryDescribeServiceOfferingResourceRequest(resourceId string, serviceId string, token string, ) *InventoryDescribeServiceOfferingResourceRequest`

NewInventoryDescribeServiceOfferingResourceRequest instantiates a new InventoryDescribeServiceOfferingResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDescribeServiceOfferingResourceRequestWithDefaults

`func NewInventoryDescribeServiceOfferingResourceRequestWithDefaults() *InventoryDescribeServiceOfferingResourceRequest`

NewInventoryDescribeServiceOfferingResourceRequestWithDefaults instantiates a new InventoryDescribeServiceOfferingResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetProductTierId

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *InventoryDescribeServiceOfferingResourceRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *InventoryDescribeServiceOfferingResourceRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetResourceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetServiceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InventoryDescribeServiceOfferingResourceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InventoryDescribeServiceOfferingResourceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


