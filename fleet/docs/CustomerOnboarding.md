# CustomerOnboarding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of a Customer Onboarding | 
**IsCompleted** | **bool** | Whether the onboarding is completed. | 
**Name** | Pointer to **string** | The name of the onboarding. | [optional] 
**OrgId** | **string** | ID of an Org | 
**ServiceId** | Pointer to **string** | The ID of the service associated with this onboarding. | [optional] 
**Stages** | [**[]OnboardingStage**](OnboardingStage.md) | The stages of the onboarding. | 
**UserId** | **string** | ID of a User | 

## Methods

### NewCustomerOnboarding

`func NewCustomerOnboarding(id string, isCompleted bool, orgId string, stages []OnboardingStage, userId string, ) *CustomerOnboarding`

NewCustomerOnboarding instantiates a new CustomerOnboarding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerOnboardingWithDefaults

`func NewCustomerOnboardingWithDefaults() *CustomerOnboarding`

NewCustomerOnboardingWithDefaults instantiates a new CustomerOnboarding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerOnboarding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerOnboarding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerOnboarding) SetId(v string)`

SetId sets Id field to given value.


### GetIsCompleted

`func (o *CustomerOnboarding) GetIsCompleted() bool`

GetIsCompleted returns the IsCompleted field if non-nil, zero value otherwise.

### GetIsCompletedOk

`func (o *CustomerOnboarding) GetIsCompletedOk() (*bool, bool)`

GetIsCompletedOk returns a tuple with the IsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompleted

`func (o *CustomerOnboarding) SetIsCompleted(v bool)`

SetIsCompleted sets IsCompleted field to given value.


### GetName

`func (o *CustomerOnboarding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerOnboarding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerOnboarding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomerOnboarding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *CustomerOnboarding) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *CustomerOnboarding) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *CustomerOnboarding) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetServiceId

`func (o *CustomerOnboarding) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CustomerOnboarding) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CustomerOnboarding) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *CustomerOnboarding) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetStages

`func (o *CustomerOnboarding) GetStages() []OnboardingStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *CustomerOnboarding) GetStagesOk() (*[]OnboardingStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *CustomerOnboarding) SetStages(v []OnboardingStage)`

SetStages sets Stages field to given value.


### GetUserId

`func (o *CustomerOnboarding) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CustomerOnboarding) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CustomerOnboarding) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


