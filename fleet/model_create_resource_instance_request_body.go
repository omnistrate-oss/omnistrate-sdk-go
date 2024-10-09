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

// checks if the CreateResourceInstanceRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateResourceInstanceRequestBody{}

// CreateResourceInstanceRequestBody struct for CreateResourceInstanceRequestBody
type CreateResourceInstanceRequestBody struct {
	// The cloud provider name
	CloudProvider *string `json:"cloud_provider,omitempty"`
	// Custom network for resource
	CustomNetworkId *string `json:"custom_network_id,omitempty"`
	// The network type
	NetworkType *string `json:"network_type,omitempty"`
	// The product tier version
	ProductTierVersion *string `json:"productTierVersion,omitempty"`
	// The region code
	Region *string `json:"region,omitempty"`
	// The request parameters
	RequestParams interface{} `json:"requestParams,omitempty"`
	// The subscription ID
	SubscriptionId *string `json:"subscriptionId,omitempty"`
}

// NewCreateResourceInstanceRequestBody instantiates a new CreateResourceInstanceRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourceInstanceRequestBody() *CreateResourceInstanceRequestBody {
	this := CreateResourceInstanceRequestBody{}
	return &this
}

// NewCreateResourceInstanceRequestBodyWithDefaults instantiates a new CreateResourceInstanceRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourceInstanceRequestBodyWithDefaults() *CreateResourceInstanceRequestBody {
	this := CreateResourceInstanceRequestBody{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *CreateResourceInstanceRequestBody) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetCustomNetworkId returns the CustomNetworkId field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetCustomNetworkId() string {
	if o == nil || IsNil(o.CustomNetworkId) {
		var ret string
		return ret
	}
	return *o.CustomNetworkId
}

// GetCustomNetworkIdOk returns a tuple with the CustomNetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetCustomNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomNetworkId) {
		return nil, false
	}
	return o.CustomNetworkId, true
}

// HasCustomNetworkId returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasCustomNetworkId() bool {
	if o != nil && !IsNil(o.CustomNetworkId) {
		return true
	}

	return false
}

// SetCustomNetworkId gets a reference to the given string and assigns it to the CustomNetworkId field.
func (o *CreateResourceInstanceRequestBody) SetCustomNetworkId(v string) {
	o.CustomNetworkId = &v
}

// GetNetworkType returns the NetworkType field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetNetworkType() string {
	if o == nil || IsNil(o.NetworkType) {
		var ret string
		return ret
	}
	return *o.NetworkType
}

// GetNetworkTypeOk returns a tuple with the NetworkType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetNetworkTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkType) {
		return nil, false
	}
	return o.NetworkType, true
}

// HasNetworkType returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasNetworkType() bool {
	if o != nil && !IsNil(o.NetworkType) {
		return true
	}

	return false
}

// SetNetworkType gets a reference to the given string and assigns it to the NetworkType field.
func (o *CreateResourceInstanceRequestBody) SetNetworkType(v string) {
	o.NetworkType = &v
}

// GetProductTierVersion returns the ProductTierVersion field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetProductTierVersion() string {
	if o == nil || IsNil(o.ProductTierVersion) {
		var ret string
		return ret
	}
	return *o.ProductTierVersion
}

// GetProductTierVersionOk returns a tuple with the ProductTierVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetProductTierVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ProductTierVersion) {
		return nil, false
	}
	return o.ProductTierVersion, true
}

// HasProductTierVersion returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasProductTierVersion() bool {
	if o != nil && !IsNil(o.ProductTierVersion) {
		return true
	}

	return false
}

// SetProductTierVersion gets a reference to the given string and assigns it to the ProductTierVersion field.
func (o *CreateResourceInstanceRequestBody) SetProductTierVersion(v string) {
	o.ProductTierVersion = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CreateResourceInstanceRequestBody) SetRegion(v string) {
	o.Region = &v
}

// GetRequestParams returns the RequestParams field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateResourceInstanceRequestBody) GetRequestParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RequestParams
}

// GetRequestParamsOk returns a tuple with the RequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateResourceInstanceRequestBody) GetRequestParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RequestParams) {
		return nil, false
	}
	return &o.RequestParams, true
}

// HasRequestParams returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasRequestParams() bool {
	if o != nil && !IsNil(o.RequestParams) {
		return true
	}

	return false
}

// SetRequestParams gets a reference to the given interface{} and assigns it to the RequestParams field.
func (o *CreateResourceInstanceRequestBody) SetRequestParams(v interface{}) {
	o.RequestParams = v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *CreateResourceInstanceRequestBody) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceInstanceRequestBody) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *CreateResourceInstanceRequestBody) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *CreateResourceInstanceRequestBody) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

func (o CreateResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateResourceInstanceRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.CustomNetworkId) {
		toSerialize["custom_network_id"] = o.CustomNetworkId
	}
	if !IsNil(o.NetworkType) {
		toSerialize["network_type"] = o.NetworkType
	}
	if !IsNil(o.ProductTierVersion) {
		toSerialize["productTierVersion"] = o.ProductTierVersion
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if o.RequestParams != nil {
		toSerialize["requestParams"] = o.RequestParams
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscriptionId"] = o.SubscriptionId
	}
	return toSerialize, nil
}

type NullableCreateResourceInstanceRequestBody struct {
	value *CreateResourceInstanceRequestBody
	isSet bool
}

func (v NullableCreateResourceInstanceRequestBody) Get() *CreateResourceInstanceRequestBody {
	return v.value
}

func (v *NullableCreateResourceInstanceRequestBody) Set(val *CreateResourceInstanceRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourceInstanceRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourceInstanceRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourceInstanceRequestBody(val *CreateResourceInstanceRequestBody) *NullableCreateResourceInstanceRequestBody {
	return &NullableCreateResourceInstanceRequestBody{value: val, isSet: true}
}

func (v NullableCreateResourceInstanceRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourceInstanceRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


