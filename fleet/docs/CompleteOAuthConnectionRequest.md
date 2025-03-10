# CompleteOAuthConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Authorization code from Stripe | [optional] 
**State** | Pointer to **string** | Random string used on the authorize URL | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCompleteOAuthConnectionRequest

`func NewCompleteOAuthConnectionRequest(token string, ) *CompleteOAuthConnectionRequest`

NewCompleteOAuthConnectionRequest instantiates a new CompleteOAuthConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteOAuthConnectionRequestWithDefaults

`func NewCompleteOAuthConnectionRequestWithDefaults() *CompleteOAuthConnectionRequest`

NewCompleteOAuthConnectionRequestWithDefaults instantiates a new CompleteOAuthConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CompleteOAuthConnectionRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CompleteOAuthConnectionRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CompleteOAuthConnectionRequest) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CompleteOAuthConnectionRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetState

`func (o *CompleteOAuthConnectionRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CompleteOAuthConnectionRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CompleteOAuthConnectionRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CompleteOAuthConnectionRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetToken

`func (o *CompleteOAuthConnectionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CompleteOAuthConnectionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CompleteOAuthConnectionRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


