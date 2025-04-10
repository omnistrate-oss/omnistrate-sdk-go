# ResourceNetworkTopologyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEndpoints** | Pointer to [**map[string]ClusterEndpoint**](ClusterEndpoint.md) | The additional endpoints for this resource | [optional] 
**AllowedIPRanges** | **[]string** | The allowed IP ranges for this resource | 
**ClusterEndpoint** | **string** | The primary cluster endpoint that load-balances across replicas of this resource | 
**ClusterPorts** | Pointer to **[]int64** | The ports that the whole cluster exposes | [optional] 
**CustomDNSEndpoint** | Pointer to [**CustomDNSEndpoint**](CustomDNSEndpoint.md) |  | [optional] 
**HasCompute** | **bool** | Whether this resource has associated compute | 
**IsJob** | Pointer to **bool** | Whether this resource is a job | [optional] [default to false]
**JobMetrics** | Pointer to [**[]JobMetric**](JobMetric.md) | The job metrics for this resource (if it&#39;s a job) | [optional] 
**Main** | **bool** | Whether this is the main resource | 
**NetworkingType** | **string** | The networking type for this resource | 
**Nodes** | Pointer to [**[]NodeNetworkTopologyResult**](NodeNetworkTopologyResult.md) | The nodes that this resource is deployed on | [optional] 
**PrivateNetworkCIDR** | Pointer to **string** | The private network CIDR for this resource | [optional] 
**PrivateNetworkID** | Pointer to **string** | ID of a Network | [optional] 
**ProxyEndpoint** | Pointer to [**ProxyEndpoint**](ProxyEndpoint.md) |  | [optional] 
**PubliclyAccessible** | **bool** | Whether this resource is publicly accessible | 
**RecentDeploymentFailure** | Pointer to [**RecentDeploymentFailureStatus**](RecentDeploymentFailureStatus.md) |  | [optional] 
**ResourceInstanceMetadata** | Pointer to **interface{}** | The resource instance metadata | [optional] 
**ResourceKey** | **string** | The key of the resource | 
**ResourceName** | **string** | The name of the resource | 
**ResourceType** | Pointer to **string** | The type of the resource | [optional] 

## Methods

### NewResourceNetworkTopologyResult

`func NewResourceNetworkTopologyResult(allowedIPRanges []string, clusterEndpoint string, hasCompute bool, main bool, networkingType string, publiclyAccessible bool, resourceKey string, resourceName string, ) *ResourceNetworkTopologyResult`

NewResourceNetworkTopologyResult instantiates a new ResourceNetworkTopologyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceNetworkTopologyResultWithDefaults

`func NewResourceNetworkTopologyResultWithDefaults() *ResourceNetworkTopologyResult`

NewResourceNetworkTopologyResultWithDefaults instantiates a new ResourceNetworkTopologyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEndpoints

`func (o *ResourceNetworkTopologyResult) GetAdditionalEndpoints() map[string]ClusterEndpoint`

GetAdditionalEndpoints returns the AdditionalEndpoints field if non-nil, zero value otherwise.

### GetAdditionalEndpointsOk

`func (o *ResourceNetworkTopologyResult) GetAdditionalEndpointsOk() (*map[string]ClusterEndpoint, bool)`

GetAdditionalEndpointsOk returns a tuple with the AdditionalEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEndpoints

`func (o *ResourceNetworkTopologyResult) SetAdditionalEndpoints(v map[string]ClusterEndpoint)`

SetAdditionalEndpoints sets AdditionalEndpoints field to given value.

### HasAdditionalEndpoints

`func (o *ResourceNetworkTopologyResult) HasAdditionalEndpoints() bool`

HasAdditionalEndpoints returns a boolean if a field has been set.

### GetAllowedIPRanges

`func (o *ResourceNetworkTopologyResult) GetAllowedIPRanges() []string`

GetAllowedIPRanges returns the AllowedIPRanges field if non-nil, zero value otherwise.

### GetAllowedIPRangesOk

`func (o *ResourceNetworkTopologyResult) GetAllowedIPRangesOk() (*[]string, bool)`

GetAllowedIPRangesOk returns a tuple with the AllowedIPRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPRanges

`func (o *ResourceNetworkTopologyResult) SetAllowedIPRanges(v []string)`

SetAllowedIPRanges sets AllowedIPRanges field to given value.


### GetClusterEndpoint

`func (o *ResourceNetworkTopologyResult) GetClusterEndpoint() string`

GetClusterEndpoint returns the ClusterEndpoint field if non-nil, zero value otherwise.

### GetClusterEndpointOk

`func (o *ResourceNetworkTopologyResult) GetClusterEndpointOk() (*string, bool)`

GetClusterEndpointOk returns a tuple with the ClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEndpoint

`func (o *ResourceNetworkTopologyResult) SetClusterEndpoint(v string)`

SetClusterEndpoint sets ClusterEndpoint field to given value.


### GetClusterPorts

`func (o *ResourceNetworkTopologyResult) GetClusterPorts() []int64`

GetClusterPorts returns the ClusterPorts field if non-nil, zero value otherwise.

### GetClusterPortsOk

`func (o *ResourceNetworkTopologyResult) GetClusterPortsOk() (*[]int64, bool)`

GetClusterPortsOk returns a tuple with the ClusterPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPorts

`func (o *ResourceNetworkTopologyResult) SetClusterPorts(v []int64)`

SetClusterPorts sets ClusterPorts field to given value.

### HasClusterPorts

`func (o *ResourceNetworkTopologyResult) HasClusterPorts() bool`

HasClusterPorts returns a boolean if a field has been set.

### GetCustomDNSEndpoint

`func (o *ResourceNetworkTopologyResult) GetCustomDNSEndpoint() CustomDNSEndpoint`

GetCustomDNSEndpoint returns the CustomDNSEndpoint field if non-nil, zero value otherwise.

### GetCustomDNSEndpointOk

`func (o *ResourceNetworkTopologyResult) GetCustomDNSEndpointOk() (*CustomDNSEndpoint, bool)`

GetCustomDNSEndpointOk returns a tuple with the CustomDNSEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDNSEndpoint

`func (o *ResourceNetworkTopologyResult) SetCustomDNSEndpoint(v CustomDNSEndpoint)`

SetCustomDNSEndpoint sets CustomDNSEndpoint field to given value.

### HasCustomDNSEndpoint

`func (o *ResourceNetworkTopologyResult) HasCustomDNSEndpoint() bool`

HasCustomDNSEndpoint returns a boolean if a field has been set.

### GetHasCompute

`func (o *ResourceNetworkTopologyResult) GetHasCompute() bool`

GetHasCompute returns the HasCompute field if non-nil, zero value otherwise.

### GetHasComputeOk

`func (o *ResourceNetworkTopologyResult) GetHasComputeOk() (*bool, bool)`

GetHasComputeOk returns a tuple with the HasCompute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompute

`func (o *ResourceNetworkTopologyResult) SetHasCompute(v bool)`

SetHasCompute sets HasCompute field to given value.


### GetIsJob

`func (o *ResourceNetworkTopologyResult) GetIsJob() bool`

GetIsJob returns the IsJob field if non-nil, zero value otherwise.

### GetIsJobOk

`func (o *ResourceNetworkTopologyResult) GetIsJobOk() (*bool, bool)`

GetIsJobOk returns a tuple with the IsJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJob

`func (o *ResourceNetworkTopologyResult) SetIsJob(v bool)`

SetIsJob sets IsJob field to given value.

### HasIsJob

`func (o *ResourceNetworkTopologyResult) HasIsJob() bool`

HasIsJob returns a boolean if a field has been set.

### GetJobMetrics

`func (o *ResourceNetworkTopologyResult) GetJobMetrics() []JobMetric`

GetJobMetrics returns the JobMetrics field if non-nil, zero value otherwise.

### GetJobMetricsOk

`func (o *ResourceNetworkTopologyResult) GetJobMetricsOk() (*[]JobMetric, bool)`

GetJobMetricsOk returns a tuple with the JobMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobMetrics

`func (o *ResourceNetworkTopologyResult) SetJobMetrics(v []JobMetric)`

SetJobMetrics sets JobMetrics field to given value.

### HasJobMetrics

`func (o *ResourceNetworkTopologyResult) HasJobMetrics() bool`

HasJobMetrics returns a boolean if a field has been set.

### GetMain

`func (o *ResourceNetworkTopologyResult) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *ResourceNetworkTopologyResult) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *ResourceNetworkTopologyResult) SetMain(v bool)`

SetMain sets Main field to given value.


### GetNetworkingType

`func (o *ResourceNetworkTopologyResult) GetNetworkingType() string`

GetNetworkingType returns the NetworkingType field if non-nil, zero value otherwise.

### GetNetworkingTypeOk

`func (o *ResourceNetworkTopologyResult) GetNetworkingTypeOk() (*string, bool)`

GetNetworkingTypeOk returns a tuple with the NetworkingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingType

`func (o *ResourceNetworkTopologyResult) SetNetworkingType(v string)`

SetNetworkingType sets NetworkingType field to given value.


### GetNodes

`func (o *ResourceNetworkTopologyResult) GetNodes() []NodeNetworkTopologyResult`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ResourceNetworkTopologyResult) GetNodesOk() (*[]NodeNetworkTopologyResult, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ResourceNetworkTopologyResult) SetNodes(v []NodeNetworkTopologyResult)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ResourceNetworkTopologyResult) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetPrivateNetworkCIDR

`func (o *ResourceNetworkTopologyResult) GetPrivateNetworkCIDR() string`

GetPrivateNetworkCIDR returns the PrivateNetworkCIDR field if non-nil, zero value otherwise.

### GetPrivateNetworkCIDROk

`func (o *ResourceNetworkTopologyResult) GetPrivateNetworkCIDROk() (*string, bool)`

GetPrivateNetworkCIDROk returns a tuple with the PrivateNetworkCIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkCIDR

`func (o *ResourceNetworkTopologyResult) SetPrivateNetworkCIDR(v string)`

SetPrivateNetworkCIDR sets PrivateNetworkCIDR field to given value.

### HasPrivateNetworkCIDR

`func (o *ResourceNetworkTopologyResult) HasPrivateNetworkCIDR() bool`

HasPrivateNetworkCIDR returns a boolean if a field has been set.

### GetPrivateNetworkID

`func (o *ResourceNetworkTopologyResult) GetPrivateNetworkID() string`

GetPrivateNetworkID returns the PrivateNetworkID field if non-nil, zero value otherwise.

### GetPrivateNetworkIDOk

`func (o *ResourceNetworkTopologyResult) GetPrivateNetworkIDOk() (*string, bool)`

GetPrivateNetworkIDOk returns a tuple with the PrivateNetworkID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkID

`func (o *ResourceNetworkTopologyResult) SetPrivateNetworkID(v string)`

SetPrivateNetworkID sets PrivateNetworkID field to given value.

### HasPrivateNetworkID

`func (o *ResourceNetworkTopologyResult) HasPrivateNetworkID() bool`

HasPrivateNetworkID returns a boolean if a field has been set.

### GetProxyEndpoint

`func (o *ResourceNetworkTopologyResult) GetProxyEndpoint() ProxyEndpoint`

GetProxyEndpoint returns the ProxyEndpoint field if non-nil, zero value otherwise.

### GetProxyEndpointOk

`func (o *ResourceNetworkTopologyResult) GetProxyEndpointOk() (*ProxyEndpoint, bool)`

GetProxyEndpointOk returns a tuple with the ProxyEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyEndpoint

`func (o *ResourceNetworkTopologyResult) SetProxyEndpoint(v ProxyEndpoint)`

SetProxyEndpoint sets ProxyEndpoint field to given value.

### HasProxyEndpoint

`func (o *ResourceNetworkTopologyResult) HasProxyEndpoint() bool`

HasProxyEndpoint returns a boolean if a field has been set.

### GetPubliclyAccessible

`func (o *ResourceNetworkTopologyResult) GetPubliclyAccessible() bool`

GetPubliclyAccessible returns the PubliclyAccessible field if non-nil, zero value otherwise.

### GetPubliclyAccessibleOk

`func (o *ResourceNetworkTopologyResult) GetPubliclyAccessibleOk() (*bool, bool)`

GetPubliclyAccessibleOk returns a tuple with the PubliclyAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubliclyAccessible

`func (o *ResourceNetworkTopologyResult) SetPubliclyAccessible(v bool)`

SetPubliclyAccessible sets PubliclyAccessible field to given value.


### GetRecentDeploymentFailure

`func (o *ResourceNetworkTopologyResult) GetRecentDeploymentFailure() RecentDeploymentFailureStatus`

GetRecentDeploymentFailure returns the RecentDeploymentFailure field if non-nil, zero value otherwise.

### GetRecentDeploymentFailureOk

`func (o *ResourceNetworkTopologyResult) GetRecentDeploymentFailureOk() (*RecentDeploymentFailureStatus, bool)`

GetRecentDeploymentFailureOk returns a tuple with the RecentDeploymentFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentDeploymentFailure

`func (o *ResourceNetworkTopologyResult) SetRecentDeploymentFailure(v RecentDeploymentFailureStatus)`

SetRecentDeploymentFailure sets RecentDeploymentFailure field to given value.

### HasRecentDeploymentFailure

`func (o *ResourceNetworkTopologyResult) HasRecentDeploymentFailure() bool`

HasRecentDeploymentFailure returns a boolean if a field has been set.

### GetResourceInstanceMetadata

`func (o *ResourceNetworkTopologyResult) GetResourceInstanceMetadata() interface{}`

GetResourceInstanceMetadata returns the ResourceInstanceMetadata field if non-nil, zero value otherwise.

### GetResourceInstanceMetadataOk

`func (o *ResourceNetworkTopologyResult) GetResourceInstanceMetadataOk() (*interface{}, bool)`

GetResourceInstanceMetadataOk returns a tuple with the ResourceInstanceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstanceMetadata

`func (o *ResourceNetworkTopologyResult) SetResourceInstanceMetadata(v interface{})`

SetResourceInstanceMetadata sets ResourceInstanceMetadata field to given value.

### HasResourceInstanceMetadata

`func (o *ResourceNetworkTopologyResult) HasResourceInstanceMetadata() bool`

HasResourceInstanceMetadata returns a boolean if a field has been set.

### SetResourceInstanceMetadataNil

`func (o *ResourceNetworkTopologyResult) SetResourceInstanceMetadataNil(b bool)`

 SetResourceInstanceMetadataNil sets the value for ResourceInstanceMetadata to be an explicit nil

### UnsetResourceInstanceMetadata
`func (o *ResourceNetworkTopologyResult) UnsetResourceInstanceMetadata()`

UnsetResourceInstanceMetadata ensures that no value is present for ResourceInstanceMetadata, not even an explicit nil
### GetResourceKey

`func (o *ResourceNetworkTopologyResult) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *ResourceNetworkTopologyResult) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *ResourceNetworkTopologyResult) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceName

`func (o *ResourceNetworkTopologyResult) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceNetworkTopologyResult) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceNetworkTopologyResult) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *ResourceNetworkTopologyResult) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ResourceNetworkTopologyResult) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ResourceNetworkTopologyResult) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ResourceNetworkTopologyResult) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


