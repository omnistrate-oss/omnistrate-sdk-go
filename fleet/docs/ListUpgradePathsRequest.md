# ListUpgradePathsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**SourceProductTierVersion** | Pointer to **string** | The source product tier version of the upgrade path. | [optional] 
**Status** | Pointer to **string** | The status of the upgrade path. | [optional] 
**TargetProductTierVersion** | Pointer to **string** | The target product tier version of the upgrade path. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Type** | Pointer to **string** | The type of the upgrade path. | [optional] 

## Methods

### NewListUpgradePathsRequest

`func NewListUpgradePathsRequest(productTierId string, serviceId string, token string, ) *ListUpgradePathsRequest`

NewListUpgradePathsRequest instantiates a new ListUpgradePathsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUpgradePathsRequestWithDefaults

`func NewListUpgradePathsRequestWithDefaults() *ListUpgradePathsRequest`

NewListUpgradePathsRequestWithDefaults instantiates a new ListUpgradePathsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ListUpgradePathsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListUpgradePathsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListUpgradePathsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListUpgradePathsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListUpgradePathsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListUpgradePathsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListUpgradePathsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListUpgradePathsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListUpgradePathsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListUpgradePathsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListUpgradePathsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ListUpgradePathsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListUpgradePathsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListUpgradePathsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSourceProductTierVersion

`func (o *ListUpgradePathsRequest) GetSourceProductTierVersion() string`

GetSourceProductTierVersion returns the SourceProductTierVersion field if non-nil, zero value otherwise.

### GetSourceProductTierVersionOk

`func (o *ListUpgradePathsRequest) GetSourceProductTierVersionOk() (*string, bool)`

GetSourceProductTierVersionOk returns a tuple with the SourceProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProductTierVersion

`func (o *ListUpgradePathsRequest) SetSourceProductTierVersion(v string)`

SetSourceProductTierVersion sets SourceProductTierVersion field to given value.

### HasSourceProductTierVersion

`func (o *ListUpgradePathsRequest) HasSourceProductTierVersion() bool`

HasSourceProductTierVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ListUpgradePathsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListUpgradePathsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListUpgradePathsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListUpgradePathsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetProductTierVersion

`func (o *ListUpgradePathsRequest) GetTargetProductTierVersion() string`

GetTargetProductTierVersion returns the TargetProductTierVersion field if non-nil, zero value otherwise.

### GetTargetProductTierVersionOk

`func (o *ListUpgradePathsRequest) GetTargetProductTierVersionOk() (*string, bool)`

GetTargetProductTierVersionOk returns a tuple with the TargetProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProductTierVersion

`func (o *ListUpgradePathsRequest) SetTargetProductTierVersion(v string)`

SetTargetProductTierVersion sets TargetProductTierVersion field to given value.

### HasTargetProductTierVersion

`func (o *ListUpgradePathsRequest) HasTargetProductTierVersion() bool`

HasTargetProductTierVersion returns a boolean if a field has been set.

### GetToken

`func (o *ListUpgradePathsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListUpgradePathsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListUpgradePathsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetType

`func (o *ListUpgradePathsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListUpgradePathsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListUpgradePathsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListUpgradePathsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


