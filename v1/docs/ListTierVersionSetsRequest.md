# ListTierVersionSetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestIncrementalVersionForMajorVersion** | Pointer to **string** | Returns the latest incremental version for the given major version. The parameter needs to be specified in isolation | [optional] 
**LatestMajorVersionOnly** | Pointer to **bool** | Indicates whether to return only the latest version set. The parameter needs to be specified in isolation. | [optional] 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListTierVersionSetsRequest

`func NewListTierVersionSetsRequest(productTierId string, serviceId string, token string, ) *ListTierVersionSetsRequest`

NewListTierVersionSetsRequest instantiates a new ListTierVersionSetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTierVersionSetsRequestWithDefaults

`func NewListTierVersionSetsRequestWithDefaults() *ListTierVersionSetsRequest`

NewListTierVersionSetsRequestWithDefaults instantiates a new ListTierVersionSetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestIncrementalVersionForMajorVersion

`func (o *ListTierVersionSetsRequest) GetLatestIncrementalVersionForMajorVersion() string`

GetLatestIncrementalVersionForMajorVersion returns the LatestIncrementalVersionForMajorVersion field if non-nil, zero value otherwise.

### GetLatestIncrementalVersionForMajorVersionOk

`func (o *ListTierVersionSetsRequest) GetLatestIncrementalVersionForMajorVersionOk() (*string, bool)`

GetLatestIncrementalVersionForMajorVersionOk returns a tuple with the LatestIncrementalVersionForMajorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestIncrementalVersionForMajorVersion

`func (o *ListTierVersionSetsRequest) SetLatestIncrementalVersionForMajorVersion(v string)`

SetLatestIncrementalVersionForMajorVersion sets LatestIncrementalVersionForMajorVersion field to given value.

### HasLatestIncrementalVersionForMajorVersion

`func (o *ListTierVersionSetsRequest) HasLatestIncrementalVersionForMajorVersion() bool`

HasLatestIncrementalVersionForMajorVersion returns a boolean if a field has been set.

### GetLatestMajorVersionOnly

`func (o *ListTierVersionSetsRequest) GetLatestMajorVersionOnly() bool`

GetLatestMajorVersionOnly returns the LatestMajorVersionOnly field if non-nil, zero value otherwise.

### GetLatestMajorVersionOnlyOk

`func (o *ListTierVersionSetsRequest) GetLatestMajorVersionOnlyOk() (*bool, bool)`

GetLatestMajorVersionOnlyOk returns a tuple with the LatestMajorVersionOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestMajorVersionOnly

`func (o *ListTierVersionSetsRequest) SetLatestMajorVersionOnly(v bool)`

SetLatestMajorVersionOnly sets LatestMajorVersionOnly field to given value.

### HasLatestMajorVersionOnly

`func (o *ListTierVersionSetsRequest) HasLatestMajorVersionOnly() bool`

HasLatestMajorVersionOnly returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListTierVersionSetsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListTierVersionSetsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListTierVersionSetsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListTierVersionSetsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetPageSize

`func (o *ListTierVersionSetsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListTierVersionSetsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListTierVersionSetsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ListTierVersionSetsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProductTierId

`func (o *ListTierVersionSetsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *ListTierVersionSetsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *ListTierVersionSetsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *ListTierVersionSetsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListTierVersionSetsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListTierVersionSetsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *ListTierVersionSetsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListTierVersionSetsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListTierVersionSetsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


