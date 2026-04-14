# NebiusAccountBindingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedPublicKeyFingerprint** | Pointer to **string** | The fingerprint derived locally from the configured Nebius private key&#39;s public key for this binding | [optional] 
**KeyExpiresAt** | Pointer to **time.Time** | When the live Nebius auth public key expires for this binding | [optional] 
**KeyFingerprint** | Pointer to **string** | The live fingerprint reported for the Nebius auth public key for this binding | [optional] 
**KeyState** | Pointer to **string** | The live Nebius auth public key state returned during verification for this binding | [optional] 
**OwnsArtifactBucket** | Pointer to **bool** | Whether this binding owns the Nebius artifact bucket for the account config | [optional] 
**ProjectID** | **string** | The Nebius project ID for this binding | 
**PublicKeyID** | **string** | The Nebius public key ID for this binding | 
**PublicKeyIDMatches** | Pointer to **bool** | Whether the configured Nebius public key ID resolves to a live auth public key whose key material matches the configured private key for this binding | [optional] 
**Region** | **string** | The Nebius region served by this binding | 
**ServiceAccountID** | **string** | The Nebius service account ID for this binding | 
**ServiceAccountKeyValidated** | Pointer to **bool** | Whether the live Nebius auth public key belongs to the configured Nebius service account for this binding | [optional] 
**Status** | Pointer to **string** | The status of the account configuration | [optional] 
**StatusMessage** | Pointer to **string** | The status message of this binding | [optional] 

## Methods

### NewNebiusAccountBindingResult

`func NewNebiusAccountBindingResult(projectID string, publicKeyID string, region string, serviceAccountID string, ) *NebiusAccountBindingResult`

NewNebiusAccountBindingResult instantiates a new NebiusAccountBindingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNebiusAccountBindingResultWithDefaults

`func NewNebiusAccountBindingResultWithDefaults() *NebiusAccountBindingResult`

NewNebiusAccountBindingResultWithDefaults instantiates a new NebiusAccountBindingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedPublicKeyFingerprint

`func (o *NebiusAccountBindingResult) GetDerivedPublicKeyFingerprint() string`

GetDerivedPublicKeyFingerprint returns the DerivedPublicKeyFingerprint field if non-nil, zero value otherwise.

### GetDerivedPublicKeyFingerprintOk

`func (o *NebiusAccountBindingResult) GetDerivedPublicKeyFingerprintOk() (*string, bool)`

GetDerivedPublicKeyFingerprintOk returns a tuple with the DerivedPublicKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedPublicKeyFingerprint

`func (o *NebiusAccountBindingResult) SetDerivedPublicKeyFingerprint(v string)`

SetDerivedPublicKeyFingerprint sets DerivedPublicKeyFingerprint field to given value.

### HasDerivedPublicKeyFingerprint

`func (o *NebiusAccountBindingResult) HasDerivedPublicKeyFingerprint() bool`

HasDerivedPublicKeyFingerprint returns a boolean if a field has been set.

### GetKeyExpiresAt

`func (o *NebiusAccountBindingResult) GetKeyExpiresAt() time.Time`

GetKeyExpiresAt returns the KeyExpiresAt field if non-nil, zero value otherwise.

### GetKeyExpiresAtOk

`func (o *NebiusAccountBindingResult) GetKeyExpiresAtOk() (*time.Time, bool)`

GetKeyExpiresAtOk returns a tuple with the KeyExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiresAt

`func (o *NebiusAccountBindingResult) SetKeyExpiresAt(v time.Time)`

SetKeyExpiresAt sets KeyExpiresAt field to given value.

### HasKeyExpiresAt

`func (o *NebiusAccountBindingResult) HasKeyExpiresAt() bool`

HasKeyExpiresAt returns a boolean if a field has been set.

### GetKeyFingerprint

`func (o *NebiusAccountBindingResult) GetKeyFingerprint() string`

GetKeyFingerprint returns the KeyFingerprint field if non-nil, zero value otherwise.

### GetKeyFingerprintOk

`func (o *NebiusAccountBindingResult) GetKeyFingerprintOk() (*string, bool)`

GetKeyFingerprintOk returns a tuple with the KeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFingerprint

`func (o *NebiusAccountBindingResult) SetKeyFingerprint(v string)`

SetKeyFingerprint sets KeyFingerprint field to given value.

### HasKeyFingerprint

`func (o *NebiusAccountBindingResult) HasKeyFingerprint() bool`

HasKeyFingerprint returns a boolean if a field has been set.

### GetKeyState

`func (o *NebiusAccountBindingResult) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *NebiusAccountBindingResult) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *NebiusAccountBindingResult) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *NebiusAccountBindingResult) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetOwnsArtifactBucket

`func (o *NebiusAccountBindingResult) GetOwnsArtifactBucket() bool`

GetOwnsArtifactBucket returns the OwnsArtifactBucket field if non-nil, zero value otherwise.

### GetOwnsArtifactBucketOk

`func (o *NebiusAccountBindingResult) GetOwnsArtifactBucketOk() (*bool, bool)`

GetOwnsArtifactBucketOk returns a tuple with the OwnsArtifactBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnsArtifactBucket

`func (o *NebiusAccountBindingResult) SetOwnsArtifactBucket(v bool)`

SetOwnsArtifactBucket sets OwnsArtifactBucket field to given value.

### HasOwnsArtifactBucket

`func (o *NebiusAccountBindingResult) HasOwnsArtifactBucket() bool`

HasOwnsArtifactBucket returns a boolean if a field has been set.

### GetProjectID

`func (o *NebiusAccountBindingResult) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *NebiusAccountBindingResult) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *NebiusAccountBindingResult) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.


### GetPublicKeyID

`func (o *NebiusAccountBindingResult) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *NebiusAccountBindingResult) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *NebiusAccountBindingResult) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.


### GetPublicKeyIDMatches

`func (o *NebiusAccountBindingResult) GetPublicKeyIDMatches() bool`

GetPublicKeyIDMatches returns the PublicKeyIDMatches field if non-nil, zero value otherwise.

### GetPublicKeyIDMatchesOk

`func (o *NebiusAccountBindingResult) GetPublicKeyIDMatchesOk() (*bool, bool)`

GetPublicKeyIDMatchesOk returns a tuple with the PublicKeyIDMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyIDMatches

`func (o *NebiusAccountBindingResult) SetPublicKeyIDMatches(v bool)`

SetPublicKeyIDMatches sets PublicKeyIDMatches field to given value.

### HasPublicKeyIDMatches

`func (o *NebiusAccountBindingResult) HasPublicKeyIDMatches() bool`

HasPublicKeyIDMatches returns a boolean if a field has been set.

### GetRegion

`func (o *NebiusAccountBindingResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NebiusAccountBindingResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NebiusAccountBindingResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetServiceAccountID

`func (o *NebiusAccountBindingResult) GetServiceAccountID() string`

GetServiceAccountID returns the ServiceAccountID field if non-nil, zero value otherwise.

### GetServiceAccountIDOk

`func (o *NebiusAccountBindingResult) GetServiceAccountIDOk() (*string, bool)`

GetServiceAccountIDOk returns a tuple with the ServiceAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountID

`func (o *NebiusAccountBindingResult) SetServiceAccountID(v string)`

SetServiceAccountID sets ServiceAccountID field to given value.


### GetServiceAccountKeyValidated

`func (o *NebiusAccountBindingResult) GetServiceAccountKeyValidated() bool`

GetServiceAccountKeyValidated returns the ServiceAccountKeyValidated field if non-nil, zero value otherwise.

### GetServiceAccountKeyValidatedOk

`func (o *NebiusAccountBindingResult) GetServiceAccountKeyValidatedOk() (*bool, bool)`

GetServiceAccountKeyValidatedOk returns a tuple with the ServiceAccountKeyValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKeyValidated

`func (o *NebiusAccountBindingResult) SetServiceAccountKeyValidated(v bool)`

SetServiceAccountKeyValidated sets ServiceAccountKeyValidated field to given value.

### HasServiceAccountKeyValidated

`func (o *NebiusAccountBindingResult) HasServiceAccountKeyValidated() bool`

HasServiceAccountKeyValidated returns a boolean if a field has been set.

### GetStatus

`func (o *NebiusAccountBindingResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NebiusAccountBindingResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NebiusAccountBindingResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NebiusAccountBindingResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *NebiusAccountBindingResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *NebiusAccountBindingResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *NebiusAccountBindingResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *NebiusAccountBindingResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


