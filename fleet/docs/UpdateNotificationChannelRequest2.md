# UpdateNotificationChannelRequest2

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

### NewUpdateNotificationChannelRequest2

`func NewUpdateNotificationChannelRequest2() *UpdateNotificationChannelRequest2`

NewUpdateNotificationChannelRequest2 instantiates a new UpdateNotificationChannelRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationChannelRequest2WithDefaults

`func NewUpdateNotificationChannelRequest2WithDefaults() *UpdateNotificationChannelRequest2`

NewUpdateNotificationChannelRequest2WithDefaults instantiates a new UpdateNotificationChannelRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateNotificationChannelRequest2) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateNotificationChannelRequest2) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateNotificationChannelRequest2) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateNotificationChannelRequest2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *UpdateNotificationChannelRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationChannelRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationChannelRequest2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationChannelRequest2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPagerDuty

`func (o *UpdateNotificationChannelRequest2) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *UpdateNotificationChannelRequest2) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *UpdateNotificationChannelRequest2) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *UpdateNotificationChannelRequest2) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *UpdateNotificationChannelRequest2) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *UpdateNotificationChannelRequest2) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *UpdateNotificationChannelRequest2) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *UpdateNotificationChannelRequest2) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *UpdateNotificationChannelRequest2) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UpdateNotificationChannelRequest2) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UpdateNotificationChannelRequest2) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UpdateNotificationChannelRequest2) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetWebhook

`func (o *UpdateNotificationChannelRequest2) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *UpdateNotificationChannelRequest2) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *UpdateNotificationChannelRequest2) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *UpdateNotificationChannelRequest2) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


