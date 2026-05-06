# OnboardingResourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerVariables** | Pointer to [**[]OnboardingCustomerVariable**](OnboardingCustomerVariable.md) | Customer-provided variables. | [optional] 
**Resources** | Pointer to [**[]OnboardingResource**](OnboardingResource.md) | The resources in this onboarding. | [optional] 
**WorkplacePath** | Pointer to **string** | The workplace path for the resource. | [optional] 

## Methods

### NewOnboardingResourceConfig

`func NewOnboardingResourceConfig() *OnboardingResourceConfig`

NewOnboardingResourceConfig instantiates a new OnboardingResourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingResourceConfigWithDefaults

`func NewOnboardingResourceConfigWithDefaults() *OnboardingResourceConfig`

NewOnboardingResourceConfigWithDefaults instantiates a new OnboardingResourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerVariables

`func (o *OnboardingResourceConfig) GetCustomerVariables() []OnboardingCustomerVariable`

GetCustomerVariables returns the CustomerVariables field if non-nil, zero value otherwise.

### GetCustomerVariablesOk

`func (o *OnboardingResourceConfig) GetCustomerVariablesOk() (*[]OnboardingCustomerVariable, bool)`

GetCustomerVariablesOk returns a tuple with the CustomerVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerVariables

`func (o *OnboardingResourceConfig) SetCustomerVariables(v []OnboardingCustomerVariable)`

SetCustomerVariables sets CustomerVariables field to given value.

### HasCustomerVariables

`func (o *OnboardingResourceConfig) HasCustomerVariables() bool`

HasCustomerVariables returns a boolean if a field has been set.

### GetResources

`func (o *OnboardingResourceConfig) GetResources() []OnboardingResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *OnboardingResourceConfig) GetResourcesOk() (*[]OnboardingResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *OnboardingResourceConfig) SetResources(v []OnboardingResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *OnboardingResourceConfig) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetWorkplacePath

`func (o *OnboardingResourceConfig) GetWorkplacePath() string`

GetWorkplacePath returns the WorkplacePath field if non-nil, zero value otherwise.

### GetWorkplacePathOk

`func (o *OnboardingResourceConfig) GetWorkplacePathOk() (*string, bool)`

GetWorkplacePathOk returns a tuple with the WorkplacePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkplacePath

`func (o *OnboardingResourceConfig) SetWorkplacePath(v string)`

SetWorkplacePath sets WorkplacePath field to given value.

### HasWorkplacePath

`func (o *OnboardingResourceConfig) HasWorkplacePath() bool`

HasWorkplacePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


