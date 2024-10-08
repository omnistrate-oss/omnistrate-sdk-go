# SlackConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUrl** | **string** | URL to send notifications to | 

## Methods

### NewSlackConfiguration

`func NewSlackConfiguration(webhookUrl string, ) *SlackConfiguration`

NewSlackConfiguration instantiates a new SlackConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackConfigurationWithDefaults

`func NewSlackConfigurationWithDefaults() *SlackConfiguration`

NewSlackConfigurationWithDefaults instantiates a new SlackConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUrl

`func (o *SlackConfiguration) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SlackConfiguration) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SlackConfiguration) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


