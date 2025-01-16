# DescribeUserBillingDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingEmbedURL** | **string** | The URL from the billing provider to embed in an iframe to show the user&#39;s billing information | 
**Name** | **string** | The name of the user | 
**PaymentConfigured** | **bool** | Whether the customer has configured their payment information. | 
**PaymentInfoPortalURL** | Pointer to **string** | The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information | [optional] 
**UserId** | **string** | ID of a User | 

## Methods

### NewDescribeUserBillingDetailsResult

`func NewDescribeUserBillingDetailsResult(billingEmbedURL string, name string, paymentConfigured bool, userId string, ) *DescribeUserBillingDetailsResult`

NewDescribeUserBillingDetailsResult instantiates a new DescribeUserBillingDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserBillingDetailsResultWithDefaults

`func NewDescribeUserBillingDetailsResultWithDefaults() *DescribeUserBillingDetailsResult`

NewDescribeUserBillingDetailsResultWithDefaults instantiates a new DescribeUserBillingDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingEmbedURL

`func (o *DescribeUserBillingDetailsResult) GetBillingEmbedURL() string`

GetBillingEmbedURL returns the BillingEmbedURL field if non-nil, zero value otherwise.

### GetBillingEmbedURLOk

`func (o *DescribeUserBillingDetailsResult) GetBillingEmbedURLOk() (*string, bool)`

GetBillingEmbedURLOk returns a tuple with the BillingEmbedURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEmbedURL

`func (o *DescribeUserBillingDetailsResult) SetBillingEmbedURL(v string)`

SetBillingEmbedURL sets BillingEmbedURL field to given value.


### GetName

`func (o *DescribeUserBillingDetailsResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeUserBillingDetailsResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeUserBillingDetailsResult) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentConfigured

`func (o *DescribeUserBillingDetailsResult) GetPaymentConfigured() bool`

GetPaymentConfigured returns the PaymentConfigured field if non-nil, zero value otherwise.

### GetPaymentConfiguredOk

`func (o *DescribeUserBillingDetailsResult) GetPaymentConfiguredOk() (*bool, bool)`

GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigured

`func (o *DescribeUserBillingDetailsResult) SetPaymentConfigured(v bool)`

SetPaymentConfigured sets PaymentConfigured field to given value.


### GetPaymentInfoPortalURL

`func (o *DescribeUserBillingDetailsResult) GetPaymentInfoPortalURL() string`

GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field if non-nil, zero value otherwise.

### GetPaymentInfoPortalURLOk

`func (o *DescribeUserBillingDetailsResult) GetPaymentInfoPortalURLOk() (*string, bool)`

GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoPortalURL

`func (o *DescribeUserBillingDetailsResult) SetPaymentInfoPortalURL(v string)`

SetPaymentInfoPortalURL sets PaymentInfoPortalURL field to given value.

### HasPaymentInfoPortalURL

`func (o *DescribeUserBillingDetailsResult) HasPaymentInfoPortalURL() bool`

HasPaymentInfoPortalURL returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeUserBillingDetailsResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeUserBillingDetailsResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeUserBillingDetailsResult) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


