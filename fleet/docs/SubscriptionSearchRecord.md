# SubscriptionSearchRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Subscription ID. | 
**OrgID** | **string** | The organization ID of the subscription root user. | 
**ProductTierID** | **string** | The product tier ID of the subscription. | 
**RootUserEmail** | **string** | The root user email of the subscription. | 
**RootUserID** | **string** | The root user ID of the subscription. | 
**RootUserName** | **string** | The root user name of the subscription. | 
**ServiceEnvironmentID** | **string** | The service environment I of the subscriptionD. | 
**ServiceEnvironmentName** | **string** | The service environment name of the subscription. | 
**ServiceEnvironmentType** | Pointer to **string** | The service environment type of the subscription. | [optional] 
**ServiceID** | **string** | The service ID of the subscription. | 
**ServiceName** | **string** | The service name of the subscription. | 
**ServicePlanName** | **string** | The service plan name of the subscription. | 
**Status** | **string** | The subscription status. | 

## Methods

### NewSubscriptionSearchRecord

`func NewSubscriptionSearchRecord(id string, orgID string, productTierID string, rootUserEmail string, rootUserID string, rootUserName string, serviceEnvironmentID string, serviceEnvironmentName string, serviceID string, serviceName string, servicePlanName string, status string, ) *SubscriptionSearchRecord`

NewSubscriptionSearchRecord instantiates a new SubscriptionSearchRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionSearchRecordWithDefaults

`func NewSubscriptionSearchRecordWithDefaults() *SubscriptionSearchRecord`

NewSubscriptionSearchRecordWithDefaults instantiates a new SubscriptionSearchRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionSearchRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionSearchRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionSearchRecord) SetId(v string)`

SetId sets Id field to given value.


### GetOrgID

`func (o *SubscriptionSearchRecord) GetOrgID() string`

GetOrgID returns the OrgID field if non-nil, zero value otherwise.

### GetOrgIDOk

`func (o *SubscriptionSearchRecord) GetOrgIDOk() (*string, bool)`

GetOrgIDOk returns a tuple with the OrgID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgID

`func (o *SubscriptionSearchRecord) SetOrgID(v string)`

SetOrgID sets OrgID field to given value.


### GetProductTierID

`func (o *SubscriptionSearchRecord) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *SubscriptionSearchRecord) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *SubscriptionSearchRecord) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.


### GetRootUserEmail

`func (o *SubscriptionSearchRecord) GetRootUserEmail() string`

GetRootUserEmail returns the RootUserEmail field if non-nil, zero value otherwise.

### GetRootUserEmailOk

`func (o *SubscriptionSearchRecord) GetRootUserEmailOk() (*string, bool)`

GetRootUserEmailOk returns a tuple with the RootUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserEmail

`func (o *SubscriptionSearchRecord) SetRootUserEmail(v string)`

SetRootUserEmail sets RootUserEmail field to given value.


### GetRootUserID

`func (o *SubscriptionSearchRecord) GetRootUserID() string`

GetRootUserID returns the RootUserID field if non-nil, zero value otherwise.

### GetRootUserIDOk

`func (o *SubscriptionSearchRecord) GetRootUserIDOk() (*string, bool)`

GetRootUserIDOk returns a tuple with the RootUserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserID

`func (o *SubscriptionSearchRecord) SetRootUserID(v string)`

SetRootUserID sets RootUserID field to given value.


### GetRootUserName

`func (o *SubscriptionSearchRecord) GetRootUserName() string`

GetRootUserName returns the RootUserName field if non-nil, zero value otherwise.

### GetRootUserNameOk

`func (o *SubscriptionSearchRecord) GetRootUserNameOk() (*string, bool)`

GetRootUserNameOk returns a tuple with the RootUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUserName

`func (o *SubscriptionSearchRecord) SetRootUserName(v string)`

SetRootUserName sets RootUserName field to given value.


### GetServiceEnvironmentID

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentID() string`

GetServiceEnvironmentID returns the ServiceEnvironmentID field if non-nil, zero value otherwise.

### GetServiceEnvironmentIDOk

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentIDOk() (*string, bool)`

GetServiceEnvironmentIDOk returns a tuple with the ServiceEnvironmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentID

`func (o *SubscriptionSearchRecord) SetServiceEnvironmentID(v string)`

SetServiceEnvironmentID sets ServiceEnvironmentID field to given value.


### GetServiceEnvironmentName

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentName() string`

GetServiceEnvironmentName returns the ServiceEnvironmentName field if non-nil, zero value otherwise.

### GetServiceEnvironmentNameOk

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentNameOk() (*string, bool)`

GetServiceEnvironmentNameOk returns a tuple with the ServiceEnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentName

`func (o *SubscriptionSearchRecord) SetServiceEnvironmentName(v string)`

SetServiceEnvironmentName sets ServiceEnvironmentName field to given value.


### GetServiceEnvironmentType

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentType() string`

GetServiceEnvironmentType returns the ServiceEnvironmentType field if non-nil, zero value otherwise.

### GetServiceEnvironmentTypeOk

`func (o *SubscriptionSearchRecord) GetServiceEnvironmentTypeOk() (*string, bool)`

GetServiceEnvironmentTypeOk returns a tuple with the ServiceEnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnvironmentType

`func (o *SubscriptionSearchRecord) SetServiceEnvironmentType(v string)`

SetServiceEnvironmentType sets ServiceEnvironmentType field to given value.

### HasServiceEnvironmentType

`func (o *SubscriptionSearchRecord) HasServiceEnvironmentType() bool`

HasServiceEnvironmentType returns a boolean if a field has been set.

### GetServiceID

`func (o *SubscriptionSearchRecord) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *SubscriptionSearchRecord) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *SubscriptionSearchRecord) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.


### GetServiceName

`func (o *SubscriptionSearchRecord) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *SubscriptionSearchRecord) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *SubscriptionSearchRecord) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetServicePlanName

`func (o *SubscriptionSearchRecord) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *SubscriptionSearchRecord) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *SubscriptionSearchRecord) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.


### GetStatus

`func (o *SubscriptionSearchRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionSearchRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionSearchRecord) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


