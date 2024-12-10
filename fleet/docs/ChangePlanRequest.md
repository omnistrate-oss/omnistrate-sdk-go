# ChangePlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlanName** | **string** | This parameter is used to select the appropriate Product Plan | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewChangePlanRequest

`func NewChangePlanRequest(planName string, token string, ) *ChangePlanRequest`

NewChangePlanRequest instantiates a new ChangePlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePlanRequestWithDefaults

`func NewChangePlanRequestWithDefaults() *ChangePlanRequest`

NewChangePlanRequestWithDefaults instantiates a new ChangePlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanName

`func (o *ChangePlanRequest) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *ChangePlanRequest) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *ChangePlanRequest) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetToken

`func (o *ChangePlanRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChangePlanRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChangePlanRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


