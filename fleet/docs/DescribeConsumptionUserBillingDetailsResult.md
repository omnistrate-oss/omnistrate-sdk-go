# DescribeConsumptionUserBillingDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEmbedURL** | Pointer to **string** | DEPRECATED | [optional] 
**Name** | Pointer to **string** | DEPRECATED: The name of the user | [optional] 
**PaymentConfigured** | **bool** | Whether the customer has configured their payment information. | 
**PaymentInfoPortalURL** | Pointer to **string** | The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information | [optional] 
**UserId** | Pointer to **string** | ID of a User | [optional] 

## Methods

### NewDescribeConsumptionUserBillingDetailsResult

`func NewDescribeConsumptionUserBillingDetailsResult(paymentConfigured bool, ) *DescribeConsumptionUserBillingDetailsResult`

NewDescribeConsumptionUserBillingDetailsResult instantiates a new DescribeConsumptionUserBillingDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeConsumptionUserBillingDetailsResultWithDefaults

`func NewDescribeConsumptionUserBillingDetailsResultWithDefaults() *DescribeConsumptionUserBillingDetailsResult`

NewDescribeConsumptionUserBillingDetailsResultWithDefaults instantiates a new DescribeConsumptionUserBillingDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEmbedURL

`func (o *DescribeConsumptionUserBillingDetailsResult) GetBillingEmbedURL() string`

GetBillingEmbedURL returns the BillingEmbedURL field if non-nil, zero value otherwise.

### GetBillingEmbedURLOk

`func (o *DescribeConsumptionUserBillingDetailsResult) GetBillingEmbedURLOk() (*string, bool)`

GetBillingEmbedURLOk returns a tuple with the BillingEmbedURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmbedURL

`func (o *DescribeConsumptionUserBillingDetailsResult) SetBillingEmbedURL(v string)`

SetBillingEmbedURL sets BillingEmbedURL field to given value.

### HasBillingEmbedURL

`func (o *DescribeConsumptionUserBillingDetailsResult) HasBillingEmbedURL() bool`

HasBillingEmbedURL returns a boolean if a field has been set.

### GetName

`func (o *DescribeConsumptionUserBillingDetailsResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeConsumptionUserBillingDetailsResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeConsumptionUserBillingDetailsResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeConsumptionUserBillingDetailsResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentConfigured

`func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentConfigured() bool`

GetPaymentConfigured returns the PaymentConfigured field if non-nil, zero value otherwise.

### GetPaymentConfiguredOk

`func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentConfiguredOk() (*bool, bool)`

GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigured

`func (o *DescribeConsumptionUserBillingDetailsResult) SetPaymentConfigured(v bool)`

SetPaymentConfigured sets PaymentConfigured field to given value.


### GetPaymentInfoPortalURL

`func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentInfoPortalURL() string`

GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field if non-nil, zero value otherwise.

### GetPaymentInfoPortalURLOk

`func (o *DescribeConsumptionUserBillingDetailsResult) GetPaymentInfoPortalURLOk() (*string, bool)`

GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoPortalURL

`func (o *DescribeConsumptionUserBillingDetailsResult) SetPaymentInfoPortalURL(v string)`

SetPaymentInfoPortalURL sets PaymentInfoPortalURL field to given value.

### HasPaymentInfoPortalURL

`func (o *DescribeConsumptionUserBillingDetailsResult) HasPaymentInfoPortalURL() bool`

HasPaymentInfoPortalURL returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeConsumptionUserBillingDetailsResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeConsumptionUserBillingDetailsResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeConsumptionUserBillingDetailsResult) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DescribeConsumptionUserBillingDetailsResult) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


