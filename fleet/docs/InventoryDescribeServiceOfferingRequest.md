# InventoryDescribeServiceOfferingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Visibility** | Pointer to **string** | This parameter is used to configure the visibility of the service control-plane APIs | [optional] 

## Methods

### NewInventoryDescribeServiceOfferingRequest

`func NewInventoryDescribeServiceOfferingRequest(serviceId string, token string, ) *InventoryDescribeServiceOfferingRequest`

NewInventoryDescribeServiceOfferingRequest instantiates a new InventoryDescribeServiceOfferingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryDescribeServiceOfferingRequestWithDefaults

`func NewInventoryDescribeServiceOfferingRequestWithDefaults() *InventoryDescribeServiceOfferingRequest`

NewInventoryDescribeServiceOfferingRequestWithDefaults instantiates a new InventoryDescribeServiceOfferingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *InventoryDescribeServiceOfferingRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *InventoryDescribeServiceOfferingRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *InventoryDescribeServiceOfferingRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *InventoryDescribeServiceOfferingRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *InventoryDescribeServiceOfferingRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *InventoryDescribeServiceOfferingRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *InventoryDescribeServiceOfferingRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *InventoryDescribeServiceOfferingRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetServiceId

`func (o *InventoryDescribeServiceOfferingRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *InventoryDescribeServiceOfferingRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *InventoryDescribeServiceOfferingRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *InventoryDescribeServiceOfferingRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InventoryDescribeServiceOfferingRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InventoryDescribeServiceOfferingRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVisibility

`func (o *InventoryDescribeServiceOfferingRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InventoryDescribeServiceOfferingRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InventoryDescribeServiceOfferingRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InventoryDescribeServiceOfferingRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


