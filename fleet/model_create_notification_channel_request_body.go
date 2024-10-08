/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateNotificationChannelRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationChannelRequestBody{}

// CreateNotificationChannelRequestBody struct for CreateNotificationChannelRequestBody
type CreateNotificationChannelRequestBody struct {
	Email *EmailConfiguration `json:"email,omitempty"`
	// Name of the channel
	Name string `json:"name"`
	PagerDuty *PagerDutyConfiguration `json:"pagerDuty,omitempty"`
	Slack *SlackConfiguration `json:"slack,omitempty"`
	Subscription ChannelSubscription `json:"subscription"`
	Webhook *WebhookConfiguration `json:"webhook,omitempty"`
}

type _CreateNotificationChannelRequestBody CreateNotificationChannelRequestBody

// NewCreateNotificationChannelRequestBody instantiates a new CreateNotificationChannelRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationChannelRequestBody(name string, subscription ChannelSubscription) *CreateNotificationChannelRequestBody {
	this := CreateNotificationChannelRequestBody{}
	this.Name = name
	this.Subscription = subscription
	return &this
}

// NewCreateNotificationChannelRequestBodyWithDefaults instantiates a new CreateNotificationChannelRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationChannelRequestBodyWithDefaults() *CreateNotificationChannelRequestBody {
	this := CreateNotificationChannelRequestBody{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateNotificationChannelRequestBody) GetEmail() EmailConfiguration {
	if o == nil || IsNil(o.Email) {
		var ret EmailConfiguration
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetEmailOk() (*EmailConfiguration, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateNotificationChannelRequestBody) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailConfiguration and assigns it to the Email field.
func (o *CreateNotificationChannelRequestBody) SetEmail(v EmailConfiguration) {
	o.Email = &v
}

// GetName returns the Name field value
func (o *CreateNotificationChannelRequestBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNotificationChannelRequestBody) SetName(v string) {
	o.Name = v
}

// GetPagerDuty returns the PagerDuty field value if set, zero value otherwise.
func (o *CreateNotificationChannelRequestBody) GetPagerDuty() PagerDutyConfiguration {
	if o == nil || IsNil(o.PagerDuty) {
		var ret PagerDutyConfiguration
		return ret
	}
	return *o.PagerDuty
}

// GetPagerDutyOk returns a tuple with the PagerDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetPagerDutyOk() (*PagerDutyConfiguration, bool) {
	if o == nil || IsNil(o.PagerDuty) {
		return nil, false
	}
	return o.PagerDuty, true
}

// HasPagerDuty returns a boolean if a field has been set.
func (o *CreateNotificationChannelRequestBody) HasPagerDuty() bool {
	if o != nil && !IsNil(o.PagerDuty) {
		return true
	}

	return false
}

// SetPagerDuty gets a reference to the given PagerDutyConfiguration and assigns it to the PagerDuty field.
func (o *CreateNotificationChannelRequestBody) SetPagerDuty(v PagerDutyConfiguration) {
	o.PagerDuty = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *CreateNotificationChannelRequestBody) GetSlack() SlackConfiguration {
	if o == nil || IsNil(o.Slack) {
		var ret SlackConfiguration
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetSlackOk() (*SlackConfiguration, bool) {
	if o == nil || IsNil(o.Slack) {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *CreateNotificationChannelRequestBody) HasSlack() bool {
	if o != nil && !IsNil(o.Slack) {
		return true
	}

	return false
}

// SetSlack gets a reference to the given SlackConfiguration and assigns it to the Slack field.
func (o *CreateNotificationChannelRequestBody) SetSlack(v SlackConfiguration) {
	o.Slack = &v
}

// GetSubscription returns the Subscription field value
func (o *CreateNotificationChannelRequestBody) GetSubscription() ChannelSubscription {
	if o == nil {
		var ret ChannelSubscription
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetSubscriptionOk() (*ChannelSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *CreateNotificationChannelRequestBody) SetSubscription(v ChannelSubscription) {
	o.Subscription = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *CreateNotificationChannelRequestBody) GetWebhook() WebhookConfiguration {
	if o == nil || IsNil(o.Webhook) {
		var ret WebhookConfiguration
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationChannelRequestBody) GetWebhookOk() (*WebhookConfiguration, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *CreateNotificationChannelRequestBody) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhookConfiguration and assigns it to the Webhook field.
func (o *CreateNotificationChannelRequestBody) SetWebhook(v WebhookConfiguration) {
	o.Webhook = &v
}

func (o CreateNotificationChannelRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNotificationChannelRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.PagerDuty) {
		toSerialize["pagerDuty"] = o.PagerDuty
	}
	if !IsNil(o.Slack) {
		toSerialize["slack"] = o.Slack
	}
	toSerialize["subscription"] = o.Subscription
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

func (o *CreateNotificationChannelRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"subscription",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateNotificationChannelRequestBody := _CreateNotificationChannelRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNotificationChannelRequestBody)

	if err != nil {
		return err
	}

	*o = CreateNotificationChannelRequestBody(varCreateNotificationChannelRequestBody)

	return err
}

type NullableCreateNotificationChannelRequestBody struct {
	value *CreateNotificationChannelRequestBody
	isSet bool
}

func (v NullableCreateNotificationChannelRequestBody) Get() *CreateNotificationChannelRequestBody {
	return v.value
}

func (v *NullableCreateNotificationChannelRequestBody) Set(val *CreateNotificationChannelRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationChannelRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationChannelRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationChannelRequestBody(val *CreateNotificationChannelRequestBody) *NullableCreateNotificationChannelRequestBody {
	return &NullableCreateNotificationChannelRequestBody{value: val, isSet: true}
}

func (v NullableCreateNotificationChannelRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotificationChannelRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


