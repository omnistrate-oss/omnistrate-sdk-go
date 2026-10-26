# JustInTimeAccessConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalMode** | Pointer to **string** | Provider-environment approval mode. Known values: AUTO, MANUAL. Omission is effectively AUTO; missing configuration remains distinguishable from explicitly configured AUTO. Supported write values are validated by the service. | [optional] 
**ApproverUsers** | Pointer to [**[]JustInTimeAccessApproverUser**](JustInTimeAccessApproverUser.md) | One inline list of active provider users allowed to approve, deny, or revoke, including self-approval if listed. Omission leaves existing membership unchanged; an empty list explicitly clears it only when the effective mode permits it. MANUAL requires a valid nonempty list. | [optional] 

## Methods

### NewJustInTimeAccessConfiguration

`func NewJustInTimeAccessConfiguration() *JustInTimeAccessConfiguration`

NewJustInTimeAccessConfiguration instantiates a new JustInTimeAccessConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJustInTimeAccessConfigurationWithDefaults

`func NewJustInTimeAccessConfigurationWithDefaults() *JustInTimeAccessConfiguration`

NewJustInTimeAccessConfigurationWithDefaults instantiates a new JustInTimeAccessConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalMode

`func (o *JustInTimeAccessConfiguration) GetApprovalMode() string`

GetApprovalMode returns the ApprovalMode field if non-nil, zero value otherwise.

### GetApprovalModeOk

`func (o *JustInTimeAccessConfiguration) GetApprovalModeOk() (*string, bool)`

GetApprovalModeOk returns a tuple with the ApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalMode

`func (o *JustInTimeAccessConfiguration) SetApprovalMode(v string)`

SetApprovalMode sets ApprovalMode field to given value.

### HasApprovalMode

`func (o *JustInTimeAccessConfiguration) HasApprovalMode() bool`

HasApprovalMode returns a boolean if a field has been set.

### GetApproverUsers

`func (o *JustInTimeAccessConfiguration) GetApproverUsers() []JustInTimeAccessApproverUser`

GetApproverUsers returns the ApproverUsers field if non-nil, zero value otherwise.

### GetApproverUsersOk

`func (o *JustInTimeAccessConfiguration) GetApproverUsersOk() (*[]JustInTimeAccessApproverUser, bool)`

GetApproverUsersOk returns a tuple with the ApproverUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverUsers

`func (o *JustInTimeAccessConfiguration) SetApproverUsers(v []JustInTimeAccessApproverUser)`

SetApproverUsers sets ApproverUsers field to given value.

### HasApproverUsers

`func (o *JustInTimeAccessConfiguration) HasApproverUsers() bool`

HasApproverUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


