# SandboxCheckoutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalRef** | **string** | The armed purchase&#39;s reference on the simulated channel, which is what the purchase control returned. Resolved against the CALLER&#39;S sandbox: a reference belonging to another organization is simply not found | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSandboxCheckoutRequest

`func NewSandboxCheckoutRequest(externalRef string, token string, ) *SandboxCheckoutRequest`

NewSandboxCheckoutRequest instantiates a new SandboxCheckoutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxCheckoutRequestWithDefaults

`func NewSandboxCheckoutRequestWithDefaults() *SandboxCheckoutRequest`

NewSandboxCheckoutRequestWithDefaults instantiates a new SandboxCheckoutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalRef

`func (o *SandboxCheckoutRequest) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *SandboxCheckoutRequest) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *SandboxCheckoutRequest) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.


### GetToken

`func (o *SandboxCheckoutRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SandboxCheckoutRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SandboxCheckoutRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


