/*
Omnistrate Fleet API

REST API for Omnistrate Fleet

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fleet

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ServiceWorkflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceWorkflow{}

// ServiceWorkflow struct for ServiceWorkflow
type ServiceWorkflow struct {
	// The number of resources updated by the workflow.
	ResourceCount *int64 `json:"ResourceCount,omitempty"`
	// The user who updated the workflow execution. - Pause, Resume, Terminate
	UpdatedBy *string `json:"UpdatedBy,omitempty"`
	// The reason the workflow execution was terminated.
	UpdatedReason *string `json:"UpdatedReason,omitempty"`
	// The type of workflow execution.
	WorkflowType string `json:"WorkflowType"`
	// The AWS account ID
	AwsAccountID *string `json:"awsAccountID,omitempty"`
	// The cloud provider the workflow executed on.
	CloudProvider string `json:"cloudProvider"`
	// The time the workflow execution ended.
	EndTime *string `json:"endTime,omitempty"`
	// The GCP project ID
	GcpProjectID *string `json:"gcpProjectID,omitempty"`
	// ID of the ServiceWorkflow
	Id string `json:"id"`
	// The name of the instance owner organization.
	OrgName string `json:"orgName"`
	// The parent workflow's id for the execution.
	ParentId *string `json:"parentId,omitempty"`
	// The plan type of the instance owner organization.
	PlanType *string `json:"planType,omitempty"`
	// The service plan name for the workflow.
	ServicePlanName *string `json:"servicePlanName,omitempty"`
	// The time the workflow execution started.
	StartTime string `json:"startTime"`
	// The status of the workflow execution.
	Status string `json:"status"`
}

type _ServiceWorkflow ServiceWorkflow

// NewServiceWorkflow instantiates a new ServiceWorkflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceWorkflow(workflowType string, cloudProvider string, id string, orgName string, startTime string, status string) *ServiceWorkflow {
	this := ServiceWorkflow{}
	this.WorkflowType = workflowType
	this.CloudProvider = cloudProvider
	this.Id = id
	this.OrgName = orgName
	this.StartTime = startTime
	this.Status = status
	return &this
}

// NewServiceWorkflowWithDefaults instantiates a new ServiceWorkflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWorkflowWithDefaults() *ServiceWorkflow {
	this := ServiceWorkflow{}
	return &this
}

// GetResourceCount returns the ResourceCount field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetResourceCount() int64 {
	if o == nil || IsNil(o.ResourceCount) {
		var ret int64
		return ret
	}
	return *o.ResourceCount
}

// GetResourceCountOk returns a tuple with the ResourceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetResourceCountOk() (*int64, bool) {
	if o == nil || IsNil(o.ResourceCount) {
		return nil, false
	}
	return o.ResourceCount, true
}

// HasResourceCount returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasResourceCount() bool {
	if o != nil && !IsNil(o.ResourceCount) {
		return true
	}

	return false
}

// SetResourceCount gets a reference to the given int64 and assigns it to the ResourceCount field.
func (o *ServiceWorkflow) SetResourceCount(v int64) {
	o.ResourceCount = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *ServiceWorkflow) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedReason returns the UpdatedReason field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetUpdatedReason() string {
	if o == nil || IsNil(o.UpdatedReason) {
		var ret string
		return ret
	}
	return *o.UpdatedReason
}

// GetUpdatedReasonOk returns a tuple with the UpdatedReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetUpdatedReasonOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedReason) {
		return nil, false
	}
	return o.UpdatedReason, true
}

// HasUpdatedReason returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasUpdatedReason() bool {
	if o != nil && !IsNil(o.UpdatedReason) {
		return true
	}

	return false
}

// SetUpdatedReason gets a reference to the given string and assigns it to the UpdatedReason field.
func (o *ServiceWorkflow) SetUpdatedReason(v string) {
	o.UpdatedReason = &v
}

// GetWorkflowType returns the WorkflowType field value
func (o *ServiceWorkflow) GetWorkflowType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowType
}

// GetWorkflowTypeOk returns a tuple with the WorkflowType field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetWorkflowTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowType, true
}

// SetWorkflowType sets field value
func (o *ServiceWorkflow) SetWorkflowType(v string) {
	o.WorkflowType = v
}

// GetAwsAccountID returns the AwsAccountID field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetAwsAccountID() string {
	if o == nil || IsNil(o.AwsAccountID) {
		var ret string
		return ret
	}
	return *o.AwsAccountID
}

// GetAwsAccountIDOk returns a tuple with the AwsAccountID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetAwsAccountIDOk() (*string, bool) {
	if o == nil || IsNil(o.AwsAccountID) {
		return nil, false
	}
	return o.AwsAccountID, true
}

// HasAwsAccountID returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasAwsAccountID() bool {
	if o != nil && !IsNil(o.AwsAccountID) {
		return true
	}

	return false
}

// SetAwsAccountID gets a reference to the given string and assigns it to the AwsAccountID field.
func (o *ServiceWorkflow) SetAwsAccountID(v string) {
	o.AwsAccountID = &v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ServiceWorkflow) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ServiceWorkflow) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *ServiceWorkflow) SetEndTime(v string) {
	o.EndTime = &v
}

// GetGcpProjectID returns the GcpProjectID field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetGcpProjectID() string {
	if o == nil || IsNil(o.GcpProjectID) {
		var ret string
		return ret
	}
	return *o.GcpProjectID
}

// GetGcpProjectIDOk returns a tuple with the GcpProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetGcpProjectIDOk() (*string, bool) {
	if o == nil || IsNil(o.GcpProjectID) {
		return nil, false
	}
	return o.GcpProjectID, true
}

// HasGcpProjectID returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasGcpProjectID() bool {
	if o != nil && !IsNil(o.GcpProjectID) {
		return true
	}

	return false
}

// SetGcpProjectID gets a reference to the given string and assigns it to the GcpProjectID field.
func (o *ServiceWorkflow) SetGcpProjectID(v string) {
	o.GcpProjectID = &v
}

// GetId returns the Id field value
func (o *ServiceWorkflow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceWorkflow) SetId(v string) {
	o.Id = v
}

// GetOrgName returns the OrgName field value
func (o *ServiceWorkflow) GetOrgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetOrgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgName, true
}

// SetOrgName sets field value
func (o *ServiceWorkflow) SetOrgName(v string) {
	o.OrgName = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *ServiceWorkflow) SetParentId(v string) {
	o.ParentId = &v
}

// GetPlanType returns the PlanType field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetPlanType() string {
	if o == nil || IsNil(o.PlanType) {
		var ret string
		return ret
	}
	return *o.PlanType
}

// GetPlanTypeOk returns a tuple with the PlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetPlanTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PlanType) {
		return nil, false
	}
	return o.PlanType, true
}

// HasPlanType returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasPlanType() bool {
	if o != nil && !IsNil(o.PlanType) {
		return true
	}

	return false
}

// SetPlanType gets a reference to the given string and assigns it to the PlanType field.
func (o *ServiceWorkflow) SetPlanType(v string) {
	o.PlanType = &v
}

// GetServicePlanName returns the ServicePlanName field value if set, zero value otherwise.
func (o *ServiceWorkflow) GetServicePlanName() string {
	if o == nil || IsNil(o.ServicePlanName) {
		var ret string
		return ret
	}
	return *o.ServicePlanName
}

// GetServicePlanNameOk returns a tuple with the ServicePlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetServicePlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePlanName) {
		return nil, false
	}
	return o.ServicePlanName, true
}

// HasServicePlanName returns a boolean if a field has been set.
func (o *ServiceWorkflow) HasServicePlanName() bool {
	if o != nil && !IsNil(o.ServicePlanName) {
		return true
	}

	return false
}

// SetServicePlanName gets a reference to the given string and assigns it to the ServicePlanName field.
func (o *ServiceWorkflow) SetServicePlanName(v string) {
	o.ServicePlanName = &v
}

// GetStartTime returns the StartTime field value
func (o *ServiceWorkflow) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ServiceWorkflow) SetStartTime(v string) {
	o.StartTime = v
}

// GetStatus returns the Status field value
func (o *ServiceWorkflow) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ServiceWorkflow) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ServiceWorkflow) SetStatus(v string) {
	o.Status = v
}

func (o ServiceWorkflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceWorkflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceCount) {
		toSerialize["ResourceCount"] = o.ResourceCount
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["UpdatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedReason) {
		toSerialize["UpdatedReason"] = o.UpdatedReason
	}
	toSerialize["WorkflowType"] = o.WorkflowType
	if !IsNil(o.AwsAccountID) {
		toSerialize["awsAccountID"] = o.AwsAccountID
	}
	toSerialize["cloudProvider"] = o.CloudProvider
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.GcpProjectID) {
		toSerialize["gcpProjectID"] = o.GcpProjectID
	}
	toSerialize["id"] = o.Id
	toSerialize["orgName"] = o.OrgName
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.PlanType) {
		toSerialize["planType"] = o.PlanType
	}
	if !IsNil(o.ServicePlanName) {
		toSerialize["servicePlanName"] = o.ServicePlanName
	}
	toSerialize["startTime"] = o.StartTime
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ServiceWorkflow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"WorkflowType",
		"cloudProvider",
		"id",
		"orgName",
		"startTime",
		"status",
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

	varServiceWorkflow := _ServiceWorkflow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceWorkflow)

	if err != nil {
		return err
	}

	*o = ServiceWorkflow(varServiceWorkflow)

	return err
}

type NullableServiceWorkflow struct {
	value *ServiceWorkflow
	isSet bool
}

func (v NullableServiceWorkflow) Get() *ServiceWorkflow {
	return v.value
}

func (v *NullableServiceWorkflow) Set(val *ServiceWorkflow) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceWorkflow(val *ServiceWorkflow) *NullableServiceWorkflow {
	return &NullableServiceWorkflow{value: val, isSet: true}
}

func (v NullableServiceWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


