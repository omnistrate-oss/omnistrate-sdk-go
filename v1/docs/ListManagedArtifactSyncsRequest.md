# ListManagedArtifactSyncsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BundleVersion** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] [default to 20]
**NextPageToken** | Pointer to **string** | Opaque token returned by the previous list response. | [optional] 
**Status** | Pointer to **string** | Lifecycle status for one provisioner-target synchronization. | [optional] 
**TargetId** | Pointer to **string** |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**UpdatedAfter** | Pointer to **time.Time** |  | [optional] 
**UpdatedBefore** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListManagedArtifactSyncsRequest

`func NewListManagedArtifactSyncsRequest(token string, ) *ListManagedArtifactSyncsRequest`

NewListManagedArtifactSyncsRequest instantiates a new ListManagedArtifactSyncsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedArtifactSyncsRequestWithDefaults

`func NewListManagedArtifactSyncsRequestWithDefaults() *ListManagedArtifactSyncsRequest`

NewListManagedArtifactSyncsRequestWithDefaults instantiates a new ListManagedArtifactSyncsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBundleVersion

`func (o *ListManagedArtifactSyncsRequest) GetBundleVersion() string`

GetBundleVersion returns the BundleVersion field if non-nil, zero value otherwise.

### GetBundleVersionOk

`func (o *ListManagedArtifactSyncsRequest) GetBundleVersionOk() (*string, bool)`

GetBundleVersionOk returns a tuple with the BundleVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleVersion

`func (o *ListManagedArtifactSyncsRequest) SetBundleVersion(v string)`

SetBundleVersion sets BundleVersion field to given value.

### HasBundleVersion

`func (o *ListManagedArtifactSyncsRequest) HasBundleVersion() bool`

HasBundleVersion returns a boolean if a field has been set.

### GetLimit

`func (o *ListManagedArtifactSyncsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListManagedArtifactSyncsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListManagedArtifactSyncsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListManagedArtifactSyncsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListManagedArtifactSyncsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListManagedArtifactSyncsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListManagedArtifactSyncsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListManagedArtifactSyncsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetStatus

`func (o *ListManagedArtifactSyncsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListManagedArtifactSyncsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListManagedArtifactSyncsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListManagedArtifactSyncsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetId

`func (o *ListManagedArtifactSyncsRequest) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ListManagedArtifactSyncsRequest) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ListManagedArtifactSyncsRequest) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ListManagedArtifactSyncsRequest) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetToken

`func (o *ListManagedArtifactSyncsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListManagedArtifactSyncsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListManagedArtifactSyncsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUpdatedAfter

`func (o *ListManagedArtifactSyncsRequest) GetUpdatedAfter() time.Time`

GetUpdatedAfter returns the UpdatedAfter field if non-nil, zero value otherwise.

### GetUpdatedAfterOk

`func (o *ListManagedArtifactSyncsRequest) GetUpdatedAfterOk() (*time.Time, bool)`

GetUpdatedAfterOk returns a tuple with the UpdatedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAfter

`func (o *ListManagedArtifactSyncsRequest) SetUpdatedAfter(v time.Time)`

SetUpdatedAfter sets UpdatedAfter field to given value.

### HasUpdatedAfter

`func (o *ListManagedArtifactSyncsRequest) HasUpdatedAfter() bool`

HasUpdatedAfter returns a boolean if a field has been set.

### GetUpdatedBefore

`func (o *ListManagedArtifactSyncsRequest) GetUpdatedBefore() time.Time`

GetUpdatedBefore returns the UpdatedBefore field if non-nil, zero value otherwise.

### GetUpdatedBeforeOk

`func (o *ListManagedArtifactSyncsRequest) GetUpdatedBeforeOk() (*time.Time, bool)`

GetUpdatedBeforeOk returns a tuple with the UpdatedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBefore

`func (o *ListManagedArtifactSyncsRequest) SetUpdatedBefore(v time.Time)`

SetUpdatedBefore sets UpdatedBefore field to given value.

### HasUpdatedBefore

`func (o *ListManagedArtifactSyncsRequest) HasUpdatedBefore() bool`

HasUpdatedBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


