# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **interface{}** | Body of the event, including payload | [optional] 
**ChannelResponse** | Pointer to **interface{}** | Response from the notification channel, if applicable | [optional] 
**Id** | **string** | ID of a Event | 
**PublicationStatus** | **string** | Status of the event publication | 
**Timestamp** | **time.Time** | Timestamp when the event occurred | 

## Methods

### NewEvent

`func NewEvent(id string, publicationStatus string, timestamp time.Time, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Event) GetBody() interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Event) GetBodyOk() (*interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Event) SetBody(v interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *Event) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Event) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Event) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetChannelResponse

`func (o *Event) GetChannelResponse() interface{}`

GetChannelResponse returns the ChannelResponse field if non-nil, zero value otherwise.

### GetChannelResponseOk

`func (o *Event) GetChannelResponseOk() (*interface{}, bool)`

GetChannelResponseOk returns a tuple with the ChannelResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelResponse

`func (o *Event) SetChannelResponse(v interface{})`

SetChannelResponse sets ChannelResponse field to given value.

### HasChannelResponse

`func (o *Event) HasChannelResponse() bool`

HasChannelResponse returns a boolean if a field has been set.

### SetChannelResponseNil

`func (o *Event) SetChannelResponseNil(b bool)`

 SetChannelResponseNil sets the value for ChannelResponse to be an explicit nil

### UnsetChannelResponse
`func (o *Event) UnsetChannelResponse()`

UnsetChannelResponse ensures that no value is present for ChannelResponse, not even an explicit nil
### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetPublicationStatus

`func (o *Event) GetPublicationStatus() string`

GetPublicationStatus returns the PublicationStatus field if non-nil, zero value otherwise.

### GetPublicationStatusOk

`func (o *Event) GetPublicationStatusOk() (*string, bool)`

GetPublicationStatusOk returns a tuple with the PublicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationStatus

`func (o *Event) SetPublicationStatus(v string)`

SetPublicationStatus sets PublicationStatus field to given value.


### GetTimestamp

`func (o *Event) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


