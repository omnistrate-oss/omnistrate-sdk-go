# EndCustomerEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | **string** | The ID of the event | 
**EventPayload** | **map[string]interface{}** | The event payload for a service provider. | 
**EventType** | **string** | The type of the event  | 
**OrgID** | **string** | Associated organization ID. | 
**OrgName** | **string** | Associated organization name. | 
**OrgURL** | **string** | Associated organization URL. | 
**Priority** | **string** | The priority of the event | 
**Time** | **string** | The event time | 
**UserEmail** | Pointer to **string** | User email associated with the event. | [optional] 
**UserID** | Pointer to **string** | User ID associated with the event. | [optional] 
**UserName** | Pointer to **string** | Name of the user associated with the event. | [optional] 

## Methods

### NewEndCustomerEvent

`func NewEndCustomerEvent(eventID string, eventPayload map[string]interface{}, eventType string, orgID string, orgName string, orgURL string, priority string, time string, ) *EndCustomerEvent`

NewEndCustomerEvent instantiates a new EndCustomerEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndCustomerEventWithDefaults

`func NewEndCustomerEventWithDefaults() *EndCustomerEvent`

NewEndCustomerEventWithDefaults instantiates a new EndCustomerEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *EndCustomerEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *EndCustomerEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *EndCustomerEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.


### GetEventPayload

`func (o *EndCustomerEvent) GetEventPayload() map[string]interface{}`

GetEventPayload returns the EventPayload field if non-nil, zero value otherwise.

### GetEventPayloadOk

`func (o *EndCustomerEvent) GetEventPayloadOk() (*map[string]interface{}, bool)`

GetEventPayloadOk returns a tuple with the EventPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPayload

`func (o *EndCustomerEvent) SetEventPayload(v map[string]interface{})`

SetEventPayload sets EventPayload field to given value.


### GetEventType

`func (o *EndCustomerEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EndCustomerEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EndCustomerEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetOrgID

`func (o *EndCustomerEvent) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *EndCustomerEvent) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *EndCustomerEvent) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.


### GetOrgName

`func (o *EndCustomerEvent) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *EndCustomerEvent) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *EndCustomerEvent) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgURL

`func (o *EndCustomerEvent) GetOrgURL() string`

GetOrgURL returns the OrgURL field if non-nil, zero value otherwise.

### GetOrgURLOk

`func (o *EndCustomerEvent) GetOrgURLOk() (*string, bool)`

GetOrgURLOk returns a tuple with the OrgURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgURL

`func (o *EndCustomerEvent) SetOrgURL(v string)`

SetOrgURL sets OrgURL field to given value.


### GetPriority

`func (o *EndCustomerEvent) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EndCustomerEvent) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EndCustomerEvent) SetPriority(v string)`

SetPriority sets Priority field to given value.


### GetTime

`func (o *EndCustomerEvent) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *EndCustomerEvent) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *EndCustomerEvent) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserEmail

`func (o *EndCustomerEvent) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *EndCustomerEvent) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *EndCustomerEvent) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *EndCustomerEvent) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserID

`func (o *EndCustomerEvent) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *EndCustomerEvent) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *EndCustomerEvent) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *EndCustomerEvent) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetUserName

`func (o *EndCustomerEvent) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *EndCustomerEvent) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *EndCustomerEvent) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *EndCustomerEvent) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


