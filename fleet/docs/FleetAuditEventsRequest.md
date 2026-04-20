# FleetAuditEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProvider** | Pointer to **string** | The billing provider on the instance&#39;s subscription. Empty when no subscription is linked. | [optional] 
**EndDate** | Pointer to **time.Time** | End time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**EnvironmentType** | Pointer to **string** | The type of service environment | [optional] 
**EventSourceTypes** | Pointer to **[]string** | The event types to filter by | [optional] 
**ExcludeWorkflowFailures** | Pointer to **bool** | Whether to exclude workflow failure details from the response. | [optional] 
**ExternalPayerId** | Pointer to **string** | The raw external payer ID on the instance&#39;s subscription. | [optional] 
**InstanceID** | Pointer to **string** | ID of a Resource Instance | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the client that caused the event (exact match) | [optional] 
**NextPageToken** | Pointer to **string** | The next token to use for pagination | [optional] 
**OrgId** | Pointer to **string** | ID of an Org | [optional] 
**PageSize** | Pointer to **int64** | The number of resources to return per page | [optional] 
**ProductTierID** | Pointer to **string** | ID of a Product Tier | [optional] 
**ServiceID** | Pointer to **string** | ID of a Service | [optional] 
**StartDate** | Pointer to **time.Time** | Start time of the window in RFC 3339 format. If omitted, the filter is open-ended at the start. | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**UserId** | Pointer to **string** | ID of a User | [optional] 

## Methods

### NewFleetAuditEventsRequest

`func NewFleetAuditEventsRequest(token string, ) *FleetAuditEventsRequest`

NewFleetAuditEventsRequest instantiates a new FleetAuditEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAuditEventsRequestWithDefaults

`func NewFleetAuditEventsRequestWithDefaults() *FleetAuditEventsRequest`

NewFleetAuditEventsRequestWithDefaults instantiates a new FleetAuditEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProvider

`func (o *FleetAuditEventsRequest) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetAuditEventsRequest) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetAuditEventsRequest) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetAuditEventsRequest) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetEndDate

`func (o *FleetAuditEventsRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *FleetAuditEventsRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *FleetAuditEventsRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *FleetAuditEventsRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *FleetAuditEventsRequest) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *FleetAuditEventsRequest) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *FleetAuditEventsRequest) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *FleetAuditEventsRequest) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetEventSourceTypes

`func (o *FleetAuditEventsRequest) GetEventSourceTypes() []string`

GetEventSourceTypes returns the EventSourceTypes field if non-nil, zero value otherwise.

### GetEventSourceTypesOk

`func (o *FleetAuditEventsRequest) GetEventSourceTypesOk() (*[]string, bool)`

GetEventSourceTypesOk returns a tuple with the EventSourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceTypes

`func (o *FleetAuditEventsRequest) SetEventSourceTypes(v []string)`

SetEventSourceTypes sets EventSourceTypes field to given value.

### HasEventSourceTypes

`func (o *FleetAuditEventsRequest) HasEventSourceTypes() bool`

HasEventSourceTypes returns a boolean if a field has been set.

### GetExcludeWorkflowFailures

`func (o *FleetAuditEventsRequest) GetExcludeWorkflowFailures() bool`

GetExcludeWorkflowFailures returns the ExcludeWorkflowFailures field if non-nil, zero value otherwise.

### GetExcludeWorkflowFailuresOk

`func (o *FleetAuditEventsRequest) GetExcludeWorkflowFailuresOk() (*bool, bool)`

GetExcludeWorkflowFailuresOk returns a tuple with the ExcludeWorkflowFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeWorkflowFailures

`func (o *FleetAuditEventsRequest) SetExcludeWorkflowFailures(v bool)`

SetExcludeWorkflowFailures sets ExcludeWorkflowFailures field to given value.

### HasExcludeWorkflowFailures

`func (o *FleetAuditEventsRequest) HasExcludeWorkflowFailures() bool`

HasExcludeWorkflowFailures returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *FleetAuditEventsRequest) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetAuditEventsRequest) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetAuditEventsRequest) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetAuditEventsRequest) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetInstanceID

`func (o *FleetAuditEventsRequest) GetInstanceID() string`

GetInstanceID returns the InstanceID field if non-nil, zero value otherwise.

### GetInstanceIDOk

`func (o *FleetAuditEventsRequest) GetInstanceIDOk() (*string, bool)`

GetInstanceIDOk returns a tuple with the InstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceID

`func (o *FleetAuditEventsRequest) SetInstanceID(v string)`

SetInstanceID sets InstanceID field to given value.

### HasInstanceID

`func (o *FleetAuditEventsRequest) HasInstanceID() bool`

HasInstanceID returns a boolean if a field has been set.

### GetIpAddress

`func (o *FleetAuditEventsRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FleetAuditEventsRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FleetAuditEventsRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *FleetAuditEventsRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetNextPageToken

`func (o *FleetAuditEventsRequest) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FleetAuditEventsRequest) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FleetAuditEventsRequest) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FleetAuditEventsRequest) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetOrgId

`func (o *FleetAuditEventsRequest) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FleetAuditEventsRequest) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FleetAuditEventsRequest) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FleetAuditEventsRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPageSize

`func (o *FleetAuditEventsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FleetAuditEventsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FleetAuditEventsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FleetAuditEventsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProductTierID

`func (o *FleetAuditEventsRequest) GetProductTierID() string`

GetProductTierID returns the ProductTierID field if non-nil, zero value otherwise.

### GetProductTierIDOk

`func (o *FleetAuditEventsRequest) GetProductTierIDOk() (*string, bool)`

GetProductTierIDOk returns a tuple with the ProductTierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierID

`func (o *FleetAuditEventsRequest) SetProductTierID(v string)`

SetProductTierID sets ProductTierID field to given value.

### HasProductTierID

`func (o *FleetAuditEventsRequest) HasProductTierID() bool`

HasProductTierID returns a boolean if a field has been set.

### GetServiceID

`func (o *FleetAuditEventsRequest) GetServiceID() string`

GetServiceID returns the ServiceID field if non-nil, zero value otherwise.

### GetServiceIDOk

`func (o *FleetAuditEventsRequest) GetServiceIDOk() (*string, bool)`

GetServiceIDOk returns a tuple with the ServiceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceID

`func (o *FleetAuditEventsRequest) SetServiceID(v string)`

SetServiceID sets ServiceID field to given value.

### HasServiceID

`func (o *FleetAuditEventsRequest) HasServiceID() bool`

HasServiceID returns a boolean if a field has been set.

### GetStartDate

`func (o *FleetAuditEventsRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *FleetAuditEventsRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *FleetAuditEventsRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *FleetAuditEventsRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FleetAuditEventsRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetAuditEventsRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetAuditEventsRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetAuditEventsRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetToken

`func (o *FleetAuditEventsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetAuditEventsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetAuditEventsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUserId

`func (o *FleetAuditEventsRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetAuditEventsRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetAuditEventsRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FleetAuditEventsRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


