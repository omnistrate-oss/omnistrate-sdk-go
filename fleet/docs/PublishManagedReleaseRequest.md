# PublishManagedReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]ManagedReleaseArtifact**](ManagedReleaseArtifact.md) |  | 
**Component** | **string** | Component pipeline that owns a managed artifact subset. | 
**ComponentVersion** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**SourceCommit** | **string** |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewPublishManagedReleaseRequest

`func NewPublishManagedReleaseRequest(artifacts []ManagedReleaseArtifact, component string, componentVersion string, idempotencyKey string, sourceCommit string, token string, ) *PublishManagedReleaseRequest`

NewPublishManagedReleaseRequest instantiates a new PublishManagedReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishManagedReleaseRequestWithDefaults

`func NewPublishManagedReleaseRequestWithDefaults() *PublishManagedReleaseRequest`

NewPublishManagedReleaseRequestWithDefaults instantiates a new PublishManagedReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *PublishManagedReleaseRequest) GetArtifacts() []ManagedReleaseArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *PublishManagedReleaseRequest) GetArtifactsOk() (*[]ManagedReleaseArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *PublishManagedReleaseRequest) SetArtifacts(v []ManagedReleaseArtifact)`

SetArtifacts sets Artifacts field to given value.


### GetComponent

`func (o *PublishManagedReleaseRequest) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PublishManagedReleaseRequest) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PublishManagedReleaseRequest) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetComponentVersion

`func (o *PublishManagedReleaseRequest) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *PublishManagedReleaseRequest) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *PublishManagedReleaseRequest) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.


### GetIdempotencyKey

`func (o *PublishManagedReleaseRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *PublishManagedReleaseRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *PublishManagedReleaseRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetSourceCommit

`func (o *PublishManagedReleaseRequest) GetSourceCommit() string`

GetSourceCommit returns the SourceCommit field if non-nil, zero value otherwise.

### GetSourceCommitOk

`func (o *PublishManagedReleaseRequest) GetSourceCommitOk() (*string, bool)`

GetSourceCommitOk returns a tuple with the SourceCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommit

`func (o *PublishManagedReleaseRequest) SetSourceCommit(v string)`

SetSourceCommit sets SourceCommit field to given value.


### GetToken

`func (o *PublishManagedReleaseRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PublishManagedReleaseRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PublishManagedReleaseRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


