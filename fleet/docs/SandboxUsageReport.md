# SandboxUsageReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempts** | **int64** |  | 
**ChannelCode** | Pointer to **string** | The channel-specific status code from the last report attempt | [optional] 
**ChannelError** | Pointer to **string** | The channel-specific error from the last report attempt | [optional] 
**CreatedAt** | **time.Time** |  | 
**Dimension** | **string** |  | 
**ExternalRef** | **string** | The sandbox channel&#39;s entitlement or contract identifier | 
**LastAttempt** | Pointer to **time.Time** |  | [optional] 
**MarketplaceContractId** | **string** |  | 
**Quantity** | **float64** |  | 
**RecordId** | **string** | The deterministic marketplace usage idempotency key | 
**Status** | **string** | What happened to one sandbox usage record on its way to the channel | 
**SubscriptionId** | Pointer to **string** | The Omnistrate subscription attached to the simulated contract, when available | [optional] 
**UpdatedAt** | **time.Time** |  | 
**WindowEnd** | **time.Time** |  | 
**WindowStart** | **time.Time** |  | 

## Methods

### NewSandboxUsageReport

`func NewSandboxUsageReport(attempts int64, createdAt time.Time, dimension string, externalRef string, marketplaceContractId string, quantity float64, recordId string, status string, updatedAt time.Time, windowEnd time.Time, windowStart time.Time, ) *SandboxUsageReport`

NewSandboxUsageReport instantiates a new SandboxUsageReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxUsageReportWithDefaults

`func NewSandboxUsageReportWithDefaults() *SandboxUsageReport`

NewSandboxUsageReportWithDefaults instantiates a new SandboxUsageReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempts

`func (o *SandboxUsageReport) GetAttempts() int64`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *SandboxUsageReport) GetAttemptsOk() (*int64, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *SandboxUsageReport) SetAttempts(v int64)`

SetAttempts sets Attempts field to given value.


### GetChannelCode

`func (o *SandboxUsageReport) GetChannelCode() string`

GetChannelCode returns the ChannelCode field if non-nil, zero value otherwise.

### GetChannelCodeOk

`func (o *SandboxUsageReport) GetChannelCodeOk() (*string, bool)`

GetChannelCodeOk returns a tuple with the ChannelCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelCode

`func (o *SandboxUsageReport) SetChannelCode(v string)`

SetChannelCode sets ChannelCode field to given value.

### HasChannelCode

`func (o *SandboxUsageReport) HasChannelCode() bool`

HasChannelCode returns a boolean if a field has been set.

### GetChannelError

`func (o *SandboxUsageReport) GetChannelError() string`

GetChannelError returns the ChannelError field if non-nil, zero value otherwise.

### GetChannelErrorOk

`func (o *SandboxUsageReport) GetChannelErrorOk() (*string, bool)`

GetChannelErrorOk returns a tuple with the ChannelError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelError

`func (o *SandboxUsageReport) SetChannelError(v string)`

SetChannelError sets ChannelError field to given value.

### HasChannelError

`func (o *SandboxUsageReport) HasChannelError() bool`

HasChannelError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SandboxUsageReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SandboxUsageReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SandboxUsageReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDimension

`func (o *SandboxUsageReport) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *SandboxUsageReport) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *SandboxUsageReport) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetExternalRef

`func (o *SandboxUsageReport) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *SandboxUsageReport) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *SandboxUsageReport) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.


### GetLastAttempt

`func (o *SandboxUsageReport) GetLastAttempt() time.Time`

GetLastAttempt returns the LastAttempt field if non-nil, zero value otherwise.

### GetLastAttemptOk

`func (o *SandboxUsageReport) GetLastAttemptOk() (*time.Time, bool)`

GetLastAttemptOk returns a tuple with the LastAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttempt

`func (o *SandboxUsageReport) SetLastAttempt(v time.Time)`

SetLastAttempt sets LastAttempt field to given value.

### HasLastAttempt

`func (o *SandboxUsageReport) HasLastAttempt() bool`

HasLastAttempt returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *SandboxUsageReport) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *SandboxUsageReport) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *SandboxUsageReport) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.


### GetQuantity

`func (o *SandboxUsageReport) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SandboxUsageReport) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SandboxUsageReport) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### GetRecordId

`func (o *SandboxUsageReport) GetRecordId() string`

GetRecordId returns the RecordId field if non-nil, zero value otherwise.

### GetRecordIdOk

`func (o *SandboxUsageReport) GetRecordIdOk() (*string, bool)`

GetRecordIdOk returns a tuple with the RecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordId

`func (o *SandboxUsageReport) SetRecordId(v string)`

SetRecordId sets RecordId field to given value.


### GetStatus

`func (o *SandboxUsageReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SandboxUsageReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SandboxUsageReport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubscriptionId

`func (o *SandboxUsageReport) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SandboxUsageReport) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SandboxUsageReport) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SandboxUsageReport) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SandboxUsageReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SandboxUsageReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SandboxUsageReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWindowEnd

`func (o *SandboxUsageReport) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *SandboxUsageReport) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *SandboxUsageReport) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetWindowStart

`func (o *SandboxUsageReport) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *SandboxUsageReport) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *SandboxUsageReport) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


