# ListSubscriptionUsageReportingAttemptsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptedAtFrom** | Pointer to **time.Time** | Only attempts at or after this RFC3339 instant. Defaults to seven days ago | [optional] 
**AttemptedAtTo** | Pointer to **time.Time** | Only attempts at or before this RFC3339 instant. Defaults to now | [optional] 
**BillingChannel** | Pointer to **string** | Filter to one downstream billing channel | [optional] 
**Dimension** | Pointer to **string** |  | [optional] 
**FailuresOnly** | Pointer to **bool** |  | [optional] 
**Limit** | Pointer to **int64** | Maximum attempts returned | [optional] 
**Status** | Pointer to **string** | The current outcome represented by one usage reporting attempt | [optional] 
**SubscriptionId** | **string** | The Omnistrate subscription identifier | 
**Token** | **string** | JWT token used to perform authorization | 
**UsageReportRecordId** | Pointer to **string** | Filter to one usage report idempotency key | [optional] 

## Methods

### NewListSubscriptionUsageReportingAttemptsRequest

`func NewListSubscriptionUsageReportingAttemptsRequest(subscriptionId string, token string, ) *ListSubscriptionUsageReportingAttemptsRequest`

NewListSubscriptionUsageReportingAttemptsRequest instantiates a new ListSubscriptionUsageReportingAttemptsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubscriptionUsageReportingAttemptsRequestWithDefaults

`func NewListSubscriptionUsageReportingAttemptsRequestWithDefaults() *ListSubscriptionUsageReportingAttemptsRequest`

NewListSubscriptionUsageReportingAttemptsRequestWithDefaults instantiates a new ListSubscriptionUsageReportingAttemptsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptedAtFrom

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetAttemptedAtFrom() time.Time`

GetAttemptedAtFrom returns the AttemptedAtFrom field if non-nil, zero value otherwise.

### GetAttemptedAtFromOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetAttemptedAtFromOk() (*time.Time, bool)`

GetAttemptedAtFromOk returns a tuple with the AttemptedAtFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAtFrom

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetAttemptedAtFrom(v time.Time)`

SetAttemptedAtFrom sets AttemptedAtFrom field to given value.

### HasAttemptedAtFrom

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasAttemptedAtFrom() bool`

HasAttemptedAtFrom returns a boolean if a field has been set.

### GetAttemptedAtTo

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetAttemptedAtTo() time.Time`

GetAttemptedAtTo returns the AttemptedAtTo field if non-nil, zero value otherwise.

### GetAttemptedAtToOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetAttemptedAtToOk() (*time.Time, bool)`

GetAttemptedAtToOk returns a tuple with the AttemptedAtTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedAtTo

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetAttemptedAtTo(v time.Time)`

SetAttemptedAtTo sets AttemptedAtTo field to given value.

### HasAttemptedAtTo

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasAttemptedAtTo() bool`

HasAttemptedAtTo returns a boolean if a field has been set.

### GetBillingChannel

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetBillingChannel() string`

GetBillingChannel returns the BillingChannel field if non-nil, zero value otherwise.

### GetBillingChannelOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetBillingChannelOk() (*string, bool)`

GetBillingChannelOk returns a tuple with the BillingChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingChannel

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetBillingChannel(v string)`

SetBillingChannel sets BillingChannel field to given value.

### HasBillingChannel

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasBillingChannel() bool`

HasBillingChannel returns a boolean if a field has been set.

### GetDimension

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetFailuresOnly

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetFailuresOnly() bool`

GetFailuresOnly returns the FailuresOnly field if non-nil, zero value otherwise.

### GetFailuresOnlyOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetFailuresOnlyOk() (*bool, bool)`

GetFailuresOnlyOk returns a tuple with the FailuresOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailuresOnly

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetFailuresOnly(v bool)`

SetFailuresOnly sets FailuresOnly field to given value.

### HasFailuresOnly

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasFailuresOnly() bool`

HasFailuresOnly returns a boolean if a field has been set.

### GetLimit

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetStatus

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetToken

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetUsageReportRecordId

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetUsageReportRecordId() string`

GetUsageReportRecordId returns the UsageReportRecordId field if non-nil, zero value otherwise.

### GetUsageReportRecordIdOk

`func (o *ListSubscriptionUsageReportingAttemptsRequest) GetUsageReportRecordIdOk() (*string, bool)`

GetUsageReportRecordIdOk returns a tuple with the UsageReportRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageReportRecordId

`func (o *ListSubscriptionUsageReportingAttemptsRequest) SetUsageReportRecordId(v string)`

SetUsageReportRecordId sets UsageReportRecordId field to given value.

### HasUsageReportRecordId

`func (o *ListSubscriptionUsageReportingAttemptsRequest) HasUsageReportRecordId() bool`

HasUsageReportRecordId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


