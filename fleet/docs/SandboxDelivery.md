# SandboxDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackoffSeconds** | Pointer to **int64** | How long we waited after THIS attempt before the next one. Absent on the last attempt, which is how a reader tells a delivery that is still being retried from one that has stopped | [optional] 
**CheckId** | Pointer to **string** | Which conformance check a result belongs to | [optional] 
**DeliveryId** | **string** |  | 
**EventId** | **string** |  | 
**EventType** | **string** | The type of a marketplace fulfillment event delivered to an ISV receiver | 
**LatencyMs** | Pointer to **int64** |  | [optional] 
**Payload** | Pointer to [**MarketplaceEvent**](MarketplaceEvent.md) |  | [optional] 
**RequestBody** | Pointer to **string** | The exact bytes sent, never re-marshalled. These are the bytes the signature covers, after the timestamp prefix, so a re-serialized copy would not verify | [optional] 
**RequestHeaders** | Pointer to **map[string]string** | The headers as sent, including the five X-Omnistrate-* per-delivery headers | [optional] 
**ResponseBody** | Pointer to **string** | Truncated at the stated limit. The most common Port B failure is a receiver that returns 2xx and never confirms, and this is usually the only clue | [optional] 
**ResponseBodyTruncated** | Pointer to **bool** | Whether responseBody was cut at the limit, stated rather than silently applied | [optional] 
**ResponseHeaders** | Pointer to **map[string]string** |  | [optional] 
**ResponseStatus** | Pointer to **int64** | Absent when no response was received | [optional] 
**SentAt** | **time.Time** |  | 
**Signature** | Pointer to **string** | The X-Omnistrate-Signature-256 value sent | [optional] 
**SignedPayload** | Pointer to **string** | The exact string the signature was computed over: the timestamp, a full stop, then the body. Rendered so an ISV can reproduce the digest locally rather than reasoning from a 401 | [optional] 
**TransportError** | Pointer to **string** |  | [optional] 

## Methods

### NewSandboxDelivery

`func NewSandboxDelivery(deliveryId string, eventId string, eventType string, sentAt time.Time, ) *SandboxDelivery`

NewSandboxDelivery instantiates a new SandboxDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxDeliveryWithDefaults

`func NewSandboxDeliveryWithDefaults() *SandboxDelivery`

NewSandboxDeliveryWithDefaults instantiates a new SandboxDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackoffSeconds

`func (o *SandboxDelivery) GetBackoffSeconds() int64`

GetBackoffSeconds returns the BackoffSeconds field if non-nil, zero value otherwise.

### GetBackoffSecondsOk

`func (o *SandboxDelivery) GetBackoffSecondsOk() (*int64, bool)`

GetBackoffSecondsOk returns a tuple with the BackoffSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackoffSeconds

`func (o *SandboxDelivery) SetBackoffSeconds(v int64)`

SetBackoffSeconds sets BackoffSeconds field to given value.

### HasBackoffSeconds

`func (o *SandboxDelivery) HasBackoffSeconds() bool`

HasBackoffSeconds returns a boolean if a field has been set.

### GetCheckId

`func (o *SandboxDelivery) GetCheckId() string`

GetCheckId returns the CheckId field if non-nil, zero value otherwise.

### GetCheckIdOk

`func (o *SandboxDelivery) GetCheckIdOk() (*string, bool)`

GetCheckIdOk returns a tuple with the CheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckId

`func (o *SandboxDelivery) SetCheckId(v string)`

SetCheckId sets CheckId field to given value.

### HasCheckId

`func (o *SandboxDelivery) HasCheckId() bool`

HasCheckId returns a boolean if a field has been set.

### GetDeliveryId

`func (o *SandboxDelivery) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *SandboxDelivery) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *SandboxDelivery) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.


### GetEventId

`func (o *SandboxDelivery) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SandboxDelivery) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SandboxDelivery) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventType

`func (o *SandboxDelivery) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SandboxDelivery) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SandboxDelivery) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetLatencyMs

`func (o *SandboxDelivery) GetLatencyMs() int64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *SandboxDelivery) GetLatencyMsOk() (*int64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *SandboxDelivery) SetLatencyMs(v int64)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *SandboxDelivery) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetPayload

`func (o *SandboxDelivery) GetPayload() MarketplaceEvent`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SandboxDelivery) GetPayloadOk() (*MarketplaceEvent, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SandboxDelivery) SetPayload(v MarketplaceEvent)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SandboxDelivery) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRequestBody

`func (o *SandboxDelivery) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *SandboxDelivery) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *SandboxDelivery) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *SandboxDelivery) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *SandboxDelivery) GetRequestHeaders() map[string]string`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SandboxDelivery) GetRequestHeadersOk() (*map[string]string, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *SandboxDelivery) SetRequestHeaders(v map[string]string)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *SandboxDelivery) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetResponseBody

`func (o *SandboxDelivery) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *SandboxDelivery) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *SandboxDelivery) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *SandboxDelivery) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseBodyTruncated

`func (o *SandboxDelivery) GetResponseBodyTruncated() bool`

GetResponseBodyTruncated returns the ResponseBodyTruncated field if non-nil, zero value otherwise.

### GetResponseBodyTruncatedOk

`func (o *SandboxDelivery) GetResponseBodyTruncatedOk() (*bool, bool)`

GetResponseBodyTruncatedOk returns a tuple with the ResponseBodyTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBodyTruncated

`func (o *SandboxDelivery) SetResponseBodyTruncated(v bool)`

SetResponseBodyTruncated sets ResponseBodyTruncated field to given value.

### HasResponseBodyTruncated

`func (o *SandboxDelivery) HasResponseBodyTruncated() bool`

HasResponseBodyTruncated returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *SandboxDelivery) GetResponseHeaders() map[string]string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *SandboxDelivery) GetResponseHeadersOk() (*map[string]string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *SandboxDelivery) SetResponseHeaders(v map[string]string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *SandboxDelivery) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseStatus

`func (o *SandboxDelivery) GetResponseStatus() int64`

GetResponseStatus returns the ResponseStatus field if non-nil, zero value otherwise.

### GetResponseStatusOk

`func (o *SandboxDelivery) GetResponseStatusOk() (*int64, bool)`

GetResponseStatusOk returns a tuple with the ResponseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatus

`func (o *SandboxDelivery) SetResponseStatus(v int64)`

SetResponseStatus sets ResponseStatus field to given value.

### HasResponseStatus

`func (o *SandboxDelivery) HasResponseStatus() bool`

HasResponseStatus returns a boolean if a field has been set.

### GetSentAt

`func (o *SandboxDelivery) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *SandboxDelivery) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *SandboxDelivery) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.


### GetSignature

`func (o *SandboxDelivery) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SandboxDelivery) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SandboxDelivery) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SandboxDelivery) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignedPayload

`func (o *SandboxDelivery) GetSignedPayload() string`

GetSignedPayload returns the SignedPayload field if non-nil, zero value otherwise.

### GetSignedPayloadOk

`func (o *SandboxDelivery) GetSignedPayloadOk() (*string, bool)`

GetSignedPayloadOk returns a tuple with the SignedPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedPayload

`func (o *SandboxDelivery) SetSignedPayload(v string)`

SetSignedPayload sets SignedPayload field to given value.

### HasSignedPayload

`func (o *SandboxDelivery) HasSignedPayload() bool`

HasSignedPayload returns a boolean if a field has been set.

### GetTransportError

`func (o *SandboxDelivery) GetTransportError() string`

GetTransportError returns the TransportError field if non-nil, zero value otherwise.

### GetTransportErrorOk

`func (o *SandboxDelivery) GetTransportErrorOk() (*string, bool)`

GetTransportErrorOk returns a tuple with the TransportError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportError

`func (o *SandboxDelivery) SetTransportError(v string)`

SetTransportError sets TransportError field to given value.

### HasTransportError

`func (o *SandboxDelivery) HasTransportError() bool`

HasTransportError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


