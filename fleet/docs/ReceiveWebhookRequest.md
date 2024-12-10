# ReceiveWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **string** | The event data | [optional] 
**Id** | **string** | The unique id per producer. | 

## Methods

### NewReceiveWebhookRequest

`func NewReceiveWebhookRequest(id string, ) *ReceiveWebhookRequest`

NewReceiveWebhookRequest instantiates a new ReceiveWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReceiveWebhookRequestWithDefaults

`func NewReceiveWebhookRequestWithDefaults() *ReceiveWebhookRequest`

NewReceiveWebhookRequestWithDefaults instantiates a new ReceiveWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ReceiveWebhookRequest) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ReceiveWebhookRequest) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ReceiveWebhookRequest) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ReceiveWebhookRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *ReceiveWebhookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReceiveWebhookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReceiveWebhookRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


