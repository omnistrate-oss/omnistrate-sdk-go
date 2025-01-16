# CreateAccountConfigRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to **string** | The AWS access key | [optional] 
**AwsAccountID** | Pointer to **string** | The AWS account ID | [optional] 
**AwsBootstrapRoleARN** | Pointer to **string** | The security role ARN or service account ARN that grants access to operate the infra | [optional] 
**AwsSecretKey** | Pointer to **string** | The AWS secret key | [optional] 
**ByoaInstanceID** | Pointer to **string** | The BYOA instance ID that this account config is tied to | [optional] 
**CloudProviderId** | **string** | Cloud Provider ID to operate on | 
**Description** | **string** | The description for the account | 
**GcpProjectID** | Pointer to **string** | The GCP project ID | [optional] 
**GcpProjectNumber** | Pointer to **string** | The GCP project number | [optional] 
**GcpServiceAccountEmail** | Pointer to **string** | The GCP service account email | [optional] 
**GcpServiceAccountKey** | Pointer to **string** | The GCP service account key | [optional] 
**Name** | **string** | The name of the account | 

## Methods

### NewCreateAccountConfigRequest2

`func NewCreateAccountConfigRequest2(cloudProviderId string, description string, name string, ) *CreateAccountConfigRequest2`

NewCreateAccountConfigRequest2 instantiates a new CreateAccountConfigRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountConfigRequest2WithDefaults

`func NewCreateAccountConfigRequest2WithDefaults() *CreateAccountConfigRequest2`

NewCreateAccountConfigRequest2WithDefaults instantiates a new CreateAccountConfigRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *CreateAccountConfigRequest2) GetAwsAccessKey() string`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *CreateAccountConfigRequest2) GetAwsAccessKeyOk() (*string, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *CreateAccountConfigRequest2) SetAwsAccessKey(v string)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *CreateAccountConfigRequest2) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsAccountID

`func (o *CreateAccountConfigRequest2) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *CreateAccountConfigRequest2) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *CreateAccountConfigRequest2) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *CreateAccountConfigRequest2) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest2) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *CreateAccountConfigRequest2) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest2) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.

### HasAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequest2) HasAwsBootstrapRoleARN() bool`

HasAwsBootstrapRoleARN returns a boolean if a field has been set.

### GetAwsSecretKey

`func (o *CreateAccountConfigRequest2) GetAwsSecretKey() string`

GetAwsSecretKey returns the AwsSecretKey field if non-nil, zero value otherwise.

### GetAwsSecretKeyOk

`func (o *CreateAccountConfigRequest2) GetAwsSecretKeyOk() (*string, bool)`

GetAwsSecretKeyOk returns a tuple with the AwsSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretKey

`func (o *CreateAccountConfigRequest2) SetAwsSecretKey(v string)`

SetAwsSecretKey sets AwsSecretKey field to given value.

### HasAwsSecretKey

`func (o *CreateAccountConfigRequest2) HasAwsSecretKey() bool`

HasAwsSecretKey returns a boolean if a field has been set.

### GetByoaInstanceID

`func (o *CreateAccountConfigRequest2) GetByoaInstanceID() string`

GetByoaInstanceID returns the ByoaInstanceID field if non-nil, zero value otherwise.

### GetByoaInstanceIDOk

`func (o *CreateAccountConfigRequest2) GetByoaInstanceIDOk() (*string, bool)`

GetByoaInstanceIDOk returns a tuple with the ByoaInstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceID

`func (o *CreateAccountConfigRequest2) SetByoaInstanceID(v string)`

SetByoaInstanceID sets ByoaInstanceID field to given value.

### HasByoaInstanceID

`func (o *CreateAccountConfigRequest2) HasByoaInstanceID() bool`

HasByoaInstanceID returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *CreateAccountConfigRequest2) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *CreateAccountConfigRequest2) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *CreateAccountConfigRequest2) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *CreateAccountConfigRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAccountConfigRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAccountConfigRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpProjectID

`func (o *CreateAccountConfigRequest2) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *CreateAccountConfigRequest2) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *CreateAccountConfigRequest2) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *CreateAccountConfigRequest2) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *CreateAccountConfigRequest2) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *CreateAccountConfigRequest2) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *CreateAccountConfigRequest2) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *CreateAccountConfigRequest2) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest2) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *CreateAccountConfigRequest2) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest2) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *CreateAccountConfigRequest2) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *CreateAccountConfigRequest2) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *CreateAccountConfigRequest2) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *CreateAccountConfigRequest2) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *CreateAccountConfigRequest2) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetName

`func (o *CreateAccountConfigRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountConfigRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountConfigRequest2) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


