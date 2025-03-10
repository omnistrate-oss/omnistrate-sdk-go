# PerUserCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | [**[]CostDataPerDate**](CostDataPerDate.md) |  | 
**Email** | **string** | The email of the user | 
**InstancesCost** | Pointer to [**[]PerInstanceCost**](PerInstanceCost.md) |  | [optional] 
**OrgID** | **string** | ID of an Org | 
**OrgName** | **string** | The name of the organization | 
**TotalCost** | **float64** | The total cost of the user | 
**UserName** | **string** | The name of the user | 

## Methods

### NewPerUserCost

`func NewPerUserCost(cost []CostDataPerDate, email string, orgID string, orgName string, totalCost float64, userName string, ) *PerUserCost`

NewPerUserCost instantiates a new PerUserCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerUserCostWithDefaults

`func NewPerUserCostWithDefaults() *PerUserCost`

NewPerUserCostWithDefaults instantiates a new PerUserCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *PerUserCost) GetCost() []CostDataPerDate`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PerUserCost) GetCostOk() (*[]CostDataPerDate, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PerUserCost) SetCost(v []CostDataPerDate)`

SetCost sets Cost field to given value.


### GetEmail

`func (o *PerUserCost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PerUserCost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PerUserCost) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInstancesCost

`func (o *PerUserCost) GetInstancesCost() []PerInstanceCost`

GetInstancesCost returns the InstancesCost field if non-nil, zero value otherwise.

### GetInstancesCostOk

`func (o *PerUserCost) GetInstancesCostOk() (*[]PerInstanceCost, bool)`

GetInstancesCostOk returns a tuple with the InstancesCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstancesCost

`func (o *PerUserCost) SetInstancesCost(v []PerInstanceCost)`

SetInstancesCost sets InstancesCost field to given value.

### HasInstancesCost

`func (o *PerUserCost) HasInstancesCost() bool`

HasInstancesCost returns a boolean if a field has been set.

### GetOrgID

`func (o *PerUserCost) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *PerUserCost) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *PerUserCost) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.


### GetOrgName

`func (o *PerUserCost) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *PerUserCost) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *PerUserCost) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetTotalCost

`func (o *PerUserCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PerUserCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PerUserCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetUserName

`func (o *PerUserCost) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *PerUserCost) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *PerUserCost) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


