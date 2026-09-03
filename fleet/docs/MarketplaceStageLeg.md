# MarketplaceStageLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | **string** | Who called whom. OUTBOUND and INBOUND are relative to the ISV; the CHANNEL_ pair are the marketplace side | 
**Duration** | Pointer to **string** | Server computed, so the console never reasons about clock skew | [optional] 
**Endpoint** | Pointer to **string** | The logical endpoint, for example channel.webhook, channel.contract.read or subscription.approve. Absent on a leg that is an emitted event rather than a call | [optional] 
**EventId** | Pointer to **string** | The interaction trail row this leg is, so the same call is never described twice in two places | [optional] 
**EventType** | Pointer to **string** | The wire event name on a leg that delivered one, for example contract.suspended | [optional] 
**OccurredAt** | **time.Time** | When the call happened. Legs are ordered by this, which is what makes the stepper a sequence rather than a set | 
**StatusCode** | Pointer to **int64** | What the other side answered, or what we answered them. Absent where the leg never completed, which is itself the thing an operator is looking for | [optional] 

## Methods

### NewMarketplaceStageLeg

`func NewMarketplaceStageLeg(direction string, occurredAt time.Time, ) *MarketplaceStageLeg`

NewMarketplaceStageLeg instantiates a new MarketplaceStageLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceStageLegWithDefaults

`func NewMarketplaceStageLegWithDefaults() *MarketplaceStageLeg`

NewMarketplaceStageLegWithDefaults instantiates a new MarketplaceStageLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *MarketplaceStageLeg) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MarketplaceStageLeg) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MarketplaceStageLeg) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetDuration

`func (o *MarketplaceStageLeg) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MarketplaceStageLeg) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MarketplaceStageLeg) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MarketplaceStageLeg) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndpoint

`func (o *MarketplaceStageLeg) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MarketplaceStageLeg) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MarketplaceStageLeg) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MarketplaceStageLeg) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEventId

`func (o *MarketplaceStageLeg) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *MarketplaceStageLeg) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *MarketplaceStageLeg) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *MarketplaceStageLeg) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *MarketplaceStageLeg) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *MarketplaceStageLeg) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *MarketplaceStageLeg) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *MarketplaceStageLeg) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetOccurredAt

`func (o *MarketplaceStageLeg) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *MarketplaceStageLeg) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *MarketplaceStageLeg) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetStatusCode

`func (o *MarketplaceStageLeg) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *MarketplaceStageLeg) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *MarketplaceStageLeg) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *MarketplaceStageLeg) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


