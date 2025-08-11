# CostDataPerDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | **float64** | The total cost of the fleet on that date | 
**CostByInstanceType** | Pointer to [**map[string]InstanceTypeCost**](InstanceTypeCost.md) | The cost of the fleet by instance type on that date | [optional] 
**Date** | **time.Time** | The date of the cost data | 

## Methods

### NewCostDataPerDate

`func NewCostDataPerDate(cost float64, date time.Time, ) *CostDataPerDate`

NewCostDataPerDate instantiates a new CostDataPerDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCostDataPerDateWithDefaults

`func NewCostDataPerDateWithDefaults() *CostDataPerDate`

NewCostDataPerDateWithDefaults instantiates a new CostDataPerDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCost

`func (o *CostDataPerDate) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CostDataPerDate) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CostDataPerDate) SetCost(v float64)`

SetCost sets Cost field to given value.


### GetCostByInstanceType

`func (o *CostDataPerDate) GetCostByInstanceType() map[string]InstanceTypeCost`

GetCostByInstanceType returns the CostByInstanceType field if non-nil, zero value otherwise.

### GetCostByInstanceTypeOk

`func (o *CostDataPerDate) GetCostByInstanceTypeOk() (*map[string]InstanceTypeCost, bool)`

GetCostByInstanceTypeOk returns a tuple with the CostByInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostByInstanceType

`func (o *CostDataPerDate) SetCostByInstanceType(v map[string]InstanceTypeCost)`

SetCostByInstanceType sets CostByInstanceType field to given value.

### HasCostByInstanceType

`func (o *CostDataPerDate) HasCostByInstanceType() bool`

HasCostByInstanceType returns a boolean if a field has been set.

### GetDate

`func (o *CostDataPerDate) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CostDataPerDate) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CostDataPerDate) SetDate(v time.Time)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


