# DescribeAuditEventResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthMethod** | Pointer to **string** | The authentication method used by the caller for the request that produced the event. One of PASSWORD or API_KEY. Empty for events not originated by an authenticated user (e.g. internal sweeper actions). | [optional] 
**BillingProvider** | Pointer to **string** | The billing provider on the instance&#39;s subscription at the time of the event. Empty when no subscription is linked. | [optional] 
**BillingProviderId** | Pointer to **string** | The resolved billing-provider customer ID of the instance owner at the time of the event. Honors the instance owner-payer override when set; otherwise falls back to the subscription&#39;s external payer ID. Empty when no billing identity has been assigned. | [optional] 
**CloudProvider** | Pointer to **string** | The cloud provider name where the instance is deployed. Derived from the instance&#39;s region; may be empty if unavailable. | [optional] 
**CloudProviderAccountId** | Pointer to **string** | The cloud provider account ID associated with the instance at the time of the event. Derived from the instance&#39;s subscription; may be empty if the instance has no subscription or the subscription has been hard-deleted. | [optional] 
**EnvironmentId** | Pointer to **string** | ID of a Service Environment | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**EventSource** | Pointer to **string** | The event source | [optional] 
**Id** | **string** | ID of a Event | 
**IpAddress** | Pointer to **string** | The IP address of the client that caused the event | [optional] 
**Message** | **string** | Resource Instance of message | 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**OrgName** | Pointer to **string** | The organization name of the user that caused the event | [optional] 
**PlanVersion** | Pointer to **string** | The version of the plan | [optional] 
**ProductTierId** | Pointer to **string** | ID of a Product Tier | [optional] 
**Region** | Pointer to **string** | The region code where the instance is deployed. Derived from the instance&#39;s region; may be empty if unavailable. | [optional] 
**ResourceInstanceId** | **string** | Instance Id of the resource instance | 
**ResourceName** | **string** | Name of the resource | 
**RoleType** | Pointer to **string** | The role asserted by the user&#39;s credential at the time of the event. Captured from the token claims so the audit log reflects the role active at the moment of the action even after the user&#39;s role assignments change. | [optional] 
**ServiceId** | Pointer to **string** | ID of a Service | [optional] 
**ServiceName** | Pointer to **string** | The service name | [optional] 
**ServicePlanName** | Pointer to **string** | The name of the service plan | [optional] 
**SessionId** | Pointer to **string** | The credential identifier used for the request that produced the event. For password-issued JWTs this is the token&#39;s unique ID (jti); for API-key issued JWTs this is the API key&#39;s persistent ID, allowing all events for that key to be correlated even across token refreshes. | [optional] 
**SubscriptionId** | **string** | The subscription ID | 
**Time** | **string** | The event time | 
**UserAgent** | Pointer to **string** | The User-Agent string of the client that caused the event | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 
**UserName** | Pointer to **string** | The user name of the user that caused the event | [optional] 
**WorkflowFailures** | Pointer to [**[]WorkflowFailure**](WorkflowFailure.md) | The list of workflow events that indicate failures | [optional] 

## Methods

### NewDescribeAuditEventResult

`func NewDescribeAuditEventResult(id string, message string, resourceInstanceId string, resourceName string, subscriptionId string, time string, ) *DescribeAuditEventResult`

NewDescribeAuditEventResult instantiates a new DescribeAuditEventResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAuditEventResultWithDefaults

`func NewDescribeAuditEventResultWithDefaults() *DescribeAuditEventResult`

NewDescribeAuditEventResultWithDefaults instantiates a new DescribeAuditEventResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthMethod

`func (o *DescribeAuditEventResult) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *DescribeAuditEventResult) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *DescribeAuditEventResult) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *DescribeAuditEventResult) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetBillingProvider

`func (o *DescribeAuditEventResult) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *DescribeAuditEventResult) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *DescribeAuditEventResult) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *DescribeAuditEventResult) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetBillingProviderId

`func (o *DescribeAuditEventResult) GetBillingProviderId() string`

GetBillingProviderId returns the BillingProviderId field if non-nil, zero value otherwise.

### GetBillingProviderIdOk

`func (o *DescribeAuditEventResult) GetBillingProviderIdOk() (*string, bool)`

GetBillingProviderIdOk returns a tuple with the BillingProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProviderId

`func (o *DescribeAuditEventResult) SetBillingProviderId(v string)`

SetBillingProviderId sets BillingProviderId field to given value.

### HasBillingProviderId

`func (o *DescribeAuditEventResult) HasBillingProviderId() bool`

HasBillingProviderId returns a boolean if a field has been set.

### GetCloudProvider

`func (o *DescribeAuditEventResult) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DescribeAuditEventResult) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DescribeAuditEventResult) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *DescribeAuditEventResult) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCloudProviderAccountId

`func (o *DescribeAuditEventResult) GetCloudProviderAccountId() string`

GetCloudProviderAccountId returns the CloudProviderAccountId field if non-nil, zero value otherwise.

### GetCloudProviderAccountIdOk

`func (o *DescribeAuditEventResult) GetCloudProviderAccountIdOk() (*string, bool)`

GetCloudProviderAccountIdOk returns a tuple with the CloudProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderAccountId

`func (o *DescribeAuditEventResult) SetCloudProviderAccountId(v string)`

SetCloudProviderAccountId sets CloudProviderAccountId field to given value.

### HasCloudProviderAccountId

`func (o *DescribeAuditEventResult) HasCloudProviderAccountId() bool`

HasCloudProviderAccountId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *DescribeAuditEventResult) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *DescribeAuditEventResult) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *DescribeAuditEventResult) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *DescribeAuditEventResult) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *DescribeAuditEventResult) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *DescribeAuditEventResult) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *DescribeAuditEventResult) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *DescribeAuditEventResult) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetEventSource

`func (o *DescribeAuditEventResult) GetEventSource() string`

GetEventSource returns the EventSource field if non-nil, zero value otherwise.

### GetEventSourceOk

`func (o *DescribeAuditEventResult) GetEventSourceOk() (*string, bool)`

GetEventSourceOk returns a tuple with the EventSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSource

`func (o *DescribeAuditEventResult) SetEventSource(v string)`

SetEventSource sets EventSource field to given value.

### HasEventSource

`func (o *DescribeAuditEventResult) HasEventSource() bool`

HasEventSource returns a boolean if a field has been set.

### GetId

`func (o *DescribeAuditEventResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAuditEventResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAuditEventResult) SetId(v string)`

SetId sets Id field to given value.


### GetIpAddress

`func (o *DescribeAuditEventResult) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *DescribeAuditEventResult) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *DescribeAuditEventResult) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *DescribeAuditEventResult) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMessage

`func (o *DescribeAuditEventResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DescribeAuditEventResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DescribeAuditEventResult) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrgId

`func (o *DescribeAuditEventResult) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *DescribeAuditEventResult) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *DescribeAuditEventResult) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *DescribeAuditEventResult) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *DescribeAuditEventResult) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *DescribeAuditEventResult) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *DescribeAuditEventResult) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *DescribeAuditEventResult) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetPlanVersion

`func (o *DescribeAuditEventResult) GetPlanVersion() string`

GetPlanVersion returns the PlanVersion field if non-nil, zero value otherwise.

### GetPlanVersionOk

`func (o *DescribeAuditEventResult) GetPlanVersionOk() (*string, bool)`

GetPlanVersionOk returns a tuple with the PlanVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanVersion

`func (o *DescribeAuditEventResult) SetPlanVersion(v string)`

SetPlanVersion sets PlanVersion field to given value.

### HasPlanVersion

`func (o *DescribeAuditEventResult) HasPlanVersion() bool`

HasPlanVersion returns a boolean if a field has been set.

### GetProductTierId

`func (o *DescribeAuditEventResult) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeAuditEventResult) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeAuditEventResult) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.

### HasProductTierId

`func (o *DescribeAuditEventResult) HasProductTierId() bool`

HasProductTierId returns a boolean if a field has been set.

### GetRegion

`func (o *DescribeAuditEventResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DescribeAuditEventResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DescribeAuditEventResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *DescribeAuditEventResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResourceInstanceId

`func (o *DescribeAuditEventResult) GetResourceInstanceId() string`

GetResourceInstanceId returns the ResourceInstanceId field if non-nil, zero value otherwise.

### GetResourceInstanceIdOk

`func (o *DescribeAuditEventResult) GetResourceInstanceIdOk() (*string, bool)`

GetResourceInstanceIdOk returns a tuple with the ResourceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInstanceId

`func (o *DescribeAuditEventResult) SetResourceInstanceId(v string)`

SetResourceInstanceId sets ResourceInstanceId field to given value.


### GetResourceName

`func (o *DescribeAuditEventResult) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *DescribeAuditEventResult) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *DescribeAuditEventResult) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetRoleType

`func (o *DescribeAuditEventResult) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *DescribeAuditEventResult) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *DescribeAuditEventResult) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *DescribeAuditEventResult) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetServiceId

`func (o *DescribeAuditEventResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeAuditEventResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeAuditEventResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *DescribeAuditEventResult) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *DescribeAuditEventResult) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *DescribeAuditEventResult) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *DescribeAuditEventResult) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *DescribeAuditEventResult) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServicePlanName

`func (o *DescribeAuditEventResult) GetServicePlanName() string`

GetServicePlanName returns the ServicePlanName field if non-nil, zero value otherwise.

### GetServicePlanNameOk

`func (o *DescribeAuditEventResult) GetServicePlanNameOk() (*string, bool)`

GetServicePlanNameOk returns a tuple with the ServicePlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanName

`func (o *DescribeAuditEventResult) SetServicePlanName(v string)`

SetServicePlanName sets ServicePlanName field to given value.

### HasServicePlanName

`func (o *DescribeAuditEventResult) HasServicePlanName() bool`

HasServicePlanName returns a boolean if a field has been set.

### GetSessionId

`func (o *DescribeAuditEventResult) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DescribeAuditEventResult) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DescribeAuditEventResult) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DescribeAuditEventResult) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *DescribeAuditEventResult) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeAuditEventResult) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeAuditEventResult) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetTime

`func (o *DescribeAuditEventResult) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *DescribeAuditEventResult) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *DescribeAuditEventResult) SetTime(v string)`

SetTime sets Time field to given value.


### GetUserAgent

`func (o *DescribeAuditEventResult) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *DescribeAuditEventResult) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *DescribeAuditEventResult) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *DescribeAuditEventResult) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeAuditEventResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeAuditEventResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeAuditEventResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DescribeAuditEventResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *DescribeAuditEventResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *DescribeAuditEventResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *DescribeAuditEventResult) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *DescribeAuditEventResult) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetWorkflowFailures

`func (o *DescribeAuditEventResult) GetWorkflowFailures() []WorkflowFailure`

GetWorkflowFailures returns the WorkflowFailures field if non-nil, zero value otherwise.

### GetWorkflowFailuresOk

`func (o *DescribeAuditEventResult) GetWorkflowFailuresOk() (*[]WorkflowFailure, bool)`

GetWorkflowFailuresOk returns a tuple with the WorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFailures

`func (o *DescribeAuditEventResult) SetWorkflowFailures(v []WorkflowFailure)`

SetWorkflowFailures sets WorkflowFailures field to given value.

### HasWorkflowFailures

`func (o *DescribeAuditEventResult) HasWorkflowFailures() bool`

HasWorkflowFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


