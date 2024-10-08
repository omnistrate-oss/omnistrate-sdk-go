# ListNotificationChannelsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]Channel**](Channel.md) | List of notification channels | [optional] 

## Methods

### NewListNotificationChannelsResult

`func NewListNotificationChannelsResult() *ListNotificationChannelsResult`

NewListNotificationChannelsResult instantiates a new ListNotificationChannelsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationChannelsResultWithDefaults

`func NewListNotificationChannelsResultWithDefaults() *ListNotificationChannelsResult`

NewListNotificationChannelsResultWithDefaults instantiates a new ListNotificationChannelsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ListNotificationChannelsResult) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ListNotificationChannelsResult) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ListNotificationChannelsResult) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ListNotificationChannelsResult) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


