# UpdateCustomerOnboardingRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** | The ID of the service associated with this onboarding. | [optional] 
**Stage** | Pointer to [**OnboardingStage**](OnboardingStage.md) |  | [optional] 

## Methods

### NewUpdateCustomerOnboardingRequestBody

`func NewUpdateCustomerOnboardingRequestBody() *UpdateCustomerOnboardingRequestBody`

NewUpdateCustomerOnboardingRequestBody instantiates a new UpdateCustomerOnboardingRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomerOnboardingRequestBodyWithDefaults

`func NewUpdateCustomerOnboardingRequestBodyWithDefaults() *UpdateCustomerOnboardingRequestBody`

NewUpdateCustomerOnboardingRequestBodyWithDefaults instantiates a new UpdateCustomerOnboardingRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UpdateCustomerOnboardingRequestBody) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UpdateCustomerOnboardingRequestBody) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UpdateCustomerOnboardingRequestBody) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *UpdateCustomerOnboardingRequestBody) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetStage

`func (o *UpdateCustomerOnboardingRequestBody) GetStage() OnboardingStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *UpdateCustomerOnboardingRequestBody) GetStageOk() (*OnboardingStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *UpdateCustomerOnboardingRequestBody) SetStage(v OnboardingStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *UpdateCustomerOnboardingRequestBody) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


