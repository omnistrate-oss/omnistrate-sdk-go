# UpdateNotificationChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailConfiguration**](EmailConfiguration.md) |  | [optional] 
**Id** | **string** | ID of a Notification Channel | 
**Name** | Pointer to **string** | Name of the channel | [optional] 
**PagerDuty** | Pointer to [**PagerDutyConfiguration**](PagerDutyConfiguration.md) |  | [optional] 
**Slack** | Pointer to [**SlackConfiguration**](SlackConfiguration.md) |  | [optional] 
**Subscription** | Pointer to [**ChannelSubscription**](ChannelSubscription.md) |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**Webhook** | Pointer to [**WebhookConfiguration**](WebhookConfiguration.md) |  | [optional] 

## Methods

### NewUpdateNotificationChannelRequest

`func NewUpdateNotificationChannelRequest(id string, token string, ) *UpdateNotificationChannelRequest`

NewUpdateNotificationChannelRequest instantiates a new UpdateNotificationChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotificationChannelRequestWithDefaults

`func NewUpdateNotificationChannelRequestWithDefaults() *UpdateNotificationChannelRequest`

NewUpdateNotificationChannelRequestWithDefaults instantiates a new UpdateNotificationChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateNotificationChannelRequest) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateNotificationChannelRequest) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateNotificationChannelRequest) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateNotificationChannelRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *UpdateNotificationChannelRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNotificationChannelRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNotificationChannelRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateNotificationChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotificationChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotificationChannelRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotificationChannelRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPagerDuty

`func (o *UpdateNotificationChannelRequest) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *UpdateNotificationChannelRequest) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *UpdateNotificationChannelRequest) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *UpdateNotificationChannelRequest) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *UpdateNotificationChannelRequest) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *UpdateNotificationChannelRequest) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *UpdateNotificationChannelRequest) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *UpdateNotificationChannelRequest) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *UpdateNotificationChannelRequest) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UpdateNotificationChannelRequest) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UpdateNotificationChannelRequest) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *UpdateNotificationChannelRequest) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetToken

`func (o *UpdateNotificationChannelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateNotificationChannelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateNotificationChannelRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWebhook

`func (o *UpdateNotificationChannelRequest) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *UpdateNotificationChannelRequest) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *UpdateNotificationChannelRequest) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *UpdateNotificationChannelRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


