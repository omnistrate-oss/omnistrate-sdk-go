# WebhookConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalBodyParameters** | Pointer to **map[string]string** | Additional parameters to include in the notification body | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP headers to include in the notification | [optional] 
**Method** | **string** | HTTP method to use for the notification | 
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

`func (o *WebhookConfiguration) GetAdditionalBodyParameters() map[string]string`

GetAdditionalBodyParameters returns the AdditionalBodyParameters field if non-nil, zero value otherwise.

### GetAdditionalBodyParametersOk

`func (o *WebhookConfiguration) GetAdditionalBodyParametersOk() (*map[string]string, bool)`

GetAdditionalBodyParametersOk returns a tuple with the AdditionalBodyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalBodyParameters

`func (o *WebhookConfiguration) SetAdditionalBodyParameters(v map[string]string)`

SetAdditionalBodyParameters sets AdditionalBodyParameters field to given value.

### HasAdditionalBodyParameters

`func (o *WebhookConfiguration) HasAdditionalBodyParameters() bool`

HasAdditionalBodyParameters returns a boolean if a field has been set.

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


