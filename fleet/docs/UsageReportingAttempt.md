# UsageReportingAttempt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAt** | **time.Time** |  | 
**BillingChannel** | **string** | The downstream billing channel that received the usage | 
**BillingChannelCode** | Pointer to **string** | The billing-channel-specific status code or mapped reason for this attempt | [optional] 
**BillingChannelError** | Pointer to **string** | The billing-channel-specific error text for this attempt | [optional] 
**Dimension** | **string** |  | 
**Id** | **string** |  | 
**LatencyMs** | **int64** |  | 
**Quantity** | **float64** |  | 
**RequestBody** | Pointer to **string** | The bounded request body sent to the billing channel | [optional] 
**RequestHeaders** | Pointer to **map[string]interface{}** | Redacted request headers | [optional] 
**RequestMethod** | Pointer to **string** |  | [optional] 
**RequestPath** | Pointer to **string** | The path sent to the billing channel, without credentials or host | [optional] 
**RequestTruncated** | **bool** |  | 
**ResponseBody** | Pointer to **string** | The bounded response body returned by the billing channel | [optional] 
**ResponseHeaders** | Pointer to **map[string]interface{}** | Redacted response headers | [optional] 
**ResponseStatusCode** | Pointer to **int64** |  | [optional] 
**ResponseTruncated** | **bool** |  | 
**Status** | **string** | The current outcome represented by one usage reporting attempt | 
**SubscriptionId** | **string** | The Omnistrate subscription whose usage was reported | 
**TransportError** | Pointer to **string** | The local HTTP or transport error when no billing channel response was received | [optional] 
**UsageReportRecordId** | **string** | The usage reporting record idempotency key | 
**WindowEnd** | **time.Time** |  | 
**WindowStart** | **time.Time** |  | 

## Methods

### NewUsageReportingAttempt

`func NewUsageReportingAttempt(attemptedAt time.Time, billingChannel string, dimension string, id string, latencyMs int64, quantity float64, requestTruncated bool, responseTruncated bool, status string, subscriptionId string, usageReportRecordId string, windowEnd time.Time, windowStart time.Time, ) *UsageReportingAttempt`

NewUsageReportingAttempt instantiates a new UsageReportingAttempt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageReportingAttemptWithDefaults

`func NewUsageReportingAttemptWithDefaults() *UsageReportingAttempt`

NewUsageReportingAttemptWithDefaults instantiates a new UsageReportingAttempt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAt

`func (o *UsageReportingAttempt) GetAttemptedAt() time.Time`

GetAttemptedAt returns the AttemptedAt field if non-nil, zero value otherwise.

### GetAttemptedAtOk

`func (o *UsageReportingAttempt) GetAttemptedAtOk() (*time.Time, bool)`

GetAttemptedAtOk returns a tuple with the AttemptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAt

`func (o *UsageReportingAttempt) SetAttemptedAt(v time.Time)`

SetAttemptedAt sets AttemptedAt field to given value.


### GetBillingChannel

`func (o *UsageReportingAttempt) GetBillingChannel() string`

GetBillingChannel returns the BillingChannel field if non-nil, zero value otherwise.

### GetBillingChannelOk

`func (o *UsageReportingAttempt) GetBillingChannelOk() (*string, bool)`

GetBillingChannelOk returns a tuple with the BillingChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingChannel

`func (o *UsageReportingAttempt) SetBillingChannel(v string)`

SetBillingChannel sets BillingChannel field to given value.


### GetBillingChannelCode

`func (o *UsageReportingAttempt) GetBillingChannelCode() string`

GetBillingChannelCode returns the BillingChannelCode field if non-nil, zero value otherwise.

### GetBillingChannelCodeOk

`func (o *UsageReportingAttempt) GetBillingChannelCodeOk() (*string, bool)`

GetBillingChannelCodeOk returns a tuple with the BillingChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingChannelCode

`func (o *UsageReportingAttempt) SetBillingChannelCode(v string)`

SetBillingChannelCode sets BillingChannelCode field to given value.

### HasBillingChannelCode

`func (o *UsageReportingAttempt) HasBillingChannelCode() bool`

HasBillingChannelCode returns a boolean if a field has been set.

### GetBillingChannelError

`func (o *UsageReportingAttempt) GetBillingChannelError() string`

GetBillingChannelError returns the BillingChannelError field if non-nil, zero value otherwise.

### GetBillingChannelErrorOk

`func (o *UsageReportingAttempt) GetBillingChannelErrorOk() (*string, bool)`

GetBillingChannelErrorOk returns a tuple with the BillingChannelError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingChannelError

`func (o *UsageReportingAttempt) SetBillingChannelError(v string)`

SetBillingChannelError sets BillingChannelError field to given value.

### HasBillingChannelError

`func (o *UsageReportingAttempt) HasBillingChannelError() bool`

HasBillingChannelError returns a boolean if a field has been set.

### GetDimension

`func (o *UsageReportingAttempt) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageReportingAttempt) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageReportingAttempt) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetId

`func (o *UsageReportingAttempt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageReportingAttempt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageReportingAttempt) SetId(v string)`

SetId sets Id field to given value.


### GetLatencyMs

`func (o *UsageReportingAttempt) GetLatencyMs() int64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *UsageReportingAttempt) GetLatencyMsOk() (*int64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *UsageReportingAttempt) SetLatencyMs(v int64)`

SetLatencyMs sets LatencyMs field to given value.


### GetQuantity

`func (o *UsageReportingAttempt) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UsageReportingAttempt) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UsageReportingAttempt) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetRequestBody

`func (o *UsageReportingAttempt) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *UsageReportingAttempt) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *UsageReportingAttempt) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *UsageReportingAttempt) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestHeaders

`func (o *UsageReportingAttempt) GetRequestHeaders() map[string]interface{}`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *UsageReportingAttempt) GetRequestHeadersOk() (*map[string]interface{}, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *UsageReportingAttempt) SetRequestHeaders(v map[string]interface{})`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *UsageReportingAttempt) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### GetRequestMethod

`func (o *UsageReportingAttempt) GetRequestMethod() string`

GetRequestMethod returns the RequestMethod field if non-nil, zero value otherwise.

### GetRequestMethodOk

`func (o *UsageReportingAttempt) GetRequestMethodOk() (*string, bool)`

GetRequestMethodOk returns a tuple with the RequestMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMethod

`func (o *UsageReportingAttempt) SetRequestMethod(v string)`

SetRequestMethod sets RequestMethod field to given value.

### HasRequestMethod

`func (o *UsageReportingAttempt) HasRequestMethod() bool`

HasRequestMethod returns a boolean if a field has been set.

### GetRequestPath

`func (o *UsageReportingAttempt) GetRequestPath() string`

GetRequestPath returns the RequestPath field if non-nil, zero value otherwise.

### GetRequestPathOk

`func (o *UsageReportingAttempt) GetRequestPathOk() (*string, bool)`

GetRequestPathOk returns a tuple with the RequestPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPath

`func (o *UsageReportingAttempt) SetRequestPath(v string)`

SetRequestPath sets RequestPath field to given value.

### HasRequestPath

`func (o *UsageReportingAttempt) HasRequestPath() bool`

HasRequestPath returns a boolean if a field has been set.

### GetRequestTruncated

`func (o *UsageReportingAttempt) GetRequestTruncated() bool`

GetRequestTruncated returns the RequestTruncated field if non-nil, zero value otherwise.

### GetRequestTruncatedOk

`func (o *UsageReportingAttempt) GetRequestTruncatedOk() (*bool, bool)`

GetRequestTruncatedOk returns a tuple with the RequestTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTruncated

`func (o *UsageReportingAttempt) SetRequestTruncated(v bool)`

SetRequestTruncated sets RequestTruncated field to given value.


### GetResponseBody

`func (o *UsageReportingAttempt) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *UsageReportingAttempt) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *UsageReportingAttempt) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *UsageReportingAttempt) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetResponseHeaders

`func (o *UsageReportingAttempt) GetResponseHeaders() map[string]interface{}`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *UsageReportingAttempt) GetResponseHeadersOk() (*map[string]interface{}, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *UsageReportingAttempt) SetResponseHeaders(v map[string]interface{})`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *UsageReportingAttempt) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### GetResponseStatusCode

`func (o *UsageReportingAttempt) GetResponseStatusCode() int64`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *UsageReportingAttempt) GetResponseStatusCodeOk() (*int64, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *UsageReportingAttempt) SetResponseStatusCode(v int64)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *UsageReportingAttempt) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### GetResponseTruncated

`func (o *UsageReportingAttempt) GetResponseTruncated() bool`

GetResponseTruncated returns the ResponseTruncated field if non-nil, zero value otherwise.

### GetResponseTruncatedOk

`func (o *UsageReportingAttempt) GetResponseTruncatedOk() (*bool, bool)`

GetResponseTruncatedOk returns a tuple with the ResponseTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTruncated

`func (o *UsageReportingAttempt) SetResponseTruncated(v bool)`

SetResponseTruncated sets ResponseTruncated field to given value.


### GetStatus

`func (o *UsageReportingAttempt) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageReportingAttempt) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageReportingAttempt) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionId

`func (o *UsageReportingAttempt) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UsageReportingAttempt) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UsageReportingAttempt) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTransportError

`func (o *UsageReportingAttempt) GetTransportError() string`

GetTransportError returns the TransportError field if non-nil, zero value otherwise.

### GetTransportErrorOk

`func (o *UsageReportingAttempt) GetTransportErrorOk() (*string, bool)`

GetTransportErrorOk returns a tuple with the TransportError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportError

`func (o *UsageReportingAttempt) SetTransportError(v string)`

SetTransportError sets TransportError field to given value.

### HasTransportError

`func (o *UsageReportingAttempt) HasTransportError() bool`

HasTransportError returns a boolean if a field has been set.

### GetUsageReportRecordId

`func (o *UsageReportingAttempt) GetUsageReportRecordId() string`

GetUsageReportRecordId returns the UsageReportRecordId field if non-nil, zero value otherwise.

### GetUsageReportRecordIdOk

`func (o *UsageReportingAttempt) GetUsageReportRecordIdOk() (*string, bool)`

GetUsageReportRecordIdOk returns a tuple with the UsageReportRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportRecordId

`func (o *UsageReportingAttempt) SetUsageReportRecordId(v string)`

SetUsageReportRecordId sets UsageReportRecordId field to given value.


### GetWindowEnd

`func (o *UsageReportingAttempt) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *UsageReportingAttempt) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *UsageReportingAttempt) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetWindowStart

`func (o *UsageReportingAttempt) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *UsageReportingAttempt) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *UsageReportingAttempt) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


