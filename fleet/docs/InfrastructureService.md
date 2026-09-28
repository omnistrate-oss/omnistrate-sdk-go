# InfrastructureService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedBy** | Pointer to [**InfrastructureAppliedBy**](InfrastructureAppliedBy.md) |  | [optional] 
**ClusterIP** | Pointer to **string** |  | [optional] 
**Controller** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**EndpointCount** | Pointer to **int64** |  | [optional] 
**Endpoints** | Pointer to [**[]InfrastructureEndpoint**](InfrastructureEndpoint.md) |  | [optional] 
**EndpointsAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**EndpointsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**ExternalIPs** | Pointer to **[]string** |  | [optional] 
**ExternalIPsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**MatchedPods** | Pointer to [**[]InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**MatchedPodsAvailability** | Pointer to [**InfrastructureAvailability**](InfrastructureAvailability.md) |  | [optional] 
**MatchedPodsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Ports** | Pointer to [**[]InfrastructurePort**](InfrastructurePort.md) |  | [optional] 
**PortsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Selector** | Pointer to [**[]InfrastructureKeyValue**](InfrastructureKeyValue.md) |  | [optional] 
**SelectorTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**SessionAffinity** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureService

`func NewInfrastructureService() *InfrastructureService`

NewInfrastructureService instantiates a new InfrastructureService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureServiceWithDefaults

`func NewInfrastructureServiceWithDefaults() *InfrastructureService`

NewInfrastructureServiceWithDefaults instantiates a new InfrastructureService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedBy

`func (o *InfrastructureService) GetAppliedBy() InfrastructureAppliedBy`

GetAppliedBy returns the AppliedBy field if non-nil, zero value otherwise.

### GetAppliedByOk

`func (o *InfrastructureService) GetAppliedByOk() (*InfrastructureAppliedBy, bool)`

GetAppliedByOk returns a tuple with the AppliedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedBy

`func (o *InfrastructureService) SetAppliedBy(v InfrastructureAppliedBy)`

SetAppliedBy sets AppliedBy field to given value.

### HasAppliedBy

`func (o *InfrastructureService) HasAppliedBy() bool`

HasAppliedBy returns a boolean if a field has been set.

### GetClusterIP

`func (o *InfrastructureService) GetClusterIP() string`

GetClusterIP returns the ClusterIP field if non-nil, zero value otherwise.

### GetClusterIPOk

`func (o *InfrastructureService) GetClusterIPOk() (*string, bool)`

GetClusterIPOk returns a tuple with the ClusterIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIP

`func (o *InfrastructureService) SetClusterIP(v string)`

SetClusterIP sets ClusterIP field to given value.

### HasClusterIP

`func (o *InfrastructureService) HasClusterIP() bool`

HasClusterIP returns a boolean if a field has been set.

### GetController

`func (o *InfrastructureService) GetController() InfrastructureObjectReference`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *InfrastructureService) GetControllerOk() (*InfrastructureObjectReference, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *InfrastructureService) SetController(v InfrastructureObjectReference)`

SetController sets Controller field to given value.

### HasController

`func (o *InfrastructureService) HasController() bool`

HasController returns a boolean if a field has been set.

### GetEndpointCount

`func (o *InfrastructureService) GetEndpointCount() int64`

GetEndpointCount returns the EndpointCount field if non-nil, zero value otherwise.

### GetEndpointCountOk

`func (o *InfrastructureService) GetEndpointCountOk() (*int64, bool)`

GetEndpointCountOk returns a tuple with the EndpointCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCount

`func (o *InfrastructureService) SetEndpointCount(v int64)`

SetEndpointCount sets EndpointCount field to given value.

### HasEndpointCount

`func (o *InfrastructureService) HasEndpointCount() bool`

HasEndpointCount returns a boolean if a field has been set.

### GetEndpoints

`func (o *InfrastructureService) GetEndpoints() []InfrastructureEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *InfrastructureService) GetEndpointsOk() (*[]InfrastructureEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *InfrastructureService) SetEndpoints(v []InfrastructureEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *InfrastructureService) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetEndpointsAvailability

`func (o *InfrastructureService) GetEndpointsAvailability() InfrastructureAvailability`

GetEndpointsAvailability returns the EndpointsAvailability field if non-nil, zero value otherwise.

### GetEndpointsAvailabilityOk

`func (o *InfrastructureService) GetEndpointsAvailabilityOk() (*InfrastructureAvailability, bool)`

GetEndpointsAvailabilityOk returns a tuple with the EndpointsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsAvailability

`func (o *InfrastructureService) SetEndpointsAvailability(v InfrastructureAvailability)`

SetEndpointsAvailability sets EndpointsAvailability field to given value.

### HasEndpointsAvailability

`func (o *InfrastructureService) HasEndpointsAvailability() bool`

HasEndpointsAvailability returns a boolean if a field has been set.

### GetEndpointsTruncation

`func (o *InfrastructureService) GetEndpointsTruncation() CheckpointTruncation`

GetEndpointsTruncation returns the EndpointsTruncation field if non-nil, zero value otherwise.

### GetEndpointsTruncationOk

`func (o *InfrastructureService) GetEndpointsTruncationOk() (*CheckpointTruncation, bool)`

GetEndpointsTruncationOk returns a tuple with the EndpointsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointsTruncation

`func (o *InfrastructureService) SetEndpointsTruncation(v CheckpointTruncation)`

SetEndpointsTruncation sets EndpointsTruncation field to given value.

### HasEndpointsTruncation

`func (o *InfrastructureService) HasEndpointsTruncation() bool`

HasEndpointsTruncation returns a boolean if a field has been set.

### GetExternalIPs

`func (o *InfrastructureService) GetExternalIPs() []string`

GetExternalIPs returns the ExternalIPs field if non-nil, zero value otherwise.

### GetExternalIPsOk

`func (o *InfrastructureService) GetExternalIPsOk() (*[]string, bool)`

GetExternalIPsOk returns a tuple with the ExternalIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIPs

`func (o *InfrastructureService) SetExternalIPs(v []string)`

SetExternalIPs sets ExternalIPs field to given value.

### HasExternalIPs

`func (o *InfrastructureService) HasExternalIPs() bool`

HasExternalIPs returns a boolean if a field has been set.

### GetExternalIPsTruncation

`func (o *InfrastructureService) GetExternalIPsTruncation() CheckpointTruncation`

GetExternalIPsTruncation returns the ExternalIPsTruncation field if non-nil, zero value otherwise.

### GetExternalIPsTruncationOk

`func (o *InfrastructureService) GetExternalIPsTruncationOk() (*CheckpointTruncation, bool)`

GetExternalIPsTruncationOk returns a tuple with the ExternalIPsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIPsTruncation

`func (o *InfrastructureService) SetExternalIPsTruncation(v CheckpointTruncation)`

SetExternalIPsTruncation sets ExternalIPsTruncation field to given value.

### HasExternalIPsTruncation

`func (o *InfrastructureService) HasExternalIPsTruncation() bool`

HasExternalIPsTruncation returns a boolean if a field has been set.

### GetMatchedPods

`func (o *InfrastructureService) GetMatchedPods() []InfrastructureObjectReference`

GetMatchedPods returns the MatchedPods field if non-nil, zero value otherwise.

### GetMatchedPodsOk

`func (o *InfrastructureService) GetMatchedPodsOk() (*[]InfrastructureObjectReference, bool)`

GetMatchedPodsOk returns a tuple with the MatchedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPods

`func (o *InfrastructureService) SetMatchedPods(v []InfrastructureObjectReference)`

SetMatchedPods sets MatchedPods field to given value.

### HasMatchedPods

`func (o *InfrastructureService) HasMatchedPods() bool`

HasMatchedPods returns a boolean if a field has been set.

### GetMatchedPodsAvailability

`func (o *InfrastructureService) GetMatchedPodsAvailability() InfrastructureAvailability`

GetMatchedPodsAvailability returns the MatchedPodsAvailability field if non-nil, zero value otherwise.

### GetMatchedPodsAvailabilityOk

`func (o *InfrastructureService) GetMatchedPodsAvailabilityOk() (*InfrastructureAvailability, bool)`

GetMatchedPodsAvailabilityOk returns a tuple with the MatchedPodsAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPodsAvailability

`func (o *InfrastructureService) SetMatchedPodsAvailability(v InfrastructureAvailability)`

SetMatchedPodsAvailability sets MatchedPodsAvailability field to given value.

### HasMatchedPodsAvailability

`func (o *InfrastructureService) HasMatchedPodsAvailability() bool`

HasMatchedPodsAvailability returns a boolean if a field has been set.

### GetMatchedPodsTruncation

`func (o *InfrastructureService) GetMatchedPodsTruncation() CheckpointTruncation`

GetMatchedPodsTruncation returns the MatchedPodsTruncation field if non-nil, zero value otherwise.

### GetMatchedPodsTruncationOk

`func (o *InfrastructureService) GetMatchedPodsTruncationOk() (*CheckpointTruncation, bool)`

GetMatchedPodsTruncationOk returns a tuple with the MatchedPodsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedPodsTruncation

`func (o *InfrastructureService) SetMatchedPodsTruncation(v CheckpointTruncation)`

SetMatchedPodsTruncation sets MatchedPodsTruncation field to given value.

### HasMatchedPodsTruncation

`func (o *InfrastructureService) HasMatchedPodsTruncation() bool`

HasMatchedPodsTruncation returns a boolean if a field has been set.

### GetPorts

`func (o *InfrastructureService) GetPorts() []InfrastructurePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InfrastructureService) GetPortsOk() (*[]InfrastructurePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InfrastructureService) SetPorts(v []InfrastructurePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InfrastructureService) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPortsTruncation

`func (o *InfrastructureService) GetPortsTruncation() CheckpointTruncation`

GetPortsTruncation returns the PortsTruncation field if non-nil, zero value otherwise.

### GetPortsTruncationOk

`func (o *InfrastructureService) GetPortsTruncationOk() (*CheckpointTruncation, bool)`

GetPortsTruncationOk returns a tuple with the PortsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsTruncation

`func (o *InfrastructureService) SetPortsTruncation(v CheckpointTruncation)`

SetPortsTruncation sets PortsTruncation field to given value.

### HasPortsTruncation

`func (o *InfrastructureService) HasPortsTruncation() bool`

HasPortsTruncation returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureService) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureService) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureService) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureService) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSelector

`func (o *InfrastructureService) GetSelector() []InfrastructureKeyValue`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *InfrastructureService) GetSelectorOk() (*[]InfrastructureKeyValue, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *InfrastructureService) SetSelector(v []InfrastructureKeyValue)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *InfrastructureService) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSelectorTruncation

`func (o *InfrastructureService) GetSelectorTruncation() CheckpointTruncation`

GetSelectorTruncation returns the SelectorTruncation field if non-nil, zero value otherwise.

### GetSelectorTruncationOk

`func (o *InfrastructureService) GetSelectorTruncationOk() (*CheckpointTruncation, bool)`

GetSelectorTruncationOk returns a tuple with the SelectorTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorTruncation

`func (o *InfrastructureService) SetSelectorTruncation(v CheckpointTruncation)`

SetSelectorTruncation sets SelectorTruncation field to given value.

### HasSelectorTruncation

`func (o *InfrastructureService) HasSelectorTruncation() bool`

HasSelectorTruncation returns a boolean if a field has been set.

### GetSessionAffinity

`func (o *InfrastructureService) GetSessionAffinity() string`

GetSessionAffinity returns the SessionAffinity field if non-nil, zero value otherwise.

### GetSessionAffinityOk

`func (o *InfrastructureService) GetSessionAffinityOk() (*string, bool)`

GetSessionAffinityOk returns a tuple with the SessionAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinity

`func (o *InfrastructureService) SetSessionAffinity(v string)`

SetSessionAffinity sets SessionAffinity field to given value.

### HasSessionAffinity

`func (o *InfrastructureService) HasSessionAffinity() bool`

HasSessionAffinity returns a boolean if a field has been set.

### GetType

`func (o *InfrastructureService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfrastructureService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfrastructureService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InfrastructureService) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


