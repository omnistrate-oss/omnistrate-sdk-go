# DescribeAccountConfigByAzureSubscriptionIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionID** | **string** | The Azure subscription ID | 
**AzureTenantID** | **string** | The Azure tenant ID | 
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigByAzureSubscriptionIDResult

`func NewDescribeAccountConfigByAzureSubscriptionIDResult(azureSubscriptionID string, azureTenantID string, cloudProviderId string, description string, id string, name string, status string, statusMessage string, ) *DescribeAccountConfigByAzureSubscriptionIDResult`

NewDescribeAccountConfigByAzureSubscriptionIDResult instantiates a new DescribeAccountConfigByAzureSubscriptionIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByAzureSubscriptionIDResultWithDefaults

`func NewDescribeAccountConfigByAzureSubscriptionIDResultWithDefaults() *DescribeAccountConfigByAzureSubscriptionIDResult`

NewDescribeAccountConfigByAzureSubscriptionIDResultWithDefaults instantiates a new DescribeAccountConfigByAzureSubscriptionIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureSubscriptionID

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.


### GetAzureTenantID

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetAzureTenantID() string`

GetAzureTenantID returns the AzureTenantID field if non-nil, zero value otherwise.

### GetAzureTenantIDOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetAzureTenantIDOk() (*string, bool)`

GetAzureTenantIDOk returns a tuple with the AzureTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantID

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetAzureTenantID(v string)`

SetAzureTenantID sets AzureTenantID field to given value.


### GetByoaInstanceIDs

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigByAzureSubscriptionIDResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


