# DescribeConsumptionBillingDetailsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentConfigured** | **bool** | Whether the customer has configured their payment information. | 
**PaymentInfoPortalURL** | Pointer to **string** | The URL from the billing provide to redirect users to so they can enter their payment information.  Only present when first adding payment information | [optional] 

## Methods

### NewDescribeConsumptionBillingDetailsResult

`func NewDescribeConsumptionBillingDetailsResult(paymentConfigured bool, ) *DescribeConsumptionBillingDetailsResult`

NewDescribeConsumptionBillingDetailsResult instantiates a new DescribeConsumptionBillingDetailsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeConsumptionBillingDetailsResultWithDefaults

`func NewDescribeConsumptionBillingDetailsResultWithDefaults() *DescribeConsumptionBillingDetailsResult`

NewDescribeConsumptionBillingDetailsResultWithDefaults instantiates a new DescribeConsumptionBillingDetailsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentConfigured

`func (o *DescribeConsumptionBillingDetailsResult) GetPaymentConfigured() bool`

GetPaymentConfigured returns the PaymentConfigured field if non-nil, zero value otherwise.

### GetPaymentConfiguredOk

`func (o *DescribeConsumptionBillingDetailsResult) GetPaymentConfiguredOk() (*bool, bool)`

GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigured

`func (o *DescribeConsumptionBillingDetailsResult) SetPaymentConfigured(v bool)`

SetPaymentConfigured sets PaymentConfigured field to given value.


### GetPaymentInfoPortalURL

`func (o *DescribeConsumptionBillingDetailsResult) GetPaymentInfoPortalURL() string`

GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field if non-nil, zero value otherwise.

### GetPaymentInfoPortalURLOk

`func (o *DescribeConsumptionBillingDetailsResult) GetPaymentInfoPortalURLOk() (*string, bool)`

GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoPortalURL

`func (o *DescribeConsumptionBillingDetailsResult) SetPaymentInfoPortalURL(v string)`

SetPaymentInfoPortalURL sets PaymentInfoPortalURL field to given value.

### HasPaymentInfoPortalURL

`func (o *DescribeConsumptionBillingDetailsResult) HasPaymentInfoPortalURL() bool`

HasPaymentInfoPortalURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


