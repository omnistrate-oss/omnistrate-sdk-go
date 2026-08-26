# PublishManagedReleaseRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | [**[]ManagedReleaseArtifact2**](ManagedReleaseArtifact2.md) |  | 
**Component** | **string** |  | 
**ComponentVersion** | **string** |  | 
**IdempotencyKey** | **string** |  | 
**SourceCommit** | **string** |  | 

## Methods

### NewPublishManagedReleaseRequest2

`func NewPublishManagedReleaseRequest2(artifacts []ManagedReleaseArtifact2, component string, componentVersion string, idempotencyKey string, sourceCommit string, ) *PublishManagedReleaseRequest2`

NewPublishManagedReleaseRequest2 instantiates a new PublishManagedReleaseRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishManagedReleaseRequest2WithDefaults

`func NewPublishManagedReleaseRequest2WithDefaults() *PublishManagedReleaseRequest2`

NewPublishManagedReleaseRequest2WithDefaults instantiates a new PublishManagedReleaseRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *PublishManagedReleaseRequest2) GetArtifacts() []ManagedReleaseArtifact2`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *PublishManagedReleaseRequest2) GetArtifactsOk() (*[]ManagedReleaseArtifact2, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *PublishManagedReleaseRequest2) SetArtifacts(v []ManagedReleaseArtifact2)`

SetArtifacts sets Artifacts field to given value.


### GetComponent

`func (o *PublishManagedReleaseRequest2) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *PublishManagedReleaseRequest2) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *PublishManagedReleaseRequest2) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetComponentVersion

`func (o *PublishManagedReleaseRequest2) GetComponentVersion() string`

GetComponentVersion returns the ComponentVersion field if non-nil, zero value otherwise.

### GetComponentVersionOk

`func (o *PublishManagedReleaseRequest2) GetComponentVersionOk() (*string, bool)`

GetComponentVersionOk returns a tuple with the ComponentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVersion

`func (o *PublishManagedReleaseRequest2) SetComponentVersion(v string)`

SetComponentVersion sets ComponentVersion field to given value.


### GetIdempotencyKey

`func (o *PublishManagedReleaseRequest2) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *PublishManagedReleaseRequest2) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *PublishManagedReleaseRequest2) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetSourceCommit

`func (o *PublishManagedReleaseRequest2) GetSourceCommit() string`

GetSourceCommit returns the SourceCommit field if non-nil, zero value otherwise.

### GetSourceCommitOk

`func (o *PublishManagedReleaseRequest2) GetSourceCommitOk() (*string, bool)`

GetSourceCommitOk returns a tuple with the SourceCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCommit

`func (o *PublishManagedReleaseRequest2) SetSourceCommit(v string)`

SetSourceCommit sets SourceCommit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


