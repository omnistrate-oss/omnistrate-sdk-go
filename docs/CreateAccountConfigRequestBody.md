# CreateAccountConfigRequestBody

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

### NewCreateAccountConfigRequestBody

`func NewCreateAccountConfigRequestBody(cloudProviderId string, description string, name string, ) *CreateAccountConfigRequestBody`

NewCreateAccountConfigRequestBody instantiates a new CreateAccountConfigRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountConfigRequestBodyWithDefaults

`func NewCreateAccountConfigRequestBodyWithDefaults() *CreateAccountConfigRequestBody`

NewCreateAccountConfigRequestBodyWithDefaults instantiates a new CreateAccountConfigRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *CreateAccountConfigRequestBody) GetAwsAccessKey() string`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *CreateAccountConfigRequestBody) GetAwsAccessKeyOk() (*string, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *CreateAccountConfigRequestBody) SetAwsAccessKey(v string)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *CreateAccountConfigRequestBody) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsAccountID

`func (o *CreateAccountConfigRequestBody) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *CreateAccountConfigRequestBody) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *CreateAccountConfigRequestBody) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.

### HasAwsAccountID

`func (o *CreateAccountConfigRequestBody) HasAwsAccountID() bool`

HasAwsAccountID returns a boolean if a field has been set.

### GetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequestBody) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *CreateAccountConfigRequestBody) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequestBody) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.

### HasAwsBootstrapRoleARN

`func (o *CreateAccountConfigRequestBody) HasAwsBootstrapRoleARN() bool`

HasAwsBootstrapRoleARN returns a boolean if a field has been set.

### GetAwsSecretKey

`func (o *CreateAccountConfigRequestBody) GetAwsSecretKey() string`

GetAwsSecretKey returns the AwsSecretKey field if non-nil, zero value otherwise.

### GetAwsSecretKeyOk

`func (o *CreateAccountConfigRequestBody) GetAwsSecretKeyOk() (*string, bool)`

GetAwsSecretKeyOk returns a tuple with the AwsSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretKey

`func (o *CreateAccountConfigRequestBody) SetAwsSecretKey(v string)`

SetAwsSecretKey sets AwsSecretKey field to given value.

### HasAwsSecretKey

`func (o *CreateAccountConfigRequestBody) HasAwsSecretKey() bool`

HasAwsSecretKey returns a boolean if a field has been set.

### GetByoaInstanceID

`func (o *CreateAccountConfigRequestBody) GetByoaInstanceID() string`

GetByoaInstanceID returns the ByoaInstanceID field if non-nil, zero value otherwise.

### GetByoaInstanceIDOk

`func (o *CreateAccountConfigRequestBody) GetByoaInstanceIDOk() (*string, bool)`

GetByoaInstanceIDOk returns a tuple with the ByoaInstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceID

`func (o *CreateAccountConfigRequestBody) SetByoaInstanceID(v string)`

SetByoaInstanceID sets ByoaInstanceID field to given value.

### HasByoaInstanceID

`func (o *CreateAccountConfigRequestBody) HasByoaInstanceID() bool`

HasByoaInstanceID returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *CreateAccountConfigRequestBody) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *CreateAccountConfigRequestBody) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *CreateAccountConfigRequestBody) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *CreateAccountConfigRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAccountConfigRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAccountConfigRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGcpProjectID

`func (o *CreateAccountConfigRequestBody) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *CreateAccountConfigRequestBody) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *CreateAccountConfigRequestBody) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.

### HasGcpProjectID

`func (o *CreateAccountConfigRequestBody) HasGcpProjectID() bool`

HasGcpProjectID returns a boolean if a field has been set.

### GetGcpProjectNumber

`func (o *CreateAccountConfigRequestBody) GetGcpProjectNumber() string`

GetGcpProjectNumber returns the GcpProjectNumber field if non-nil, zero value otherwise.

### GetGcpProjectNumberOk

`func (o *CreateAccountConfigRequestBody) GetGcpProjectNumberOk() (*string, bool)`

GetGcpProjectNumberOk returns a tuple with the GcpProjectNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectNumber

`func (o *CreateAccountConfigRequestBody) SetGcpProjectNumber(v string)`

SetGcpProjectNumber sets GcpProjectNumber field to given value.

### HasGcpProjectNumber

`func (o *CreateAccountConfigRequestBody) HasGcpProjectNumber() bool`

HasGcpProjectNumber returns a boolean if a field has been set.

### GetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequestBody) GetGcpServiceAccountEmail() string`

GetGcpServiceAccountEmail returns the GcpServiceAccountEmail field if non-nil, zero value otherwise.

### GetGcpServiceAccountEmailOk

`func (o *CreateAccountConfigRequestBody) GetGcpServiceAccountEmailOk() (*string, bool)`

GetGcpServiceAccountEmailOk returns a tuple with the GcpServiceAccountEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountEmail

`func (o *CreateAccountConfigRequestBody) SetGcpServiceAccountEmail(v string)`

SetGcpServiceAccountEmail sets GcpServiceAccountEmail field to given value.

### HasGcpServiceAccountEmail

`func (o *CreateAccountConfigRequestBody) HasGcpServiceAccountEmail() bool`

HasGcpServiceAccountEmail returns a boolean if a field has been set.

### GetGcpServiceAccountKey

`func (o *CreateAccountConfigRequestBody) GetGcpServiceAccountKey() string`

GetGcpServiceAccountKey returns the GcpServiceAccountKey field if non-nil, zero value otherwise.

### GetGcpServiceAccountKeyOk

`func (o *CreateAccountConfigRequestBody) GetGcpServiceAccountKeyOk() (*string, bool)`

GetGcpServiceAccountKeyOk returns a tuple with the GcpServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpServiceAccountKey

`func (o *CreateAccountConfigRequestBody) SetGcpServiceAccountKey(v string)`

SetGcpServiceAccountKey sets GcpServiceAccountKey field to given value.

### HasGcpServiceAccountKey

`func (o *CreateAccountConfigRequestBody) HasGcpServiceAccountKey() bool`

HasGcpServiceAccountKey returns a boolean if a field has been set.

### GetName

`func (o *CreateAccountConfigRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountConfigRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountConfigRequestBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


