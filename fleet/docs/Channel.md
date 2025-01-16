# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelType** | **string** |  | 
**Email** | Pointer to [**EmailConfiguration**](EmailConfiguration.md) |  | [optional] 
**Id** | **string** | ID of a Notification Channel | 
**Name** | **string** | Name of the channel | 
**PagerDuty** | Pointer to [**PagerDutyConfiguration**](PagerDutyConfiguration.md) |  | [optional] 
**Slack** | Pointer to [**SlackConfiguration**](SlackConfiguration.md) |  | [optional] 
**Subscription** | [**ChannelSubscription**](ChannelSubscription.md) |  | 
**Webhook** | Pointer to [**WebhookConfiguration**](WebhookConfiguration.md) |  | [optional] 

## Methods

### NewChannel

`func NewChannel(channelType string, id string, name string, subscription ChannelSubscription, ) *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelType

`func (o *Channel) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *Channel) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *Channel) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.


### GetEmail

`func (o *Channel) GetEmail() EmailConfiguration`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Channel) GetEmailOk() (*EmailConfiguration, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Channel) SetEmail(v EmailConfiguration)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Channel) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *Channel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.


### GetPagerDuty

`func (o *Channel) GetPagerDuty() PagerDutyConfiguration`

GetPagerDuty returns the PagerDuty field if non-nil, zero value otherwise.

### GetPagerDutyOk

`func (o *Channel) GetPagerDutyOk() (*PagerDutyConfiguration, bool)`

GetPagerDutyOk returns a tuple with the PagerDuty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagerDuty

`func (o *Channel) SetPagerDuty(v PagerDutyConfiguration)`

SetPagerDuty sets PagerDuty field to given value.

### HasPagerDuty

`func (o *Channel) HasPagerDuty() bool`

HasPagerDuty returns a boolean if a field has been set.

### GetSlack

`func (o *Channel) GetSlack() SlackConfiguration`

GetSlack returns the Slack field if non-nil, zero value otherwise.

### GetSlackOk

`func (o *Channel) GetSlackOk() (*SlackConfiguration, bool)`

GetSlackOk returns a tuple with the Slack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlack

`func (o *Channel) SetSlack(v SlackConfiguration)`

SetSlack sets Slack field to given value.

### HasSlack

`func (o *Channel) HasSlack() bool`

HasSlack returns a boolean if a field has been set.

### GetSubscription

`func (o *Channel) GetSubscription() ChannelSubscription`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *Channel) GetSubscriptionOk() (*ChannelSubscription, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *Channel) SetSubscription(v ChannelSubscription)`

SetSubscription sets Subscription field to given value.


### GetWebhook

`func (o *Channel) GetWebhook() WebhookConfiguration`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *Channel) GetWebhookOk() (*WebhookConfiguration, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *Channel) SetWebhook(v WebhookConfiguration)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *Channel) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


