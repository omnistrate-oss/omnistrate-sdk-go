# PublicNetworkingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableClusterLoadBalancer** | Pointer to **bool** | Enable a single load balancer for all replicas | [optional] 
**EnableNodeLoadBalancer** | Pointer to **bool** | Create an external node load balancer per node rather than exposing the node ip directly | [optional] 

## Methods

### NewPublicNetworkingConfiguration

`func NewPublicNetworkingConfiguration() *PublicNetworkingConfiguration`

NewPublicNetworkingConfiguration instantiates a new PublicNetworkingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNetworkingConfigurationWithDefaults

`func NewPublicNetworkingConfigurationWithDefaults() *PublicNetworkingConfiguration`

NewPublicNetworkingConfigurationWithDefaults instantiates a new PublicNetworkingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableClusterLoadBalancer

`func (o *PublicNetworkingConfiguration) GetEnableClusterLoadBalancer() bool`

GetEnableClusterLoadBalancer returns the EnableClusterLoadBalancer field if non-nil, zero value otherwise.

### GetEnableClusterLoadBalancerOk

`func (o *PublicNetworkingConfiguration) GetEnableClusterLoadBalancerOk() (*bool, bool)`

GetEnableClusterLoadBalancerOk returns a tuple with the EnableClusterLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClusterLoadBalancer

`func (o *PublicNetworkingConfiguration) SetEnableClusterLoadBalancer(v bool)`

SetEnableClusterLoadBalancer sets EnableClusterLoadBalancer field to given value.

### HasEnableClusterLoadBalancer

`func (o *PublicNetworkingConfiguration) HasEnableClusterLoadBalancer() bool`

HasEnableClusterLoadBalancer returns a boolean if a field has been set.

### GetEnableNodeLoadBalancer

`func (o *PublicNetworkingConfiguration) GetEnableNodeLoadBalancer() bool`

GetEnableNodeLoadBalancer returns the EnableNodeLoadBalancer field if non-nil, zero value otherwise.

### GetEnableNodeLoadBalancerOk

`func (o *PublicNetworkingConfiguration) GetEnableNodeLoadBalancerOk() (*bool, bool)`

GetEnableNodeLoadBalancerOk returns a tuple with the EnableNodeLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNodeLoadBalancer

`func (o *PublicNetworkingConfiguration) SetEnableNodeLoadBalancer(v bool)`

SetEnableNodeLoadBalancer sets EnableNodeLoadBalancer field to given value.

### HasEnableNodeLoadBalancer

`func (o *PublicNetworkingConfiguration) HasEnableNodeLoadBalancer() bool`

HasEnableNodeLoadBalancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


