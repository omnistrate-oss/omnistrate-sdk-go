# RotateSandboxSecretResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningSecret** | **string** | The new secret, in plaintext. This is the ONLY time it is ever returned. During rotation both the old and the new secret are accepted on verification, so an ISV can install the new one without a flag day | 
**SigningSecretId** | **string** |  | 

## Methods

### NewRotateSandboxSecretResult

`func NewRotateSandboxSecretResult(signingSecret string, signingSecretId string, ) *RotateSandboxSecretResult`

NewRotateSandboxSecretResult instantiates a new RotateSandboxSecretResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateSandboxSecretResultWithDefaults

`func NewRotateSandboxSecretResultWithDefaults() *RotateSandboxSecretResult`

NewRotateSandboxSecretResultWithDefaults instantiates a new RotateSandboxSecretResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningSecret

`func (o *RotateSandboxSecretResult) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *RotateSandboxSecretResult) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *RotateSandboxSecretResult) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.


### GetSigningSecretId

`func (o *RotateSandboxSecretResult) GetSigningSecretId() string`

GetSigningSecretId returns the SigningSecretId field if non-nil, zero value otherwise.

### GetSigningSecretIdOk

`func (o *RotateSandboxSecretResult) GetSigningSecretIdOk() (*string, bool)`

GetSigningSecretIdOk returns a tuple with the SigningSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecretId

`func (o *RotateSandboxSecretResult) SetSigningSecretId(v string)`

SetSigningSecretId sets SigningSecretId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


