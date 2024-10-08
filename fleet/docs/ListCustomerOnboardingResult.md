# ListCustomerOnboardingResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Onboardings** | [**[]CustomerOnboarding**](CustomerOnboarding.md) | List of customer onboardings. | 

## Methods

### NewListCustomerOnboardingResult

`func NewListCustomerOnboardingResult(onboardings []CustomerOnboarding, ) *ListCustomerOnboardingResult`

NewListCustomerOnboardingResult instantiates a new ListCustomerOnboardingResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerOnboardingResultWithDefaults

`func NewListCustomerOnboardingResultWithDefaults() *ListCustomerOnboardingResult`

NewListCustomerOnboardingResultWithDefaults instantiates a new ListCustomerOnboardingResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnboardings

`func (o *ListCustomerOnboardingResult) GetOnboardings() []CustomerOnboarding`

GetOnboardings returns the Onboardings field if non-nil, zero value otherwise.

### GetOnboardingsOk

`func (o *ListCustomerOnboardingResult) GetOnboardingsOk() (*[]CustomerOnboarding, bool)`

GetOnboardingsOk returns a tuple with the Onboardings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardings

`func (o *ListCustomerOnboardingResult) SetOnboardings(v []CustomerOnboarding)`

SetOnboardings sets Onboardings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


