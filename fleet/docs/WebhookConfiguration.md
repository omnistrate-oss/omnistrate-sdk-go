# WebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalBodyParameters** | Pointer to **map[string]interface{}** | Additional parameters to include in the notification body | [optional] 
**DeliveryTimeoutSeconds** | Pointer to **int64** | How long a single attempt may take before it counts as failed and the backoff applies. Defaults to 10. This is a per-attempt limit and is unrelated to the total budget; a receiver that answers slowly burns attempts rather than extending them | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP headers to include in the notification | [optional] 
**IdempotencyHeader** | Pointer to **string** | Read only, and always X-Omnistrate-Event-Id. The header carrying the deduplication token for this channel&#39;s deliveries. The value is stable across retries AND across an operator redelivering the same event, so recording processed ids makes a repeat delivery a no-op. Reported so a receiver knows the token exists rather than having to discover it from a delivery | [optional] 
**MaxDeliveryAttempts** | Pointer to **int64** | How many attempts the budget is spread across, including the first. Defaults to 7, which covers a day with backoff. More attempts inside the same budget means a shorter gap between each, which helps a receiver that flaps and does nothing for one that is simply down | [optional] 
**MaxDeliveryDurationSeconds** | Pointer to **int64** | Total budget for delivering one event, across all attempts, in seconds. Retries use exponential backoff within it. Defaults to 86400, a full day, and 86400 is also the ceiling: past a day an event is usually stale enough that reporting the channel as broken is more useful than delivering it. Set it lower for a channel where a late notification is worse than none | [optional] 
**Method** | **string** | HTTP method to use for the notification | 
**SigningSecret** | Pointer to **string** | Write only, and never returned by any read. The HMAC-SHA256 secret deliveries to this webhook are signed with, and the only place a signing secret is set: there is no channel-level equivalent. You supply it; Omnistrate does not generate one. Omit it and this channel&#39;s deliveries are unsigned. On update, a new value rotates the secret and both the previous and the new one are accepted for an overlap window, so a receiver is not cut off mid-rotation; omit it to leave the secret unchanged. Reads report signingSecretSet instead | [optional] 
**SigningSecretId** | Pointer to **string** | Read only. Identifies WHICH secret is active, so a rotation can be confirmed without revealing anything. An identifier is not a credential | [optional] 
**SigningSecretSet** | Pointer to **bool** | Read only. Whether a signing secret is configured, which is the whole of what a read says about it. False means deliveries to this webhook are unsigned | [optional] 
**Url** | **string** | URL to send notifications to | 

## Methods

### NewWebhookConfiguration

`func NewWebhookConfiguration(method string, url string, ) *WebhookConfiguration`

NewWebhookConfiguration instantiates a new WebhookConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookConfigurationWithDefaults

`func NewWebhookConfigurationWithDefaults() *WebhookConfiguration`

NewWebhookConfigurationWithDefaults instantiates a new WebhookConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalBodyParameters

`func (o *WebhookConfiguration) GetAdditionalBodyParameters() map[string]interface{}`

GetAdditionalBodyParameters returns the AdditionalBodyParameters field if non-nil, zero value otherwise.

### GetAdditionalBodyParametersOk

`func (o *WebhookConfiguration) GetAdditionalBodyParametersOk() (*map[string]interface{}, bool)`

GetAdditionalBodyParametersOk returns a tuple with the AdditionalBodyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBodyParameters

`func (o *WebhookConfiguration) SetAdditionalBodyParameters(v map[string]interface{})`

SetAdditionalBodyParameters sets AdditionalBodyParameters field to given value.

### HasAdditionalBodyParameters

`func (o *WebhookConfiguration) HasAdditionalBodyParameters() bool`

HasAdditionalBodyParameters returns a boolean if a field has been set.

### GetDeliveryTimeoutSeconds

`func (o *WebhookConfiguration) GetDeliveryTimeoutSeconds() int64`

GetDeliveryTimeoutSeconds returns the DeliveryTimeoutSeconds field if non-nil, zero value otherwise.

### GetDeliveryTimeoutSecondsOk

`func (o *WebhookConfiguration) GetDeliveryTimeoutSecondsOk() (*int64, bool)`

GetDeliveryTimeoutSecondsOk returns a tuple with the DeliveryTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeoutSeconds

`func (o *WebhookConfiguration) SetDeliveryTimeoutSeconds(v int64)`

SetDeliveryTimeoutSeconds sets DeliveryTimeoutSeconds field to given value.

### HasDeliveryTimeoutSeconds

`func (o *WebhookConfiguration) HasDeliveryTimeoutSeconds() bool`

HasDeliveryTimeoutSeconds returns a boolean if a field has been set.

### GetHeaders

`func (o *WebhookConfiguration) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookConfiguration) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookConfiguration) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookConfiguration) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetIdempotencyHeader

`func (o *WebhookConfiguration) GetIdempotencyHeader() string`

GetIdempotencyHeader returns the IdempotencyHeader field if non-nil, zero value otherwise.

### GetIdempotencyHeaderOk

`func (o *WebhookConfiguration) GetIdempotencyHeaderOk() (*string, bool)`

GetIdempotencyHeaderOk returns a tuple with the IdempotencyHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyHeader

`func (o *WebhookConfiguration) SetIdempotencyHeader(v string)`

SetIdempotencyHeader sets IdempotencyHeader field to given value.

### HasIdempotencyHeader

`func (o *WebhookConfiguration) HasIdempotencyHeader() bool`

HasIdempotencyHeader returns a boolean if a field has been set.

### GetMaxDeliveryAttempts

`func (o *WebhookConfiguration) GetMaxDeliveryAttempts() int64`

GetMaxDeliveryAttempts returns the MaxDeliveryAttempts field if non-nil, zero value otherwise.

### GetMaxDeliveryAttemptsOk

`func (o *WebhookConfiguration) GetMaxDeliveryAttemptsOk() (*int64, bool)`

GetMaxDeliveryAttemptsOk returns a tuple with the MaxDeliveryAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryAttempts

`func (o *WebhookConfiguration) SetMaxDeliveryAttempts(v int64)`

SetMaxDeliveryAttempts sets MaxDeliveryAttempts field to given value.

### HasMaxDeliveryAttempts

`func (o *WebhookConfiguration) HasMaxDeliveryAttempts() bool`

HasMaxDeliveryAttempts returns a boolean if a field has been set.

### GetMaxDeliveryDurationSeconds

`func (o *WebhookConfiguration) GetMaxDeliveryDurationSeconds() int64`

GetMaxDeliveryDurationSeconds returns the MaxDeliveryDurationSeconds field if non-nil, zero value otherwise.

### GetMaxDeliveryDurationSecondsOk

`func (o *WebhookConfiguration) GetMaxDeliveryDurationSecondsOk() (*int64, bool)`

GetMaxDeliveryDurationSecondsOk returns a tuple with the MaxDeliveryDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDeliveryDurationSeconds

`func (o *WebhookConfiguration) SetMaxDeliveryDurationSeconds(v int64)`

SetMaxDeliveryDurationSeconds sets MaxDeliveryDurationSeconds field to given value.

### HasMaxDeliveryDurationSeconds

`func (o *WebhookConfiguration) HasMaxDeliveryDurationSeconds() bool`

HasMaxDeliveryDurationSeconds returns a boolean if a field has been set.

### GetMethod

`func (o *WebhookConfiguration) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookConfiguration) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookConfiguration) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetSigningSecret

`func (o *WebhookConfiguration) GetSigningSecret() string`

GetSigningSecret returns the SigningSecret field if non-nil, zero value otherwise.

### GetSigningSecretOk

`func (o *WebhookConfiguration) GetSigningSecretOk() (*string, bool)`

GetSigningSecretOk returns a tuple with the SigningSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecret

`func (o *WebhookConfiguration) SetSigningSecret(v string)`

SetSigningSecret sets SigningSecret field to given value.

### HasSigningSecret

`func (o *WebhookConfiguration) HasSigningSecret() bool`

HasSigningSecret returns a boolean if a field has been set.

### GetSigningSecretId

`func (o *WebhookConfiguration) GetSigningSecretId() string`

GetSigningSecretId returns the SigningSecretId field if non-nil, zero value otherwise.

### GetSigningSecretIdOk

`func (o *WebhookConfiguration) GetSigningSecretIdOk() (*string, bool)`

GetSigningSecretIdOk returns a tuple with the SigningSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecretId

`func (o *WebhookConfiguration) SetSigningSecretId(v string)`

SetSigningSecretId sets SigningSecretId field to given value.

### HasSigningSecretId

`func (o *WebhookConfiguration) HasSigningSecretId() bool`

HasSigningSecretId returns a boolean if a field has been set.

### GetSigningSecretSet

`func (o *WebhookConfiguration) GetSigningSecretSet() bool`

GetSigningSecretSet returns the SigningSecretSet field if non-nil, zero value otherwise.

### GetSigningSecretSetOk

`func (o *WebhookConfiguration) GetSigningSecretSetOk() (*bool, bool)`

GetSigningSecretSetOk returns a tuple with the SigningSecretSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecretSet

`func (o *WebhookConfiguration) SetSigningSecretSet(v bool)`

SetSigningSecretSet sets SigningSecretSet field to given value.

### HasSigningSecretSet

`func (o *WebhookConfiguration) HasSigningSecretSet() bool`

HasSigningSecretSet returns a boolean if a field has been set.

### GetUrl

`func (o *WebhookConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


