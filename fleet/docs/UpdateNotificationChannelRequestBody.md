# UpdateNotificationChannelRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailConfiguration**](EmailConfiguration.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the channel | [optional] 
**PagerDuty** | Pointer to [**PagerDutyConfiguration**](PagerDutyConfiguration.md) |  | [optional] 
**Slack** | Pointer to [**SlackConfiguration**](SlackConfiguration.md) |  | [optional] 
**Subscription** | Pointer to [**ChannelSubscription**](ChannelSubscription.md) |  | [optional] 
**Webhook** | Pointer to [**WebhookConfiguration**](WebhookConfiguration.md) |  | [optional] 

## Methods

### NewUpdateNotificationChannelRequestBody

`func NewUpdateNotificationChannelRequestBody() *UpdateNotificationChannelRequestBody`

NewUpdateNotificationChannelRequestBody instantiates a new UpdateNotificationChannelRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationChannelRequestBodyWithDefaults

`func NewUpdateNotificationChannelRequestBodyWithDefaults() *UpdateNotificationChannelRequestBody`

NewUpdateNotificationChannelRequestBodyWithDefaults instantiates a new UpdateNotificationChannelRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateNotificationChannelRequestBody) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateNotificationChannelRequestBody) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateNotificationChannelRequestBody) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateNotificationChannelRequestBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateNotificationChannelRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationChannelRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationChannelRequestBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationChannelRequestBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPagerDuty

`func (o *UpdateNotificationChannelRequestBody) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *UpdateNotificationChannelRequestBody) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *UpdateNotificationChannelRequestBody) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *UpdateNotificationChannelRequestBody) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *UpdateNotificationChannelRequestBody) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *UpdateNotificationChannelRequestBody) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *UpdateNotificationChannelRequestBody) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *UpdateNotificationChannelRequestBody) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *UpdateNotificationChannelRequestBody) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UpdateNotificationChannelRequestBody) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UpdateNotificationChannelRequestBody) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UpdateNotificationChannelRequestBody) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetWebhook

`func (o *UpdateNotificationChannelRequestBody) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *UpdateNotificationChannelRequestBody) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *UpdateNotificationChannelRequestBody) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *UpdateNotificationChannelRequestBody) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


