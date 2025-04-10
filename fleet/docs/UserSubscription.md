# UserSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderNames** | Pointer to **[]string** | List of cloud provider names | [optional] 
**DefaultSubscription** | Pointer to **bool** | Whether this is the default subscription for the user | [optional] 
**Email** | Pointer to **string** | [DEPRECATED] The email of the user | [optional] 
**InstanceCount** | Pointer to **int64** | The number of active instances in the subscription | [optional] 
**Name** | Pointer to **string** | [DEPRECATED] The name of the user | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierName** | Pointer to **string** | The name of the product tier | [optional] 
**RoleType** | Pointer to **string** | Type of the role | [optional] 
**ServiceEnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 
**ServiceName** | Pointer to **string** | The name of the service | [optional] 
**SubscriptionDate** | Pointer to **string** | The date the user joined the subscription | [optional] 
**SubscriptionId** | Pointer to **string** | ID of a Subscription | [optional] 
**SubscriptionOwnerName** | Pointer to **string** | The name of the subscription owner user | [optional] 
**UserId** | Pointer to **string** | [DEPRECATED] The User ID | [optional] 

## Methods

### NewUserSubscription

`func NewUserSubscription() *UserSubscription`

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

### HasCloudProviderNames

`func (o *UserSubscription) HasCloudProviderNames() bool`

HasCloudProviderNames returns a boolean if a field has been set.

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

### HasDefaultSubscription

`func (o *UserSubscription) HasDefaultSubscription() bool`

HasDefaultSubscription returns a boolean if a field has been set.

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

### HasEmail

`func (o *UserSubscription) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

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

### HasInstanceCount

`func (o *UserSubscription) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

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

### HasName

`func (o *UserSubscription) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasProductTierId

`func (o *UserSubscription) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

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

### HasProductTierName

`func (o *UserSubscription) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

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

### HasRoleType

`func (o *UserSubscription) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

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

### HasServiceEnvironmentId

`func (o *UserSubscription) HasServiceEnvironmentId() bool`

HasServiceEnvironmentId returns a boolean if a field has been set.

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

### HasServiceId

`func (o *UserSubscription) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

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

### HasServiceName

`func (o *UserSubscription) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

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

### HasSubscriptionDate

`func (o *UserSubscription) HasSubscriptionDate() bool`

HasSubscriptionDate returns a boolean if a field has been set.

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

### HasSubscriptionId

`func (o *UserSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

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

### HasSubscriptionOwnerName

`func (o *UserSubscription) HasSubscriptionOwnerName() bool`

HasSubscriptionOwnerName returns a boolean if a field has been set.

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

### HasUserId

`func (o *UserSubscription) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


