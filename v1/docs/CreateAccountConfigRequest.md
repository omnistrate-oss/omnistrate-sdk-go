# CreateAccountConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to **string** | The AWS access key | [optional] 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AwsBootstrapRoleARN** | Pointer to **string** | The security role ARN or service account ARN that grants access to operate the infra | [optional] 
**AwsSecretKey** | Pointer to **string** | The AWS secret key | [optional] 
**AzureSubscriptionID** | Pointer to **string** | The Azure subscription ID | [optional] 
**AzureTenantID** | Pointer to **string** | The Azure tenant ID | [optional] 
**ByoaInstanceID** | Pointer to **string** | The BYOA instance ID that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**GcpProjectNumber** | Pointer to **string** | The GCP project number | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** | The GCP service account email | [optional] 
**GcpServiceAccountKey** | Pointer to **string** | The GCP service account key | [optional] 
**Name** | **string** | The name of the account | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateAccountConfigRequest

`func NewCreateAccountConfigRequest(cloudProviderId string, description string, name string, token string, ) *CreateAccountConfigRequest`

NewCreateAccountConfigRequest instantiates a new CreateAccountConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountConfigRequestWithDefaults

`func NewCreateAccountConfigRequestWithDefaults() *CreateAccountConfigRequest`

NewCreateAccountConfigRequestWithDefaults instantiates a new CreateAccountConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *CreateAccountConfigRequest) GetAwsAccessKey() string`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *CreateAccountConfigRequest) GetAwsAccessKeyOk() (*string, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *CreateAccountConfigRequest) SetAwsAccessKey(v string)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *CreateAccountConfigRequest) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsAccountID

`func (o *CreateAccountConfigRequest) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *CreateAccountConfigRequest) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *CreateAccountConfigRequest) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *CreateAccountConfigRequest) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *CreateAccountConfigRequest) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.

### HasAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest) HasAwsBootstrapRoleARN() bool`

HasAwsBootstrapRoleARN returns a boolean if a field has been set.

### GetAwsSecretKey

`func (o *CreateAccountConfigRequest) GetAwsSecretKey() string`

GetAwsSecretKey returns the AwsSecretKey field if non-nil, zero value otherwise.

### GetAwsSecretKeyOk

`func (o *CreateAccountConfigRequest) GetAwsSecretKeyOk() (*string, bool)`

GetAwsSecretKeyOk returns a tuple with the AwsSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretKey

`func (o *CreateAccountConfigRequest) SetAwsSecretKey(v string)`

SetAwsSecretKey sets AwsSecretKey field to given value.

### HasAwsSecretKey

`func (o *CreateAccountConfigRequest) HasAwsSecretKey() bool`

HasAwsSecretKey returns a boolean if a field has been set.

### GetAzureSubscriptionID

`func (o *CreateAccountConfigRequest) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *CreateAccountConfigRequest) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *CreateAccountConfigRequest) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.

### HasAzureSubscriptionID

`func (o *CreateAccountConfigRequest) HasAzureSubscriptionID() bool`

HasAzureSubscriptionID returns a boolean if a field has been set.

### GetAzureTenantID

`func (o *CreateAccountConfigRequest) GetAzureTenantID() string`

GetAzureTenantID returns the AzureTenantID field if non-nil, zero value otherwise.

### GetAzureTenantIDOk

`func (o *CreateAccountConfigRequest) GetAzureTenantIDOk() (*string, bool)`

GetAzureTenantIDOk returns a tuple with the AzureTenantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantID

`func (o *CreateAccountConfigRequest) SetAzureTenantID(v string)`

SetAzureTenantID sets AzureTenantID field to given value.

### HasAzureTenantID

`func (o *CreateAccountConfigRequest) HasAzureTenantID() bool`

HasAzureTenantID returns a boolean if a field has been set.

### GetByoaInstanceID

`func (o *CreateAccountConfigRequest) GetByoaInstanceID() string`

GetByoaInstanceID returns the ByoaInstanceID field if non-nil, zero value otherwise.

### GetByoaInstanceIDOk

`func (o *CreateAccountConfigRequest) GetByoaInstanceIDOk() (*string, bool)`

GetByoaInstanceIDOk returns a tuple with the ByoaInstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceID

`func (o *CreateAccountConfigRequest) SetByoaInstanceID(v string)`

SetByoaInstanceID sets ByoaInstanceID field to given value.

### HasByoaInstanceID

`func (o *CreateAccountConfigRequest) HasByoaInstanceID() bool`

HasByoaInstanceID returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *CreateAccountConfigRequest) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *CreateAccountConfigRequest) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *CreateAccountConfigRequest) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *CreateAccountConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAccountConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAccountConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpProjectID

`func (o *CreateAccountConfigRequest) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *CreateAccountConfigRequest) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *CreateAccountConfigRequest) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *CreateAccountConfigRequest) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *CreateAccountConfigRequest) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *CreateAccountConfigRequest) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *CreateAccountConfigRequest) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *CreateAccountConfigRequest) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *CreateAccountConfigRequest) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *CreateAccountConfigRequest) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *CreateAccountConfigRequest) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *CreateAccountConfigRequest) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *CreateAccountConfigRequest) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetName

`func (o *CreateAccountConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountConfigRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *CreateAccountConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateAccountConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateAccountConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


