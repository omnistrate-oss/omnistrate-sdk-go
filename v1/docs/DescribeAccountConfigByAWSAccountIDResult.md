# DescribeAccountConfigByAWSAccountIDResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | **string** | The AWS account ID | 
**AwsBootstrapRoleARN** | **string** | The security role ARN or service account ARN that grants access to operate the infra | 
**AwsCloudFormationNoLBTemplateURL** | Pointer to **string** | The URL to the CloudFormation template (no LoadBalancer policy version) | [optional] 
**AwsCloudFormationTemplateURL** | Pointer to **string** | The URL to the CloudFormation template | [optional] 
**ByoaInstanceIDs** | Pointer to **[]string** | The BYOA instance IDs that this account config is tied to | [optional] 
**CloudProviderId** | **string** | ID of an CloudProvider | 
**Description** | **string** | The description for the account | 
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account | 
**Status** | **string** | The status of the account configuration | 
**StatusMessage** | **string** | The status message of the account | 

## Methods

### NewDescribeAccountConfigByAWSAccountIDResult

`func NewDescribeAccountConfigByAWSAccountIDResult(awsAccountID string, awsBootstrapRoleARN string, cloudProviderId string, description string, id string, name string, status string, statusMessage string, ) *DescribeAccountConfigByAWSAccountIDResult`

NewDescribeAccountConfigByAWSAccountIDResult instantiates a new DescribeAccountConfigByAWSAccountIDResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByAWSAccountIDResultWithDefaults

`func NewDescribeAccountConfigByAWSAccountIDResultWithDefaults() *DescribeAccountConfigByAWSAccountIDResult`

NewDescribeAccountConfigByAWSAccountIDResultWithDefaults instantiates a new DescribeAccountConfigByAWSAccountIDResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.


### GetAwsBootstrapRoleARN

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsBootstrapRoleARN() string`

GetAwsBootstrapRoleARN returns the AwsBootstrapRoleARN field if non-nil, zero value otherwise.

### GetAwsBootstrapRoleARNOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsBootstrapRoleARNOk() (*string, bool)`

GetAwsBootstrapRoleARNOk returns a tuple with the AwsBootstrapRoleARN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsBootstrapRoleARN

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsBootstrapRoleARN(v string)`

SetAwsBootstrapRoleARN sets AwsBootstrapRoleARN field to given value.


### GetAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationNoLBTemplateURL() string`

GetAwsCloudFormationNoLBTemplateURL returns the AwsCloudFormationNoLBTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationNoLBTemplateURLOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationNoLBTemplateURLOk() (*string, bool)`

GetAwsCloudFormationNoLBTemplateURLOk returns a tuple with the AwsCloudFormationNoLBTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsCloudFormationNoLBTemplateURL(v string)`

SetAwsCloudFormationNoLBTemplateURL sets AwsCloudFormationNoLBTemplateURL field to given value.

### HasAwsCloudFormationNoLBTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) HasAwsCloudFormationNoLBTemplateURL() bool`

HasAwsCloudFormationNoLBTemplateURL returns a boolean if a field has been set.

### GetAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationTemplateURL() string`

GetAwsCloudFormationTemplateURL returns the AwsCloudFormationTemplateURL field if non-nil, zero value otherwise.

### GetAwsCloudFormationTemplateURLOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetAwsCloudFormationTemplateURLOk() (*string, bool)`

GetAwsCloudFormationTemplateURLOk returns a tuple with the AwsCloudFormationTemplateURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetAwsCloudFormationTemplateURL(v string)`

SetAwsCloudFormationTemplateURL sets AwsCloudFormationTemplateURL field to given value.

### HasAwsCloudFormationTemplateURL

`func (o *DescribeAccountConfigByAWSAccountIDResult) HasAwsCloudFormationTemplateURL() bool`

HasAwsCloudFormationTemplateURL returns a boolean if a field has been set.

### GetByoaInstanceIDs

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetByoaInstanceIDs() []string`

GetByoaInstanceIDs returns the ByoaInstanceIDs field if non-nil, zero value otherwise.

### GetByoaInstanceIDsOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetByoaInstanceIDsOk() (*[]string, bool)`

GetByoaInstanceIDsOk returns a tuple with the ByoaInstanceIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByoaInstanceIDs

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetByoaInstanceIDs(v []string)`

SetByoaInstanceIDs sets ByoaInstanceIDs field to given value.

### HasByoaInstanceIDs

`func (o *DescribeAccountConfigByAWSAccountIDResult) HasByoaInstanceIDs() bool`

HasByoaInstanceIDs returns a boolean if a field has been set.

### GetCloudProviderId

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetCloudProviderId() string`

GetCloudProviderId returns the CloudProviderId field if non-nil, zero value otherwise.

### GetCloudProviderIdOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetCloudProviderIdOk() (*string, bool)`

GetCloudProviderIdOk returns a tuple with the CloudProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderId

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetCloudProviderId(v string)`

SetCloudProviderId sets CloudProviderId field to given value.


### GetDescription

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeAccountConfigByAWSAccountIDResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeAccountConfigByAWSAccountIDResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


