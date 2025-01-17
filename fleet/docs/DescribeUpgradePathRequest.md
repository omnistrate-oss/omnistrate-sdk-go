# DescribeUpgradePathRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewDescribeUpgradePathRequest

`func NewDescribeUpgradePathRequest(productTierId string, serviceId string, token string, upgradePathId string, ) *DescribeUpgradePathRequest`

NewDescribeUpgradePathRequest instantiates a new DescribeUpgradePathRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUpgradePathRequestWithDefaults

`func NewDescribeUpgradePathRequestWithDefaults() *DescribeUpgradePathRequest`

NewDescribeUpgradePathRequestWithDefaults instantiates a new DescribeUpgradePathRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *DescribeUpgradePathRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeUpgradePathRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeUpgradePathRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *DescribeUpgradePathRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeUpgradePathRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeUpgradePathRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DescribeUpgradePathRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeUpgradePathRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeUpgradePathRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpgradePathId

`func (o *DescribeUpgradePathRequest) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *DescribeUpgradePathRequest) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *DescribeUpgradePathRequest) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


