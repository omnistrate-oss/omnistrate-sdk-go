# ListManagedReleaseRevisionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeArtifacts** | Pointer to **bool** |  | [optional] [default to false]
**Limit** | Pointer to **int64** |  | [optional] [default to 20]
**ReleaseComponent** | Pointer to **string** | Component pipeline that owns a managed artifact subset. | [optional] 
**Status** | Pointer to **string** | Lifecycle status for a managed release revision or provisioner-target sync. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListManagedReleaseRevisionsRequest

`func NewListManagedReleaseRevisionsRequest(token string, ) *ListManagedReleaseRevisionsRequest`

NewListManagedReleaseRevisionsRequest instantiates a new ListManagedReleaseRevisionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListManagedReleaseRevisionsRequestWithDefaults

`func NewListManagedReleaseRevisionsRequestWithDefaults() *ListManagedReleaseRevisionsRequest`

NewListManagedReleaseRevisionsRequestWithDefaults instantiates a new ListManagedReleaseRevisionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeArtifacts

`func (o *ListManagedReleaseRevisionsRequest) GetIncludeArtifacts() bool`

GetIncludeArtifacts returns the IncludeArtifacts field if non-nil, zero value otherwise.

### GetIncludeArtifactsOk

`func (o *ListManagedReleaseRevisionsRequest) GetIncludeArtifactsOk() (*bool, bool)`

GetIncludeArtifactsOk returns a tuple with the IncludeArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeArtifacts

`func (o *ListManagedReleaseRevisionsRequest) SetIncludeArtifacts(v bool)`

SetIncludeArtifacts sets IncludeArtifacts field to given value.

### HasIncludeArtifacts

`func (o *ListManagedReleaseRevisionsRequest) HasIncludeArtifacts() bool`

HasIncludeArtifacts returns a boolean if a field has been set.

### GetLimit

`func (o *ListManagedReleaseRevisionsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListManagedReleaseRevisionsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListManagedReleaseRevisionsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListManagedReleaseRevisionsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetReleaseComponent

`func (o *ListManagedReleaseRevisionsRequest) GetReleaseComponent() string`

GetReleaseComponent returns the ReleaseComponent field if non-nil, zero value otherwise.

### GetReleaseComponentOk

`func (o *ListManagedReleaseRevisionsRequest) GetReleaseComponentOk() (*string, bool)`

GetReleaseComponentOk returns a tuple with the ReleaseComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseComponent

`func (o *ListManagedReleaseRevisionsRequest) SetReleaseComponent(v string)`

SetReleaseComponent sets ReleaseComponent field to given value.

### HasReleaseComponent

`func (o *ListManagedReleaseRevisionsRequest) HasReleaseComponent() bool`

HasReleaseComponent returns a boolean if a field has been set.

### GetStatus

`func (o *ListManagedReleaseRevisionsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListManagedReleaseRevisionsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListManagedReleaseRevisionsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListManagedReleaseRevisionsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListManagedReleaseRevisionsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListManagedReleaseRevisionsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListManagedReleaseRevisionsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


