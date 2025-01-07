# ServiceProviderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertType** | **string** | The type of the alert | 
**EventCategory** | **string** | The category of the event | 
**EventID** | **string** | The ID of the event | 
**EventPayload** | **map[string]interface{}** | The event payload for a service provider | 
**EventType** | **string** | The type of the event | 
**ExpiryTime** | **string** | The expiry time of the event | 
**InstanceID** | Pointer to **string** | The ID of the instance | [optional] 
**PlanVersion** | Pointer to **string** | The version of the plan | [optional] 
**Priority** | **string** | The priority of the event | 
**ResourceName** | Pointer to **string** | The name of the resource | [optional] 
**Scope** | **string** | The scope of the event | 
**ServiceEnvironmentID** | Pointer to **string** | The ID of the service environment | [optional] 
**ServiceID** | Pointer to **string** | The ID of the service | [optional] 
**ServiceName** | Pointer to **string** | The name of the service | [optional] 
**ServicePlanName** | Pointer to **string** | The name of the service plan | [optional] 
**Time** | **string** | The event time | 

## Methods

### NewServiceProviderEvent

`func NewServiceProviderEvent(alertType string, eventCategory string, eventID string, eventPayload map[string]interface{}, eventType string, expiryTime string, priority string, scope string, time string, ) *ServiceProviderEvent`

NewServiceProviderEvent instantiates a new ServiceProviderEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderEventWithDefaults

`func NewServiceProviderEventWithDefaults() *ServiceProviderEvent`

NewServiceProviderEventWithDefaults instantiates a new ServiceProviderEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertType

`func (o *ServiceProviderEvent) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *ServiceProviderEvent) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *ServiceProviderEvent) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.


### GetEventCategory

`func (o *ServiceProviderEvent) GetEventCategory() string`

GetEventCategory returns the EventCategory field if non-nil, zero value otherwise.

### GetEventCategoryOk

`func (o *ServiceProviderEvent) GetEventCategoryOk() (*string, bool)`

GetEventCategoryOk returns a tuple with the EventCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategory

`func (o *ServiceProviderEvent) SetEventCategory(v string)`

SetEventCategory sets EventCategory field to given value.


### GetEventID

`func (o *ServiceProviderEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *ServiceProviderEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *ServiceProviderEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetEventPayload

`func (o *ServiceProviderEvent) GetEventPayload() map[string]interface{}`

GetEventPayload returns the EventPayload field if non-nil, zero value otherwise.

### GetEventPayloadOk

`func (o *ServiceProviderEvent) GetEventPayloadOk() (*map[string]interface{}, bool)`

GetEventPayloadOk returns a tuple with the EventPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayload

`func (o *ServiceProviderEvent) SetEventPayload(v map[string]interface{})`

SetEventPayload sets EventPayload field to given value.


### GetEventType

`func (o *ServiceProviderEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ServiceProviderEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ServiceProviderEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetExpiryTime

`func (o *ServiceProviderEvent) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ServiceProviderEvent) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ServiceProviderEvent) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.


### GetInstanceID

`func (o *ServiceProviderEvent) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *ServiceProviderEvent) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *ServiceProviderEvent) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *ServiceProviderEvent) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetPlanVersion

`func (o *ServiceProviderEvent) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *ServiceProviderEvent) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *ServiceProviderEvent) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *ServiceProviderEvent) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetPriority

`func (o *ServiceProviderEvent) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ServiceProviderEvent) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ServiceProviderEvent) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetResourceName

`func (o *ServiceProviderEvent) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ServiceProviderEvent) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ServiceProviderEvent) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ServiceProviderEvent) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetScope

`func (o *ServiceProviderEvent) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ServiceProviderEvent) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ServiceProviderEvent) SetScope(v string)`

SetScope sets Scope field to given value.


### GetServiceEnvironmentID

`func (o *ServiceProviderEvent) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *ServiceProviderEvent) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *ServiceProviderEvent) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.

### HasServiceEnvironmentID

`func (o *ServiceProviderEvent) HasServiceEnvironmentID() bool`

HasServiceEnvironmentID returns a boolean if a field has been set.

### GetServiceID

`func (o *ServiceProviderEvent) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *ServiceProviderEvent) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *ServiceProviderEvent) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *ServiceProviderEvent) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceProviderEvent) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceProviderEvent) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceProviderEvent) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceProviderEvent) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServicePlanName

`func (o *ServiceProviderEvent) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *ServiceProviderEvent) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *ServiceProviderEvent) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *ServiceProviderEvent) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetTime

`func (o *ServiceProviderEvent) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ServiceProviderEvent) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ServiceProviderEvent) SetTime(v string)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


