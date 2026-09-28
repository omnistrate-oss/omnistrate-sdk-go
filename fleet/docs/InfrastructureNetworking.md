# InfrastructureNetworking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**Ingresses** | Pointer to [**[]InfrastructureIngress**](InfrastructureIngress.md) |  | [optional] 
**IngressesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**LoadBalancers** | Pointer to [**[]InfrastructureLoadBalancer**](InfrastructureLoadBalancer.md) |  | [optional] 
**LoadBalancersTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**NetworkPolicies** | Pointer to [**[]InfrastructureNetworkPolicy**](InfrastructureNetworkPolicy.md) |  | [optional] 
**NetworkPoliciesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Services** | Pointer to [**[]InfrastructureService**](InfrastructureService.md) |  | [optional] 
**ServicesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**SortOrder** | Pointer to **string** |  | [optional] 
**StringsTruncated** | Pointer to **bool** |  | [optional] 

## Methods

### NewInfrastructureNetworking

`func NewInfrastructureNetworking() *InfrastructureNetworking`

NewInfrastructureNetworking instantiates a new InfrastructureNetworking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureNetworkingWithDefaults

`func NewInfrastructureNetworkingWithDefaults() *InfrastructureNetworking`

NewInfrastructureNetworkingWithDefaults instantiates a new InfrastructureNetworking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *InfrastructureNetworking) GetAvailability() InfrastructureAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *InfrastructureNetworking) GetAvailabilityOk() (*InfrastructureAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *InfrastructureNetworking) SetAvailability(v InfrastructureAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *InfrastructureNetworking) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetIngresses

`func (o *InfrastructureNetworking) GetIngresses() []InfrastructureIngress`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *InfrastructureNetworking) GetIngressesOk() (*[]InfrastructureIngress, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *InfrastructureNetworking) SetIngresses(v []InfrastructureIngress)`

SetIngresses sets Ingresses field to given value.

### HasIngresses

`func (o *InfrastructureNetworking) HasIngresses() bool`

HasIngresses returns a boolean if a field has been set.

### GetIngressesTruncation

`func (o *InfrastructureNetworking) GetIngressesTruncation() CheckpointTruncation`

GetIngressesTruncation returns the IngressesTruncation field if non-nil, zero value otherwise.

### GetIngressesTruncationOk

`func (o *InfrastructureNetworking) GetIngressesTruncationOk() (*CheckpointTruncation, bool)`

GetIngressesTruncationOk returns a tuple with the IngressesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressesTruncation

`func (o *InfrastructureNetworking) SetIngressesTruncation(v CheckpointTruncation)`

SetIngressesTruncation sets IngressesTruncation field to given value.

### HasIngressesTruncation

`func (o *InfrastructureNetworking) HasIngressesTruncation() bool`

HasIngressesTruncation returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *InfrastructureNetworking) GetLoadBalancers() []InfrastructureLoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *InfrastructureNetworking) GetLoadBalancersOk() (*[]InfrastructureLoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *InfrastructureNetworking) SetLoadBalancers(v []InfrastructureLoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *InfrastructureNetworking) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetLoadBalancersTruncation

`func (o *InfrastructureNetworking) GetLoadBalancersTruncation() CheckpointTruncation`

GetLoadBalancersTruncation returns the LoadBalancersTruncation field if non-nil, zero value otherwise.

### GetLoadBalancersTruncationOk

`func (o *InfrastructureNetworking) GetLoadBalancersTruncationOk() (*CheckpointTruncation, bool)`

GetLoadBalancersTruncationOk returns a tuple with the LoadBalancersTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancersTruncation

`func (o *InfrastructureNetworking) SetLoadBalancersTruncation(v CheckpointTruncation)`

SetLoadBalancersTruncation sets LoadBalancersTruncation field to given value.

### HasLoadBalancersTruncation

`func (o *InfrastructureNetworking) HasLoadBalancersTruncation() bool`

HasLoadBalancersTruncation returns a boolean if a field has been set.

### GetNetworkPolicies

`func (o *InfrastructureNetworking) GetNetworkPolicies() []InfrastructureNetworkPolicy`

GetNetworkPolicies returns the NetworkPolicies field if non-nil, zero value otherwise.

### GetNetworkPoliciesOk

`func (o *InfrastructureNetworking) GetNetworkPoliciesOk() (*[]InfrastructureNetworkPolicy, bool)`

GetNetworkPoliciesOk returns a tuple with the NetworkPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicies

`func (o *InfrastructureNetworking) SetNetworkPolicies(v []InfrastructureNetworkPolicy)`

SetNetworkPolicies sets NetworkPolicies field to given value.

### HasNetworkPolicies

`func (o *InfrastructureNetworking) HasNetworkPolicies() bool`

HasNetworkPolicies returns a boolean if a field has been set.

### GetNetworkPoliciesTruncation

`func (o *InfrastructureNetworking) GetNetworkPoliciesTruncation() CheckpointTruncation`

GetNetworkPoliciesTruncation returns the NetworkPoliciesTruncation field if non-nil, zero value otherwise.

### GetNetworkPoliciesTruncationOk

`func (o *InfrastructureNetworking) GetNetworkPoliciesTruncationOk() (*CheckpointTruncation, bool)`

GetNetworkPoliciesTruncationOk returns a tuple with the NetworkPoliciesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoliciesTruncation

`func (o *InfrastructureNetworking) SetNetworkPoliciesTruncation(v CheckpointTruncation)`

SetNetworkPoliciesTruncation sets NetworkPoliciesTruncation field to given value.

### HasNetworkPoliciesTruncation

`func (o *InfrastructureNetworking) HasNetworkPoliciesTruncation() bool`

HasNetworkPoliciesTruncation returns a boolean if a field has been set.

### GetServices

`func (o *InfrastructureNetworking) GetServices() []InfrastructureService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InfrastructureNetworking) GetServicesOk() (*[]InfrastructureService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InfrastructureNetworking) SetServices(v []InfrastructureService)`

SetServices sets Services field to given value.

### HasServices

`func (o *InfrastructureNetworking) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetServicesTruncation

`func (o *InfrastructureNetworking) GetServicesTruncation() CheckpointTruncation`

GetServicesTruncation returns the ServicesTruncation field if non-nil, zero value otherwise.

### GetServicesTruncationOk

`func (o *InfrastructureNetworking) GetServicesTruncationOk() (*CheckpointTruncation, bool)`

GetServicesTruncationOk returns a tuple with the ServicesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicesTruncation

`func (o *InfrastructureNetworking) SetServicesTruncation(v CheckpointTruncation)`

SetServicesTruncation sets ServicesTruncation field to given value.

### HasServicesTruncation

`func (o *InfrastructureNetworking) HasServicesTruncation() bool`

HasServicesTruncation returns a boolean if a field has been set.

### GetSortOrder

`func (o *InfrastructureNetworking) GetSortOrder() string`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InfrastructureNetworking) GetSortOrderOk() (*string, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InfrastructureNetworking) SetSortOrder(v string)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InfrastructureNetworking) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetStringsTruncated

`func (o *InfrastructureNetworking) GetStringsTruncated() bool`

GetStringsTruncated returns the StringsTruncated field if non-nil, zero value otherwise.

### GetStringsTruncatedOk

`func (o *InfrastructureNetworking) GetStringsTruncatedOk() (*bool, bool)`

GetStringsTruncatedOk returns a tuple with the StringsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringsTruncated

`func (o *InfrastructureNetworking) SetStringsTruncated(v bool)`

SetStringsTruncated sets StringsTruncated field to given value.

### HasStringsTruncated

`func (o *InfrastructureNetworking) HasStringsTruncated() bool`

HasStringsTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


