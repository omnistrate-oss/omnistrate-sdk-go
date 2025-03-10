# DescribeAccountConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AwsBootstrapRoleARN** | Pointer to **string** | The security role ARN or service account ARN that grants access to operate the infra | [optional] 
**AwsCloudFormationNoLBTemplateURL** | Pointer to **string** | The URL to the CloudFormation template (no LoadBalancer policy version) | [optional] 
**AwsCloudFormationTemplateURL** | Pointer to **string** | The URL to the CloudFormation template | [optional] 
**AzureBootstrapShellCommand** | Pointer to **string** | The Azure bootstrap shell command | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**AzureTenantID** | Pointer to **string** | The Azure tenant ID | [optional] 
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**GcpBootstrapShellCommand** | Pointer to **string** | The GCP bootstrap shell command | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**GcpProjectNumber** | Pointer to **string** | The GCP project number | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** | The GCP service account email | [optional] 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigResult

`func NewDescribeAccountConfigResult(cloudProviderId string, description string, id string, name string, status string, statusMessage string, ) *DescribeAccountConfigResult`

NewDescribeAccountConfigResult instantiates a new DescribeAccountConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigResultWithDefaults

`func NewDescribeAccountConfigResultWithDefaults() *DescribeAccountConfigResult`

NewDescribeAccountConfigResultWithDefaults instantiates a new DescribeAccountConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *DescribeAccountConfigResult) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *DescribeAccountConfigResult) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *DescribeAccountConfigResult) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *DescribeAccountConfigResult) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsBootstrapRoleARN

`func (o *DescribeAccountConfigResult) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *DescribeAccountConfigResult) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *DescribeAccountConfigResult) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.

### HasAwsBootstrapRoleARN

`func (o *DescribeAccountConfigResult) HasAwsBootstrapRoleARN() bool`

HasAwsBootstrapRoleARN returns a boolean if a field has been set.

### GetAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURL() string`

GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationNoLBTemplateURLOk

`func (o *DescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool)`

GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigResult) SetAwsCloudFormationNoLBTemplateURL(v string)`

SetAwsCloudFormationNoLBTemplateURL sets AwsCloudFormationNoLBTemplateURL field to given value.

### HasAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigResult) HasAwsCloudFormationNoLBTemplateURL() bool`

HasAwsCloudFormationNoLBTemplateURL returns a boolean if a field has been set.

### GetAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigResult) GetAwsCloudFormationTemplateURL() string`

GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationTemplateURLOk

`func (o *DescribeAccountConfigResult) GetAwsCloudFormationTemplateURLOk() (*string, bool)`

GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigResult) SetAwsCloudFormationTemplateURL(v string)`

SetAwsCloudFormationTemplateURL sets AwsCloudFormationTemplateURL field to given value.

### HasAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigResult) HasAwsCloudFormationTemplateURL() bool`

HasAwsCloudFormationTemplateURL returns a boolean if a field has been set.

### GetAzureBootstrapShellCommand

`func (o *DescribeAccountConfigResult) GetAzureBootstrapShellCommand() string`

GetAzureBootstrapShellCommand returns the AzureBootstrapShellCommand field if non-nil, zero value otherwise.

### GetAzureBootstrapShellCommandOk

`func (o *DescribeAccountConfigResult) GetAzureBootstrapShellCommandOk() (*string, bool)`

GetAzureBootstrapShellCommandOk returns a tuple with the AzureBootstrapShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBootstrapShellCommand

`func (o *DescribeAccountConfigResult) SetAzureBootstrapShellCommand(v string)`

SetAzureBootstrapShellCommand sets AzureBootstrapShellCommand field to given value.

### HasAzureBootstrapShellCommand

`func (o *DescribeAccountConfigResult) HasAzureBootstrapShellCommand() bool`

HasAzureBootstrapShellCommand returns a boolean if a field has been set.

### GetAzureSubscriptionID

`func (o *DescribeAccountConfigResult) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *DescribeAccountConfigResult) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *DescribeAccountConfigResult) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *DescribeAccountConfigResult) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

### GetAzureTenantID

`func (o *DescribeAccountConfigResult) GetAzureTenantID() string`

GetAzureTenantID returns the AzureTenantID field if non-nil, zero value otherwise.

### GetAzureTenantIDOk

`func (o *DescribeAccountConfigResult) GetAzureTenantIDOk() (*string, bool)`

GetAzureTenantIDOk returns a tuple with the AzureTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantID

`func (o *DescribeAccountConfigResult) SetAzureTenantID(v string)`

SetAzureTenantID sets AzureTenantID field to given value.

### HasAzureTenantID

`func (o *DescribeAccountConfigResult) HasAzureTenantID() bool`

HasAzureTenantID returns a boolean if a field has been set.

### GetByoaInstanceIDs

`func (o *DescribeAccountConfigResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpBootstrapShellCommand

`func (o *DescribeAccountConfigResult) GetGcpBootstrapShellCommand() string`

GetGcpBootstrapShellCommand returns the GcpBootstrapShellCommand field if non-nil, zero value otherwise.

### GetGcpBootstrapShellCommandOk

`func (o *DescribeAccountConfigResult) GetGcpBootstrapShellCommandOk() (*string, bool)`

GetGcpBootstrapShellCommandOk returns a tuple with the GcpBootstrapShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBootstrapShellCommand

`func (o *DescribeAccountConfigResult) SetGcpBootstrapShellCommand(v string)`

SetGcpBootstrapShellCommand sets GcpBootstrapShellCommand field to given value.

### HasGcpBootstrapShellCommand

`func (o *DescribeAccountConfigResult) HasGcpBootstrapShellCommand() bool`

HasGcpBootstrapShellCommand returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *DescribeAccountConfigResult) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *DescribeAccountConfigResult) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *DescribeAccountConfigResult) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *DescribeAccountConfigResult) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *DescribeAccountConfigResult) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *DescribeAccountConfigResult) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *DescribeAccountConfigResult) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *DescribeAccountConfigResult) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *DescribeAccountConfigResult) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *DescribeAccountConfigResult) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *DescribeAccountConfigResult) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *DescribeAccountConfigResult) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetId

`func (o *DescribeAccountConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DescribeAccountConfigResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


