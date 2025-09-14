# DescribeCloudProviderCostResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderCosts** | Pointer to [**map[string]PerCloudProviderCost**](PerCloudProviderCost.md) | The cost data for each cloud provider | [optional] 

## Methods

### NewDescribeCloudProviderCostResult

`func NewDescribeCloudProviderCostResult() *DescribeCloudProviderCostResult`

NewDescribeCloudProviderCostResult instantiates a new DescribeCloudProviderCostResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCloudProviderCostResultWithDefaults

`func NewDescribeCloudProviderCostResultWithDefaults() *DescribeCloudProviderCostResult`

NewDescribeCloudProviderCostResultWithDefaults instantiates a new DescribeCloudProviderCostResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderCosts

`func (o *DescribeCloudProviderCostResult) GetCloudProviderCosts() map[string]PerCloudProviderCost`

GetCloudProviderCosts returns the CloudProviderCosts field if non-nil, zero value otherwise.

### GetCloudProviderCostsOk

`func (o *DescribeCloudProviderCostResult) GetCloudProviderCostsOk() (*map[string]PerCloudProviderCost, bool)`

GetCloudProviderCostsOk returns a tuple with the CloudProviderCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderCosts

`func (o *DescribeCloudProviderCostResult) SetCloudProviderCosts(v map[string]PerCloudProviderCost)`

SetCloudProviderCosts sets CloudProviderCosts field to given value.

### HasCloudProviderCosts

`func (o *DescribeCloudProviderCostResult) HasCloudProviderCosts() bool`

HasCloudProviderCosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


