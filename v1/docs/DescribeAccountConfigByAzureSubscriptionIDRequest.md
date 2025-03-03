# DescribeAccountConfigByAzureSubscriptionIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureSubscriptionID** | **string** | The Azure subscription ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccountConfigByAzureSubscriptionIDRequest

`func NewDescribeAccountConfigByAzureSubscriptionIDRequest(azureSubscriptionID string, token string, ) *DescribeAccountConfigByAzureSubscriptionIDRequest`

NewDescribeAccountConfigByAzureSubscriptionIDRequest instantiates a new DescribeAccountConfigByAzureSubscriptionIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByAzureSubscriptionIDRequestWithDefaults

`func NewDescribeAccountConfigByAzureSubscriptionIDRequestWithDefaults() *DescribeAccountConfigByAzureSubscriptionIDRequest`

NewDescribeAccountConfigByAzureSubscriptionIDRequestWithDefaults instantiates a new DescribeAccountConfigByAzureSubscriptionIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureSubscriptionID

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) GetAzureSubscriptionID() string`

GetAzureSubscriptionID returns the AzureSubscriptionID field if non-nil, zero value otherwise.

### GetAzureSubscriptionIDOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) GetAzureSubscriptionIDOk() (*string, bool)`

GetAzureSubscriptionIDOk returns a tuple with the AzureSubscriptionID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSubscriptionID

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) SetAzureSubscriptionID(v string)`

SetAzureSubscriptionID sets AzureSubscriptionID field to given value.


### GetToken

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccountConfigByAzureSubscriptionIDRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


