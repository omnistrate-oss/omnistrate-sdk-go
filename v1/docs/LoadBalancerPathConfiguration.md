# LoadBalancerPathConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedKubernetesServiceName** | Pointer to **string** | Override the default target Kubernetes service name with this value | [optional] 
**AssociatedResourceID** | **string** | The ID of the resource associated with the path | 
**Path** | **string** | The REST API path backed by the load balancer | 
**Port** | **int64** | The port to forward traffic to | 

## Methods

### NewLoadBalancerPathConfiguration

`func NewLoadBalancerPathConfiguration(associatedResourceID string, path string, port int64, ) *LoadBalancerPathConfiguration`

NewLoadBalancerPathConfiguration instantiates a new LoadBalancerPathConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPathConfigurationWithDefaults

`func NewLoadBalancerPathConfigurationWithDefaults() *LoadBalancerPathConfiguration`

NewLoadBalancerPathConfigurationWithDefaults instantiates a new LoadBalancerPathConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedKubernetesServiceName

`func (o *LoadBalancerPathConfiguration) GetAssociatedKubernetesServiceName() string`

GetAssociatedKubernetesServiceName returns the AssociatedKubernetesServiceName field if non-nil, zero value otherwise.

### GetAssociatedKubernetesServiceNameOk

`func (o *LoadBalancerPathConfiguration) GetAssociatedKubernetesServiceNameOk() (*string, bool)`

GetAssociatedKubernetesServiceNameOk returns a tuple with the AssociatedKubernetesServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedKubernetesServiceName

`func (o *LoadBalancerPathConfiguration) SetAssociatedKubernetesServiceName(v string)`

SetAssociatedKubernetesServiceName sets AssociatedKubernetesServiceName field to given value.

### HasAssociatedKubernetesServiceName

`func (o *LoadBalancerPathConfiguration) HasAssociatedKubernetesServiceName() bool`

HasAssociatedKubernetesServiceName returns a boolean if a field has been set.

### GetAssociatedResourceID

`func (o *LoadBalancerPathConfiguration) GetAssociatedResourceID() string`

GetAssociatedResourceID returns the AssociatedResourceID field if non-nil, zero value otherwise.

### GetAssociatedResourceIDOk

`func (o *LoadBalancerPathConfiguration) GetAssociatedResourceIDOk() (*string, bool)`

GetAssociatedResourceIDOk returns a tuple with the AssociatedResourceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResourceID

`func (o *LoadBalancerPathConfiguration) SetAssociatedResourceID(v string)`

SetAssociatedResourceID sets AssociatedResourceID field to given value.


### GetPath

`func (o *LoadBalancerPathConfiguration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoadBalancerPathConfiguration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoadBalancerPathConfiguration) SetPath(v string)`

SetPath sets Path field to given value.


### GetPort

`func (o *LoadBalancerPathConfiguration) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerPathConfiguration) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerPathConfiguration) SetPort(v int64)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


