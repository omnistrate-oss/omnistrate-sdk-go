# ManagedReleaseSyncArtifactResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactKey** | **string** |  | 
**BackingChecksum** | Pointer to **string** |  | [optional] 
**BackingDigest** | Pointer to **string** |  | [optional] 
**BackingRef** | Pointer to **string** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**GatewayRef** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedReleaseSyncArtifactResult

`func NewManagedReleaseSyncArtifactResult(artifactKey string, ) *ManagedReleaseSyncArtifactResult`

NewManagedReleaseSyncArtifactResult instantiates a new ManagedReleaseSyncArtifactResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedReleaseSyncArtifactResultWithDefaults

`func NewManagedReleaseSyncArtifactResultWithDefaults() *ManagedReleaseSyncArtifactResult`

NewManagedReleaseSyncArtifactResultWithDefaults instantiates a new ManagedReleaseSyncArtifactResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactKey

`func (o *ManagedReleaseSyncArtifactResult) GetArtifactKey() string`

GetArtifactKey returns the ArtifactKey field if non-nil, zero value otherwise.

### GetArtifactKeyOk

`func (o *ManagedReleaseSyncArtifactResult) GetArtifactKeyOk() (*string, bool)`

GetArtifactKeyOk returns a tuple with the ArtifactKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactKey

`func (o *ManagedReleaseSyncArtifactResult) SetArtifactKey(v string)`

SetArtifactKey sets ArtifactKey field to given value.


### GetBackingChecksum

`func (o *ManagedReleaseSyncArtifactResult) GetBackingChecksum() string`

GetBackingChecksum returns the BackingChecksum field if non-nil, zero value otherwise.

### GetBackingChecksumOk

`func (o *ManagedReleaseSyncArtifactResult) GetBackingChecksumOk() (*string, bool)`

GetBackingChecksumOk returns a tuple with the BackingChecksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingChecksum

`func (o *ManagedReleaseSyncArtifactResult) SetBackingChecksum(v string)`

SetBackingChecksum sets BackingChecksum field to given value.

### HasBackingChecksum

`func (o *ManagedReleaseSyncArtifactResult) HasBackingChecksum() bool`

HasBackingChecksum returns a boolean if a field has been set.

### GetBackingDigest

`func (o *ManagedReleaseSyncArtifactResult) GetBackingDigest() string`

GetBackingDigest returns the BackingDigest field if non-nil, zero value otherwise.

### GetBackingDigestOk

`func (o *ManagedReleaseSyncArtifactResult) GetBackingDigestOk() (*string, bool)`

GetBackingDigestOk returns a tuple with the BackingDigest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingDigest

`func (o *ManagedReleaseSyncArtifactResult) SetBackingDigest(v string)`

SetBackingDigest sets BackingDigest field to given value.

### HasBackingDigest

`func (o *ManagedReleaseSyncArtifactResult) HasBackingDigest() bool`

HasBackingDigest returns a boolean if a field has been set.

### GetBackingRef

`func (o *ManagedReleaseSyncArtifactResult) GetBackingRef() string`

GetBackingRef returns the BackingRef field if non-nil, zero value otherwise.

### GetBackingRefOk

`func (o *ManagedReleaseSyncArtifactResult) GetBackingRefOk() (*string, bool)`

GetBackingRefOk returns a tuple with the BackingRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackingRef

`func (o *ManagedReleaseSyncArtifactResult) SetBackingRef(v string)`

SetBackingRef sets BackingRef field to given value.

### HasBackingRef

`func (o *ManagedReleaseSyncArtifactResult) HasBackingRef() bool`

HasBackingRef returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ManagedReleaseSyncArtifactResult) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ManagedReleaseSyncArtifactResult) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ManagedReleaseSyncArtifactResult) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ManagedReleaseSyncArtifactResult) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetError

`func (o *ManagedReleaseSyncArtifactResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ManagedReleaseSyncArtifactResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ManagedReleaseSyncArtifactResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ManagedReleaseSyncArtifactResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetGatewayRef

`func (o *ManagedReleaseSyncArtifactResult) GetGatewayRef() string`

GetGatewayRef returns the GatewayRef field if non-nil, zero value otherwise.

### GetGatewayRefOk

`func (o *ManagedReleaseSyncArtifactResult) GetGatewayRefOk() (*string, bool)`

GetGatewayRefOk returns a tuple with the GatewayRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayRef

`func (o *ManagedReleaseSyncArtifactResult) SetGatewayRef(v string)`

SetGatewayRef sets GatewayRef field to given value.

### HasGatewayRef

`func (o *ManagedReleaseSyncArtifactResult) HasGatewayRef() bool`

HasGatewayRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


