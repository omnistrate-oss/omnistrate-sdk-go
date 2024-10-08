# ChannelSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertTypes** | **[]string** | Types of alerts to route to this channel (optional) | 
**EnvironmentTypes** | **[]string** | Types of environments to route to this channel (optional) | 
**EventCategories** | **[]string** | Categories of events to route to this channel | 
**EventPriorities** | **[]string** | Priorities of events to route to this channel | 
**EventTypes** | **[]string** | Types of events to route to this channel (optional) | 

## Methods

### NewChannelSubscription

`func NewChannelSubscription(alertTypes []string, environmentTypes []string, eventCategories []string, eventPriorities []string, eventTypes []string, ) *ChannelSubscription`

NewChannelSubscription instantiates a new ChannelSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSubscriptionWithDefaults

`func NewChannelSubscriptionWithDefaults() *ChannelSubscription`

NewChannelSubscriptionWithDefaults instantiates a new ChannelSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertTypes

`func (o *ChannelSubscription) GetAlertTypes() []string`

GetAlertTypes returns the AlertTypes field if non-nil, zero value otherwise.

### GetAlertTypesOk

`func (o *ChannelSubscription) GetAlertTypesOk() (*[]string, bool)`

GetAlertTypesOk returns a tuple with the AlertTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTypes

`func (o *ChannelSubscription) SetAlertTypes(v []string)`

SetAlertTypes sets AlertTypes field to given value.


### GetEnvironmentTypes

`func (o *ChannelSubscription) GetEnvironmentTypes() []string`

GetEnvironmentTypes returns the EnvironmentTypes field if non-nil, zero value otherwise.

### GetEnvironmentTypesOk

`func (o *ChannelSubscription) GetEnvironmentTypesOk() (*[]string, bool)`

GetEnvironmentTypesOk returns a tuple with the EnvironmentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentTypes

`func (o *ChannelSubscription) SetEnvironmentTypes(v []string)`

SetEnvironmentTypes sets EnvironmentTypes field to given value.


### GetEventCategories

`func (o *ChannelSubscription) GetEventCategories() []string`

GetEventCategories returns the EventCategories field if non-nil, zero value otherwise.

### GetEventCategoriesOk

`func (o *ChannelSubscription) GetEventCategoriesOk() (*[]string, bool)`

GetEventCategoriesOk returns a tuple with the EventCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCategories

`func (o *ChannelSubscription) SetEventCategories(v []string)`

SetEventCategories sets EventCategories field to given value.


### GetEventPriorities

`func (o *ChannelSubscription) GetEventPriorities() []string`

GetEventPriorities returns the EventPriorities field if non-nil, zero value otherwise.

### GetEventPrioritiesOk

`func (o *ChannelSubscription) GetEventPrioritiesOk() (*[]string, bool)`

GetEventPrioritiesOk returns a tuple with the EventPriorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventPriorities

`func (o *ChannelSubscription) SetEventPriorities(v []string)`

SetEventPriorities sets EventPriorities field to given value.


### GetEventTypes

`func (o *ChannelSubscription) GetEventTypes() []string`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *ChannelSubscription) GetEventTypesOk() (*[]string, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *ChannelSubscription) SetEventTypes(v []string)`

SetEventTypes sets EventTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


