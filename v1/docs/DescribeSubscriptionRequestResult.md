# DescribeSubscriptionRequestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time that this subscription request was issued | 
**Id** | **string** | ID of a Subscription Request | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The name of the product tier | 
**RootUserEmail** | **string** | The email of the user that issued the subscription request | 
**RootUserId** | **string** | ID of a User | 
**RootUserName** | **string** | The name of the user that issued the subscription request | 
**ServiceId** | **string** | ID of a Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceName** | **string** | The name of the service | 
**Status** | **string** | The status of the subscription request | 
**UpdatedAt** | **string** | The time that this subscription request was issued | 
**UpdatedByUserId** | **string** | ID of a User | 
**UpdatedByUserName** | **string** | The user that last updated the subscription request | 

## Methods

### NewDescribeSubscriptionRequestResult

`func NewDescribeSubscriptionRequestResult(createdAt string, id string, productTierId string, productTierName string, rootUserEmail string, rootUserId string, rootUserName string, serviceId string, serviceName string, status string, updatedAt string, updatedByUserId string, updatedByUserName string, ) *DescribeSubscriptionRequestResult`

NewDescribeSubscriptionRequestResult instantiates a new DescribeSubscriptionRequestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSubscriptionRequestResultWithDefaults

`func NewDescribeSubscriptionRequestResultWithDefaults() *DescribeSubscriptionRequestResult`

NewDescribeSubscriptionRequestResultWithDefaults instantiates a new DescribeSubscriptionRequestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DescribeSubscriptionRequestResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeSubscriptionRequestResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeSubscriptionRequestResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *DescribeSubscriptionRequestResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeSubscriptionRequestResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeSubscriptionRequestResult) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierId

`func (o *DescribeSubscriptionRequestResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeSubscriptionRequestResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeSubscriptionRequestResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *DescribeSubscriptionRequestResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *DescribeSubscriptionRequestResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *DescribeSubscriptionRequestResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetRootUserEmail

`func (o *DescribeSubscriptionRequestResult) GetRootUserEmail() string`

GetRootUserEmail returns the RootUserEmail field if non-nil, zero value otherwise.

### GetRootUserEmailOk

`func (o *DescribeSubscriptionRequestResult) GetRootUserEmailOk() (*string, bool)`

GetRootUserEmailOk returns a tuple with the RootUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserEmail

`func (o *DescribeSubscriptionRequestResult) SetRootUserEmail(v string)`

SetRootUserEmail sets RootUserEmail field to given value.


### GetRootUserId

`func (o *DescribeSubscriptionRequestResult) GetRootUserId() string`

GetRootUserId returns the RootUserId field if non-nil, zero value otherwise.

### GetRootUserIdOk

`func (o *DescribeSubscriptionRequestResult) GetRootUserIdOk() (*string, bool)`

GetRootUserIdOk returns a tuple with the RootUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserId

`func (o *DescribeSubscriptionRequestResult) SetRootUserId(v string)`

SetRootUserId sets RootUserId field to given value.


### GetRootUserName

`func (o *DescribeSubscriptionRequestResult) GetRootUserName() string`

GetRootUserName returns the RootUserName field if non-nil, zero value otherwise.

### GetRootUserNameOk

`func (o *DescribeSubscriptionRequestResult) GetRootUserNameOk() (*string, bool)`

GetRootUserNameOk returns a tuple with the RootUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserName

`func (o *DescribeSubscriptionRequestResult) SetRootUserName(v string)`

SetRootUserName sets RootUserName field to given value.


### GetServiceId

`func (o *DescribeSubscriptionRequestResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeSubscriptionRequestResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeSubscriptionRequestResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceLogoURL

`func (o *DescribeSubscriptionRequestResult) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *DescribeSubscriptionRequestResult) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *DescribeSubscriptionRequestResult) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *DescribeSubscriptionRequestResult) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetServiceName

`func (o *DescribeSubscriptionRequestResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DescribeSubscriptionRequestResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DescribeSubscriptionRequestResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *DescribeSubscriptionRequestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeSubscriptionRequestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeSubscriptionRequestResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *DescribeSubscriptionRequestResult) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DescribeSubscriptionRequestResult) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DescribeSubscriptionRequestResult) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedByUserId

`func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserId() string`

GetUpdatedByUserId returns the UpdatedByUserId field if non-nil, zero value otherwise.

### GetUpdatedByUserIdOk

`func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserIdOk() (*string, bool)`

GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserId

`func (o *DescribeSubscriptionRequestResult) SetUpdatedByUserId(v string)`

SetUpdatedByUserId sets UpdatedByUserId field to given value.


### GetUpdatedByUserName

`func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserName() string`

GetUpdatedByUserName returns the UpdatedByUserName field if non-nil, zero value otherwise.

### GetUpdatedByUserNameOk

`func (o *DescribeSubscriptionRequestResult) GetUpdatedByUserNameOk() (*string, bool)`

GetUpdatedByUserNameOk returns a tuple with the UpdatedByUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByUserName

`func (o *DescribeSubscriptionRequestResult) SetUpdatedByUserName(v string)`

SetUpdatedByUserName sets UpdatedByUserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


