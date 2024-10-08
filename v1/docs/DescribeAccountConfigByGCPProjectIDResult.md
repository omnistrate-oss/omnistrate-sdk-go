# DescribeAccountConfigByGCPProjectIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | Cloud Provider ID to operate on | 
**Description** | **string** | The description for the account | 
**GcpProjectID** | **string** | The GCP project ID | 
**GcpProjectNumber** | **string** | The GCP project number | 
**GcpServiceAccountEmail** | **string** | The GCP service account email | 
**Id** | **string** | Account Config ID to operate on | 
**Name** | **string** | The name of the account | 
**Status** | **string** | The status of the account | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigByGCPProjectIDResult

`func NewDescribeAccountConfigByGCPProjectIDResult(cloudProviderId string, description string, gcpProjectID string, gcpProjectNumber string, gcpServiceAccountEmail string, id string, name string, status string, statusMessage string, ) *DescribeAccountConfigByGCPProjectIDResult`

NewDescribeAccountConfigByGCPProjectIDResult instantiates a new DescribeAccountConfigByGCPProjectIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByGCPProjectIDResultWithDefaults

`func NewDescribeAccountConfigByGCPProjectIDResultWithDefaults() *DescribeAccountConfigByGCPProjectIDResult`

NewDescribeAccountConfigByGCPProjectIDResultWithDefaults instantiates a new DescribeAccountConfigByGCPProjectIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByoaInstanceIDs

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigByGCPProjectIDResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpProjectID

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.


### GetGcpProjectNumber

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.


### GetGcpServiceAccountEmail

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.


### GetId

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigByGCPProjectIDResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigByGCPProjectIDResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


