# DescribeAccountConfigByAWSAccountIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccountID** | **string** | The AWS account ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccountConfigByAWSAccountIDRequest

`func NewDescribeAccountConfigByAWSAccountIDRequest(awsAccountID string, token string, ) *DescribeAccountConfigByAWSAccountIDRequest`

NewDescribeAccountConfigByAWSAccountIDRequest instantiates a new DescribeAccountConfigByAWSAccountIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByAWSAccountIDRequestWithDefaults

`func NewDescribeAccountConfigByAWSAccountIDRequestWithDefaults() *DescribeAccountConfigByAWSAccountIDRequest`

NewDescribeAccountConfigByAWSAccountIDRequestWithDefaults instantiates a new DescribeAccountConfigByAWSAccountIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccountID

`func (o *DescribeAccountConfigByAWSAccountIDRequest) GetAwsAccountID() string`

GetAwsAccountID returns the AwsAccountID field if non-nil, zero value otherwise.

### GetAwsAccountIDOk

`func (o *DescribeAccountConfigByAWSAccountIDRequest) GetAwsAccountIDOk() (*string, bool)`

GetAwsAccountIDOk returns a tuple with the AwsAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccountID

`func (o *DescribeAccountConfigByAWSAccountIDRequest) SetAwsAccountID(v string)`

SetAwsAccountID sets AwsAccountID field to given value.


### GetToken

`func (o *DescribeAccountConfigByAWSAccountIDRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccountConfigByAWSAccountIDRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccountConfigByAWSAccountIDRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


