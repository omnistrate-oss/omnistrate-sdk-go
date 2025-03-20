# DescribeConsumptionBillingDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnUrl** | Pointer to **string** | Return Url used to configure payment methods links | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeConsumptionBillingDetailsRequest

`func NewDescribeConsumptionBillingDetailsRequest(token string, ) *DescribeConsumptionBillingDetailsRequest`

NewDescribeConsumptionBillingDetailsRequest instantiates a new DescribeConsumptionBillingDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeConsumptionBillingDetailsRequestWithDefaults

`func NewDescribeConsumptionBillingDetailsRequestWithDefaults() *DescribeConsumptionBillingDetailsRequest`

NewDescribeConsumptionBillingDetailsRequestWithDefaults instantiates a new DescribeConsumptionBillingDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnUrl

`func (o *DescribeConsumptionBillingDetailsRequest) GetReturnUrl() string`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *DescribeConsumptionBillingDetailsRequest) GetReturnUrlOk() (*string, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *DescribeConsumptionBillingDetailsRequest) SetReturnUrl(v string)`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *DescribeConsumptionBillingDetailsRequest) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### GetToken

`func (o *DescribeConsumptionBillingDetailsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeConsumptionBillingDetailsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeConsumptionBillingDetailsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


