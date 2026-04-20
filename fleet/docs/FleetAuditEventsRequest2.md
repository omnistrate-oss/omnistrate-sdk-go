# FleetAuditEventsRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingProvider** | Pointer to **string** | The billing provider on the instance&#39;s subscription. Empty when no subscription is linked. | [optional] 
**ExternalPayerId** | Pointer to **string** | The raw external payer ID on the instance&#39;s subscription. | [optional] 
**IpAddress** | Pointer to **string** | The IP address of the client that caused the event (exact match) | [optional] 
**OrgId** | Pointer to **string** | The organization ID of the user that caused the event | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription ID | [optional] 
**UserId** | Pointer to **string** | The ID of the user that caused the event | [optional] 

## Methods

### NewFleetAuditEventsRequest2

`func NewFleetAuditEventsRequest2() *FleetAuditEventsRequest2`

NewFleetAuditEventsRequest2 instantiates a new FleetAuditEventsRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetAuditEventsRequest2WithDefaults

`func NewFleetAuditEventsRequest2WithDefaults() *FleetAuditEventsRequest2`

NewFleetAuditEventsRequest2WithDefaults instantiates a new FleetAuditEventsRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingProvider

`func (o *FleetAuditEventsRequest2) GetBillingProvider() string`

GetBillingProvider returns the BillingProvider field if non-nil, zero value otherwise.

### GetBillingProviderOk

`func (o *FleetAuditEventsRequest2) GetBillingProviderOk() (*string, bool)`

GetBillingProviderOk returns a tuple with the BillingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingProvider

`func (o *FleetAuditEventsRequest2) SetBillingProvider(v string)`

SetBillingProvider sets BillingProvider field to given value.

### HasBillingProvider

`func (o *FleetAuditEventsRequest2) HasBillingProvider() bool`

HasBillingProvider returns a boolean if a field has been set.

### GetExternalPayerId

`func (o *FleetAuditEventsRequest2) GetExternalPayerId() string`

GetExternalPayerId returns the ExternalPayerId field if non-nil, zero value otherwise.

### GetExternalPayerIdOk

`func (o *FleetAuditEventsRequest2) GetExternalPayerIdOk() (*string, bool)`

GetExternalPayerIdOk returns a tuple with the ExternalPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPayerId

`func (o *FleetAuditEventsRequest2) SetExternalPayerId(v string)`

SetExternalPayerId sets ExternalPayerId field to given value.

### HasExternalPayerId

`func (o *FleetAuditEventsRequest2) HasExternalPayerId() bool`

HasExternalPayerId returns a boolean if a field has been set.

### GetIpAddress

`func (o *FleetAuditEventsRequest2) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FleetAuditEventsRequest2) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FleetAuditEventsRequest2) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *FleetAuditEventsRequest2) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetOrgId

`func (o *FleetAuditEventsRequest2) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FleetAuditEventsRequest2) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FleetAuditEventsRequest2) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FleetAuditEventsRequest2) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *FleetAuditEventsRequest2) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *FleetAuditEventsRequest2) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *FleetAuditEventsRequest2) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *FleetAuditEventsRequest2) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUserId

`func (o *FleetAuditEventsRequest2) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FleetAuditEventsRequest2) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FleetAuditEventsRequest2) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FleetAuditEventsRequest2) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


