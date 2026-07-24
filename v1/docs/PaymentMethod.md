# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | Pointer to **string** | Bank name when the payment method is a bank account | [optional] 
**Brand** | Pointer to **string** | Card brand when the payment method is a card | [optional] 
**DisplayName** | **string** | Display-safe payment method label | 
**ExpMonth** | Pointer to **int64** | Card expiration month when the payment method is a card | [optional] 
**ExpYear** | Pointer to **int64** | Card expiration year when the payment method is a card | [optional] 
**Id** | **string** | Stripe payment method ID | 
**IsDefault** | **bool** | Whether this is the default payment method | 
**Last4** | Pointer to **string** | Last four digits or characters when Stripe exposes them | [optional] 
**Type** | **string** | Stripe payment method type (card, us_bank_account, sepa_debit, etc.) | 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(displayName string, id string, isDefault bool, type_ string, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankName

`func (o *PaymentMethod) GetBankName() string`

GetBankName returns the BankName field if non-nil, zero value otherwise.

### GetBankNameOk

`func (o *PaymentMethod) GetBankNameOk() (*string, bool)`

GetBankNameOk returns a tuple with the BankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankName

`func (o *PaymentMethod) SetBankName(v string)`

SetBankName sets BankName field to given value.

### HasBankName

`func (o *PaymentMethod) HasBankName() bool`

HasBankName returns a boolean if a field has been set.

### GetBrand

`func (o *PaymentMethod) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *PaymentMethod) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *PaymentMethod) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *PaymentMethod) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetDisplayName

`func (o *PaymentMethod) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PaymentMethod) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PaymentMethod) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpMonth

`func (o *PaymentMethod) GetExpMonth() int64`

GetExpMonth returns the ExpMonth field if non-nil, zero value otherwise.

### GetExpMonthOk

`func (o *PaymentMethod) GetExpMonthOk() (*int64, bool)`

GetExpMonthOk returns a tuple with the ExpMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpMonth

`func (o *PaymentMethod) SetExpMonth(v int64)`

SetExpMonth sets ExpMonth field to given value.

### HasExpMonth

`func (o *PaymentMethod) HasExpMonth() bool`

HasExpMonth returns a boolean if a field has been set.

### GetExpYear

`func (o *PaymentMethod) GetExpYear() int64`

GetExpYear returns the ExpYear field if non-nil, zero value otherwise.

### GetExpYearOk

`func (o *PaymentMethod) GetExpYearOk() (*int64, bool)`

GetExpYearOk returns a tuple with the ExpYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpYear

`func (o *PaymentMethod) SetExpYear(v int64)`

SetExpYear sets ExpYear field to given value.

### HasExpYear

`func (o *PaymentMethod) HasExpYear() bool`

HasExpYear returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.


### GetIsDefault

`func (o *PaymentMethod) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethod) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetLast4

`func (o *PaymentMethod) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *PaymentMethod) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *PaymentMethod) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *PaymentMethod) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


