# OnboardingStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the stage. | 
**Status** | Pointer to **string** | The status of the stage. | [optional] 

## Methods

### NewOnboardingStage

`func NewOnboardingStage(name string, ) *OnboardingStage`

NewOnboardingStage instantiates a new OnboardingStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingStageWithDefaults

`func NewOnboardingStageWithDefaults() *OnboardingStage`

NewOnboardingStageWithDefaults instantiates a new OnboardingStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OnboardingStage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnboardingStage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnboardingStage) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *OnboardingStage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OnboardingStage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OnboardingStage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OnboardingStage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


