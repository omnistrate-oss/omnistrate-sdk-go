# ListCustomerOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PendingOnly** | Pointer to **bool** | Whether to return only pending onboardings. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListCustomerOnboardingRequest

`func NewListCustomerOnboardingRequest(token string, ) *ListCustomerOnboardingRequest`

NewListCustomerOnboardingRequest instantiates a new ListCustomerOnboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerOnboardingRequestWithDefaults

`func NewListCustomerOnboardingRequestWithDefaults() *ListCustomerOnboardingRequest`

NewListCustomerOnboardingRequestWithDefaults instantiates a new ListCustomerOnboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPendingOnly

`func (o *ListCustomerOnboardingRequest) GetPendingOnly() bool`

GetPendingOnly returns the PendingOnly field if non-nil, zero value otherwise.

### GetPendingOnlyOk

`func (o *ListCustomerOnboardingRequest) GetPendingOnlyOk() (*bool, bool)`

GetPendingOnlyOk returns a tuple with the PendingOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingOnly

`func (o *ListCustomerOnboardingRequest) SetPendingOnly(v bool)`

SetPendingOnly sets PendingOnly field to given value.

### HasPendingOnly

`func (o *ListCustomerOnboardingRequest) HasPendingOnly() bool`

HasPendingOnly returns a boolean if a field has been set.

### GetToken

`func (o *ListCustomerOnboardingRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListCustomerOnboardingRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListCustomerOnboardingRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


