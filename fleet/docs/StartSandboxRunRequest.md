# StartSandboxRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiverUrl** | **string** | Where to deliver. Validated on write: https only, no embedded credentials, and not a private, loopback, link-local or metadata address | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewStartSandboxRunRequest

`func NewStartSandboxRunRequest(receiverUrl string, token string, ) *StartSandboxRunRequest`

NewStartSandboxRunRequest instantiates a new StartSandboxRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartSandboxRunRequestWithDefaults

`func NewStartSandboxRunRequestWithDefaults() *StartSandboxRunRequest`

NewStartSandboxRunRequestWithDefaults instantiates a new StartSandboxRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiverUrl

`func (o *StartSandboxRunRequest) GetReceiverUrl() string`

GetReceiverUrl returns the ReceiverUrl field if non-nil, zero value otherwise.

### GetReceiverUrlOk

`func (o *StartSandboxRunRequest) GetReceiverUrlOk() (*string, bool)`

GetReceiverUrlOk returns a tuple with the ReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUrl

`func (o *StartSandboxRunRequest) SetReceiverUrl(v string)`

SetReceiverUrl sets ReceiverUrl field to given value.


### GetToken

`func (o *StartSandboxRunRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StartSandboxRunRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StartSandboxRunRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


