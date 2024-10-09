/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
)

// checks if the UpdateNotificationChannelRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNotificationChannelRequestBody{}

// UpdateNotificationChannelRequestBody struct for UpdateNotificationChannelRequestBody
type UpdateNotificationChannelRequestBody struct {
	Email *EmailConfiguration `json:"email,omitempty"`
	// Name of the channel
	Name *string `json:"name,omitempty"`
	PagerDuty *PagerDutyConfiguration `json:"pagerDuty,omitempty"`
	Slack *SlackConfiguration `json:"slack,omitempty"`
	Subscription *ChannelSubscription `json:"subscription,omitempty"`
	Webhook *WebhookConfiguration `json:"webhook,omitempty"`
}

// NewUpdateNotificationChannelRequestBody instantiates a new UpdateNotificationChannelRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNotificationChannelRequestBody() *UpdateNotificationChannelRequestBody {
	this := UpdateNotificationChannelRequestBody{}
	return &this
}

// NewUpdateNotificationChannelRequestBodyWithDefaults instantiates a new UpdateNotificationChannelRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNotificationChannelRequestBodyWithDefaults() *UpdateNotificationChannelRequestBody {
	this := UpdateNotificationChannelRequestBody{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetEmail() EmailConfiguration {
	if o == nil || IsNil(o.Email) {
		var ret EmailConfiguration
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetEmailOk() (*EmailConfiguration, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailConfiguration and assigns it to the Email field.
func (o *UpdateNotificationChannelRequestBody) SetEmail(v EmailConfiguration) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNotificationChannelRequestBody) SetName(v string) {
	o.Name = &v
}

// GetPagerDuty returns the PagerDuty field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetPagerDuty() PagerDutyConfiguration {
	if o == nil || IsNil(o.PagerDuty) {
		var ret PagerDutyConfiguration
		return ret
	}
	return *o.PagerDuty
}

// GetPagerDutyOk returns a tuple with the PagerDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetPagerDutyOk() (*PagerDutyConfiguration, bool) {
	if o == nil || IsNil(o.PagerDuty) {
		return nil, false
	}
	return o.PagerDuty, true
}

// HasPagerDuty returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasPagerDuty() bool {
	if o != nil && !IsNil(o.PagerDuty) {
		return true
	}

	return false
}

// SetPagerDuty gets a reference to the given PagerDutyConfiguration and assigns it to the PagerDuty field.
func (o *UpdateNotificationChannelRequestBody) SetPagerDuty(v PagerDutyConfiguration) {
	o.PagerDuty = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetSlack() SlackConfiguration {
	if o == nil || IsNil(o.Slack) {
		var ret SlackConfiguration
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetSlackOk() (*SlackConfiguration, bool) {
	if o == nil || IsNil(o.Slack) {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasSlack() bool {
	if o != nil && !IsNil(o.Slack) {
		return true
	}

	return false
}

// SetSlack gets a reference to the given SlackConfiguration and assigns it to the Slack field.
func (o *UpdateNotificationChannelRequestBody) SetSlack(v SlackConfiguration) {
	o.Slack = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetSubscription() ChannelSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret ChannelSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetSubscriptionOk() (*ChannelSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ChannelSubscription and assigns it to the Subscription field.
func (o *UpdateNotificationChannelRequestBody) SetSubscription(v ChannelSubscription) {
	o.Subscription = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequestBody) GetWebhook() WebhookConfiguration {
	if o == nil || IsNil(o.Webhook) {
		var ret WebhookConfiguration
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequestBody) GetWebhookOk() (*WebhookConfiguration, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequestBody) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhookConfiguration and assigns it to the Webhook field.
func (o *UpdateNotificationChannelRequestBody) SetWebhook(v WebhookConfiguration) {
	o.Webhook = &v
}

func (o UpdateNotificationChannelRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNotificationChannelRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PagerDuty) {
		toSerialize["pagerDuty"] = o.PagerDuty
	}
	if !IsNil(o.Slack) {
		toSerialize["slack"] = o.Slack
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableUpdateNotificationChannelRequestBody struct {
	value *UpdateNotificationChannelRequestBody
	isSet bool
}

func (v NullableUpdateNotificationChannelRequestBody) Get() *UpdateNotificationChannelRequestBody {
	return v.value
}

func (v *NullableUpdateNotificationChannelRequestBody) Set(val *UpdateNotificationChannelRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNotificationChannelRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNotificationChannelRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNotificationChannelRequestBody(val *UpdateNotificationChannelRequestBody) *NullableUpdateNotificationChannelRequestBody {
	return &NullableUpdateNotificationChannelRequestBody{value: val, isSet: true}
}

func (v NullableUpdateNotificationChannelRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNotificationChannelRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

