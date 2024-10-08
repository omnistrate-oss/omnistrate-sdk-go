# CreateSubscriptionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | The product tier ID | 
**ServiceId** | **string** | The service ID | 

## Methods

### NewCreateSubscriptionRequestBody

`func NewCreateSubscriptionRequestBody(productTierId string, serviceId string, ) *CreateSubscriptionRequestBody`

NewCreateSubscriptionRequestBody instantiates a new CreateSubscriptionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubscriptionRequestBodyWithDefaults

`func NewCreateSubscriptionRequestBodyWithDefaults() *CreateSubscriptionRequestBody`

NewCreateSubscriptionRequestBodyWithDefaults instantiates a new CreateSubscriptionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *CreateSubscriptionRequestBody) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *CreateSubscriptionRequestBody) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *CreateSubscriptionRequestBody) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *CreateSubscriptionRequestBody) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CreateSubscriptionRequestBody) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CreateSubscriptionRequestBody) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


