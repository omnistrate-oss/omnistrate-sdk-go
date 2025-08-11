# PerInstanceCost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**CostByInstanceType** | Pointer to [**map[string]InstanceTypeCost**](InstanceTypeCost.md) | The cost of the fleet by instance type on that date | [optional] 
**InstanceID** | **string** | ID of a Resource Instance | 
**IsDeleted** | **bool** | Whether the instance is deleted | 
**ProductTierID** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The name of the product tier | 
**ProductTierTenancyType** | **string** | The tenancy type of the product tier | 
**RegionName** | **string** | The name of the region | 
**ServiceEnvironmentID** | **string** | ID of a Service Environment | 
**ServiceID** | **string** | ID of a Service | 
**SubscriptionID** | **string** | ID of a Subscription | 
**TotalCost** | **float64** | The total cost of the instance | 
**Utilization** | **float64** | Percentage of the instance resource utilization | 

## Methods

### NewPerInstanceCost

`func NewPerInstanceCost(cloudProviderName string, instanceID string, isDeleted bool, productTierID string, productTierName string, productTierTenancyType string, regionName string, serviceEnvironmentID string, serviceID string, subscriptionID string, totalCost float64, utilization float64, ) *PerInstanceCost`

NewPerInstanceCost instantiates a new PerInstanceCost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerInstanceCostWithDefaults

`func NewPerInstanceCostWithDefaults() *PerInstanceCost`

NewPerInstanceCostWithDefaults instantiates a new PerInstanceCost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *PerInstanceCost) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *PerInstanceCost) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *PerInstanceCost) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCostByInstanceType

`func (o *PerInstanceCost) GetCostByInstanceType() map[string]InstanceTypeCost`

GetCostByInstanceType returns the CostByInstanceType field if non-nil, zero value otherwise.

### GetCostByInstanceTypeOk

`func (o *PerInstanceCost) GetCostByInstanceTypeOk() (*map[string]InstanceTypeCost, bool)`

GetCostByInstanceTypeOk returns a tuple with the CostByInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostByInstanceType

`func (o *PerInstanceCost) SetCostByInstanceType(v map[string]InstanceTypeCost)`

SetCostByInstanceType sets CostByInstanceType field to given value.

### HasCostByInstanceType

`func (o *PerInstanceCost) HasCostByInstanceType() bool`

HasCostByInstanceType returns a boolean if a field has been set.

### GetInstanceID

`func (o *PerInstanceCost) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *PerInstanceCost) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *PerInstanceCost) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.


### GetIsDeleted

`func (o *PerInstanceCost) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PerInstanceCost) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PerInstanceCost) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetProductTierID

`func (o *PerInstanceCost) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *PerInstanceCost) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *PerInstanceCost) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetProductTierName

`func (o *PerInstanceCost) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *PerInstanceCost) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *PerInstanceCost) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetProductTierTenancyType

`func (o *PerInstanceCost) GetProductTierTenancyType() string`

GetProductTierTenancyType returns the ProductTierTenancyType field if non-nil, zero value otherwise.

### GetProductTierTenancyTypeOk

`func (o *PerInstanceCost) GetProductTierTenancyTypeOk() (*string, bool)`

GetProductTierTenancyTypeOk returns a tuple with the ProductTierTenancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierTenancyType

`func (o *PerInstanceCost) SetProductTierTenancyType(v string)`

SetProductTierTenancyType sets ProductTierTenancyType field to given value.


### GetRegionName

`func (o *PerInstanceCost) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### GetRegionNameOk

`func (o *PerInstanceCost) GetRegionNameOk() (*string, bool)`

GetRegionNameOk returns a tuple with the RegionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionName

`func (o *PerInstanceCost) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.


### GetServiceEnvironmentID

`func (o *PerInstanceCost) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *PerInstanceCost) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *PerInstanceCost) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceID

`func (o *PerInstanceCost) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *PerInstanceCost) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *PerInstanceCost) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetSubscriptionID

`func (o *PerInstanceCost) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *PerInstanceCost) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *PerInstanceCost) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.


### GetTotalCost

`func (o *PerInstanceCost) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *PerInstanceCost) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *PerInstanceCost) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.


### GetUtilization

`func (o *PerInstanceCost) GetUtilization() float64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *PerInstanceCost) GetUtilizationOk() (*float64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *PerInstanceCost) SetUtilization(v float64)`

SetUtilization sets Utilization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


