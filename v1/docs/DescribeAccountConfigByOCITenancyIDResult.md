# DescribeAccountConfigByOCITenancyIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**OciDomainID** | **string** | The Domain OCID for Oracle Cloud Infrastructure | 
**OciTenancyID** | **string** | The Tenancy OCID for Oracle Cloud Infrastructure | 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigByOCITenancyIDResult

`func NewDescribeAccountConfigByOCITenancyIDResult(cloudProviderId string, description string, id string, name string, ociDomainID string, ociTenancyID string, status string, statusMessage string, ) *DescribeAccountConfigByOCITenancyIDResult`

NewDescribeAccountConfigByOCITenancyIDResult instantiates a new DescribeAccountConfigByOCITenancyIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByOCITenancyIDResultWithDefaults

`func NewDescribeAccountConfigByOCITenancyIDResultWithDefaults() *DescribeAccountConfigByOCITenancyIDResult`

NewDescribeAccountConfigByOCITenancyIDResultWithDefaults instantiates a new DescribeAccountConfigByOCITenancyIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByoaInstanceIDs

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigByOCITenancyIDResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetName(v string)`

SetName sets Name field to given value.


### GetOciDomainID

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetOciDomainID() string`

GetOciDomainID returns the OciDomainID field if non-nil, zero value otherwise.

### GetOciDomainIDOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetOciDomainIDOk() (*string, bool)`

GetOciDomainIDOk returns a tuple with the OciDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciDomainID

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetOciDomainID(v string)`

SetOciDomainID sets OciDomainID field to given value.


### GetOciTenancyID

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetOciTenancyID() string`

GetOciTenancyID returns the OciTenancyID field if non-nil, zero value otherwise.

### GetOciTenancyIDOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetOciTenancyIDOk() (*string, bool)`

GetOciTenancyIDOk returns a tuple with the OciTenancyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciTenancyID

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetOciTenancyID(v string)`

SetOciTenancyID sets OciTenancyID field to given value.


### GetStatus

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigByOCITenancyIDResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigByOCITenancyIDResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


