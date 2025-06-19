# DescribeSubscriptionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountConfigIdentityId** | **string** | ID of an Org | 
**CloudProviderNames** | **[]string** | List of cloud provider names | 
**CreatedAt** | **string** | The time that this subscription was created | 
**DefaultSubscription** | **bool** | Whether this is the default subscription for the user | 
**Id** | **string** | ID of a Subscription | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The name of the product tier | 
**RoleType** | **string** | Type of the role | 
**RootUserId** | **string** | ID of a User | 
**ServiceId** | **string** | ID of a Service | 
**ServiceLogoURL** | **string** | The logo for the service | 
**ServiceName** | **string** | The name of the service | 
**ServiceOrgId** | **string** | ID of an Org | 
**ServiceOrgName** | **string** | The name of the org that owns the service | 
**Status** | **string** | Subscription Status | 
**SubscriptionOwnerName** | **string** | The name of the subscription owner user | 

## Methods

### NewDescribeSubscriptionResult

`func NewDescribeSubscriptionResult(accountConfigIdentityId string, cloudProviderNames []string, createdAt string, defaultSubscription bool, id string, productTierId string, productTierName string, roleType string, rootUserId string, serviceId string, serviceLogoURL string, serviceName string, serviceOrgId string, serviceOrgName string, status string, subscriptionOwnerName string, ) *DescribeSubscriptionResult`

NewDescribeSubscriptionResult instantiates a new DescribeSubscriptionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeSubscriptionResultWithDefaults

`func NewDescribeSubscriptionResultWithDefaults() *DescribeSubscriptionResult`

NewDescribeSubscriptionResultWithDefaults instantiates a new DescribeSubscriptionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountConfigIdentityId

`func (o *DescribeSubscriptionResult) GetAccountConfigIdentityId() string`

GetAccountConfigIdentityId returns the AccountConfigIdentityId field if non-nil, zero value otherwise.

### GetAccountConfigIdentityIdOk

`func (o *DescribeSubscriptionResult) GetAccountConfigIdentityIdOk() (*string, bool)`

GetAccountConfigIdentityIdOk returns a tuple with the AccountConfigIdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigIdentityId

`func (o *DescribeSubscriptionResult) SetAccountConfigIdentityId(v string)`

SetAccountConfigIdentityId sets AccountConfigIdentityId field to given value.


### GetCloudProviderNames

`func (o *DescribeSubscriptionResult) GetCloudProviderNames() []string`

GetCloudProviderNames returns the CloudProviderNames field if non-nil, zero value otherwise.

### GetCloudProviderNamesOk

`func (o *DescribeSubscriptionResult) GetCloudProviderNamesOk() (*[]string, bool)`

GetCloudProviderNamesOk returns a tuple with the CloudProviderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderNames

`func (o *DescribeSubscriptionResult) SetCloudProviderNames(v []string)`

SetCloudProviderNames sets CloudProviderNames field to given value.


### GetCreatedAt

`func (o *DescribeSubscriptionResult) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DescribeSubscriptionResult) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DescribeSubscriptionResult) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultSubscription

`func (o *DescribeSubscriptionResult) GetDefaultSubscription() bool`

GetDefaultSubscription returns the DefaultSubscription field if non-nil, zero value otherwise.

### GetDefaultSubscriptionOk

`func (o *DescribeSubscriptionResult) GetDefaultSubscriptionOk() (*bool, bool)`

GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubscription

`func (o *DescribeSubscriptionResult) SetDefaultSubscription(v bool)`

SetDefaultSubscription sets DefaultSubscription field to given value.


### GetId

`func (o *DescribeSubscriptionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeSubscriptionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeSubscriptionResult) SetId(v string)`

SetId sets Id field to given value.


### GetProductTierId

`func (o *DescribeSubscriptionResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeSubscriptionResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeSubscriptionResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *DescribeSubscriptionResult) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *DescribeSubscriptionResult) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *DescribeSubscriptionResult) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetRoleType

`func (o *DescribeSubscriptionResult) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DescribeSubscriptionResult) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DescribeSubscriptionResult) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetRootUserId

`func (o *DescribeSubscriptionResult) GetRootUserId() string`

GetRootUserId returns the RootUserId field if non-nil, zero value otherwise.

### GetRootUserIdOk

`func (o *DescribeSubscriptionResult) GetRootUserIdOk() (*string, bool)`

GetRootUserIdOk returns a tuple with the RootUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserId

`func (o *DescribeSubscriptionResult) SetRootUserId(v string)`

SetRootUserId sets RootUserId field to given value.


### GetServiceId

`func (o *DescribeSubscriptionResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeSubscriptionResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeSubscriptionResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceLogoURL

`func (o *DescribeSubscriptionResult) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *DescribeSubscriptionResult) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *DescribeSubscriptionResult) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.


### GetServiceName

`func (o *DescribeSubscriptionResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DescribeSubscriptionResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DescribeSubscriptionResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServiceOrgId

`func (o *DescribeSubscriptionResult) GetServiceOrgId() string`

GetServiceOrgId returns the ServiceOrgId field if non-nil, zero value otherwise.

### GetServiceOrgIdOk

`func (o *DescribeSubscriptionResult) GetServiceOrgIdOk() (*string, bool)`

GetServiceOrgIdOk returns a tuple with the ServiceOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOrgId

`func (o *DescribeSubscriptionResult) SetServiceOrgId(v string)`

SetServiceOrgId sets ServiceOrgId field to given value.


### GetServiceOrgName

`func (o *DescribeSubscriptionResult) GetServiceOrgName() string`

GetServiceOrgName returns the ServiceOrgName field if non-nil, zero value otherwise.

### GetServiceOrgNameOk

`func (o *DescribeSubscriptionResult) GetServiceOrgNameOk() (*string, bool)`

GetServiceOrgNameOk returns a tuple with the ServiceOrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceOrgName

`func (o *DescribeSubscriptionResult) SetServiceOrgName(v string)`

SetServiceOrgName sets ServiceOrgName field to given value.


### GetStatus

`func (o *DescribeSubscriptionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeSubscriptionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeSubscriptionResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionOwnerName

`func (o *DescribeSubscriptionResult) GetSubscriptionOwnerName() string`

GetSubscriptionOwnerName returns the SubscriptionOwnerName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerNameOk

`func (o *DescribeSubscriptionResult) GetSubscriptionOwnerNameOk() (*string, bool)`

GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerName

`func (o *DescribeSubscriptionResult) SetSubscriptionOwnerName(v string)`

SetSubscriptionOwnerName sets SubscriptionOwnerName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


