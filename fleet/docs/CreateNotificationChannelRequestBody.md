# CreateNotificationChannelRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailConfiguration**](EmailConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the channel | 
**PagerDuty** | Pointer to [**PagerDutyConfiguration**](PagerDutyConfiguration.md) |  | [optional] 
**Slack** | Pointer to [**SlackConfiguration**](SlackConfiguration.md) |  | [optional] 
**Subscription** | [**ChannelSubscription**](ChannelSubscription.md) |  | 
**Webhook** | Pointer to [**WebhookConfiguration**](WebhookConfiguration.md) |  | [optional] 

## Methods

### NewCreateNotificationChannelRequestBody

`func NewCreateNotificationChannelRequestBody(name string, subscription ChannelSubscription, ) *CreateNotificationChannelRequestBody`

NewCreateNotificationChannelRequestBody instantiates a new CreateNotificationChannelRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationChannelRequestBodyWithDefaults

`func NewCreateNotificationChannelRequestBodyWithDefaults() *CreateNotificationChannelRequestBody`

NewCreateNotificationChannelRequestBodyWithDefaults instantiates a new CreateNotificationChannelRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateNotificationChannelRequestBody) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateNotificationChannelRequestBody) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateNotificationChannelRequestBody) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateNotificationChannelRequestBody) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CreateNotificationChannelRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationChannelRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationChannelRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPagerDuty

`func (o *CreateNotificationChannelRequestBody) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *CreateNotificationChannelRequestBody) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *CreateNotificationChannelRequestBody) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *CreateNotificationChannelRequestBody) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *CreateNotificationChannelRequestBody) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *CreateNotificationChannelRequestBody) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *CreateNotificationChannelRequestBody) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *CreateNotificationChannelRequestBody) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *CreateNotificationChannelRequestBody) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CreateNotificationChannelRequestBody) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CreateNotificationChannelRequestBody) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.


### GetWebhook

`func (o *CreateNotificationChannelRequestBody) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CreateNotificationChannelRequestBody) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CreateNotificationChannelRequestBody) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CreateNotificationChannelRequestBody) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


