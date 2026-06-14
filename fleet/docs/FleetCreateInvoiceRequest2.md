# FleetCreateInvoiceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingPeriod** | **string** | Billing month to invoice in YYYY-MM format | 
**CustomerId** | **string** | Organization ID of the customer to invoice | 

## Methods

### NewFleetCreateInvoiceRequest2

`func NewFleetCreateInvoiceRequest2(billingPeriod string, customerId string, ) *FleetCreateInvoiceRequest2`

NewFleetCreateInvoiceRequest2 instantiates a new FleetCreateInvoiceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateInvoiceRequest2WithDefaults

`func NewFleetCreateInvoiceRequest2WithDefaults() *FleetCreateInvoiceRequest2`

NewFleetCreateInvoiceRequest2WithDefaults instantiates a new FleetCreateInvoiceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingPeriod

`func (o *FleetCreateInvoiceRequest2) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *FleetCreateInvoiceRequest2) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *FleetCreateInvoiceRequest2) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.


### GetCustomerId

`func (o *FleetCreateInvoiceRequest2) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *FleetCreateInvoiceRequest2) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *FleetCreateInvoiceRequest2) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


