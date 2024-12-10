# CreateCustomerOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | DEPRECATED: Name will be generated automatically. | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateCustomerOnboardingRequest

`func NewCreateCustomerOnboardingRequest(token string, ) *CreateCustomerOnboardingRequest`

NewCreateCustomerOnboardingRequest instantiates a new CreateCustomerOnboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerOnboardingRequestWithDefaults

`func NewCreateCustomerOnboardingRequestWithDefaults() *CreateCustomerOnboardingRequest`

NewCreateCustomerOnboardingRequestWithDefaults instantiates a new CreateCustomerOnboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCustomerOnboardingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomerOnboardingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomerOnboardingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCustomerOnboardingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *CreateCustomerOnboardingRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateCustomerOnboardingRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateCustomerOnboardingRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


