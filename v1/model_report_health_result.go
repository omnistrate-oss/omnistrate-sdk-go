/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the ReportHealthResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportHealthResult{}

// ReportHealthResult struct for ReportHealthResult
type ReportHealthResult struct {
	// The health of the service
	Health *string `json:"health,omitempty"`
	// The ID of the service
	Id *string `json:"id,omitempty"`
	// The health of each environment under this service
	Report *map[string]EnvironmentHealthReport `json:"report,omitempty"`
}

// NewReportHealthResult instantiates a new ReportHealthResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportHealthResult() *ReportHealthResult {
	this := ReportHealthResult{}
	return &this
}

// NewReportHealthResultWithDefaults instantiates a new ReportHealthResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportHealthResultWithDefaults() *ReportHealthResult {
	this := ReportHealthResult{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *ReportHealthResult) GetHealth() string {
	if o == nil || IsNil(o.Health) {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportHealthResult) GetHealthOk() (*string, bool) {
	if o == nil || IsNil(o.Health) {
		return nil, false
	}
	return o.Health, true
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *ReportHealthResult) SetHealth(v string) {
	o.Health = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportHealthResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportHealthResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportHealthResult) SetId(v string) {
	o.Id = &v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *ReportHealthResult) GetReport() map[string]EnvironmentHealthReport {
	if o == nil || IsNil(o.Report) {
		var ret map[string]EnvironmentHealthReport
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportHealthResult) GetReportOk() (*map[string]EnvironmentHealthReport, bool) {
	if o == nil || IsNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// SetReport gets a reference to the given map[string]EnvironmentHealthReport and assigns it to the Report field.
func (o *ReportHealthResult) SetReport(v map[string]EnvironmentHealthReport) {
	o.Report = &v
}

func (o ReportHealthResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportHealthResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Health) {
		toSerialize["health"] = o.Health
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	return toSerialize, nil
}

type NullableReportHealthResult struct {
	value *ReportHealthResult
	isSet bool
}

func (v NullableReportHealthResult) Get() *ReportHealthResult {
	return v.value
}

func (v *NullableReportHealthResult) Set(val *ReportHealthResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReportHealthResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReportHealthResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportHealthResult(val *ReportHealthResult) *NullableReportHealthResult {
	return &NullableReportHealthResult{value: val, isSet: true}
}

func (v NullableReportHealthResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportHealthResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

