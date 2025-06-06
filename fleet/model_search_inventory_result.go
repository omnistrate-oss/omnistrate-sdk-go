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

// checks if the SearchInventoryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchInventoryResult{}

// SearchInventoryResult struct for SearchInventoryResult
type SearchInventoryResult struct {
	// The deployment cell search results
	DeploymentCellResults []DeploymentCellSearchRecord `json:"deploymentCellResults,omitempty"`
	// The notification search results
	NotificationResults []NotificationSearchRecord `json:"notificationResults,omitempty"`
	// The proxy instance search results
	ProxyInstanceResults []ProxyInstanceSearchRecord `json:"proxyInstanceResults,omitempty"`
	// The resource instance search results
	ResourceInstanceResults []ResourceInstanceSearchRecord `json:"resourceInstanceResults,omitempty"`
	// The resource search results
	ResourceResults []ResourceSearchRecord `json:"resourceResults,omitempty"`
	// The serverless proxy search results
	ServerlessProxyResults []ServerlessProxySearchRecord `json:"serverlessProxyResults,omitempty"`
	// The service plan search results
	ServicePlanResults []ServicePlanSearchRecord `json:"servicePlanResults,omitempty"`
	// The service search results
	ServiceResults []ServiceSearchRecord `json:"serviceResults,omitempty"`
	// The subscription search results
	SubscriptionResults []SubscriptionSearchRecord `json:"subscriptionResults,omitempty"`
	// The upgrade path search results
	UpgradePathResults []UpgradePathSearchRecord `json:"upgradePathResults,omitempty"`
	// The user search results
	UserResults []UserSearchRecord `json:"userResults,omitempty"`
	// The workflow search results
	WorkflowResults []WorkflowSearchRecord `json:"workflowResults,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchInventoryResult SearchInventoryResult

// NewSearchInventoryResult instantiates a new SearchInventoryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchInventoryResult() *SearchInventoryResult {
	this := SearchInventoryResult{}
	return &this
}

// NewSearchInventoryResultWithDefaults instantiates a new SearchInventoryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchInventoryResultWithDefaults() *SearchInventoryResult {
	this := SearchInventoryResult{}
	return &this
}

// GetDeploymentCellResults returns the DeploymentCellResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetDeploymentCellResults() []DeploymentCellSearchRecord {
	if o == nil || IsNil(o.DeploymentCellResults) {
		var ret []DeploymentCellSearchRecord
		return ret
	}
	return o.DeploymentCellResults
}

// GetDeploymentCellResultsOk returns a tuple with the DeploymentCellResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetDeploymentCellResultsOk() ([]DeploymentCellSearchRecord, bool) {
	if o == nil || IsNil(o.DeploymentCellResults) {
		return nil, false
	}
	return o.DeploymentCellResults, true
}

// HasDeploymentCellResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasDeploymentCellResults() bool {
	if o != nil && !IsNil(o.DeploymentCellResults) {
		return true
	}

	return false
}

// SetDeploymentCellResults gets a reference to the given []DeploymentCellSearchRecord and assigns it to the DeploymentCellResults field.
func (o *SearchInventoryResult) SetDeploymentCellResults(v []DeploymentCellSearchRecord) {
	o.DeploymentCellResults = v
}

// GetNotificationResults returns the NotificationResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetNotificationResults() []NotificationSearchRecord {
	if o == nil || IsNil(o.NotificationResults) {
		var ret []NotificationSearchRecord
		return ret
	}
	return o.NotificationResults
}

// GetNotificationResultsOk returns a tuple with the NotificationResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetNotificationResultsOk() ([]NotificationSearchRecord, bool) {
	if o == nil || IsNil(o.NotificationResults) {
		return nil, false
	}
	return o.NotificationResults, true
}

// HasNotificationResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasNotificationResults() bool {
	if o != nil && !IsNil(o.NotificationResults) {
		return true
	}

	return false
}

// SetNotificationResults gets a reference to the given []NotificationSearchRecord and assigns it to the NotificationResults field.
func (o *SearchInventoryResult) SetNotificationResults(v []NotificationSearchRecord) {
	o.NotificationResults = v
}

// GetProxyInstanceResults returns the ProxyInstanceResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetProxyInstanceResults() []ProxyInstanceSearchRecord {
	if o == nil || IsNil(o.ProxyInstanceResults) {
		var ret []ProxyInstanceSearchRecord
		return ret
	}
	return o.ProxyInstanceResults
}

// GetProxyInstanceResultsOk returns a tuple with the ProxyInstanceResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetProxyInstanceResultsOk() ([]ProxyInstanceSearchRecord, bool) {
	if o == nil || IsNil(o.ProxyInstanceResults) {
		return nil, false
	}
	return o.ProxyInstanceResults, true
}

// HasProxyInstanceResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasProxyInstanceResults() bool {
	if o != nil && !IsNil(o.ProxyInstanceResults) {
		return true
	}

	return false
}

// SetProxyInstanceResults gets a reference to the given []ProxyInstanceSearchRecord and assigns it to the ProxyInstanceResults field.
func (o *SearchInventoryResult) SetProxyInstanceResults(v []ProxyInstanceSearchRecord) {
	o.ProxyInstanceResults = v
}

// GetResourceInstanceResults returns the ResourceInstanceResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetResourceInstanceResults() []ResourceInstanceSearchRecord {
	if o == nil || IsNil(o.ResourceInstanceResults) {
		var ret []ResourceInstanceSearchRecord
		return ret
	}
	return o.ResourceInstanceResults
}

// GetResourceInstanceResultsOk returns a tuple with the ResourceInstanceResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetResourceInstanceResultsOk() ([]ResourceInstanceSearchRecord, bool) {
	if o == nil || IsNil(o.ResourceInstanceResults) {
		return nil, false
	}
	return o.ResourceInstanceResults, true
}

// HasResourceInstanceResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasResourceInstanceResults() bool {
	if o != nil && !IsNil(o.ResourceInstanceResults) {
		return true
	}

	return false
}

// SetResourceInstanceResults gets a reference to the given []ResourceInstanceSearchRecord and assigns it to the ResourceInstanceResults field.
func (o *SearchInventoryResult) SetResourceInstanceResults(v []ResourceInstanceSearchRecord) {
	o.ResourceInstanceResults = v
}

// GetResourceResults returns the ResourceResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetResourceResults() []ResourceSearchRecord {
	if o == nil || IsNil(o.ResourceResults) {
		var ret []ResourceSearchRecord
		return ret
	}
	return o.ResourceResults
}

// GetResourceResultsOk returns a tuple with the ResourceResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetResourceResultsOk() ([]ResourceSearchRecord, bool) {
	if o == nil || IsNil(o.ResourceResults) {
		return nil, false
	}
	return o.ResourceResults, true
}

// HasResourceResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasResourceResults() bool {
	if o != nil && !IsNil(o.ResourceResults) {
		return true
	}

	return false
}

// SetResourceResults gets a reference to the given []ResourceSearchRecord and assigns it to the ResourceResults field.
func (o *SearchInventoryResult) SetResourceResults(v []ResourceSearchRecord) {
	o.ResourceResults = v
}

// GetServerlessProxyResults returns the ServerlessProxyResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetServerlessProxyResults() []ServerlessProxySearchRecord {
	if o == nil || IsNil(o.ServerlessProxyResults) {
		var ret []ServerlessProxySearchRecord
		return ret
	}
	return o.ServerlessProxyResults
}

// GetServerlessProxyResultsOk returns a tuple with the ServerlessProxyResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetServerlessProxyResultsOk() ([]ServerlessProxySearchRecord, bool) {
	if o == nil || IsNil(o.ServerlessProxyResults) {
		return nil, false
	}
	return o.ServerlessProxyResults, true
}

// HasServerlessProxyResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasServerlessProxyResults() bool {
	if o != nil && !IsNil(o.ServerlessProxyResults) {
		return true
	}

	return false
}

// SetServerlessProxyResults gets a reference to the given []ServerlessProxySearchRecord and assigns it to the ServerlessProxyResults field.
func (o *SearchInventoryResult) SetServerlessProxyResults(v []ServerlessProxySearchRecord) {
	o.ServerlessProxyResults = v
}

// GetServicePlanResults returns the ServicePlanResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetServicePlanResults() []ServicePlanSearchRecord {
	if o == nil || IsNil(o.ServicePlanResults) {
		var ret []ServicePlanSearchRecord
		return ret
	}
	return o.ServicePlanResults
}

// GetServicePlanResultsOk returns a tuple with the ServicePlanResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetServicePlanResultsOk() ([]ServicePlanSearchRecord, bool) {
	if o == nil || IsNil(o.ServicePlanResults) {
		return nil, false
	}
	return o.ServicePlanResults, true
}

// HasServicePlanResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasServicePlanResults() bool {
	if o != nil && !IsNil(o.ServicePlanResults) {
		return true
	}

	return false
}

// SetServicePlanResults gets a reference to the given []ServicePlanSearchRecord and assigns it to the ServicePlanResults field.
func (o *SearchInventoryResult) SetServicePlanResults(v []ServicePlanSearchRecord) {
	o.ServicePlanResults = v
}

// GetServiceResults returns the ServiceResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetServiceResults() []ServiceSearchRecord {
	if o == nil || IsNil(o.ServiceResults) {
		var ret []ServiceSearchRecord
		return ret
	}
	return o.ServiceResults
}

// GetServiceResultsOk returns a tuple with the ServiceResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetServiceResultsOk() ([]ServiceSearchRecord, bool) {
	if o == nil || IsNil(o.ServiceResults) {
		return nil, false
	}
	return o.ServiceResults, true
}

// HasServiceResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasServiceResults() bool {
	if o != nil && !IsNil(o.ServiceResults) {
		return true
	}

	return false
}

// SetServiceResults gets a reference to the given []ServiceSearchRecord and assigns it to the ServiceResults field.
func (o *SearchInventoryResult) SetServiceResults(v []ServiceSearchRecord) {
	o.ServiceResults = v
}

// GetSubscriptionResults returns the SubscriptionResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetSubscriptionResults() []SubscriptionSearchRecord {
	if o == nil || IsNil(o.SubscriptionResults) {
		var ret []SubscriptionSearchRecord
		return ret
	}
	return o.SubscriptionResults
}

// GetSubscriptionResultsOk returns a tuple with the SubscriptionResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetSubscriptionResultsOk() ([]SubscriptionSearchRecord, bool) {
	if o == nil || IsNil(o.SubscriptionResults) {
		return nil, false
	}
	return o.SubscriptionResults, true
}

// HasSubscriptionResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasSubscriptionResults() bool {
	if o != nil && !IsNil(o.SubscriptionResults) {
		return true
	}

	return false
}

// SetSubscriptionResults gets a reference to the given []SubscriptionSearchRecord and assigns it to the SubscriptionResults field.
func (o *SearchInventoryResult) SetSubscriptionResults(v []SubscriptionSearchRecord) {
	o.SubscriptionResults = v
}

// GetUpgradePathResults returns the UpgradePathResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetUpgradePathResults() []UpgradePathSearchRecord {
	if o == nil || IsNil(o.UpgradePathResults) {
		var ret []UpgradePathSearchRecord
		return ret
	}
	return o.UpgradePathResults
}

// GetUpgradePathResultsOk returns a tuple with the UpgradePathResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetUpgradePathResultsOk() ([]UpgradePathSearchRecord, bool) {
	if o == nil || IsNil(o.UpgradePathResults) {
		return nil, false
	}
	return o.UpgradePathResults, true
}

// HasUpgradePathResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasUpgradePathResults() bool {
	if o != nil && !IsNil(o.UpgradePathResults) {
		return true
	}

	return false
}

// SetUpgradePathResults gets a reference to the given []UpgradePathSearchRecord and assigns it to the UpgradePathResults field.
func (o *SearchInventoryResult) SetUpgradePathResults(v []UpgradePathSearchRecord) {
	o.UpgradePathResults = v
}

// GetUserResults returns the UserResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetUserResults() []UserSearchRecord {
	if o == nil || IsNil(o.UserResults) {
		var ret []UserSearchRecord
		return ret
	}
	return o.UserResults
}

// GetUserResultsOk returns a tuple with the UserResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetUserResultsOk() ([]UserSearchRecord, bool) {
	if o == nil || IsNil(o.UserResults) {
		return nil, false
	}
	return o.UserResults, true
}

// HasUserResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasUserResults() bool {
	if o != nil && !IsNil(o.UserResults) {
		return true
	}

	return false
}

// SetUserResults gets a reference to the given []UserSearchRecord and assigns it to the UserResults field.
func (o *SearchInventoryResult) SetUserResults(v []UserSearchRecord) {
	o.UserResults = v
}

// GetWorkflowResults returns the WorkflowResults field value if set, zero value otherwise.
func (o *SearchInventoryResult) GetWorkflowResults() []WorkflowSearchRecord {
	if o == nil || IsNil(o.WorkflowResults) {
		var ret []WorkflowSearchRecord
		return ret
	}
	return o.WorkflowResults
}

// GetWorkflowResultsOk returns a tuple with the WorkflowResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchInventoryResult) GetWorkflowResultsOk() ([]WorkflowSearchRecord, bool) {
	if o == nil || IsNil(o.WorkflowResults) {
		return nil, false
	}
	return o.WorkflowResults, true
}

// HasWorkflowResults returns a boolean if a field has been set.
func (o *SearchInventoryResult) HasWorkflowResults() bool {
	if o != nil && !IsNil(o.WorkflowResults) {
		return true
	}

	return false
}

// SetWorkflowResults gets a reference to the given []WorkflowSearchRecord and assigns it to the WorkflowResults field.
func (o *SearchInventoryResult) SetWorkflowResults(v []WorkflowSearchRecord) {
	o.WorkflowResults = v
}

func (o SearchInventoryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchInventoryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeploymentCellResults) {
		toSerialize["deploymentCellResults"] = o.DeploymentCellResults
	}
	if !IsNil(o.NotificationResults) {
		toSerialize["notificationResults"] = o.NotificationResults
	}
	if !IsNil(o.ProxyInstanceResults) {
		toSerialize["proxyInstanceResults"] = o.ProxyInstanceResults
	}
	if !IsNil(o.ResourceInstanceResults) {
		toSerialize["resourceInstanceResults"] = o.ResourceInstanceResults
	}
	if !IsNil(o.ResourceResults) {
		toSerialize["resourceResults"] = o.ResourceResults
	}
	if !IsNil(o.ServerlessProxyResults) {
		toSerialize["serverlessProxyResults"] = o.ServerlessProxyResults
	}
	if !IsNil(o.ServicePlanResults) {
		toSerialize["servicePlanResults"] = o.ServicePlanResults
	}
	if !IsNil(o.ServiceResults) {
		toSerialize["serviceResults"] = o.ServiceResults
	}
	if !IsNil(o.SubscriptionResults) {
		toSerialize["subscriptionResults"] = o.SubscriptionResults
	}
	if !IsNil(o.UpgradePathResults) {
		toSerialize["upgradePathResults"] = o.UpgradePathResults
	}
	if !IsNil(o.UserResults) {
		toSerialize["userResults"] = o.UserResults
	}
	if !IsNil(o.WorkflowResults) {
		toSerialize["workflowResults"] = o.WorkflowResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchInventoryResult) UnmarshalJSON(data []byte) (err error) {
	varSearchInventoryResult := _SearchInventoryResult{}

	err = json.Unmarshal(data, &varSearchInventoryResult)

	if err != nil {
		return err
	}

	*o = SearchInventoryResult(varSearchInventoryResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deploymentCellResults")
		delete(additionalProperties, "notificationResults")
		delete(additionalProperties, "proxyInstanceResults")
		delete(additionalProperties, "resourceInstanceResults")
		delete(additionalProperties, "resourceResults")
		delete(additionalProperties, "serverlessProxyResults")
		delete(additionalProperties, "servicePlanResults")
		delete(additionalProperties, "serviceResults")
		delete(additionalProperties, "subscriptionResults")
		delete(additionalProperties, "upgradePathResults")
		delete(additionalProperties, "userResults")
		delete(additionalProperties, "workflowResults")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchInventoryResult struct {
	value *SearchInventoryResult
	isSet bool
}

func (v NullableSearchInventoryResult) Get() *SearchInventoryResult {
	return v.value
}

func (v *NullableSearchInventoryResult) Set(val *SearchInventoryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchInventoryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchInventoryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchInventoryResult(val *SearchInventoryResult) *NullableSearchInventoryResult {
	return &NullableSearchInventoryResult{value: val, isSet: true}
}

func (v NullableSearchInventoryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchInventoryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


