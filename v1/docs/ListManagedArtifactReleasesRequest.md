# ListManagedArtifactReleasesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleVersion** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] [default to 20]
**NextPageToken** | Pointer to **string** | Opaque token returned by the previous list response. | [optional] 
**ReleasedAfter** | Pointer to **time.Time** |  | [optional] 
**ReleasedBefore** | Pointer to **time.Time** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListManagedArtifactReleasesRequest

`func NewListManagedArtifactReleasesRequest(token string, ) *ListManagedArtifactReleasesRequest`

NewListManagedArtifactReleasesRequest instantiates a new ListManagedArtifactReleasesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedArtifactReleasesRequestWithDefaults

`func NewListManagedArtifactReleasesRequestWithDefaults() *ListManagedArtifactReleasesRequest`

NewListManagedArtifactReleasesRequestWithDefaults instantiates a new ListManagedArtifactReleasesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleVersion

`func (o *ListManagedArtifactReleasesRequest) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ListManagedArtifactReleasesRequest) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ListManagedArtifactReleasesRequest) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *ListManagedArtifactReleasesRequest) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetLimit

`func (o *ListManagedArtifactReleasesRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListManagedArtifactReleasesRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListManagedArtifactReleasesRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListManagedArtifactReleasesRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListManagedArtifactReleasesRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListManagedArtifactReleasesRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListManagedArtifactReleasesRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListManagedArtifactReleasesRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetReleasedAfter

`func (o *ListManagedArtifactReleasesRequest) GetReleasedAfter() time.Time`

GetReleasedAfter returns the ReleasedAfter field if non-nil, zero value otherwise.

### GetReleasedAfterOk

`func (o *ListManagedArtifactReleasesRequest) GetReleasedAfterOk() (*time.Time, bool)`

GetReleasedAfterOk returns a tuple with the ReleasedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAfter

`func (o *ListManagedArtifactReleasesRequest) SetReleasedAfter(v time.Time)`

SetReleasedAfter sets ReleasedAfter field to given value.

### HasReleasedAfter

`func (o *ListManagedArtifactReleasesRequest) HasReleasedAfter() bool`

HasReleasedAfter returns a boolean if a field has been set.

### GetReleasedBefore

`func (o *ListManagedArtifactReleasesRequest) GetReleasedBefore() time.Time`

GetReleasedBefore returns the ReleasedBefore field if non-nil, zero value otherwise.

### GetReleasedBeforeOk

`func (o *ListManagedArtifactReleasesRequest) GetReleasedBeforeOk() (*time.Time, bool)`

GetReleasedBeforeOk returns a tuple with the ReleasedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedBefore

`func (o *ListManagedArtifactReleasesRequest) SetReleasedBefore(v time.Time)`

SetReleasedBefore sets ReleasedBefore field to given value.

### HasReleasedBefore

`func (o *ListManagedArtifactReleasesRequest) HasReleasedBefore() bool`

HasReleasedBefore returns a boolean if a field has been set.

### GetToken

`func (o *ListManagedArtifactReleasesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListManagedArtifactReleasesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListManagedArtifactReleasesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


