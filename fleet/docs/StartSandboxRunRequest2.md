# StartSandboxRunRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReceiverUrl** | **string** | Where to deliver. Validated on write: https only, no embedded credentials, and not a private, loopback, link-local or metadata address | 

## Methods

### NewStartSandboxRunRequest2

`func NewStartSandboxRunRequest2(receiverUrl string, ) *StartSandboxRunRequest2`

NewStartSandboxRunRequest2 instantiates a new StartSandboxRunRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartSandboxRunRequest2WithDefaults

`func NewStartSandboxRunRequest2WithDefaults() *StartSandboxRunRequest2`

NewStartSandboxRunRequest2WithDefaults instantiates a new StartSandboxRunRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiverUrl

`func (o *StartSandboxRunRequest2) GetReceiverUrl() string`

GetReceiverUrl returns the ReceiverUrl field if non-nil, zero value otherwise.

### GetReceiverUrlOk

`func (o *StartSandboxRunRequest2) GetReceiverUrlOk() (*string, bool)`

GetReceiverUrlOk returns a tuple with the ReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUrl

`func (o *StartSandboxRunRequest2) SetReceiverUrl(v string)`

SetReceiverUrl sets ReceiverUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


