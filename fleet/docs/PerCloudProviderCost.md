# PerCloudProviderCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Cost** | [**[]CostDataPerDate**](CostDataPerDate.md) |  | 
**TotalCost** | **float64** | The total cost of the fleet on the cloud provider | 

## Methods

### NewPerCloudProviderCost

`func NewPerCloudProviderCost(cloudProviderName string, cost []CostDataPerDate, totalCost float64, ) *PerCloudProviderCost`

NewPerCloudProviderCost instantiates a new PerCloudProviderCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerCloudProviderCostWithDefaults

`func NewPerCloudProviderCostWithDefaults() *PerCloudProviderCost`

NewPerCloudProviderCostWithDefaults instantiates a new PerCloudProviderCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *PerCloudProviderCost) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *PerCloudProviderCost) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *PerCloudProviderCost) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCost

`func (o *PerCloudProviderCost) GetCost() []CostDataPerDate`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *PerCloudProviderCost) GetCostOk() (*[]CostDataPerDate, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *PerCloudProviderCost) SetCost(v []CostDataPerDate)`

SetCost sets Cost field to given value.


### GetTotalCost

`func (o *PerCloudProviderCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PerCloudProviderCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PerCloudProviderCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


