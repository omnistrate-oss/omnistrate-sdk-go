# CreateNotificationChannelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**EmailConfiguration**](EmailConfiguration.md) |  | [optional] 
**Name** | **string** | Name of the channel | 
**PagerDuty** | Pointer to [**PagerDutyConfiguration**](PagerDutyConfiguration.md) |  | [optional] 
**Slack** | Pointer to [**SlackConfiguration**](SlackConfiguration.md) |  | [optional] 
**Subscription** | [**ChannelSubscription**](ChannelSubscription.md) |  | 
**Token** | **string** | JWT token used to perform authorization | 
**Webhook** | Pointer to [**WebhookConfiguration**](WebhookConfiguration.md) |  | [optional] 

## Methods

### NewCreateNotificationChannelRequest

`func NewCreateNotificationChannelRequest(name string, subscription ChannelSubscription, token string, ) *CreateNotificationChannelRequest`

NewCreateNotificationChannelRequest instantiates a new CreateNotificationChannelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotificationChannelRequestWithDefaults

`func NewCreateNotificationChannelRequestWithDefaults() *CreateNotificationChannelRequest`

NewCreateNotificationChannelRequestWithDefaults instantiates a new CreateNotificationChannelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *CreateNotificationChannelRequest) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateNotificationChannelRequest) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateNotificationChannelRequest) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateNotificationChannelRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *CreateNotificationChannelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotificationChannelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotificationChannelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPagerDuty

`func (o *CreateNotificationChannelRequest) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *CreateNotificationChannelRequest) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *CreateNotificationChannelRequest) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *CreateNotificationChannelRequest) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *CreateNotificationChannelRequest) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *CreateNotificationChannelRequest) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *CreateNotificationChannelRequest) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *CreateNotificationChannelRequest) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *CreateNotificationChannelRequest) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *CreateNotificationChannelRequest) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *CreateNotificationChannelRequest) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.


### GetToken

`func (o *CreateNotificationChannelRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateNotificationChannelRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateNotificationChannelRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWebhook

`func (o *CreateNotificationChannelRequest) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *CreateNotificationChannelRequest) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *CreateNotificationChannelRequest) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *CreateNotificationChannelRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


