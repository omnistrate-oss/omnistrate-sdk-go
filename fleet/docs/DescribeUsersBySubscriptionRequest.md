# DescribeUsersBySubscriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The subscription ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeUsersBySubscriptionRequest

`func NewDescribeUsersBySubscriptionRequest(subscriptionId string, token string, ) *DescribeUsersBySubscriptionRequest`

NewDescribeUsersBySubscriptionRequest instantiates a new DescribeUsersBySubscriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersBySubscriptionRequestWithDefaults

`func NewDescribeUsersBySubscriptionRequestWithDefaults() *DescribeUsersBySubscriptionRequest`

NewDescribeUsersBySubscriptionRequestWithDefaults instantiates a new DescribeUsersBySubscriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriptionId

`func (o *DescribeUsersBySubscriptionRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *DescribeUsersBySubscriptionRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *DescribeUsersBySubscriptionRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetToken

`func (o *DescribeUsersBySubscriptionRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeUsersBySubscriptionRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeUsersBySubscriptionRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


