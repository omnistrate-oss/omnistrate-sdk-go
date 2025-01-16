# ListEligibleInstancesPerUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The next token to use for pagination. | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**UpgradePathId** | **string** | ID of an Upgrade Path | 

## Methods

### NewListEligibleInstancesPerUpgradeRequest

`func NewListEligibleInstancesPerUpgradeRequest(productTierId string, serviceId string, token string, upgradePathId string, ) *ListEligibleInstancesPerUpgradeRequest`

NewListEligibleInstancesPerUpgradeRequest instantiates a new ListEligibleInstancesPerUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEligibleInstancesPerUpgradeRequestWithDefaults

`func NewListEligibleInstancesPerUpgradeRequestWithDefaults() *ListEligibleInstancesPerUpgradeRequest`

NewListEligibleInstancesPerUpgradeRequestWithDefaults instantiates a new ListEligibleInstancesPerUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListEligibleInstancesPerUpgradeRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListEligibleInstancesPerUpgradeRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListEligibleInstancesPerUpgradeRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListEligibleInstancesPerUpgradeRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListEligibleInstancesPerUpgradeRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListEligibleInstancesPerUpgradeRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListEligibleInstancesPerUpgradeRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ListEligibleInstancesPerUpgradeRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListEligibleInstancesPerUpgradeRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListEligibleInstancesPerUpgradeRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListEligibleInstancesPerUpgradeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListEligibleInstancesPerUpgradeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListEligibleInstancesPerUpgradeRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpgradePathId

`func (o *ListEligibleInstancesPerUpgradeRequest) GetUpgradePathId() string`

GetUpgradePathId returns the UpgradePathId field if non-nil, zero value otherwise.

### GetUpgradePathIdOk

`func (o *ListEligibleInstancesPerUpgradeRequest) GetUpgradePathIdOk() (*string, bool)`

GetUpgradePathIdOk returns a tuple with the UpgradePathId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePathId

`func (o *ListEligibleInstancesPerUpgradeRequest) SetUpgradePathId(v string)`

SetUpgradePathId sets UpgradePathId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


