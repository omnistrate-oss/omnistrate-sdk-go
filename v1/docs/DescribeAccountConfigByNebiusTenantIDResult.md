# DescribeAccountConfigByNebiusTenantIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**NebiusBindings** | Pointer to [**[]NebiusAccountBindingResult**](NebiusAccountBindingResult.md) | The safe per-region Nebius bindings configured under this tenant account configuration | [optional] 
**NebiusTenantID** | **string** | The Nebius tenant ID for the tenant-scoped Nebius account configuration | 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigByNebiusTenantIDResult

`func NewDescribeAccountConfigByNebiusTenantIDResult(cloudProviderId string, description string, id string, name string, nebiusTenantID string, status string, statusMessage string, ) *DescribeAccountConfigByNebiusTenantIDResult`

NewDescribeAccountConfigByNebiusTenantIDResult instantiates a new DescribeAccountConfigByNebiusTenantIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByNebiusTenantIDResultWithDefaults

`func NewDescribeAccountConfigByNebiusTenantIDResultWithDefaults() *DescribeAccountConfigByNebiusTenantIDResult`

NewDescribeAccountConfigByNebiusTenantIDResultWithDefaults instantiates a new DescribeAccountConfigByNebiusTenantIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByoaInstanceIDs

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigByNebiusTenantIDResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetName(v string)`

SetName sets Name field to given value.


### GetNebiusBindings

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetNebiusBindings() []NebiusAccountBindingResult`

GetNebiusBindings returns the NebiusBindings field if non-nil, zero value otherwise.

### GetNebiusBindingsOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetNebiusBindingsOk() (*[]NebiusAccountBindingResult, bool)`

GetNebiusBindingsOk returns a tuple with the NebiusBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusBindings

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetNebiusBindings(v []NebiusAccountBindingResult)`

SetNebiusBindings sets NebiusBindings field to given value.

### HasNebiusBindings

`func (o *DescribeAccountConfigByNebiusTenantIDResult) HasNebiusBindings() bool`

HasNebiusBindings returns a boolean if a field has been set.

### GetNebiusTenantID

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetNebiusTenantID() string`

GetNebiusTenantID returns the NebiusTenantID field if non-nil, zero value otherwise.

### GetNebiusTenantIDOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetNebiusTenantIDOk() (*string, bool)`

GetNebiusTenantIDOk returns a tuple with the NebiusTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNebiusTenantID

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetNebiusTenantID(v string)`

SetNebiusTenantID sets NebiusTenantID field to given value.


### GetStatus

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigByNebiusTenantIDResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigByNebiusTenantIDResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


