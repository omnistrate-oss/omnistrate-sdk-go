# RedeliverMarketplaceDeliveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryId** | **string** | Which delivery to send again. Must be an outbound one: there is nothing to redeliver about a call the ISV made. The redelivery carries the SAME eventId, so a receiver that already processed this event correctly treats it as a duplicate and does nothing | 
**Reason** | Pointer to **string** | Why, recorded on the audit event. A redelivery is a manual intervention on a contract that is usually already billing | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewRedeliverMarketplaceDeliveryRequest

`func NewRedeliverMarketplaceDeliveryRequest(deliveryId string, token string, ) *RedeliverMarketplaceDeliveryRequest`

NewRedeliverMarketplaceDeliveryRequest instantiates a new RedeliverMarketplaceDeliveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeliverMarketplaceDeliveryRequestWithDefaults

`func NewRedeliverMarketplaceDeliveryRequestWithDefaults() *RedeliverMarketplaceDeliveryRequest`

NewRedeliverMarketplaceDeliveryRequestWithDefaults instantiates a new RedeliverMarketplaceDeliveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryId

`func (o *RedeliverMarketplaceDeliveryRequest) GetDeliveryId() string`

GetDeliveryId returns the DeliveryId field if non-nil, zero value otherwise.

### GetDeliveryIdOk

`func (o *RedeliverMarketplaceDeliveryRequest) GetDeliveryIdOk() (*string, bool)`

GetDeliveryIdOk returns a tuple with the DeliveryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryId

`func (o *RedeliverMarketplaceDeliveryRequest) SetDeliveryId(v string)`

SetDeliveryId sets DeliveryId field to given value.


### GetReason

`func (o *RedeliverMarketplaceDeliveryRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RedeliverMarketplaceDeliveryRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RedeliverMarketplaceDeliveryRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RedeliverMarketplaceDeliveryRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetToken

`func (o *RedeliverMarketplaceDeliveryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RedeliverMarketplaceDeliveryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RedeliverMarketplaceDeliveryRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


