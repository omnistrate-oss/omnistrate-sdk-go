# BackendPortConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedResourceIDs** | **[]string** | The IDs of the resource associated with the backend port | 
**BackendPort** | **int64** | The port to forward traffic to | 
**IngressPort** | **int64** | The ingress port to configure on the load balancer | 

## Methods

### NewBackendPortConfiguration

`func NewBackendPortConfiguration(associatedResourceIDs []string, backendPort int64, ingressPort int64, ) *BackendPortConfiguration`

NewBackendPortConfiguration instantiates a new BackendPortConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendPortConfigurationWithDefaults

`func NewBackendPortConfigurationWithDefaults() *BackendPortConfiguration`

NewBackendPortConfigurationWithDefaults instantiates a new BackendPortConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedResourceIDs

`func (o *BackendPortConfiguration) GetAssociatedResourceIDs() []string`

GetAssociatedResourceIDs returns the AssociatedResourceIDs field if non-nil, zero value otherwise.

### GetAssociatedResourceIDsOk

`func (o *BackendPortConfiguration) GetAssociatedResourceIDsOk() (*[]string, bool)`

GetAssociatedResourceIDsOk returns a tuple with the AssociatedResourceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedResourceIDs

`func (o *BackendPortConfiguration) SetAssociatedResourceIDs(v []string)`

SetAssociatedResourceIDs sets AssociatedResourceIDs field to given value.


### GetBackendPort

`func (o *BackendPortConfiguration) GetBackendPort() int64`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *BackendPortConfiguration) GetBackendPortOk() (*int64, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *BackendPortConfiguration) SetBackendPort(v int64)`

SetBackendPort sets BackendPort field to given value.


### GetIngressPort

`func (o *BackendPortConfiguration) GetIngressPort() int64`

GetIngressPort returns the IngressPort field if non-nil, zero value otherwise.

### GetIngressPortOk

`func (o *BackendPortConfiguration) GetIngressPortOk() (*int64, bool)`

GetIngressPortOk returns a tuple with the IngressPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPort

`func (o *BackendPortConfiguration) SetIngressPort(v int64)`

SetIngressPort sets IngressPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


