# L4LoadBalancerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngressPortConfiguration** | Pointer to [**[]BackendPortConfiguration**](BackendPortConfiguration.md) | The ingress port configuration to configure on the load balancer | [optional] 

## Methods

### NewL4LoadBalancerConfiguration

`func NewL4LoadBalancerConfiguration() *L4LoadBalancerConfiguration`

NewL4LoadBalancerConfiguration instantiates a new L4LoadBalancerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL4LoadBalancerConfigurationWithDefaults

`func NewL4LoadBalancerConfigurationWithDefaults() *L4LoadBalancerConfiguration`

NewL4LoadBalancerConfigurationWithDefaults instantiates a new L4LoadBalancerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngressPortConfiguration

`func (o *L4LoadBalancerConfiguration) GetIngressPortConfiguration() []BackendPortConfiguration`

GetIngressPortConfiguration returns the IngressPortConfiguration field if non-nil, zero value otherwise.

### GetIngressPortConfigurationOk

`func (o *L4LoadBalancerConfiguration) GetIngressPortConfigurationOk() (*[]BackendPortConfiguration, bool)`

GetIngressPortConfigurationOk returns a tuple with the IngressPortConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPortConfiguration

`func (o *L4LoadBalancerConfiguration) SetIngressPortConfiguration(v []BackendPortConfiguration)`

SetIngressPortConfiguration sets IngressPortConfiguration field to given value.

### HasIngressPortConfiguration

`func (o *L4LoadBalancerConfiguration) HasIngressPortConfiguration() bool`

HasIngressPortConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


