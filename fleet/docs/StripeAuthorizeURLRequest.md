# StripeAuthorizeURLRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewStripeAuthorizeURLRequest

`func NewStripeAuthorizeURLRequest(token string, ) *StripeAuthorizeURLRequest`

NewStripeAuthorizeURLRequest instantiates a new StripeAuthorizeURLRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStripeAuthorizeURLRequestWithDefaults

`func NewStripeAuthorizeURLRequestWithDefaults() *StripeAuthorizeURLRequest`

NewStripeAuthorizeURLRequestWithDefaults instantiates a new StripeAuthorizeURLRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *StripeAuthorizeURLRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StripeAuthorizeURLRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StripeAuthorizeURLRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


