# L7LoadBalancerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**[]LoadBalancerPathConfiguration**](LoadBalancerPathConfiguration.md) | The paths to configure on the load balancer | [optional] 

## Methods

### NewL7LoadBalancerConfiguration

`func NewL7LoadBalancerConfiguration() *L7LoadBalancerConfiguration`

NewL7LoadBalancerConfiguration instantiates a new L7LoadBalancerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL7LoadBalancerConfigurationWithDefaults

`func NewL7LoadBalancerConfigurationWithDefaults() *L7LoadBalancerConfiguration`

NewL7LoadBalancerConfigurationWithDefaults instantiates a new L7LoadBalancerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *L7LoadBalancerConfiguration) GetPaths() []LoadBalancerPathConfiguration`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *L7LoadBalancerConfiguration) GetPathsOk() (*[]LoadBalancerPathConfiguration, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *L7LoadBalancerConfiguration) SetPaths(v []LoadBalancerPathConfiguration)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *L7LoadBalancerConfiguration) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


