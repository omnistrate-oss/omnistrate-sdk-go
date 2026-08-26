# SendSandboxEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The type of a marketplace fulfillment event delivered to an ISV receiver | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSendSandboxEventRequest

`func NewSendSandboxEventRequest(eventType string, token string, ) *SendSandboxEventRequest`

NewSendSandboxEventRequest instantiates a new SendSandboxEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSandboxEventRequestWithDefaults

`func NewSendSandboxEventRequestWithDefaults() *SendSandboxEventRequest`

NewSendSandboxEventRequestWithDefaults instantiates a new SendSandboxEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *SendSandboxEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SendSandboxEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SendSandboxEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetToken

`func (o *SendSandboxEventRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SendSandboxEventRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SendSandboxEventRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


