# SandboxRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]SandboxCheck**](SandboxCheck.md) | All seven conformance checks and their results | [optional] 
**Deliveries** | Pointer to [**[]SandboxDelivery**](SandboxDelivery.md) | Every delivery made during the run, newest first | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**MarketplaceContractId** | Pointer to **string** | The simulated contract the run drove. Its IsSimulated flag is true and it is excluded from every revenue rollup by construction | [optional] 
**ReceiverUrl** | **string** | Where the sandbox delivers. HTTPS only, and private, loopback, link-local and cloud metadata addresses are refused, at registration and again at connect time | 
**RunId** | **string** |  | 
**SigningSecretId** | Pointer to **string** | Identifier of the secret deliveries are signed with. The secret itself is shown once when it is created or rotated and is never returned again by any endpoint | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | **string** |  | 

## Methods

### NewSandboxRun

`func NewSandboxRun(receiverUrl string, runId string, status string, ) *SandboxRun`

NewSandboxRun instantiates a new SandboxRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxRunWithDefaults

`func NewSandboxRunWithDefaults() *SandboxRun`

NewSandboxRunWithDefaults instantiates a new SandboxRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *SandboxRun) GetChecks() []SandboxCheck`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *SandboxRun) GetChecksOk() (*[]SandboxCheck, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *SandboxRun) SetChecks(v []SandboxCheck)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *SandboxRun) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetDeliveries

`func (o *SandboxRun) GetDeliveries() []SandboxDelivery`

GetDeliveries returns the Deliveries field if non-nil, zero value otherwise.

### GetDeliveriesOk

`func (o *SandboxRun) GetDeliveriesOk() (*[]SandboxDelivery, bool)`

GetDeliveriesOk returns a tuple with the Deliveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveries

`func (o *SandboxRun) SetDeliveries(v []SandboxDelivery)`

SetDeliveries sets Deliveries field to given value.

### HasDeliveries

`func (o *SandboxRun) HasDeliveries() bool`

HasDeliveries returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SandboxRun) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SandboxRun) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SandboxRun) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SandboxRun) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMarketplaceContractId

`func (o *SandboxRun) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *SandboxRun) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *SandboxRun) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *SandboxRun) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetReceiverUrl

`func (o *SandboxRun) GetReceiverUrl() string`

GetReceiverUrl returns the ReceiverUrl field if non-nil, zero value otherwise.

### GetReceiverUrlOk

`func (o *SandboxRun) GetReceiverUrlOk() (*string, bool)`

GetReceiverUrlOk returns a tuple with the ReceiverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverUrl

`func (o *SandboxRun) SetReceiverUrl(v string)`

SetReceiverUrl sets ReceiverUrl field to given value.


### GetRunId

`func (o *SandboxRun) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *SandboxRun) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *SandboxRun) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetSigningSecretId

`func (o *SandboxRun) GetSigningSecretId() string`

GetSigningSecretId returns the SigningSecretId field if non-nil, zero value otherwise.

### GetSigningSecretIdOk

`func (o *SandboxRun) GetSigningSecretIdOk() (*string, bool)`

GetSigningSecretIdOk returns a tuple with the SigningSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningSecretId

`func (o *SandboxRun) SetSigningSecretId(v string)`

SetSigningSecretId sets SigningSecretId field to given value.

### HasSigningSecretId

`func (o *SandboxRun) HasSigningSecretId() bool`

HasSigningSecretId returns a boolean if a field has been set.

### GetStartedAt

`func (o *SandboxRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SandboxRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SandboxRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SandboxRun) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SandboxRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SandboxRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SandboxRun) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


