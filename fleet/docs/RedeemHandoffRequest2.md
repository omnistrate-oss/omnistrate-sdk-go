# RedeemHandoffRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HandoffToken** | **string** | The value your buyer&#39;s arrival gave you: either the one appended to your callback URL when a buyer was redirected there, or the handoffToken from a contract.discovered webhook body. They are the same kind of value with the same seven day lifetime and the same non-consuming behaviour, so there is nothing to tell apart and one field takes either | 

## Methods

### NewRedeemHandoffRequest2

`func NewRedeemHandoffRequest2(handoffToken string, ) *RedeemHandoffRequest2`

NewRedeemHandoffRequest2 instantiates a new RedeemHandoffRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemHandoffRequest2WithDefaults

`func NewRedeemHandoffRequest2WithDefaults() *RedeemHandoffRequest2`

NewRedeemHandoffRequest2WithDefaults instantiates a new RedeemHandoffRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandoffToken

`func (o *RedeemHandoffRequest2) GetHandoffToken() string`

GetHandoffToken returns the HandoffToken field if non-nil, zero value otherwise.

### GetHandoffTokenOk

`func (o *RedeemHandoffRequest2) GetHandoffTokenOk() (*string, bool)`

GetHandoffTokenOk returns a tuple with the HandoffToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffToken

`func (o *RedeemHandoffRequest2) SetHandoffToken(v string)`

SetHandoffToken sets HandoffToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


