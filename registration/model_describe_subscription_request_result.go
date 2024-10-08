/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registration

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DescribeSubscriptionRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeSubscriptionRequestResult{}

// DescribeSubscriptionRequestResult struct for DescribeSubscriptionRequestResult
type DescribeSubscriptionRequestResult struct {
	// The time that this subscription request was issued
	CreatedAt string `json:"createdAt"`
	// The subscription ID
	Id string `json:"id"`
	// The product tier ID that this subscription is tied to
	ProductTierId string `json:"productTierId"`
	// The name of the product tier
	ProductTierName string `json:"productTierName"`
	// The email of the user that issued the subscription request
	RootUserEmail string `json:"rootUserEmail"`
	// The ID of the user that issued the subscription request
	RootUserId string `json:"rootUserId"`
	// The name of the user that issued the subscription request
	RootUserName string `json:"rootUserName"`
	// The service ID that this subscription is tied to
	ServiceId string `json:"serviceId"`
	// The logo for the service
	ServiceLogoURL *string `json:"serviceLogoURL,omitempty"`
	// The name of the service
	ServiceName string `json:"serviceName"`
	// The status of the subscription request
	Status string `json:"status"`
	// The time that this subscription request was issued
	UpdatedAt string `json:"updatedAt"`
	// The user ID that last updated the subscription request
	UpdatedByUserId string `json:"updatedByUserId"`
	// The user that last updated the subscription request
	UpdatedByUserName string `json:"updatedByUserName"`
}

type _DescribeSubscriptionRequestResult DescribeSubscriptionRequestResult

// NewDescribeSubscriptionRequestResult instantiates a new DescribeSubscriptionRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSubscriptionRequestResult(createdAt string, id string, productTierId string, productTierName string, rootUserEmail string, rootUserId string, rootUserName string, serviceId string, serviceName string, status string, updatedAt string, updatedByUserId string, updatedByUserName string) *DescribeSubscriptionRequestResult {
	this := DescribeSubscriptionRequestResult{}
	this.CreatedAt = createdAt
	this.Id = id
	this.ProductTierId = productTierId
	this.ProductTierName = productTierName
	this.RootUserEmail = rootUserEmail
	this.RootUserId = rootUserId
	this.RootUserName = rootUserName
	this.ServiceId = serviceId
	this.ServiceName = serviceName
	this.Status = status
	this.UpdatedAt = updatedAt
	this.UpdatedByUserId = updatedByUserId
	this.UpdatedByUserName = updatedByUserName
	return &this
}

// NewDescribeSubscriptionRequestResultWithDefaults instantiates a new DescribeSubscriptionRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSubscriptionRequestResultWithDefaults() *DescribeSubscriptionRequestResult {
	this := DescribeSubscriptionRequestResult{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *DescribeSubscriptionRequestResult) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DescribeSubscriptionRequestResult) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *DescribeSubscriptionRequestResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeSubscriptionRequestResult) SetId(v string) {
	o.Id = v
}

// GetProductTierId returns the ProductTierId field value
func (o *DescribeSubscriptionRequestResult) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *DescribeSubscriptionRequestResult) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProductTierName returns the ProductTierName field value
func (o *DescribeSubscriptionRequestResult) GetProductTierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierName
}

// GetProductTierNameOk returns a tuple with the ProductTierName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetProductTierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierName, true
}

// SetProductTierName sets field value
func (o *DescribeSubscriptionRequestResult) SetProductTierName(v string) {
	o.ProductTierName = v
}

// GetRootUserEmail returns the RootUserEmail field value
func (o *DescribeSubscriptionRequestResult) GetRootUserEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootUserEmail
}

// GetRootUserEmailOk returns a tuple with the RootUserEmail field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetRootUserEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootUserEmail, true
}

// SetRootUserEmail sets field value
func (o *DescribeSubscriptionRequestResult) SetRootUserEmail(v string) {
	o.RootUserEmail = v
}

// GetRootUserId returns the RootUserId field value
func (o *DescribeSubscriptionRequestResult) GetRootUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootUserId
}

// GetRootUserIdOk returns a tuple with the RootUserId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetRootUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootUserId, true
}

// SetRootUserId sets field value
func (o *DescribeSubscriptionRequestResult) SetRootUserId(v string) {
	o.RootUserId = v
}

// GetRootUserName returns the RootUserName field value
func (o *DescribeSubscriptionRequestResult) GetRootUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootUserName
}

// GetRootUserNameOk returns a tuple with the RootUserName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetRootUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootUserName, true
}

// SetRootUserName sets field value
func (o *DescribeSubscriptionRequestResult) SetRootUserName(v string) {
	o.RootUserName = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeSubscriptionRequestResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeSubscriptionRequestResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetServiceLogoURL returns the ServiceLogoURL field value if set, zero value otherwise.
func (o *DescribeSubscriptionRequestResult) GetServiceLogoURL() string {
	if o == nil || IsNil(o.ServiceLogoURL) {
		var ret string
		return ret
	}
	return *o.ServiceLogoURL
}

// GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetServiceLogoURLOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceLogoURL) {
		return nil, false
	}
	return o.ServiceLogoURL, true
}

// HasServiceLogoURL returns a boolean if a field has been set.
func (o *DescribeSubscriptionRequestResult) HasServiceLogoURL() bool {
	if o != nil && !IsNil(o.ServiceLogoURL) {
		return true
	}

	return false
}

// SetServiceLogoURL gets a reference to the given string and assigns it to the ServiceLogoURL field.
func (o *DescribeSubscriptionRequestResult) SetServiceLogoURL(v string) {
	o.ServiceLogoURL = &v
}

// GetServiceName returns the ServiceName field value
func (o *DescribeSubscriptionRequestResult) GetServiceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetServiceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *DescribeSubscriptionRequestResult) SetServiceName(v string) {
	o.ServiceName = v
}

// GetStatus returns the Status field value
func (o *DescribeSubscriptionRequestResult) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DescribeSubscriptionRequestResult) SetStatus(v string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *DescribeSubscriptionRequestResult) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *DescribeSubscriptionRequestResult) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value
func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByUserId, true
}

// SetUpdatedByUserId sets field value
func (o *DescribeSubscriptionRequestResult) SetUpdatedByUserId(v string) {
	o.UpdatedByUserId = v
}

// GetUpdatedByUserName returns the UpdatedByUserName field value
func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedByUserName
}

// GetUpdatedByUserNameOk returns a tuple with the UpdatedByUserName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByUserName, true
}

// SetUpdatedByUserName sets field value
func (o *DescribeSubscriptionRequestResult) SetUpdatedByUserName(v string) {
	o.UpdatedByUserName = v
}

func (o DescribeSubscriptionRequestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSubscriptionRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["id"] = o.Id
	toSerialize["productTierId"] = o.ProductTierId
	toSerialize["productTierName"] = o.ProductTierName
	toSerialize["rootUserEmail"] = o.RootUserEmail
	toSerialize["rootUserId"] = o.RootUserId
	toSerialize["rootUserName"] = o.RootUserName
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.ServiceLogoURL) {
		toSerialize["serviceLogoURL"] = o.ServiceLogoURL
	}
	toSerialize["serviceName"] = o.ServiceName
	toSerialize["status"] = o.Status
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["updatedByUserId"] = o.UpdatedByUserId
	toSerialize["updatedByUserName"] = o.UpdatedByUserName
	return toSerialize, nil
}

func (o *DescribeSubscriptionRequestResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"productTierId",
		"productTierName",
		"rootUserEmail",
		"rootUserId",
		"rootUserName",
		"serviceId",
		"serviceName",
		"status",
		"updatedAt",
		"updatedByUserId",
		"updatedByUserName",
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

	varDescribeSubscriptionRequestResult := _DescribeSubscriptionRequestResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeSubscriptionRequestResult)

	if err != nil {
		return err
	}

	*o = DescribeSubscriptionRequestResult(varDescribeSubscriptionRequestResult)

	return err
}

type NullableDescribeSubscriptionRequestResult struct {
	value *DescribeSubscriptionRequestResult
	isSet bool
}

func (v NullableDescribeSubscriptionRequestResult) Get() *DescribeSubscriptionRequestResult {
	return v.value
}

func (v *NullableDescribeSubscriptionRequestResult) Set(val *DescribeSubscriptionRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSubscriptionRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSubscriptionRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSubscriptionRequestResult(val *DescribeSubscriptionRequestResult) *NullableDescribeSubscriptionRequestResult {
	return &NullableDescribeSubscriptionRequestResult{value: val, isSet: true}
}

func (v NullableDescribeSubscriptionRequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSubscriptionRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


