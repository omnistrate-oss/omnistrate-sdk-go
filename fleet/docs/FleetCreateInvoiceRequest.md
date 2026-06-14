# FleetCreateInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPeriod** | **string** | Billing month to invoice in YYYY-MM format | 
**CustomerId** | **string** | ID of an Org | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewFleetCreateInvoiceRequest

`func NewFleetCreateInvoiceRequest(billingPeriod string, customerId string, token string, ) *FleetCreateInvoiceRequest`

NewFleetCreateInvoiceRequest instantiates a new FleetCreateInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateInvoiceRequestWithDefaults

`func NewFleetCreateInvoiceRequestWithDefaults() *FleetCreateInvoiceRequest`

NewFleetCreateInvoiceRequestWithDefaults instantiates a new FleetCreateInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingPeriod

`func (o *FleetCreateInvoiceRequest) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *FleetCreateInvoiceRequest) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *FleetCreateInvoiceRequest) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetCustomerId

`func (o *FleetCreateInvoiceRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FleetCreateInvoiceRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FleetCreateInvoiceRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetToken

`func (o *FleetCreateInvoiceRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FleetCreateInvoiceRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FleetCreateInvoiceRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


