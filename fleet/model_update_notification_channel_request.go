/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"fmt"
)

// checks if the UpdateNotificationChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNotificationChannelRequest{}

// UpdateNotificationChannelRequest struct for UpdateNotificationChannelRequest
type UpdateNotificationChannelRequest struct {
	Email *EmailConfiguration `json:"email,omitempty"`
	// ID of a Notification Channel
	Id string `json:"id"`
	// Name of the channel
	Name *string `json:"name,omitempty"`
	PagerDuty *PagerDutyConfiguration `json:"pagerDuty,omitempty"`
	Slack *SlackConfiguration `json:"slack,omitempty"`
	Subscription *ChannelSubscription `json:"subscription,omitempty"`
	// JWT token used to perform authorization
	Token string `json:"token"`
	Webhook *WebhookConfiguration `json:"webhook,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateNotificationChannelRequest UpdateNotificationChannelRequest

// NewUpdateNotificationChannelRequest instantiates a new UpdateNotificationChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNotificationChannelRequest(id string, token string) *UpdateNotificationChannelRequest {
	this := UpdateNotificationChannelRequest{}
	this.Id = id
	this.Token = token
	return &this
}

// NewUpdateNotificationChannelRequestWithDefaults instantiates a new UpdateNotificationChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNotificationChannelRequestWithDefaults() *UpdateNotificationChannelRequest {
	this := UpdateNotificationChannelRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetEmail() EmailConfiguration {
	if o == nil || IsNil(o.Email) {
		var ret EmailConfiguration
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetEmailOk() (*EmailConfiguration, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailConfiguration and assigns it to the Email field.
func (o *UpdateNotificationChannelRequest) SetEmail(v EmailConfiguration) {
	o.Email = &v
}

// GetId returns the Id field value
func (o *UpdateNotificationChannelRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UpdateNotificationChannelRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNotificationChannelRequest) SetName(v string) {
	o.Name = &v
}

// GetPagerDuty returns the PagerDuty field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetPagerDuty() PagerDutyConfiguration {
	if o == nil || IsNil(o.PagerDuty) {
		var ret PagerDutyConfiguration
		return ret
	}
	return *o.PagerDuty
}

// GetPagerDutyOk returns a tuple with the PagerDuty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetPagerDutyOk() (*PagerDutyConfiguration, bool) {
	if o == nil || IsNil(o.PagerDuty) {
		return nil, false
	}
	return o.PagerDuty, true
}

// HasPagerDuty returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasPagerDuty() bool {
	if o != nil && !IsNil(o.PagerDuty) {
		return true
	}

	return false
}

// SetPagerDuty gets a reference to the given PagerDutyConfiguration and assigns it to the PagerDuty field.
func (o *UpdateNotificationChannelRequest) SetPagerDuty(v PagerDutyConfiguration) {
	o.PagerDuty = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetSlack() SlackConfiguration {
	if o == nil || IsNil(o.Slack) {
		var ret SlackConfiguration
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetSlackOk() (*SlackConfiguration, bool) {
	if o == nil || IsNil(o.Slack) {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasSlack() bool {
	if o != nil && !IsNil(o.Slack) {
		return true
	}

	return false
}

// SetSlack gets a reference to the given SlackConfiguration and assigns it to the Slack field.
func (o *UpdateNotificationChannelRequest) SetSlack(v SlackConfiguration) {
	o.Slack = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetSubscription() ChannelSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret ChannelSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetSubscriptionOk() (*ChannelSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given ChannelSubscription and assigns it to the Subscription field.
func (o *UpdateNotificationChannelRequest) SetSubscription(v ChannelSubscription) {
	o.Subscription = &v
}

// GetToken returns the Token field value
func (o *UpdateNotificationChannelRequest) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UpdateNotificationChannelRequest) SetToken(v string) {
	o.Token = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *UpdateNotificationChannelRequest) GetWebhook() WebhookConfiguration {
	if o == nil || IsNil(o.Webhook) {
		var ret WebhookConfiguration
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationChannelRequest) GetWebhookOk() (*WebhookConfiguration, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *UpdateNotificationChannelRequest) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given WebhookConfiguration and assigns it to the Webhook field.
func (o *UpdateNotificationChannelRequest) SetWebhook(v WebhookConfiguration) {
	o.Webhook = &v
}

func (o UpdateNotificationChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNotificationChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	toSerialize["id"] = o.Id
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
	toSerialize["token"] = o.Token
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateNotificationChannelRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"token",
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

	varUpdateNotificationChannelRequest := _UpdateNotificationChannelRequest{}

	err = json.Unmarshal(data, &varUpdateNotificationChannelRequest)

	if err != nil {
		return err
	}

	*o = UpdateNotificationChannelRequest(varUpdateNotificationChannelRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "pagerDuty")
		delete(additionalProperties, "slack")
		delete(additionalProperties, "subscription")
		delete(additionalProperties, "token")
		delete(additionalProperties, "webhook")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateNotificationChannelRequest struct {
	value *UpdateNotificationChannelRequest
	isSet bool
}

func (v NullableUpdateNotificationChannelRequest) Get() *UpdateNotificationChannelRequest {
	return v.value
}

func (v *NullableUpdateNotificationChannelRequest) Set(val *UpdateNotificationChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNotificationChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNotificationChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNotificationChannelRequest(val *UpdateNotificationChannelRequest) *NullableUpdateNotificationChannelRequest {
	return &NullableUpdateNotificationChannelRequest{value: val, isSet: true}
}

func (v NullableUpdateNotificationChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNotificationChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


