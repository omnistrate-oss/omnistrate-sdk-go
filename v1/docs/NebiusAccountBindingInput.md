# NebiusAccountBindingInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OwnsArtifactBucket** | Pointer to **bool** | Whether this binding owns the Nebius artifact bucket for the account config | [optional] 
**PrivateKeyPEM** | **string** | The Nebius private key PEM for this binding. This is input-only and must never be returned in safe describe/list responses | 
**ProjectID** | **string** | The Nebius project ID for this binding | 
**PublicKeyID** | **string** | The Nebius public key ID for this binding | 
**ServiceAccountID** | **string** | The Nebius service account ID for this binding | 

## Methods

### NewNebiusAccountBindingInput

`func NewNebiusAccountBindingInput(privateKeyPEM string, projectID string, publicKeyID string, serviceAccountID string, ) *NebiusAccountBindingInput`

NewNebiusAccountBindingInput instantiates a new NebiusAccountBindingInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNebiusAccountBindingInputWithDefaults

`func NewNebiusAccountBindingInputWithDefaults() *NebiusAccountBindingInput`

NewNebiusAccountBindingInputWithDefaults instantiates a new NebiusAccountBindingInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwnsArtifactBucket

`func (o *NebiusAccountBindingInput) GetOwnsArtifactBucket() bool`

GetOwnsArtifactBucket returns the OwnsArtifactBucket field if non-nil, zero value otherwise.

### GetOwnsArtifactBucketOk

`func (o *NebiusAccountBindingInput) GetOwnsArtifactBucketOk() (*bool, bool)`

GetOwnsArtifactBucketOk returns a tuple with the OwnsArtifactBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnsArtifactBucket

`func (o *NebiusAccountBindingInput) SetOwnsArtifactBucket(v bool)`

SetOwnsArtifactBucket sets OwnsArtifactBucket field to given value.

### HasOwnsArtifactBucket

`func (o *NebiusAccountBindingInput) HasOwnsArtifactBucket() bool`

HasOwnsArtifactBucket returns a boolean if a field has been set.

### GetPrivateKeyPEM

`func (o *NebiusAccountBindingInput) GetPrivateKeyPEM() string`

GetPrivateKeyPEM returns the PrivateKeyPEM field if non-nil, zero value otherwise.

### GetPrivateKeyPEMOk

`func (o *NebiusAccountBindingInput) GetPrivateKeyPEMOk() (*string, bool)`

GetPrivateKeyPEMOk returns a tuple with the PrivateKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPEM

`func (o *NebiusAccountBindingInput) SetPrivateKeyPEM(v string)`

SetPrivateKeyPEM sets PrivateKeyPEM field to given value.


### GetProjectID

`func (o *NebiusAccountBindingInput) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *NebiusAccountBindingInput) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *NebiusAccountBindingInput) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.


### GetPublicKeyID

`func (o *NebiusAccountBindingInput) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *NebiusAccountBindingInput) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *NebiusAccountBindingInput) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.


### GetServiceAccountID

`func (o *NebiusAccountBindingInput) GetServiceAccountID() string`

GetServiceAccountID returns the ServiceAccountID field if non-nil, zero value otherwise.

### GetServiceAccountIDOk

`func (o *NebiusAccountBindingInput) GetServiceAccountIDOk() (*string, bool)`

GetServiceAccountIDOk returns a tuple with the ServiceAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountID

`func (o *NebiusAccountBindingInput) SetServiceAccountID(v string)`

SetServiceAccountID sets ServiceAccountID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


