# UserSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderNames** | **[]string** | List of cloud provider names | 
**DefaultSubscription** | **bool** | Whether this is the default subscription for the user | 
**Email** | **string** | The email of the user | 
**InstanceCount** | **int64** | The number of active instances in the subscription | 
**Name** | **string** | The name of the user | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ProductTierName** | **string** | The name of the product tier | 
**RoleType** | **string** | Type of the role | 
**ServiceEnvironmentId** | **string** | ID of a Service Environment | 
**ServiceId** | **string** | ID of a Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceName** | **string** | The name of the service | 
**SubscriptionDate** | **string** | The date the user joined the subscription | 
**SubscriptionId** | **string** | ID of a Subscription | 
**SubscriptionOwnerName** | **string** | The name of the subscription owner user | 
**UserId** | **string** | The User ID | 

## Methods

### NewUserSubscription

`func NewUserSubscription(cloudProviderNames []string, defaultSubscription bool, email string, instanceCount int64, name string, productTierId string, productTierName string, roleType string, serviceEnvironmentId string, serviceId string, serviceName string, subscriptionDate string, subscriptionId string, subscriptionOwnerName string, userId string, ) *UserSubscription`

NewUserSubscription instantiates a new UserSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSubscriptionWithDefaults

`func NewUserSubscriptionWithDefaults() *UserSubscription`

NewUserSubscriptionWithDefaults instantiates a new UserSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderNames

`func (o *UserSubscription) GetCloudProviderNames() []string`

GetCloudProviderNames returns the CloudProviderNames field if non-nil, zero value otherwise.

### GetCloudProviderNamesOk

`func (o *UserSubscription) GetCloudProviderNamesOk() (*[]string, bool)`

GetCloudProviderNamesOk returns a tuple with the CloudProviderNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderNames

`func (o *UserSubscription) SetCloudProviderNames(v []string)`

SetCloudProviderNames sets CloudProviderNames field to given value.


### GetDefaultSubscription

`func (o *UserSubscription) GetDefaultSubscription() bool`

GetDefaultSubscription returns the DefaultSubscription field if non-nil, zero value otherwise.

### GetDefaultSubscriptionOk

`func (o *UserSubscription) GetDefaultSubscriptionOk() (*bool, bool)`

GetDefaultSubscriptionOk returns a tuple with the DefaultSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubscription

`func (o *UserSubscription) SetDefaultSubscription(v bool)`

SetDefaultSubscription sets DefaultSubscription field to given value.


### GetEmail

`func (o *UserSubscription) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSubscription) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSubscription) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetInstanceCount

`func (o *UserSubscription) GetInstanceCount() int64`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *UserSubscription) GetInstanceCountOk() (*int64, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *UserSubscription) SetInstanceCount(v int64)`

SetInstanceCount sets InstanceCount field to given value.


### GetName

`func (o *UserSubscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserSubscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserSubscription) SetName(v string)`

SetName sets Name field to given value.


### GetProductTierId

`func (o *UserSubscription) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *UserSubscription) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *UserSubscription) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetProductTierName

`func (o *UserSubscription) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *UserSubscription) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *UserSubscription) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.


### GetRoleType

`func (o *UserSubscription) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *UserSubscription) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *UserSubscription) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.


### GetServiceEnvironmentId

`func (o *UserSubscription) GetServiceEnvironmentId() string`

GetServiceEnvironmentId returns the ServiceEnvironmentId field if non-nil, zero value otherwise.

### GetServiceEnvironmentIdOk

`func (o *UserSubscription) GetServiceEnvironmentIdOk() (*string, bool)`

GetServiceEnvironmentIdOk returns a tuple with the ServiceEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentId

`func (o *UserSubscription) SetServiceEnvironmentId(v string)`

SetServiceEnvironmentId sets ServiceEnvironmentId field to given value.


### GetServiceId

`func (o *UserSubscription) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UserSubscription) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UserSubscription) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetServiceLogoURL

`func (o *UserSubscription) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *UserSubscription) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *UserSubscription) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *UserSubscription) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.

### GetServiceName

`func (o *UserSubscription) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UserSubscription) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UserSubscription) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetSubscriptionDate

`func (o *UserSubscription) GetSubscriptionDate() string`

GetSubscriptionDate returns the SubscriptionDate field if non-nil, zero value otherwise.

### GetSubscriptionDateOk

`func (o *UserSubscription) GetSubscriptionDateOk() (*string, bool)`

GetSubscriptionDateOk returns a tuple with the SubscriptionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionDate

`func (o *UserSubscription) SetSubscriptionDate(v string)`

SetSubscriptionDate sets SubscriptionDate field to given value.


### GetSubscriptionId

`func (o *UserSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *UserSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *UserSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetSubscriptionOwnerName

`func (o *UserSubscription) GetSubscriptionOwnerName() string`

GetSubscriptionOwnerName returns the SubscriptionOwnerName field if non-nil, zero value otherwise.

### GetSubscriptionOwnerNameOk

`func (o *UserSubscription) GetSubscriptionOwnerNameOk() (*string, bool)`

GetSubscriptionOwnerNameOk returns a tuple with the SubscriptionOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionOwnerName

`func (o *UserSubscription) SetSubscriptionOwnerName(v string)`

SetSubscriptionOwnerName sets SubscriptionOwnerName field to given value.


### GetUserId

`func (o *UserSubscription) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSubscription) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSubscription) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


