# UsagePerSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **string** | End timestamp of usage | [optional] 
**EnvironmentID** | Pointer to **string** | ID of a Service Environment | [optional] 
**EnvironmentName** | Pointer to **string** | The name of the environment the subscription is for | [optional] 
**EnvironmentType** | Pointer to **string** | The type of the environment the subscription is for | [optional] 
**OrganizationID** | Pointer to **string** | ID of an Org | [optional] 
**OrganizationName** | Pointer to **string** | The name of the organization the subscription belongs to | [optional] 
**ProductTierID** | Pointer to **string** | ID of a Product Tier | [optional] 
**ProductTierName** | Pointer to **string** | The name of the product tier the subscription is for | [optional] 
**ServiceID** | Pointer to **string** | ID of a Service | [optional] 
**ServiceName** | Pointer to **string** | The name of the service the subscription is for | [optional] 
**StartTime** | Pointer to **string** | Start timestamp of usage | [optional] 
**SubscriptionID** | Pointer to **string** | ID of a Subscription | [optional] 
**Usage** | Pointer to [**[]UsagePerDimension**](UsagePerDimension.md) | Usage per dimension for the subscription | [optional] 
**UserEmail** | Pointer to **string** | The email of the owner of the subscription | [optional] 
**UserID** | Pointer to **string** | ID of a User | [optional] 

## Methods

### NewUsagePerSubscription

`func NewUsagePerSubscription() *UsagePerSubscription`

NewUsagePerSubscription instantiates a new UsagePerSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagePerSubscriptionWithDefaults

`func NewUsagePerSubscriptionWithDefaults() *UsagePerSubscription`

NewUsagePerSubscriptionWithDefaults instantiates a new UsagePerSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *UsagePerSubscription) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *UsagePerSubscription) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *UsagePerSubscription) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *UsagePerSubscription) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEnvironmentID

`func (o *UsagePerSubscription) GetEnvironmentID() string`

GetEnvironmentID returns the EnvironmentID field if non-nil, zero value otherwise.

### GetEnvironmentIDOk

`func (o *UsagePerSubscription) GetEnvironmentIDOk() (*string, bool)`

GetEnvironmentIDOk returns a tuple with the EnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentID

`func (o *UsagePerSubscription) SetEnvironmentID(v string)`

SetEnvironmentID sets EnvironmentID field to given value.

### HasEnvironmentID

`func (o *UsagePerSubscription) HasEnvironmentID() bool`

HasEnvironmentID returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *UsagePerSubscription) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *UsagePerSubscription) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *UsagePerSubscription) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *UsagePerSubscription) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *UsagePerSubscription) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *UsagePerSubscription) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *UsagePerSubscription) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *UsagePerSubscription) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetOrganizationID

`func (o *UsagePerSubscription) GetOrganizationID() string`

GetOrganizationID returns the OrganizationID field if non-nil, zero value otherwise.

### GetOrganizationIDOk

`func (o *UsagePerSubscription) GetOrganizationIDOk() (*string, bool)`

GetOrganizationIDOk returns a tuple with the OrganizationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationID

`func (o *UsagePerSubscription) SetOrganizationID(v string)`

SetOrganizationID sets OrganizationID field to given value.

### HasOrganizationID

`func (o *UsagePerSubscription) HasOrganizationID() bool`

HasOrganizationID returns a boolean if a field has been set.

### GetOrganizationName

`func (o *UsagePerSubscription) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *UsagePerSubscription) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *UsagePerSubscription) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *UsagePerSubscription) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetProductTierID

`func (o *UsagePerSubscription) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *UsagePerSubscription) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *UsagePerSubscription) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *UsagePerSubscription) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetProductTierName

`func (o *UsagePerSubscription) GetProductTierName() string`

GetProductTierName returns the ProductTierName field if non-nil, zero value otherwise.

### GetProductTierNameOk

`func (o *UsagePerSubscription) GetProductTierNameOk() (*string, bool)`

GetProductTierNameOk returns a tuple with the ProductTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierName

`func (o *UsagePerSubscription) SetProductTierName(v string)`

SetProductTierName sets ProductTierName field to given value.

### HasProductTierName

`func (o *UsagePerSubscription) HasProductTierName() bool`

HasProductTierName returns a boolean if a field has been set.

### GetServiceID

`func (o *UsagePerSubscription) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *UsagePerSubscription) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *UsagePerSubscription) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *UsagePerSubscription) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetServiceName

`func (o *UsagePerSubscription) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *UsagePerSubscription) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *UsagePerSubscription) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *UsagePerSubscription) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetStartTime

`func (o *UsagePerSubscription) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UsagePerSubscription) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UsagePerSubscription) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UsagePerSubscription) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSubscriptionID

`func (o *UsagePerSubscription) GetSubscriptionID() string`

GetSubscriptionID returns the SubscriptionID field if non-nil, zero value otherwise.

### GetSubscriptionIDOk

`func (o *UsagePerSubscription) GetSubscriptionIDOk() (*string, bool)`

GetSubscriptionIDOk returns a tuple with the SubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionID

`func (o *UsagePerSubscription) SetSubscriptionID(v string)`

SetSubscriptionID sets SubscriptionID field to given value.

### HasSubscriptionID

`func (o *UsagePerSubscription) HasSubscriptionID() bool`

HasSubscriptionID returns a boolean if a field has been set.

### GetUsage

`func (o *UsagePerSubscription) GetUsage() []UsagePerDimension`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsagePerSubscription) GetUsageOk() (*[]UsagePerDimension, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsagePerSubscription) SetUsage(v []UsagePerDimension)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsagePerSubscription) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetUserEmail

`func (o *UsagePerSubscription) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *UsagePerSubscription) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *UsagePerSubscription) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *UsagePerSubscription) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### GetUserID

`func (o *UsagePerSubscription) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *UsagePerSubscription) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *UsagePerSubscription) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *UsagePerSubscription) HasUserID() bool`

HasUserID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


