# UpdateAccountConfigNebiusBindingInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKeyPEM** | **string** | The Nebius private key PEM for this binding. This is input-only and must never be returned in safe describe/list responses | 
**ProjectID** | **string** | The Nebius project ID for this binding | 
**PublicKeyID** | **string** | The Nebius public key ID for this binding | 
**ServiceAccountID** | **string** | The Nebius service account ID for this binding | 

## Methods

### NewUpdateAccountConfigNebiusBindingInput

`func NewUpdateAccountConfigNebiusBindingInput(privateKeyPEM string, projectID string, publicKeyID string, serviceAccountID string, ) *UpdateAccountConfigNebiusBindingInput`

NewUpdateAccountConfigNebiusBindingInput instantiates a new UpdateAccountConfigNebiusBindingInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountConfigNebiusBindingInputWithDefaults

`func NewUpdateAccountConfigNebiusBindingInputWithDefaults() *UpdateAccountConfigNebiusBindingInput`

NewUpdateAccountConfigNebiusBindingInputWithDefaults instantiates a new UpdateAccountConfigNebiusBindingInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKeyPEM

`func (o *UpdateAccountConfigNebiusBindingInput) GetPrivateKeyPEM() string`

GetPrivateKeyPEM returns the PrivateKeyPEM field if non-nil, zero value otherwise.

### GetPrivateKeyPEMOk

`func (o *UpdateAccountConfigNebiusBindingInput) GetPrivateKeyPEMOk() (*string, bool)`

GetPrivateKeyPEMOk returns a tuple with the PrivateKeyPEM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPEM

`func (o *UpdateAccountConfigNebiusBindingInput) SetPrivateKeyPEM(v string)`

SetPrivateKeyPEM sets PrivateKeyPEM field to given value.


### GetProjectID

`func (o *UpdateAccountConfigNebiusBindingInput) GetProjectID() string`

GetProjectID returns the ProjectID field if non-nil, zero value otherwise.

### GetProjectIDOk

`func (o *UpdateAccountConfigNebiusBindingInput) GetProjectIDOk() (*string, bool)`

GetProjectIDOk returns a tuple with the ProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectID

`func (o *UpdateAccountConfigNebiusBindingInput) SetProjectID(v string)`

SetProjectID sets ProjectID field to given value.


### GetPublicKeyID

`func (o *UpdateAccountConfigNebiusBindingInput) GetPublicKeyID() string`

GetPublicKeyID returns the PublicKeyID field if non-nil, zero value otherwise.

### GetPublicKeyIDOk

`func (o *UpdateAccountConfigNebiusBindingInput) GetPublicKeyIDOk() (*string, bool)`

GetPublicKeyIDOk returns a tuple with the PublicKeyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyID

`func (o *UpdateAccountConfigNebiusBindingInput) SetPublicKeyID(v string)`

SetPublicKeyID sets PublicKeyID field to given value.


### GetServiceAccountID

`func (o *UpdateAccountConfigNebiusBindingInput) GetServiceAccountID() string`

GetServiceAccountID returns the ServiceAccountID field if non-nil, zero value otherwise.

### GetServiceAccountIDOk

`func (o *UpdateAccountConfigNebiusBindingInput) GetServiceAccountIDOk() (*string, bool)`

GetServiceAccountIDOk returns a tuple with the ServiceAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountID

`func (o *UpdateAccountConfigNebiusBindingInput) SetServiceAccountID(v string)`

SetServiceAccountID sets ServiceAccountID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


