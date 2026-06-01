# ListPaymentMethodsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethods** | Pointer to [**[]PaymentMethod**](PaymentMethod.md) | List of payment methods attached to the customer | [optional] 

## Methods

### NewListPaymentMethodsResult

`func NewListPaymentMethodsResult() *ListPaymentMethodsResult`

NewListPaymentMethodsResult instantiates a new ListPaymentMethodsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPaymentMethodsResultWithDefaults

`func NewListPaymentMethodsResultWithDefaults() *ListPaymentMethodsResult`

NewListPaymentMethodsResultWithDefaults instantiates a new ListPaymentMethodsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethods

`func (o *ListPaymentMethodsResult) GetPaymentMethods() []PaymentMethod`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *ListPaymentMethodsResult) GetPaymentMethodsOk() (*[]PaymentMethod, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *ListPaymentMethodsResult) SetPaymentMethods(v []PaymentMethod)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *ListPaymentMethodsResult) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


