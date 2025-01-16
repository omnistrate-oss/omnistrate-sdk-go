# DescribePlanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModifiedAt** | Pointer to **string** | The time the plan was last modified | [optional] 
**NextChargeDate** | Pointer to **string** | The next day stripe will process a charge for this plan | [optional] 
**PaymentConfigured** | Pointer to **bool** | Whether the customer has configured their payment information. | [optional] 
**PaymentInfoPortalURL** | Pointer to **string** | The URL from Paigo to redirect users to so they can enter their payment information.  Only present when first adding payment information | [optional] 
**PlanCoreHourCost** | Pointer to **float64** | The cost per core hour of this plan | [optional] 
**PlanDescription** | Pointer to **string** | The description of the plan | [optional] 
**PlanFrequency** | Pointer to **string** | This parameter tells you if the plan is charged monthly or yearly | [optional] 
**PlanMonthlyCost** | Pointer to **float64** | The minimum monthly cost of this plan | [optional] 
**PlanName** | **string** | This parameter is used to select the appropriate Product Plan | 
**RemainingCredits** | Pointer to **string** | The credits remaining for the customer for the customer in Paigo | [optional] 
**StartDate** | Pointer to **string** | The date that the plan starts | [optional] 

## Methods

### NewDescribePlanResult

`func NewDescribePlanResult(planName string, ) *DescribePlanResult`

NewDescribePlanResult instantiates a new DescribePlanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePlanResultWithDefaults

`func NewDescribePlanResultWithDefaults() *DescribePlanResult`

NewDescribePlanResultWithDefaults instantiates a new DescribePlanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModifiedAt

`func (o *DescribePlanResult) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *DescribePlanResult) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *DescribePlanResult) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *DescribePlanResult) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetNextChargeDate

`func (o *DescribePlanResult) GetNextChargeDate() string`

GetNextChargeDate returns the NextChargeDate field if non-nil, zero value otherwise.

### GetNextChargeDateOk

`func (o *DescribePlanResult) GetNextChargeDateOk() (*string, bool)`

GetNextChargeDateOk returns a tuple with the NextChargeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextChargeDate

`func (o *DescribePlanResult) SetNextChargeDate(v string)`

SetNextChargeDate sets NextChargeDate field to given value.

### HasNextChargeDate

`func (o *DescribePlanResult) HasNextChargeDate() bool`

HasNextChargeDate returns a boolean if a field has been set.

### GetPaymentConfigured

`func (o *DescribePlanResult) GetPaymentConfigured() bool`

GetPaymentConfigured returns the PaymentConfigured field if non-nil, zero value otherwise.

### GetPaymentConfiguredOk

`func (o *DescribePlanResult) GetPaymentConfiguredOk() (*bool, bool)`

GetPaymentConfiguredOk returns a tuple with the PaymentConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentConfigured

`func (o *DescribePlanResult) SetPaymentConfigured(v bool)`

SetPaymentConfigured sets PaymentConfigured field to given value.

### HasPaymentConfigured

`func (o *DescribePlanResult) HasPaymentConfigured() bool`

HasPaymentConfigured returns a boolean if a field has been set.

### GetPaymentInfoPortalURL

`func (o *DescribePlanResult) GetPaymentInfoPortalURL() string`

GetPaymentInfoPortalURL returns the PaymentInfoPortalURL field if non-nil, zero value otherwise.

### GetPaymentInfoPortalURLOk

`func (o *DescribePlanResult) GetPaymentInfoPortalURLOk() (*string, bool)`

GetPaymentInfoPortalURLOk returns a tuple with the PaymentInfoPortalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfoPortalURL

`func (o *DescribePlanResult) SetPaymentInfoPortalURL(v string)`

SetPaymentInfoPortalURL sets PaymentInfoPortalURL field to given value.

### HasPaymentInfoPortalURL

`func (o *DescribePlanResult) HasPaymentInfoPortalURL() bool`

HasPaymentInfoPortalURL returns a boolean if a field has been set.

### GetPlanCoreHourCost

`func (o *DescribePlanResult) GetPlanCoreHourCost() float64`

GetPlanCoreHourCost returns the PlanCoreHourCost field if non-nil, zero value otherwise.

### GetPlanCoreHourCostOk

`func (o *DescribePlanResult) GetPlanCoreHourCostOk() (*float64, bool)`

GetPlanCoreHourCostOk returns a tuple with the PlanCoreHourCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCoreHourCost

`func (o *DescribePlanResult) SetPlanCoreHourCost(v float64)`

SetPlanCoreHourCost sets PlanCoreHourCost field to given value.

### HasPlanCoreHourCost

`func (o *DescribePlanResult) HasPlanCoreHourCost() bool`

HasPlanCoreHourCost returns a boolean if a field has been set.

### GetPlanDescription

`func (o *DescribePlanResult) GetPlanDescription() string`

GetPlanDescription returns the PlanDescription field if non-nil, zero value otherwise.

### GetPlanDescriptionOk

`func (o *DescribePlanResult) GetPlanDescriptionOk() (*string, bool)`

GetPlanDescriptionOk returns a tuple with the PlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanDescription

`func (o *DescribePlanResult) SetPlanDescription(v string)`

SetPlanDescription sets PlanDescription field to given value.

### HasPlanDescription

`func (o *DescribePlanResult) HasPlanDescription() bool`

HasPlanDescription returns a boolean if a field has been set.

### GetPlanFrequency

`func (o *DescribePlanResult) GetPlanFrequency() string`

GetPlanFrequency returns the PlanFrequency field if non-nil, zero value otherwise.

### GetPlanFrequencyOk

`func (o *DescribePlanResult) GetPlanFrequencyOk() (*string, bool)`

GetPlanFrequencyOk returns a tuple with the PlanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanFrequency

`func (o *DescribePlanResult) SetPlanFrequency(v string)`

SetPlanFrequency sets PlanFrequency field to given value.

### HasPlanFrequency

`func (o *DescribePlanResult) HasPlanFrequency() bool`

HasPlanFrequency returns a boolean if a field has been set.

### GetPlanMonthlyCost

`func (o *DescribePlanResult) GetPlanMonthlyCost() float64`

GetPlanMonthlyCost returns the PlanMonthlyCost field if non-nil, zero value otherwise.

### GetPlanMonthlyCostOk

`func (o *DescribePlanResult) GetPlanMonthlyCostOk() (*float64, bool)`

GetPlanMonthlyCostOk returns a tuple with the PlanMonthlyCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanMonthlyCost

`func (o *DescribePlanResult) SetPlanMonthlyCost(v float64)`

SetPlanMonthlyCost sets PlanMonthlyCost field to given value.

### HasPlanMonthlyCost

`func (o *DescribePlanResult) HasPlanMonthlyCost() bool`

HasPlanMonthlyCost returns a boolean if a field has been set.

### GetPlanName

`func (o *DescribePlanResult) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *DescribePlanResult) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *DescribePlanResult) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetRemainingCredits

`func (o *DescribePlanResult) GetRemainingCredits() string`

GetRemainingCredits returns the RemainingCredits field if non-nil, zero value otherwise.

### GetRemainingCreditsOk

`func (o *DescribePlanResult) GetRemainingCreditsOk() (*string, bool)`

GetRemainingCreditsOk returns a tuple with the RemainingCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCredits

`func (o *DescribePlanResult) SetRemainingCredits(v string)`

SetRemainingCredits sets RemainingCredits field to given value.

### HasRemainingCredits

`func (o *DescribePlanResult) HasRemainingCredits() bool`

HasRemainingCredits returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribePlanResult) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribePlanResult) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribePlanResult) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribePlanResult) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


