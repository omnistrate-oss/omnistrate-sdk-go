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

// checks if the ResourceNetworkTopologyResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceNetworkTopologyResult{}

// ResourceNetworkTopologyResult struct for ResourceNetworkTopologyResult
type ResourceNetworkTopologyResult struct {
	// The allowed IP ranges for this resource
	AllowedIPRanges []string `json:"allowedIPRanges"`
	// The cluster endpoint that load-balances across replicas of this resource
	ClusterEndpoint string `json:"clusterEndpoint"`
	// The ports that the whole cluster exposes
	ClusterPorts []int64 `json:"clusterPorts,omitempty"`
	CustomDNSEndpoint *CustomDNSEndpoint `json:"customDNSEndpoint,omitempty"`
	// Whether this resource has associated compute
	HasCompute bool `json:"hasCompute"`
	// Whether this is the main resource
	Main bool `json:"main"`
	// The networking type for this resource
	NetworkingType string `json:"networkingType"`
	// The nodes that this resource is deployed on
	Nodes []NodeNetworkTopologyResult `json:"nodes,omitempty"`
	// The private network CIDR for this resource
	PrivateNetworkCIDR *string `json:"privateNetworkCIDR,omitempty"`
	// The private network ID for this resource
	PrivateNetworkID *string `json:"privateNetworkID,omitempty"`
	ProxyEndpoint *ProxyEndpoint `json:"proxyEndpoint,omitempty"`
	// Whether this resource is publicly accessible
	PubliclyAccessible bool `json:"publiclyAccessible"`
	RecentDeploymentFailure *RecentDeploymentFailureStatus `json:"recentDeploymentFailure,omitempty"`
	// The resource instance metadata
	ResourceInstanceMetadata interface{} `json:"resourceInstanceMetadata,omitempty"`
	// The key of the resource
	ResourceKey string `json:"resourceKey"`
	// The name of the resource
	ResourceName string `json:"resourceName"`
	// The type of the resource
	ResourceType *string `json:"resourceType,omitempty"`
}

type _ResourceNetworkTopologyResult ResourceNetworkTopologyResult

// NewResourceNetworkTopologyResult instantiates a new ResourceNetworkTopologyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceNetworkTopologyResult(allowedIPRanges []string, clusterEndpoint string, hasCompute bool, main bool, networkingType string, publiclyAccessible bool, resourceKey string, resourceName string) *ResourceNetworkTopologyResult {
	this := ResourceNetworkTopologyResult{}
	this.AllowedIPRanges = allowedIPRanges
	this.ClusterEndpoint = clusterEndpoint
	this.HasCompute = hasCompute
	this.Main = main
	this.NetworkingType = networkingType
	this.PubliclyAccessible = publiclyAccessible
	this.ResourceKey = resourceKey
	this.ResourceName = resourceName
	return &this
}

// NewResourceNetworkTopologyResultWithDefaults instantiates a new ResourceNetworkTopologyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceNetworkTopologyResultWithDefaults() *ResourceNetworkTopologyResult {
	this := ResourceNetworkTopologyResult{}
	return &this
}

// GetAllowedIPRanges returns the AllowedIPRanges field value
func (o *ResourceNetworkTopologyResult) GetAllowedIPRanges() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AllowedIPRanges
}

// GetAllowedIPRangesOk returns a tuple with the AllowedIPRanges field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetAllowedIPRangesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowedIPRanges, true
}

// SetAllowedIPRanges sets field value
func (o *ResourceNetworkTopologyResult) SetAllowedIPRanges(v []string) {
	o.AllowedIPRanges = v
}

// GetClusterEndpoint returns the ClusterEndpoint field value
func (o *ResourceNetworkTopologyResult) GetClusterEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterEndpoint
}

// GetClusterEndpointOk returns a tuple with the ClusterEndpoint field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetClusterEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterEndpoint, true
}

// SetClusterEndpoint sets field value
func (o *ResourceNetworkTopologyResult) SetClusterEndpoint(v string) {
	o.ClusterEndpoint = v
}

// GetClusterPorts returns the ClusterPorts field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetClusterPorts() []int64 {
	if o == nil || IsNil(o.ClusterPorts) {
		var ret []int64
		return ret
	}
	return o.ClusterPorts
}

// GetClusterPortsOk returns a tuple with the ClusterPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetClusterPortsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ClusterPorts) {
		return nil, false
	}
	return o.ClusterPorts, true
}

// HasClusterPorts returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasClusterPorts() bool {
	if o != nil && !IsNil(o.ClusterPorts) {
		return true
	}

	return false
}

// SetClusterPorts gets a reference to the given []int64 and assigns it to the ClusterPorts field.
func (o *ResourceNetworkTopologyResult) SetClusterPorts(v []int64) {
	o.ClusterPorts = v
}

// GetCustomDNSEndpoint returns the CustomDNSEndpoint field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetCustomDNSEndpoint() CustomDNSEndpoint {
	if o == nil || IsNil(o.CustomDNSEndpoint) {
		var ret CustomDNSEndpoint
		return ret
	}
	return *o.CustomDNSEndpoint
}

// GetCustomDNSEndpointOk returns a tuple with the CustomDNSEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetCustomDNSEndpointOk() (*CustomDNSEndpoint, bool) {
	if o == nil || IsNil(o.CustomDNSEndpoint) {
		return nil, false
	}
	return o.CustomDNSEndpoint, true
}

// HasCustomDNSEndpoint returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasCustomDNSEndpoint() bool {
	if o != nil && !IsNil(o.CustomDNSEndpoint) {
		return true
	}

	return false
}

// SetCustomDNSEndpoint gets a reference to the given CustomDNSEndpoint and assigns it to the CustomDNSEndpoint field.
func (o *ResourceNetworkTopologyResult) SetCustomDNSEndpoint(v CustomDNSEndpoint) {
	o.CustomDNSEndpoint = &v
}

// GetHasCompute returns the HasCompute field value
func (o *ResourceNetworkTopologyResult) GetHasCompute() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasCompute
}

// GetHasComputeOk returns a tuple with the HasCompute field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetHasComputeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasCompute, true
}

// SetHasCompute sets field value
func (o *ResourceNetworkTopologyResult) SetHasCompute(v bool) {
	o.HasCompute = v
}

// GetMain returns the Main field value
func (o *ResourceNetworkTopologyResult) GetMain() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Main
}

// GetMainOk returns a tuple with the Main field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetMainOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Main, true
}

// SetMain sets field value
func (o *ResourceNetworkTopologyResult) SetMain(v bool) {
	o.Main = v
}

// GetNetworkingType returns the NetworkingType field value
func (o *ResourceNetworkTopologyResult) GetNetworkingType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkingType
}

// GetNetworkingTypeOk returns a tuple with the NetworkingType field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetNetworkingTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkingType, true
}

// SetNetworkingType sets field value
func (o *ResourceNetworkTopologyResult) SetNetworkingType(v string) {
	o.NetworkingType = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetNodes() []NodeNetworkTopologyResult {
	if o == nil || IsNil(o.Nodes) {
		var ret []NodeNetworkTopologyResult
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetNodesOk() ([]NodeNetworkTopologyResult, bool) {
	if o == nil || IsNil(o.Nodes) {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasNodes() bool {
	if o != nil && !IsNil(o.Nodes) {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []NodeNetworkTopologyResult and assigns it to the Nodes field.
func (o *ResourceNetworkTopologyResult) SetNodes(v []NodeNetworkTopologyResult) {
	o.Nodes = v
}

// GetPrivateNetworkCIDR returns the PrivateNetworkCIDR field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetPrivateNetworkCIDR() string {
	if o == nil || IsNil(o.PrivateNetworkCIDR) {
		var ret string
		return ret
	}
	return *o.PrivateNetworkCIDR
}

// GetPrivateNetworkCIDROk returns a tuple with the PrivateNetworkCIDR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetPrivateNetworkCIDROk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNetworkCIDR) {
		return nil, false
	}
	return o.PrivateNetworkCIDR, true
}

// HasPrivateNetworkCIDR returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasPrivateNetworkCIDR() bool {
	if o != nil && !IsNil(o.PrivateNetworkCIDR) {
		return true
	}

	return false
}

// SetPrivateNetworkCIDR gets a reference to the given string and assigns it to the PrivateNetworkCIDR field.
func (o *ResourceNetworkTopologyResult) SetPrivateNetworkCIDR(v string) {
	o.PrivateNetworkCIDR = &v
}

// GetPrivateNetworkID returns the PrivateNetworkID field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetPrivateNetworkID() string {
	if o == nil || IsNil(o.PrivateNetworkID) {
		var ret string
		return ret
	}
	return *o.PrivateNetworkID
}

// GetPrivateNetworkIDOk returns a tuple with the PrivateNetworkID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetPrivateNetworkIDOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNetworkID) {
		return nil, false
	}
	return o.PrivateNetworkID, true
}

// HasPrivateNetworkID returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasPrivateNetworkID() bool {
	if o != nil && !IsNil(o.PrivateNetworkID) {
		return true
	}

	return false
}

// SetPrivateNetworkID gets a reference to the given string and assigns it to the PrivateNetworkID field.
func (o *ResourceNetworkTopologyResult) SetPrivateNetworkID(v string) {
	o.PrivateNetworkID = &v
}

// GetProxyEndpoint returns the ProxyEndpoint field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetProxyEndpoint() ProxyEndpoint {
	if o == nil || IsNil(o.ProxyEndpoint) {
		var ret ProxyEndpoint
		return ret
	}
	return *o.ProxyEndpoint
}

// GetProxyEndpointOk returns a tuple with the ProxyEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetProxyEndpointOk() (*ProxyEndpoint, bool) {
	if o == nil || IsNil(o.ProxyEndpoint) {
		return nil, false
	}
	return o.ProxyEndpoint, true
}

// HasProxyEndpoint returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasProxyEndpoint() bool {
	if o != nil && !IsNil(o.ProxyEndpoint) {
		return true
	}

	return false
}

// SetProxyEndpoint gets a reference to the given ProxyEndpoint and assigns it to the ProxyEndpoint field.
func (o *ResourceNetworkTopologyResult) SetProxyEndpoint(v ProxyEndpoint) {
	o.ProxyEndpoint = &v
}

// GetPubliclyAccessible returns the PubliclyAccessible field value
func (o *ResourceNetworkTopologyResult) GetPubliclyAccessible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PubliclyAccessible
}

// GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetPubliclyAccessibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PubliclyAccessible, true
}

// SetPubliclyAccessible sets field value
func (o *ResourceNetworkTopologyResult) SetPubliclyAccessible(v bool) {
	o.PubliclyAccessible = v
}

// GetRecentDeploymentFailure returns the RecentDeploymentFailure field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetRecentDeploymentFailure() RecentDeploymentFailureStatus {
	if o == nil || IsNil(o.RecentDeploymentFailure) {
		var ret RecentDeploymentFailureStatus
		return ret
	}
	return *o.RecentDeploymentFailure
}

// GetRecentDeploymentFailureOk returns a tuple with the RecentDeploymentFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetRecentDeploymentFailureOk() (*RecentDeploymentFailureStatus, bool) {
	if o == nil || IsNil(o.RecentDeploymentFailure) {
		return nil, false
	}
	return o.RecentDeploymentFailure, true
}

// HasRecentDeploymentFailure returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasRecentDeploymentFailure() bool {
	if o != nil && !IsNil(o.RecentDeploymentFailure) {
		return true
	}

	return false
}

// SetRecentDeploymentFailure gets a reference to the given RecentDeploymentFailureStatus and assigns it to the RecentDeploymentFailure field.
func (o *ResourceNetworkTopologyResult) SetRecentDeploymentFailure(v RecentDeploymentFailureStatus) {
	o.RecentDeploymentFailure = &v
}

// GetResourceInstanceMetadata returns the ResourceInstanceMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceNetworkTopologyResult) GetResourceInstanceMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResourceInstanceMetadata
}

// GetResourceInstanceMetadataOk returns a tuple with the ResourceInstanceMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceNetworkTopologyResult) GetResourceInstanceMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResourceInstanceMetadata) {
		return nil, false
	}
	return &o.ResourceInstanceMetadata, true
}

// HasResourceInstanceMetadata returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasResourceInstanceMetadata() bool {
	if o != nil && !IsNil(o.ResourceInstanceMetadata) {
		return true
	}

	return false
}

// SetResourceInstanceMetadata gets a reference to the given interface{} and assigns it to the ResourceInstanceMetadata field.
func (o *ResourceNetworkTopologyResult) SetResourceInstanceMetadata(v interface{}) {
	o.ResourceInstanceMetadata = v
}

// GetResourceKey returns the ResourceKey field value
func (o *ResourceNetworkTopologyResult) GetResourceKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceKey
}

// GetResourceKeyOk returns a tuple with the ResourceKey field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetResourceKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceKey, true
}

// SetResourceKey sets field value
func (o *ResourceNetworkTopologyResult) SetResourceKey(v string) {
	o.ResourceKey = v
}

// GetResourceName returns the ResourceName field value
func (o *ResourceNetworkTopologyResult) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *ResourceNetworkTopologyResult) SetResourceName(v string) {
	o.ResourceName = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourceNetworkTopologyResult) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceNetworkTopologyResult) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourceNetworkTopologyResult) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourceNetworkTopologyResult) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o ResourceNetworkTopologyResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceNetworkTopologyResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowedIPRanges"] = o.AllowedIPRanges
	toSerialize["clusterEndpoint"] = o.ClusterEndpoint
	if !IsNil(o.ClusterPorts) {
		toSerialize["clusterPorts"] = o.ClusterPorts
	}
	if !IsNil(o.CustomDNSEndpoint) {
		toSerialize["customDNSEndpoint"] = o.CustomDNSEndpoint
	}
	toSerialize["hasCompute"] = o.HasCompute
	toSerialize["main"] = o.Main
	toSerialize["networkingType"] = o.NetworkingType
	if !IsNil(o.Nodes) {
		toSerialize["nodes"] = o.Nodes
	}
	if !IsNil(o.PrivateNetworkCIDR) {
		toSerialize["privateNetworkCIDR"] = o.PrivateNetworkCIDR
	}
	if !IsNil(o.PrivateNetworkID) {
		toSerialize["privateNetworkID"] = o.PrivateNetworkID
	}
	if !IsNil(o.ProxyEndpoint) {
		toSerialize["proxyEndpoint"] = o.ProxyEndpoint
	}
	toSerialize["publiclyAccessible"] = o.PubliclyAccessible
	if !IsNil(o.RecentDeploymentFailure) {
		toSerialize["recentDeploymentFailure"] = o.RecentDeploymentFailure
	}
	if o.ResourceInstanceMetadata != nil {
		toSerialize["resourceInstanceMetadata"] = o.ResourceInstanceMetadata
	}
	toSerialize["resourceKey"] = o.ResourceKey
	toSerialize["resourceName"] = o.ResourceName
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	return toSerialize, nil
}

func (o *ResourceNetworkTopologyResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"allowedIPRanges",
		"clusterEndpoint",
		"hasCompute",
		"main",
		"networkingType",
		"publiclyAccessible",
		"resourceKey",
		"resourceName",
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

	varResourceNetworkTopologyResult := _ResourceNetworkTopologyResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourceNetworkTopologyResult)

	if err != nil {
		return err
	}

	*o = ResourceNetworkTopologyResult(varResourceNetworkTopologyResult)

	return err
}

type NullableResourceNetworkTopologyResult struct {
	value *ResourceNetworkTopologyResult
	isSet bool
}

func (v NullableResourceNetworkTopologyResult) Get() *ResourceNetworkTopologyResult {
	return v.value
}

func (v *NullableResourceNetworkTopologyResult) Set(val *ResourceNetworkTopologyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceNetworkTopologyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceNetworkTopologyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceNetworkTopologyResult(val *ResourceNetworkTopologyResult) *NullableResourceNetworkTopologyResult {
	return &NullableResourceNetworkTopologyResult{value: val, isSet: true}
}

func (v NullableResourceNetworkTopologyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceNetworkTopologyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

