# FleetDescribeSubscriptionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSubscriptionPricing** | Pointer to **map[string]interface{}** | The active pricing for the subscription at the time of the request. | [optional] 
**CreatedAt** | **string** | The time that this subscription was created | 
**HistorySubscriptionPricing** | Pointer to [**[]HistorySubscriptionPricing**](HistorySubscriptionPricing.md) | The history pricing for the subscription. | [optional] 
**Id** | **string** | ID of a Subscription | 
**InheritServicePlanPrice** | Pointer to **bool** | Whether to inherit the service plan price | [optional] 
**InstanceCount** | **int64** | The number of active instances in the subscription | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The name of the product tier | 
**RootUserEmail** | **string** | The email of the user that owns the subscription | 
**RootUserId** | **string** | ID of a User | 
**RootUserName** | **string** | The name of the user that owns the subscription | 
**RootUserOrgId** | Pointer to **string** | ID of an Org | [optional] 
**ServiceId** | **string** | ID of a Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceName** | **string** | The name of the service | 
**Status** | **string** | Subscription Status | 
**UpdatedAt** | **string** | The time that this subscription was last updated | 
**UpdatedByUserId** | **string** | ID of a User | 
**UpdatedByUserName** | **string** | The name of the user that last updated the subscription | 
**UpdatedByUserOrgId** | Pointer to **string** | ID of an Org | [optional] 

## Methods

### NewFleetDescribeSubscriptionResult

`func NewFleetDescribeSubscriptionResult(createdAt string, id string, instanceCount int64, productTierId string, productTierName string, rootUserEmail string, rootUserId string, rootUserName string, serviceId string, serviceName string, status string, updatedAt string, updatedByUserId string, updatedByUserName string, ) *FleetDescribeSubscriptionResult`

NewFleetDescribeSubscriptionResult instantiates a new FleetDescribeSubscriptionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeSubscriptionResultWithDefaults

`func NewFleetDescribeSubscriptionResultWithDefaults() *FleetDescribeSubscriptionResult`

NewFleetDescribeSubscriptionResultWithDefaults instantiates a new FleetDescribeSubscriptionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) GetActiveSubscriptionPricing() map[string]interface{}`

GetActiveSubscriptionPricing returns the ActiveSubscriptionPricing field if non-nil, zero value otherwise.

### GetActiveSubscriptionPricingOk

`func (o *FleetDescribeSubscriptionResult) GetActiveSubscriptionPricingOk() (*map[string]interface{}, bool)`

GetActiveSubscriptionPricingOk returns a tuple with the ActiveSubscriptionPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) SetActiveSubscriptionPricing(v map[string]interface{})`

SetActiveSubscriptionPricing sets ActiveSubscriptionPricing field to given value.

### HasActiveSubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) HasActiveSubscriptionPricing() bool`

HasActiveSubscriptionPricing returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FleetDescribeSubscriptionResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetDescribeSubscriptionResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetDescribeSubscriptionResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHistorySubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) GetHistorySubscriptionPricing() []HistorySubscriptionPricing`

GetHistorySubscriptionPricing returns the HistorySubscriptionPricing field if non-nil, zero value otherwise.

### GetHistorySubscriptionPricingOk

`func (o *FleetDescribeSubscriptionResult) GetHistorySubscriptionPricingOk() (*[]HistorySubscriptionPricing, bool)`

GetHistorySubscriptionPricingOk returns a tuple with the HistorySubscriptionPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistorySubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) SetHistorySubscriptionPricing(v []HistorySubscriptionPricing)`

SetHistorySubscriptionPricing sets HistorySubscriptionPricing field to given value.

### HasHistorySubscriptionPricing

`func (o *FleetDescribeSubscriptionResult) HasHistorySubscriptionPricing() bool`

HasHistorySubscriptionPricing returns a boolean if a field has been set.

### GetId

`func (o *FleetDescribeSubscriptionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetDescribeSubscriptionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetDescribeSubscriptionResult) SetId(v string)`

SetId sets Id field to given value.


### GetInheritServicePlanPrice

`func (o *FleetDescribeSubscriptionResult) GetInheritServicePlanPrice() bool`

GetInheritServicePlanPrice returns the InheritServicePlanPrice field if non-nil, zero value otherwise.

### GetInheritServicePlanPriceOk

`func (o *FleetDescribeSubscriptionResult) GetInheritServicePlanPriceOk() (*bool, bool)`

GetInheritServicePlanPriceOk returns a tuple with the InheritServicePlanPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritServicePlanPrice

`func (o *FleetDescribeSubscriptionResult) SetInheritServicePlanPrice(v bool)`

SetInheritServicePlanPrice sets InheritServicePlanPrice field to given value.

### HasInheritServicePlanPrice

`func (o *FleetDescribeSubscriptionResult) HasInheritServicePlanPrice() bool`

HasInheritServicePlanPrice returns a boolean if a field has been set.

### GetInstanceCount

`func (o *FleetDescribeSubscriptionResult) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *FleetDescribeSubscriptionResult) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *FleetDescribeSubscriptionResult) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.


### GetProductTierId

`func (o *FleetDescribeSubscriptionResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *FleetDescribeSubscriptionResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *FleetDescribeSubscriptionResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *FleetDescribeSubscriptionResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *FleetDescribeSubscriptionResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *FleetDescribeSubscriptionResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetRootUserEmail

`func (o *FleetDescribeSubscriptionResult) GetRootUserEmail() string`

GetRootUserEmail returns the RootUserEmail field if non-nil, zero value otherwise.

### GetRootUserEmailOk

`func (o *FleetDescribeSubscriptionResult) GetRootUserEmailOk() (*string, bool)`

GetRootUserEmailOk returns a tuple with the RootUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserEmail

`func (o *FleetDescribeSubscriptionResult) SetRootUserEmail(v string)`

SetRootUserEmail sets RootUserEmail field to given value.


### GetRootUserId

`func (o *FleetDescribeSubscriptionResult) GetRootUserId() string`

GetRootUserId returns the RootUserId field if non-nil, zero value otherwise.

### GetRootUserIdOk

`func (o *FleetDescribeSubscriptionResult) GetRootUserIdOk() (*string, bool)`

GetRootUserIdOk returns a tuple with the RootUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserId

`func (o *FleetDescribeSubscriptionResult) SetRootUserId(v string)`

SetRootUserId sets RootUserId field to given value.


### GetRootUserName

`func (o *FleetDescribeSubscriptionResult) GetRootUserName() string`

GetRootUserName returns the RootUserName field if non-nil, zero value otherwise.

### GetRootUserNameOk

`func (o *FleetDescribeSubscriptionResult) GetRootUserNameOk() (*string, bool)`

GetRootUserNameOk returns a tuple with the RootUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserName

`func (o *FleetDescribeSubscriptionResult) SetRootUserName(v string)`

SetRootUserName sets RootUserName field to given value.


### GetRootUserOrgId

`func (o *FleetDescribeSubscriptionResult) GetRootUserOrgId() string`

GetRootUserOrgId returns the RootUserOrgId field if non-nil, zero value otherwise.

### GetRootUserOrgIdOk

`func (o *FleetDescribeSubscriptionResult) GetRootUserOrgIdOk() (*string, bool)`

GetRootUserOrgIdOk returns a tuple with the RootUserOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserOrgId

`func (o *FleetDescribeSubscriptionResult) SetRootUserOrgId(v string)`

SetRootUserOrgId sets RootUserOrgId field to given value.

### HasRootUserOrgId

`func (o *FleetDescribeSubscriptionResult) HasRootUserOrgId() bool`

HasRootUserOrgId returns a boolean if a field has been set.

### GetServiceId

`func (o *FleetDescribeSubscriptionResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *FleetDescribeSubscriptionResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *FleetDescribeSubscriptionResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceLogoURL

`func (o *FleetDescribeSubscriptionResult) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *FleetDescribeSubscriptionResult) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *FleetDescribeSubscriptionResult) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *FleetDescribeSubscriptionResult) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetServiceName

`func (o *FleetDescribeSubscriptionResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *FleetDescribeSubscriptionResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *FleetDescribeSubscriptionResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *FleetDescribeSubscriptionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeSubscriptionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeSubscriptionResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *FleetDescribeSubscriptionResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FleetDescribeSubscriptionResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FleetDescribeSubscriptionResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedByUserId

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserId() string`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserIdOk() (*string, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *FleetDescribeSubscriptionResult) SetUpdatedByUserId(v string)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.


### GetUpdatedByUserName

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserName() string`

GetUpdatedByUserName returns the UpdatedByUserName field if non-nil, zero value otherwise.

### GetUpdatedByUserNameOk

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserNameOk() (*string, bool)`

GetUpdatedByUserNameOk returns a tuple with the UpdatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserName

`func (o *FleetDescribeSubscriptionResult) SetUpdatedByUserName(v string)`

SetUpdatedByUserName sets UpdatedByUserName field to given value.


### GetUpdatedByUserOrgId

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserOrgId() string`

GetUpdatedByUserOrgId returns the UpdatedByUserOrgId field if non-nil, zero value otherwise.

### GetUpdatedByUserOrgIdOk

`func (o *FleetDescribeSubscriptionResult) GetUpdatedByUserOrgIdOk() (*string, bool)`

GetUpdatedByUserOrgIdOk returns a tuple with the UpdatedByUserOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserOrgId

`func (o *FleetDescribeSubscriptionResult) SetUpdatedByUserOrgId(v string)`

SetUpdatedByUserOrgId sets UpdatedByUserOrgId field to given value.

### HasUpdatedByUserOrgId

`func (o *FleetDescribeSubscriptionResult) HasUpdatedByUserOrgId() bool`

HasUpdatedByUserOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


