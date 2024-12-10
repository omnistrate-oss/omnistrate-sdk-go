# UpdateCustomerOnboardingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Customer Onboarding | 
**ServiceId** | Pointer to **string** | The ID of the service associated with this onboarding. | [optional] 
**Stage** | Pointer to [**OnboardingStage**](OnboardingStage.md) |  | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewUpdateCustomerOnboardingRequest

`func NewUpdateCustomerOnboardingRequest(id string, token string, ) *UpdateCustomerOnboardingRequest`

NewUpdateCustomerOnboardingRequest instantiates a new UpdateCustomerOnboardingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerOnboardingRequestWithDefaults

`func NewUpdateCustomerOnboardingRequestWithDefaults() *UpdateCustomerOnboardingRequest`

NewUpdateCustomerOnboardingRequestWithDefaults instantiates a new UpdateCustomerOnboardingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCustomerOnboardingRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCustomerOnboardingRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCustomerOnboardingRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServiceId

`func (o *UpdateCustomerOnboardingRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateCustomerOnboardingRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateCustomerOnboardingRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *UpdateCustomerOnboardingRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetStage

`func (o *UpdateCustomerOnboardingRequest) GetStage() OnboardingStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *UpdateCustomerOnboardingRequest) GetStageOk() (*OnboardingStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *UpdateCustomerOnboardingRequest) SetStage(v OnboardingStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *UpdateCustomerOnboardingRequest) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetToken

`func (o *UpdateCustomerOnboardingRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateCustomerOnboardingRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateCustomerOnboardingRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


