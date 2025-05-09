/*
Omnistrate Registration API

REST API for Omnistrate Service Registration

API version: 2022-09-01-00
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// checks if the DescribeResourceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeResourceResult{}

// DescribeResourceResult struct for DescribeResourceResult
type DescribeResourceResult struct {
	// The action hooks that this resource supports
	ActionHooks []ActionHook `json:"actionHooks,omitempty"`
	AdditionalSecurityContext *AdditionalSecurityContext `json:"additionalSecurityContext,omitempty"`
	BackupConfiguration *BackupConfiguration `json:"backupConfiguration,omitempty"`
	BlobStorageConfiguration *BlobStorageConfiguration `json:"blobStorageConfiguration,omitempty"`
	// The capabilities enabled for the resource
	Capabilities []ResourceCapability `json:"capabilities,omitempty"`
	// Custom labels for the resource
	CustomLabels *map[string]string `json:"customLabels,omitempty"`
	// Custom sysctl settings for the resource
	CustomSysCTLs *map[string]string `json:"customSysCTLs,omitempty"`
	// Custom ulimits for the resource
	CustomULimits []CustomULimits `json:"customULimits,omitempty"`
	Dependencies []ResourceDependency `json:"dependencies,omitempty"`
	// A brief description of the resource
	Description string `json:"description"`
	// The environment variables that this resource requires
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables,omitempty"`
	FileSystemConfiguration *FileSystemConfiguration `json:"fileSystemConfiguration,omitempty"`
	HelmChartConfiguration *HelmChartConfiguration `json:"helmChartConfiguration,omitempty"`
	// ID of a resource
	Id string `json:"id"`
	// ID of an Image Config
	ImageConfigId *string `json:"imageConfigId,omitempty"`
	// ID of an Infra Config
	InfraConfigId *string `json:"infraConfigId,omitempty"`
	// Whether this resource is internal or not
	Internal bool `json:"internal"`
	// Whether this resource is deprecated or not
	IsDeprecated bool `json:"isDeprecated"`
	JobConfig *JobConfig `json:"jobConfig,omitempty"`
	// The key of the resource
	Key string `json:"key"`
	KustomizeConfiguration *KustomizeConfiguration `json:"kustomizeConfiguration,omitempty"`
	L4LoadBalancerConfiguration *L4LoadBalancerConfiguration `json:"l4LoadBalancerConfiguration,omitempty"`
	L7LoadBalancerConfiguration *L7LoadBalancerConfiguration `json:"l7LoadBalancerConfiguration,omitempty"`
	// Name of the resource
	Name string `json:"name"`
	OperatorCRDConfiguration *OperatorCRDConfiguration `json:"operatorCRDConfiguration,omitempty"`
	// ID of a Product Tier
	ProductTierId string `json:"productTierId"`
	// The proxy type of instance
	ProxyType *string `json:"proxyType,omitempty"`
	// The type of the resource
	ResourceType string `json:"resourceType"`
	// ID of a Service
	ServiceId string `json:"serviceId"`
	// The Terraform configurations for cloud providers
	TerraformConfigurations map[string]interface{} `json:"terraformConfigurations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DescribeResourceResult DescribeResourceResult

// NewDescribeResourceResult instantiates a new DescribeResourceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeResourceResult(description string, id string, internal bool, isDeprecated bool, key string, name string, productTierId string, resourceType string, serviceId string) *DescribeResourceResult {
	this := DescribeResourceResult{}
	this.Description = description
	this.Id = id
	this.Internal = internal
	this.IsDeprecated = isDeprecated
	this.Key = key
	this.Name = name
	this.ProductTierId = productTierId
	this.ResourceType = resourceType
	this.ServiceId = serviceId
	return &this
}

// NewDescribeResourceResultWithDefaults instantiates a new DescribeResourceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeResourceResultWithDefaults() *DescribeResourceResult {
	this := DescribeResourceResult{}
	var internal bool = false
	this.Internal = internal
	var isDeprecated bool = false
	this.IsDeprecated = isDeprecated
	return &this
}

// GetActionHooks returns the ActionHooks field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetActionHooks() []ActionHook {
	if o == nil || IsNil(o.ActionHooks) {
		var ret []ActionHook
		return ret
	}
	return o.ActionHooks
}

// GetActionHooksOk returns a tuple with the ActionHooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetActionHooksOk() ([]ActionHook, bool) {
	if o == nil || IsNil(o.ActionHooks) {
		return nil, false
	}
	return o.ActionHooks, true
}

// SetActionHooks gets a reference to the given []ActionHook and assigns it to the ActionHooks field.
func (o *DescribeResourceResult) SetActionHooks(v []ActionHook) {
	o.ActionHooks = v
}

// GetAdditionalSecurityContext returns the AdditionalSecurityContext field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetAdditionalSecurityContext() AdditionalSecurityContext {
	if o == nil || IsNil(o.AdditionalSecurityContext) {
		var ret AdditionalSecurityContext
		return ret
	}
	return *o.AdditionalSecurityContext
}

// GetAdditionalSecurityContextOk returns a tuple with the AdditionalSecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetAdditionalSecurityContextOk() (*AdditionalSecurityContext, bool) {
	if o == nil || IsNil(o.AdditionalSecurityContext) {
		return nil, false
	}
	return o.AdditionalSecurityContext, true
}

// SetAdditionalSecurityContext gets a reference to the given AdditionalSecurityContext and assigns it to the AdditionalSecurityContext field.
func (o *DescribeResourceResult) SetAdditionalSecurityContext(v AdditionalSecurityContext) {
	o.AdditionalSecurityContext = &v
}

// GetBackupConfiguration returns the BackupConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetBackupConfiguration() BackupConfiguration {
	if o == nil || IsNil(o.BackupConfiguration) {
		var ret BackupConfiguration
		return ret
	}
	return *o.BackupConfiguration
}

// GetBackupConfigurationOk returns a tuple with the BackupConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetBackupConfigurationOk() (*BackupConfiguration, bool) {
	if o == nil || IsNil(o.BackupConfiguration) {
		return nil, false
	}
	return o.BackupConfiguration, true
}

// SetBackupConfiguration gets a reference to the given BackupConfiguration and assigns it to the BackupConfiguration field.
func (o *DescribeResourceResult) SetBackupConfiguration(v BackupConfiguration) {
	o.BackupConfiguration = &v
}

// GetBlobStorageConfiguration returns the BlobStorageConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetBlobStorageConfiguration() BlobStorageConfiguration {
	if o == nil || IsNil(o.BlobStorageConfiguration) {
		var ret BlobStorageConfiguration
		return ret
	}
	return *o.BlobStorageConfiguration
}

// GetBlobStorageConfigurationOk returns a tuple with the BlobStorageConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetBlobStorageConfigurationOk() (*BlobStorageConfiguration, bool) {
	if o == nil || IsNil(o.BlobStorageConfiguration) {
		return nil, false
	}
	return o.BlobStorageConfiguration, true
}

// SetBlobStorageConfiguration gets a reference to the given BlobStorageConfiguration and assigns it to the BlobStorageConfiguration field.
func (o *DescribeResourceResult) SetBlobStorageConfiguration(v BlobStorageConfiguration) {
	o.BlobStorageConfiguration = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetCapabilities() []ResourceCapability {
	if o == nil || IsNil(o.Capabilities) {
		var ret []ResourceCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetCapabilitiesOk() ([]ResourceCapability, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// SetCapabilities gets a reference to the given []ResourceCapability and assigns it to the Capabilities field.
func (o *DescribeResourceResult) SetCapabilities(v []ResourceCapability) {
	o.Capabilities = v
}

// GetCustomLabels returns the CustomLabels field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetCustomLabels() map[string]string {
	if o == nil || IsNil(o.CustomLabels) {
		var ret map[string]string
		return ret
	}
	return *o.CustomLabels
}

// GetCustomLabelsOk returns a tuple with the CustomLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetCustomLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomLabels) {
		return nil, false
	}
	return o.CustomLabels, true
}

// SetCustomLabels gets a reference to the given map[string]string and assigns it to the CustomLabels field.
func (o *DescribeResourceResult) SetCustomLabels(v map[string]string) {
	o.CustomLabels = &v
}

// GetCustomSysCTLs returns the CustomSysCTLs field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetCustomSysCTLs() map[string]string {
	if o == nil || IsNil(o.CustomSysCTLs) {
		var ret map[string]string
		return ret
	}
	return *o.CustomSysCTLs
}

// GetCustomSysCTLsOk returns a tuple with the CustomSysCTLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetCustomSysCTLsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomSysCTLs) {
		return nil, false
	}
	return o.CustomSysCTLs, true
}

// SetCustomSysCTLs gets a reference to the given map[string]string and assigns it to the CustomSysCTLs field.
func (o *DescribeResourceResult) SetCustomSysCTLs(v map[string]string) {
	o.CustomSysCTLs = &v
}

// GetCustomULimits returns the CustomULimits field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetCustomULimits() []CustomULimits {
	if o == nil || IsNil(o.CustomULimits) {
		var ret []CustomULimits
		return ret
	}
	return o.CustomULimits
}

// GetCustomULimitsOk returns a tuple with the CustomULimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetCustomULimitsOk() ([]CustomULimits, bool) {
	if o == nil || IsNil(o.CustomULimits) {
		return nil, false
	}
	return o.CustomULimits, true
}

// SetCustomULimits gets a reference to the given []CustomULimits and assigns it to the CustomULimits field.
func (o *DescribeResourceResult) SetCustomULimits(v []CustomULimits) {
	o.CustomULimits = v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetDependencies() []ResourceDependency {
	if o == nil || IsNil(o.Dependencies) {
		var ret []ResourceDependency
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetDependenciesOk() ([]ResourceDependency, bool) {
	if o == nil || IsNil(o.Dependencies) {
		return nil, false
	}
	return o.Dependencies, true
}

// SetDependencies gets a reference to the given []ResourceDependency and assigns it to the Dependencies field.
func (o *DescribeResourceResult) SetDependencies(v []ResourceDependency) {
	o.Dependencies = v
}

// GetDescription returns the Description field value
func (o *DescribeResourceResult) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DescribeResourceResult) SetDescription(v string) {
	o.Description = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetEnvironmentVariables() []EnvironmentVariable {
	if o == nil || IsNil(o.EnvironmentVariables) {
		var ret []EnvironmentVariable
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetEnvironmentVariablesOk() ([]EnvironmentVariable, bool) {
	if o == nil || IsNil(o.EnvironmentVariables) {
		return nil, false
	}
	return o.EnvironmentVariables, true
}

// SetEnvironmentVariables gets a reference to the given []EnvironmentVariable and assigns it to the EnvironmentVariables field.
func (o *DescribeResourceResult) SetEnvironmentVariables(v []EnvironmentVariable) {
	o.EnvironmentVariables = v
}

// GetFileSystemConfiguration returns the FileSystemConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetFileSystemConfiguration() FileSystemConfiguration {
	if o == nil || IsNil(o.FileSystemConfiguration) {
		var ret FileSystemConfiguration
		return ret
	}
	return *o.FileSystemConfiguration
}

// GetFileSystemConfigurationOk returns a tuple with the FileSystemConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetFileSystemConfigurationOk() (*FileSystemConfiguration, bool) {
	if o == nil || IsNil(o.FileSystemConfiguration) {
		return nil, false
	}
	return o.FileSystemConfiguration, true
}

// SetFileSystemConfiguration gets a reference to the given FileSystemConfiguration and assigns it to the FileSystemConfiguration field.
func (o *DescribeResourceResult) SetFileSystemConfiguration(v FileSystemConfiguration) {
	o.FileSystemConfiguration = &v
}

// GetHelmChartConfiguration returns the HelmChartConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetHelmChartConfiguration() HelmChartConfiguration {
	if o == nil || IsNil(o.HelmChartConfiguration) {
		var ret HelmChartConfiguration
		return ret
	}
	return *o.HelmChartConfiguration
}

// GetHelmChartConfigurationOk returns a tuple with the HelmChartConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetHelmChartConfigurationOk() (*HelmChartConfiguration, bool) {
	if o == nil || IsNil(o.HelmChartConfiguration) {
		return nil, false
	}
	return o.HelmChartConfiguration, true
}

// SetHelmChartConfiguration gets a reference to the given HelmChartConfiguration and assigns it to the HelmChartConfiguration field.
func (o *DescribeResourceResult) SetHelmChartConfiguration(v HelmChartConfiguration) {
	o.HelmChartConfiguration = &v
}

// GetId returns the Id field value
func (o *DescribeResourceResult) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeResourceResult) SetId(v string) {
	o.Id = v
}

// GetImageConfigId returns the ImageConfigId field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetImageConfigId() string {
	if o == nil || IsNil(o.ImageConfigId) {
		var ret string
		return ret
	}
	return *o.ImageConfigId
}

// GetImageConfigIdOk returns a tuple with the ImageConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetImageConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageConfigId) {
		return nil, false
	}
	return o.ImageConfigId, true
}

// SetImageConfigId gets a reference to the given string and assigns it to the ImageConfigId field.
func (o *DescribeResourceResult) SetImageConfigId(v string) {
	o.ImageConfigId = &v
}

// GetInfraConfigId returns the InfraConfigId field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetInfraConfigId() string {
	if o == nil || IsNil(o.InfraConfigId) {
		var ret string
		return ret
	}
	return *o.InfraConfigId
}

// GetInfraConfigIdOk returns a tuple with the InfraConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetInfraConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.InfraConfigId) {
		return nil, false
	}
	return o.InfraConfigId, true
}

// SetInfraConfigId gets a reference to the given string and assigns it to the InfraConfigId field.
func (o *DescribeResourceResult) SetInfraConfigId(v string) {
	o.InfraConfigId = &v
}

// GetInternal returns the Internal field value
func (o *DescribeResourceResult) GetInternal() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Internal
}

// GetInternalOk returns a tuple with the Internal field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetInternalOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Internal, true
}

// SetInternal sets field value
func (o *DescribeResourceResult) SetInternal(v bool) {
	o.Internal = v
}

// GetIsDeprecated returns the IsDeprecated field value
func (o *DescribeResourceResult) GetIsDeprecated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeprecated
}

// GetIsDeprecatedOk returns a tuple with the IsDeprecated field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetIsDeprecatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeprecated, true
}

// SetIsDeprecated sets field value
func (o *DescribeResourceResult) SetIsDeprecated(v bool) {
	o.IsDeprecated = v
}

// GetJobConfig returns the JobConfig field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetJobConfig() JobConfig {
	if o == nil || IsNil(o.JobConfig) {
		var ret JobConfig
		return ret
	}
	return *o.JobConfig
}

// GetJobConfigOk returns a tuple with the JobConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetJobConfigOk() (*JobConfig, bool) {
	if o == nil || IsNil(o.JobConfig) {
		return nil, false
	}
	return o.JobConfig, true
}

// SetJobConfig gets a reference to the given JobConfig and assigns it to the JobConfig field.
func (o *DescribeResourceResult) SetJobConfig(v JobConfig) {
	o.JobConfig = &v
}

// GetKey returns the Key field value
func (o *DescribeResourceResult) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *DescribeResourceResult) SetKey(v string) {
	o.Key = v
}

// GetKustomizeConfiguration returns the KustomizeConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetKustomizeConfiguration() KustomizeConfiguration {
	if o == nil || IsNil(o.KustomizeConfiguration) {
		var ret KustomizeConfiguration
		return ret
	}
	return *o.KustomizeConfiguration
}

// GetKustomizeConfigurationOk returns a tuple with the KustomizeConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetKustomizeConfigurationOk() (*KustomizeConfiguration, bool) {
	if o == nil || IsNil(o.KustomizeConfiguration) {
		return nil, false
	}
	return o.KustomizeConfiguration, true
}

// SetKustomizeConfiguration gets a reference to the given KustomizeConfiguration and assigns it to the KustomizeConfiguration field.
func (o *DescribeResourceResult) SetKustomizeConfiguration(v KustomizeConfiguration) {
	o.KustomizeConfiguration = &v
}

// GetL4LoadBalancerConfiguration returns the L4LoadBalancerConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetL4LoadBalancerConfiguration() L4LoadBalancerConfiguration {
	if o == nil || IsNil(o.L4LoadBalancerConfiguration) {
		var ret L4LoadBalancerConfiguration
		return ret
	}
	return *o.L4LoadBalancerConfiguration
}

// GetL4LoadBalancerConfigurationOk returns a tuple with the L4LoadBalancerConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetL4LoadBalancerConfigurationOk() (*L4LoadBalancerConfiguration, bool) {
	if o == nil || IsNil(o.L4LoadBalancerConfiguration) {
		return nil, false
	}
	return o.L4LoadBalancerConfiguration, true
}

// SetL4LoadBalancerConfiguration gets a reference to the given L4LoadBalancerConfiguration and assigns it to the L4LoadBalancerConfiguration field.
func (o *DescribeResourceResult) SetL4LoadBalancerConfiguration(v L4LoadBalancerConfiguration) {
	o.L4LoadBalancerConfiguration = &v
}

// GetL7LoadBalancerConfiguration returns the L7LoadBalancerConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetL7LoadBalancerConfiguration() L7LoadBalancerConfiguration {
	if o == nil || IsNil(o.L7LoadBalancerConfiguration) {
		var ret L7LoadBalancerConfiguration
		return ret
	}
	return *o.L7LoadBalancerConfiguration
}

// GetL7LoadBalancerConfigurationOk returns a tuple with the L7LoadBalancerConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetL7LoadBalancerConfigurationOk() (*L7LoadBalancerConfiguration, bool) {
	if o == nil || IsNil(o.L7LoadBalancerConfiguration) {
		return nil, false
	}
	return o.L7LoadBalancerConfiguration, true
}

// SetL7LoadBalancerConfiguration gets a reference to the given L7LoadBalancerConfiguration and assigns it to the L7LoadBalancerConfiguration field.
func (o *DescribeResourceResult) SetL7LoadBalancerConfiguration(v L7LoadBalancerConfiguration) {
	o.L7LoadBalancerConfiguration = &v
}

// GetName returns the Name field value
func (o *DescribeResourceResult) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DescribeResourceResult) SetName(v string) {
	o.Name = v
}

// GetOperatorCRDConfiguration returns the OperatorCRDConfiguration field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetOperatorCRDConfiguration() OperatorCRDConfiguration {
	if o == nil || IsNil(o.OperatorCRDConfiguration) {
		var ret OperatorCRDConfiguration
		return ret
	}
	return *o.OperatorCRDConfiguration
}

// GetOperatorCRDConfigurationOk returns a tuple with the OperatorCRDConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetOperatorCRDConfigurationOk() (*OperatorCRDConfiguration, bool) {
	if o == nil || IsNil(o.OperatorCRDConfiguration) {
		return nil, false
	}
	return o.OperatorCRDConfiguration, true
}

// SetOperatorCRDConfiguration gets a reference to the given OperatorCRDConfiguration and assigns it to the OperatorCRDConfiguration field.
func (o *DescribeResourceResult) SetOperatorCRDConfiguration(v OperatorCRDConfiguration) {
	o.OperatorCRDConfiguration = &v
}

// GetProductTierId returns the ProductTierId field value
func (o *DescribeResourceResult) GetProductTierId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductTierId
}

// GetProductTierIdOk returns a tuple with the ProductTierId field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetProductTierIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductTierId, true
}

// SetProductTierId sets field value
func (o *DescribeResourceResult) SetProductTierId(v string) {
	o.ProductTierId = v
}

// GetProxyType returns the ProxyType field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetProxyType() string {
	if o == nil || IsNil(o.ProxyType) {
		var ret string
		return ret
	}
	return *o.ProxyType
}

// GetProxyTypeOk returns a tuple with the ProxyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetProxyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProxyType) {
		return nil, false
	}
	return o.ProxyType, true
}

// SetProxyType gets a reference to the given string and assigns it to the ProxyType field.
func (o *DescribeResourceResult) SetProxyType(v string) {
	o.ProxyType = &v
}

// GetResourceType returns the ResourceType field value
func (o *DescribeResourceResult) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *DescribeResourceResult) SetResourceType(v string) {
	o.ResourceType = v
}

// GetServiceId returns the ServiceId field value
func (o *DescribeResourceResult) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *DescribeResourceResult) SetServiceId(v string) {
	o.ServiceId = v
}

// GetTerraformConfigurations returns the TerraformConfigurations field value if set, zero value otherwise.
func (o *DescribeResourceResult) GetTerraformConfigurations() map[string]interface{} {
	if o == nil || IsNil(o.TerraformConfigurations) {
		var ret map[string]interface{}
		return ret
	}
	return o.TerraformConfigurations
}

// GetTerraformConfigurationsOk returns a tuple with the TerraformConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeResourceResult) GetTerraformConfigurationsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.TerraformConfigurations) {
		return map[string]interface{}{}, false
	}
	return o.TerraformConfigurations, true
}

// SetTerraformConfigurations gets a reference to the given map[string]interface{} and assigns it to the TerraformConfigurations field.
func (o *DescribeResourceResult) SetTerraformConfigurations(v map[string]interface{}) {
	o.TerraformConfigurations = v
}

func (o DescribeResourceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeResourceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionHooks) {
		toSerialize["actionHooks"] = o.ActionHooks
	}
	if !IsNil(o.AdditionalSecurityContext) {
		toSerialize["additionalSecurityContext"] = o.AdditionalSecurityContext
	}
	if !IsNil(o.BackupConfiguration) {
		toSerialize["backupConfiguration"] = o.BackupConfiguration
	}
	if !IsNil(o.BlobStorageConfiguration) {
		toSerialize["blobStorageConfiguration"] = o.BlobStorageConfiguration
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.CustomLabels) {
		toSerialize["customLabels"] = o.CustomLabels
	}
	if !IsNil(o.CustomSysCTLs) {
		toSerialize["customSysCTLs"] = o.CustomSysCTLs
	}
	if !IsNil(o.CustomULimits) {
		toSerialize["customULimits"] = o.CustomULimits
	}
	if !IsNil(o.Dependencies) {
		toSerialize["dependencies"] = o.Dependencies
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.EnvironmentVariables) {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	if !IsNil(o.FileSystemConfiguration) {
		toSerialize["fileSystemConfiguration"] = o.FileSystemConfiguration
	}
	if !IsNil(o.HelmChartConfiguration) {
		toSerialize["helmChartConfiguration"] = o.HelmChartConfiguration
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.ImageConfigId) {
		toSerialize["imageConfigId"] = o.ImageConfigId
	}
	if !IsNil(o.InfraConfigId) {
		toSerialize["infraConfigId"] = o.InfraConfigId
	}
	toSerialize["internal"] = o.Internal
	toSerialize["isDeprecated"] = o.IsDeprecated
	if !IsNil(o.JobConfig) {
		toSerialize["jobConfig"] = o.JobConfig
	}
	toSerialize["key"] = o.Key
	if !IsNil(o.KustomizeConfiguration) {
		toSerialize["kustomizeConfiguration"] = o.KustomizeConfiguration
	}
	if !IsNil(o.L4LoadBalancerConfiguration) {
		toSerialize["l4LoadBalancerConfiguration"] = o.L4LoadBalancerConfiguration
	}
	if !IsNil(o.L7LoadBalancerConfiguration) {
		toSerialize["l7LoadBalancerConfiguration"] = o.L7LoadBalancerConfiguration
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OperatorCRDConfiguration) {
		toSerialize["operatorCRDConfiguration"] = o.OperatorCRDConfiguration
	}
	toSerialize["productTierId"] = o.ProductTierId
	if !IsNil(o.ProxyType) {
		toSerialize["proxyType"] = o.ProxyType
	}
	toSerialize["resourceType"] = o.ResourceType
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.TerraformConfigurations) {
		toSerialize["terraformConfigurations"] = o.TerraformConfigurations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DescribeResourceResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"internal",
		"isDeprecated",
		"key",
		"name",
		"productTierId",
		"resourceType",
		"serviceId",
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

	varDescribeResourceResult := _DescribeResourceResult{}

	err = json.Unmarshal(data, &varDescribeResourceResult)

	if err != nil {
		return err
	}

	*o = DescribeResourceResult(varDescribeResourceResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actionHooks")
		delete(additionalProperties, "additionalSecurityContext")
		delete(additionalProperties, "backupConfiguration")
		delete(additionalProperties, "blobStorageConfiguration")
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "customLabels")
		delete(additionalProperties, "customSysCTLs")
		delete(additionalProperties, "customULimits")
		delete(additionalProperties, "dependencies")
		delete(additionalProperties, "description")
		delete(additionalProperties, "environmentVariables")
		delete(additionalProperties, "fileSystemConfiguration")
		delete(additionalProperties, "helmChartConfiguration")
		delete(additionalProperties, "id")
		delete(additionalProperties, "imageConfigId")
		delete(additionalProperties, "infraConfigId")
		delete(additionalProperties, "internal")
		delete(additionalProperties, "isDeprecated")
		delete(additionalProperties, "jobConfig")
		delete(additionalProperties, "key")
		delete(additionalProperties, "kustomizeConfiguration")
		delete(additionalProperties, "l4LoadBalancerConfiguration")
		delete(additionalProperties, "l7LoadBalancerConfiguration")
		delete(additionalProperties, "name")
		delete(additionalProperties, "operatorCRDConfiguration")
		delete(additionalProperties, "productTierId")
		delete(additionalProperties, "proxyType")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "serviceId")
		delete(additionalProperties, "terraformConfigurations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDescribeResourceResult struct {
	value *DescribeResourceResult
	isSet bool
}

func (v NullableDescribeResourceResult) Get() *DescribeResourceResult {
	return v.value
}

func (v *NullableDescribeResourceResult) Set(val *DescribeResourceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeResourceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeResourceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeResourceResult(val *DescribeResourceResult) *NullableDescribeResourceResult {
	return &NullableDescribeResourceResult{value: val, isSet: true}
}

func (v NullableDescribeResourceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeResourceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


