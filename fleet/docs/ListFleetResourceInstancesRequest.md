# ListFleetResourceInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDetail** | Pointer to **bool** | Whether to exclude detailed information about the resource instances | [optional] [default to false]
**Filter** | Pointer to **string** | Filter to list resources. | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierVersion** | Pointer to **string** | The product tier version of the infra config to describe. If not specified, the latest version is described. | [optional] 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 
**EnvironmentId** | **string** | ID of a Service Environment | 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListFleetResourceInstancesRequest

`func NewListFleetResourceInstancesRequest(environmentId string, serviceId string, token string, ) *ListFleetResourceInstancesRequest`

NewListFleetResourceInstancesRequest instantiates a new ListFleetResourceInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFleetResourceInstancesRequestWithDefaults

`func NewListFleetResourceInstancesRequestWithDefaults() *ListFleetResourceInstancesRequest`

NewListFleetResourceInstancesRequestWithDefaults instantiates a new ListFleetResourceInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDetail

`func (o *ListFleetResourceInstancesRequest) GetExcludeDetail() bool`

GetExcludeDetail returns the ExcludeDetail field if non-nil, zero value otherwise.

### GetExcludeDetailOk

`func (o *ListFleetResourceInstancesRequest) GetExcludeDetailOk() (*bool, bool)`

GetExcludeDetailOk returns a tuple with the ExcludeDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDetail

`func (o *ListFleetResourceInstancesRequest) SetExcludeDetail(v bool)`

SetExcludeDetail sets ExcludeDetail field to given value.

### HasExcludeDetail

`func (o *ListFleetResourceInstancesRequest) HasExcludeDetail() bool`

HasExcludeDetail returns a boolean if a field has been set.

### GetFilter

`func (o *ListFleetResourceInstancesRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListFleetResourceInstancesRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListFleetResourceInstancesRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListFleetResourceInstancesRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListFleetResourceInstancesRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListFleetResourceInstancesRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListFleetResourceInstancesRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *ListFleetResourceInstancesRequest) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetProductTierVersion

`func (o *ListFleetResourceInstancesRequest) GetProductTierVersion() string`

GetProductTierVersion returns the ProductTierVersion field if non-nil, zero value otherwise.

### GetProductTierVersionOk

`func (o *ListFleetResourceInstancesRequest) GetProductTierVersionOk() (*string, bool)`

GetProductTierVersionOk returns a tuple with the ProductTierVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierVersion

`func (o *ListFleetResourceInstancesRequest) SetProductTierVersion(v string)`

SetProductTierVersion sets ProductTierVersion field to given value.

### HasProductTierVersion

`func (o *ListFleetResourceInstancesRequest) HasProductTierVersion() bool`

HasProductTierVersion returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ListFleetResourceInstancesRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ListFleetResourceInstancesRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ListFleetResourceInstancesRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *ListFleetResourceInstancesRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ListFleetResourceInstancesRequest) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ListFleetResourceInstancesRequest) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ListFleetResourceInstancesRequest) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetNextPageToken

`func (o *ListFleetResourceInstancesRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListFleetResourceInstancesRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListFleetResourceInstancesRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListFleetResourceInstancesRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListFleetResourceInstancesRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListFleetResourceInstancesRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListFleetResourceInstancesRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListFleetResourceInstancesRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetServiceId

`func (o *ListFleetResourceInstancesRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListFleetResourceInstancesRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListFleetResourceInstancesRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListFleetResourceInstancesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListFleetResourceInstancesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListFleetResourceInstancesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


