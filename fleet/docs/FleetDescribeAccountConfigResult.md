# FleetDescribeAccountConfigResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AwsBootstrapRoleARN** | Pointer to **string** | The security role ARN or service account ARN that grants access to operate the infra | [optional] 
**AwsCloudFormationNoLBTemplateURL** | Pointer to **string** | The URL to the CloudFormation template (no LoadBalancer policy version) | [optional] 
**AwsCloudFormationTemplateURL** | Pointer to **string** | The URL to the CloudFormation template | [optional] 
**AzureBootstrapShellCommand** | Pointer to **string** | The Azure bootstrap shell command | [optional] 
**AzureDisconnectShellCommand** | Pointer to **string** | The Azure disconnect shell command | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**AzureTenantID** | Pointer to **string** | The Azure tenant ID | [optional] 
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**GcpBootstrapShellCommand** | Pointer to **string** | The GCP bootstrap shell command | [optional] 
**GcpDisconnectShellCommand** | Pointer to **string** | The GCP disconnect shell command | [optional] 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**GcpProjectNumber** | Pointer to **string** | The GCP project number | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** | The GCP service account email | [optional] 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**OciBootstrapShellCommand** | Pointer to **string** | The Azure bootstrap shell command | [optional] 
**OciDisconnectShellCommand** | Pointer to **string** | The Azure disconnect shell command | [optional] 
**OciDomainID** | Pointer to **string** | The Domain OCID for Oracle Cloud Infrastructure | [optional] 
**OciTenancyID** | Pointer to **string** | The Tenancy OCID for Oracle Cloud Infrastructure | [optional] 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewFleetDescribeAccountConfigResult

`func NewFleetDescribeAccountConfigResult(cloudProviderId string, description string, id string, name string, status string, statusMessage string, ) *FleetDescribeAccountConfigResult`

NewFleetDescribeAccountConfigResult instantiates a new FleetDescribeAccountConfigResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetDescribeAccountConfigResultWithDefaults

`func NewFleetDescribeAccountConfigResultWithDefaults() *FleetDescribeAccountConfigResult`

NewFleetDescribeAccountConfigResultWithDefaults instantiates a new FleetDescribeAccountConfigResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *FleetDescribeAccountConfigResult) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *FleetDescribeAccountConfigResult) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *FleetDescribeAccountConfigResult) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *FleetDescribeAccountConfigResult) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsBootstrapRoleARN

`func (o *FleetDescribeAccountConfigResult) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *FleetDescribeAccountConfigResult) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *FleetDescribeAccountConfigResult) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.

### HasAwsBootstrapRoleARN

`func (o *FleetDescribeAccountConfigResult) HasAwsBootstrapRoleARN() bool`

HasAwsBootstrapRoleARN returns a boolean if a field has been set.

### GetAwsCloudFormationNoLBTemplateURL

`func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURL() string`

GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationNoLBTemplateURLOk

`func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool)`

GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationNoLBTemplateURL

`func (o *FleetDescribeAccountConfigResult) SetAwsCloudFormationNoLBTemplateURL(v string)`

SetAwsCloudFormationNoLBTemplateURL sets AwsCloudFormationNoLBTemplateURL field to given value.

### HasAwsCloudFormationNoLBTemplateURL

`func (o *FleetDescribeAccountConfigResult) HasAwsCloudFormationNoLBTemplateURL() bool`

HasAwsCloudFormationNoLBTemplateURL returns a boolean if a field has been set.

### GetAwsCloudFormationTemplateURL

`func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationTemplateURL() string`

GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationTemplateURLOk

`func (o *FleetDescribeAccountConfigResult) GetAwsCloudFormationTemplateURLOk() (*string, bool)`

GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationTemplateURL

`func (o *FleetDescribeAccountConfigResult) SetAwsCloudFormationTemplateURL(v string)`

SetAwsCloudFormationTemplateURL sets AwsCloudFormationTemplateURL field to given value.

### HasAwsCloudFormationTemplateURL

`func (o *FleetDescribeAccountConfigResult) HasAwsCloudFormationTemplateURL() bool`

HasAwsCloudFormationTemplateURL returns a boolean if a field has been set.

### GetAzureBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) GetAzureBootstrapShellCommand() string`

GetAzureBootstrapShellCommand returns the AzureBootstrapShellCommand field if non-nil, zero value otherwise.

### GetAzureBootstrapShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetAzureBootstrapShellCommandOk() (*string, bool)`

GetAzureBootstrapShellCommandOk returns a tuple with the AzureBootstrapShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) SetAzureBootstrapShellCommand(v string)`

SetAzureBootstrapShellCommand sets AzureBootstrapShellCommand field to given value.

### HasAzureBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) HasAzureBootstrapShellCommand() bool`

HasAzureBootstrapShellCommand returns a boolean if a field has been set.

### GetAzureDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) GetAzureDisconnectShellCommand() string`

GetAzureDisconnectShellCommand returns the AzureDisconnectShellCommand field if non-nil, zero value otherwise.

### GetAzureDisconnectShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetAzureDisconnectShellCommandOk() (*string, bool)`

GetAzureDisconnectShellCommandOk returns a tuple with the AzureDisconnectShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) SetAzureDisconnectShellCommand(v string)`

SetAzureDisconnectShellCommand sets AzureDisconnectShellCommand field to given value.

### HasAzureDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) HasAzureDisconnectShellCommand() bool`

HasAzureDisconnectShellCommand returns a boolean if a field has been set.

### GetAzureSubscriptionID

`func (o *FleetDescribeAccountConfigResult) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *FleetDescribeAccountConfigResult) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *FleetDescribeAccountConfigResult) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *FleetDescribeAccountConfigResult) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

### GetAzureTenantID

`func (o *FleetDescribeAccountConfigResult) GetAzureTenantID() string`

GetAzureTenantID returns the AzureTenantID field if non-nil, zero value otherwise.

### GetAzureTenantIDOk

`func (o *FleetDescribeAccountConfigResult) GetAzureTenantIDOk() (*string, bool)`

GetAzureTenantIDOk returns a tuple with the AzureTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantID

`func (o *FleetDescribeAccountConfigResult) SetAzureTenantID(v string)`

SetAzureTenantID sets AzureTenantID field to given value.

### HasAzureTenantID

`func (o *FleetDescribeAccountConfigResult) HasAzureTenantID() bool`

HasAzureTenantID returns a boolean if a field has been set.

### GetByoaInstanceIDs

`func (o *FleetDescribeAccountConfigResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *FleetDescribeAccountConfigResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *FleetDescribeAccountConfigResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *FleetDescribeAccountConfigResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *FleetDescribeAccountConfigResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *FleetDescribeAccountConfigResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *FleetDescribeAccountConfigResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *FleetDescribeAccountConfigResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FleetDescribeAccountConfigResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FleetDescribeAccountConfigResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) GetGcpBootstrapShellCommand() string`

GetGcpBootstrapShellCommand returns the GcpBootstrapShellCommand field if non-nil, zero value otherwise.

### GetGcpBootstrapShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetGcpBootstrapShellCommandOk() (*string, bool)`

GetGcpBootstrapShellCommandOk returns a tuple with the GcpBootstrapShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) SetGcpBootstrapShellCommand(v string)`

SetGcpBootstrapShellCommand sets GcpBootstrapShellCommand field to given value.

### HasGcpBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) HasGcpBootstrapShellCommand() bool`

HasGcpBootstrapShellCommand returns a boolean if a field has been set.

### GetGcpDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) GetGcpDisconnectShellCommand() string`

GetGcpDisconnectShellCommand returns the GcpDisconnectShellCommand field if non-nil, zero value otherwise.

### GetGcpDisconnectShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetGcpDisconnectShellCommandOk() (*string, bool)`

GetGcpDisconnectShellCommandOk returns a tuple with the GcpDisconnectShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) SetGcpDisconnectShellCommand(v string)`

SetGcpDisconnectShellCommand sets GcpDisconnectShellCommand field to given value.

### HasGcpDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) HasGcpDisconnectShellCommand() bool`

HasGcpDisconnectShellCommand returns a boolean if a field has been set.

### GetGcpProjectID

`func (o *FleetDescribeAccountConfigResult) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *FleetDescribeAccountConfigResult) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *FleetDescribeAccountConfigResult) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *FleetDescribeAccountConfigResult) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *FleetDescribeAccountConfigResult) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *FleetDescribeAccountConfigResult) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *FleetDescribeAccountConfigResult) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *FleetDescribeAccountConfigResult) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *FleetDescribeAccountConfigResult) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *FleetDescribeAccountConfigResult) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *FleetDescribeAccountConfigResult) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *FleetDescribeAccountConfigResult) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetId

`func (o *FleetDescribeAccountConfigResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetDescribeAccountConfigResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetDescribeAccountConfigResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *FleetDescribeAccountConfigResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetDescribeAccountConfigResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetDescribeAccountConfigResult) SetName(v string)`

SetName sets Name field to given value.


### GetOciBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) GetOciBootstrapShellCommand() string`

GetOciBootstrapShellCommand returns the OciBootstrapShellCommand field if non-nil, zero value otherwise.

### GetOciBootstrapShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetOciBootstrapShellCommandOk() (*string, bool)`

GetOciBootstrapShellCommandOk returns a tuple with the OciBootstrapShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) SetOciBootstrapShellCommand(v string)`

SetOciBootstrapShellCommand sets OciBootstrapShellCommand field to given value.

### HasOciBootstrapShellCommand

`func (o *FleetDescribeAccountConfigResult) HasOciBootstrapShellCommand() bool`

HasOciBootstrapShellCommand returns a boolean if a field has been set.

### GetOciDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) GetOciDisconnectShellCommand() string`

GetOciDisconnectShellCommand returns the OciDisconnectShellCommand field if non-nil, zero value otherwise.

### GetOciDisconnectShellCommandOk

`func (o *FleetDescribeAccountConfigResult) GetOciDisconnectShellCommandOk() (*string, bool)`

GetOciDisconnectShellCommandOk returns a tuple with the OciDisconnectShellCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) SetOciDisconnectShellCommand(v string)`

SetOciDisconnectShellCommand sets OciDisconnectShellCommand field to given value.

### HasOciDisconnectShellCommand

`func (o *FleetDescribeAccountConfigResult) HasOciDisconnectShellCommand() bool`

HasOciDisconnectShellCommand returns a boolean if a field has been set.

### GetOciDomainID

`func (o *FleetDescribeAccountConfigResult) GetOciDomainID() string`

GetOciDomainID returns the OciDomainID field if non-nil, zero value otherwise.

### GetOciDomainIDOk

`func (o *FleetDescribeAccountConfigResult) GetOciDomainIDOk() (*string, bool)`

GetOciDomainIDOk returns a tuple with the OciDomainID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciDomainID

`func (o *FleetDescribeAccountConfigResult) SetOciDomainID(v string)`

SetOciDomainID sets OciDomainID field to given value.

### HasOciDomainID

`func (o *FleetDescribeAccountConfigResult) HasOciDomainID() bool`

HasOciDomainID returns a boolean if a field has been set.

### GetOciTenancyID

`func (o *FleetDescribeAccountConfigResult) GetOciTenancyID() string`

GetOciTenancyID returns the OciTenancyID field if non-nil, zero value otherwise.

### GetOciTenancyIDOk

`func (o *FleetDescribeAccountConfigResult) GetOciTenancyIDOk() (*string, bool)`

GetOciTenancyIDOk returns a tuple with the OciTenancyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciTenancyID

`func (o *FleetDescribeAccountConfigResult) SetOciTenancyID(v string)`

SetOciTenancyID sets OciTenancyID field to given value.

### HasOciTenancyID

`func (o *FleetDescribeAccountConfigResult) HasOciTenancyID() bool`

HasOciTenancyID returns a boolean if a field has been set.

### GetStatus

`func (o *FleetDescribeAccountConfigResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FleetDescribeAccountConfigResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FleetDescribeAccountConfigResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *FleetDescribeAccountConfigResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *FleetDescribeAccountConfigResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *FleetDescribeAccountConfigResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


