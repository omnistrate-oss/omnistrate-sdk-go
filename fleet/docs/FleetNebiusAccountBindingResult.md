# FleetNebiusAccountBindingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DerivedPublicKeyFingerprint** | Pointer to **string** | The fingerprint derived locally from the configured Nebius private key&#39;s public key for this binding | [optional] 
**KeyExpiresAt** | Pointer to **time.Time** | When the live Nebius auth public key expires for this binding | [optional] 
**KeyFingerprint** | Pointer to **string** | The Nebius fingerprint reported by the live auth public key during verification for this binding | [optional] 
**KeyState** | Pointer to **string** | The live Nebius auth public key state returned during verification for this binding | [optional] 
**ProjectID** | Pointer to **string** | The Nebius project ID for this binding | [optional] 
**PublicKeyID** | Pointer to **string** | The Nebius public key ID for this binding | [optional] 
**PublicKeyIDMatches** | Pointer to **bool** | Whether the configured Nebius public key ID resolves to a live auth public key whose key material matches the configured private key for this binding | [optional] 
**Region** | Pointer to **string** | The Nebius region served by this binding | [optional] 
**ServiceAccountID** | Pointer to **string** | The Nebius service account ID for this binding | [optional] 
**ServiceAccountKeyValidated** | Pointer to **bool** | Whether the live Nebius auth public key belongs to the configured Nebius service account for this binding | [optional] 
**Status** | Pointer to **string** | The status of the account configuration | [optional] 
**StatusMessage** | Pointer to **string** | The status message of this binding | [optional] 

## Methods

### NewFleetNebiusAccountBindingResult

`func NewFleetNebiusAccountBindingResult() *FleetNebiusAccountBindingResult`

NewFleetNebiusAccountBindingResult instantiates a new FleetNebiusAccountBindingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetNebiusAccountBindingResultWithDefaults

`func NewFleetNebiusAccountBindingResultWithDefaults() *FleetNebiusAccountBindingResult`

NewFleetNebiusAccountBindingResultWithDefaults instantiates a new FleetNebiusAccountBindingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDerivedPublicKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) GetDerivedPublicKeyFingerprint() string`

GetDerivedPublicKeyFingerprint returns the DerivedPublicKeyFingerprint field if non-nil, zero value otherwise.

### GetDerivedPublicKeyFingerprintOk

`func (o *FleetNebiusAccountBindingResult) GetDerivedPublicKeyFingerprintOk() (*string, bool)`

GetDerivedPublicKeyFingerprintOk returns a tuple with the DerivedPublicKeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedPublicKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) SetDerivedPublicKeyFingerprint(v string)`

SetDerivedPublicKeyFingerprint sets DerivedPublicKeyFingerprint field to given value.

### HasDerivedPublicKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) HasDerivedPublicKeyFingerprint() bool`

HasDerivedPublicKeyFingerprint returns a boolean if a field has been set.

### GetKeyExpiresAt

`func (o *FleetNebiusAccountBindingResult) GetKeyExpiresAt() time.Time`

GetKeyExpiresAt returns the KeyExpiresAt field if non-nil, zero value otherwise.

### GetKeyExpiresAtOk

`func (o *FleetNebiusAccountBindingResult) GetKeyExpiresAtOk() (*time.Time, bool)`

GetKeyExpiresAtOk returns a tuple with the KeyExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiresAt

`func (o *FleetNebiusAccountBindingResult) SetKeyExpiresAt(v time.Time)`

SetKeyExpiresAt sets KeyExpiresAt field to given value.

### HasKeyExpiresAt

`func (o *FleetNebiusAccountBindingResult) HasKeyExpiresAt() bool`

HasKeyExpiresAt returns a boolean if a field has been set.

### GetKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) GetKeyFingerprint() string`

GetKeyFingerprint returns the KeyFingerprint field if non-nil, zero value otherwise.

### GetKeyFingerprintOk

`func (o *FleetNebiusAccountBindingResult) GetKeyFingerprintOk() (*string, bool)`

GetKeyFingerprintOk returns a tuple with the KeyFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) SetKeyFingerprint(v string)`

SetKeyFingerprint sets KeyFingerprint field to given value.

### HasKeyFingerprint

`func (o *FleetNebiusAccountBindingResult) HasKeyFingerprint() bool`

HasKeyFingerprint returns a boolean if a field has been set.

### GetKeyState

`func (o *FleetNebiusAccountBindingResult) GetKeyState() string`

GetKeyState returns the KeyState field if non-nil, zero value otherwise.

### GetKeyStateOk

`func (o *FleetNebiusAccountBindingResult) GetKeyStateOk() (*string, bool)`

GetKeyStateOk returns a tuple with the KeyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyState

`func (o *FleetNebiusAccountBindingResult) SetKeyState(v string)`

SetKeyState sets KeyState field to given value.

### HasKeyState

`func (o *FleetNebiusAccountBindingResult) HasKeyState() bool`

HasKeyState returns a boolean if a field has been set.

### GetProjectID

`func (o *FleetNebiusAccountBindingResult) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *FleetNebiusAccountBindingResult) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *FleetNebiusAccountBindingResult) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.

### HasProjectID

`func (o *FleetNebiusAccountBindingResult) HasProjectID() bool`

HasProjectID returns a boolean if a field has been set.

### GetPublicKeyID

`func (o *FleetNebiusAccountBindingResult) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *FleetNebiusAccountBindingResult) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *FleetNebiusAccountBindingResult) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.

### HasPublicKeyID

`func (o *FleetNebiusAccountBindingResult) HasPublicKeyID() bool`

HasPublicKeyID returns a boolean if a field has been set.

### GetPublicKeyIDMatches

`func (o *FleetNebiusAccountBindingResult) GetPublicKeyIDMatches() bool`

GetPublicKeyIDMatches returns the PublicKeyIDMatches field if non-nil, zero value otherwise.

### GetPublicKeyIDMatchesOk

`func (o *FleetNebiusAccountBindingResult) GetPublicKeyIDMatchesOk() (*bool, bool)`

GetPublicKeyIDMatchesOk returns a tuple with the PublicKeyIDMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyIDMatches

`func (o *FleetNebiusAccountBindingResult) SetPublicKeyIDMatches(v bool)`

SetPublicKeyIDMatches sets PublicKeyIDMatches field to given value.

### HasPublicKeyIDMatches

`func (o *FleetNebiusAccountBindingResult) HasPublicKeyIDMatches() bool`

HasPublicKeyIDMatches returns a boolean if a field has been set.

### GetRegion

`func (o *FleetNebiusAccountBindingResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FleetNebiusAccountBindingResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FleetNebiusAccountBindingResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FleetNebiusAccountBindingResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceAccountID

`func (o *FleetNebiusAccountBindingResult) GetServiceAccountID() string`

GetServiceAccountID returns the ServiceAccountID field if non-nil, zero value otherwise.

### GetServiceAccountIDOk

`func (o *FleetNebiusAccountBindingResult) GetServiceAccountIDOk() (*string, bool)`

GetServiceAccountIDOk returns a tuple with the ServiceAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountID

`func (o *FleetNebiusAccountBindingResult) SetServiceAccountID(v string)`

SetServiceAccountID sets ServiceAccountID field to given value.

### HasServiceAccountID

`func (o *FleetNebiusAccountBindingResult) HasServiceAccountID() bool`

HasServiceAccountID returns a boolean if a field has been set.

### GetServiceAccountKeyValidated

`func (o *FleetNebiusAccountBindingResult) GetServiceAccountKeyValidated() bool`

GetServiceAccountKeyValidated returns the ServiceAccountKeyValidated field if non-nil, zero value otherwise.

### GetServiceAccountKeyValidatedOk

`func (o *FleetNebiusAccountBindingResult) GetServiceAccountKeyValidatedOk() (*bool, bool)`

GetServiceAccountKeyValidatedOk returns a tuple with the ServiceAccountKeyValidated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKeyValidated

`func (o *FleetNebiusAccountBindingResult) SetServiceAccountKeyValidated(v bool)`

SetServiceAccountKeyValidated sets ServiceAccountKeyValidated field to given value.

### HasServiceAccountKeyValidated

`func (o *FleetNebiusAccountBindingResult) HasServiceAccountKeyValidated() bool`

HasServiceAccountKeyValidated returns a boolean if a field has been set.

### GetStatus

`func (o *FleetNebiusAccountBindingResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetNebiusAccountBindingResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetNebiusAccountBindingResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FleetNebiusAccountBindingResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *FleetNebiusAccountBindingResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *FleetNebiusAccountBindingResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *FleetNebiusAccountBindingResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *FleetNebiusAccountBindingResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


