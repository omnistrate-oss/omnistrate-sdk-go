# ListSandboxUsageReportsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketplaceContractId** | Pointer to **string** | Filter to one simulated sandbox contract | [optional] 
**Status** | Pointer to **string** | What happened to one sandbox usage record on its way to the channel | [optional] 
**Token** | **string** | JWT token used to perform authorization | 
**WindowStartFrom** | Pointer to **time.Time** | Only usage windows starting at or after this RFC3339 instant | [optional] 
**WindowStartTo** | Pointer to **time.Time** | Only usage windows starting at or before this RFC3339 instant | [optional] 

## Methods

### NewListSandboxUsageReportsRequest

`func NewListSandboxUsageReportsRequest(token string, ) *ListSandboxUsageReportsRequest`

NewListSandboxUsageReportsRequest instantiates a new ListSandboxUsageReportsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSandboxUsageReportsRequestWithDefaults

`func NewListSandboxUsageReportsRequestWithDefaults() *ListSandboxUsageReportsRequest`

NewListSandboxUsageReportsRequestWithDefaults instantiates a new ListSandboxUsageReportsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketplaceContractId

`func (o *ListSandboxUsageReportsRequest) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *ListSandboxUsageReportsRequest) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *ListSandboxUsageReportsRequest) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *ListSandboxUsageReportsRequest) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetStatus

`func (o *ListSandboxUsageReportsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSandboxUsageReportsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSandboxUsageReportsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSandboxUsageReportsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ListSandboxUsageReportsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListSandboxUsageReportsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListSandboxUsageReportsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWindowStartFrom

`func (o *ListSandboxUsageReportsRequest) GetWindowStartFrom() time.Time`

GetWindowStartFrom returns the WindowStartFrom field if non-nil, zero value otherwise.

### GetWindowStartFromOk

`func (o *ListSandboxUsageReportsRequest) GetWindowStartFromOk() (*time.Time, bool)`

GetWindowStartFromOk returns a tuple with the WindowStartFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartFrom

`func (o *ListSandboxUsageReportsRequest) SetWindowStartFrom(v time.Time)`

SetWindowStartFrom sets WindowStartFrom field to given value.

### HasWindowStartFrom

`func (o *ListSandboxUsageReportsRequest) HasWindowStartFrom() bool`

HasWindowStartFrom returns a boolean if a field has been set.

### GetWindowStartTo

`func (o *ListSandboxUsageReportsRequest) GetWindowStartTo() time.Time`

GetWindowStartTo returns the WindowStartTo field if non-nil, zero value otherwise.

### GetWindowStartToOk

`func (o *ListSandboxUsageReportsRequest) GetWindowStartToOk() (*time.Time, bool)`

GetWindowStartToOk returns a tuple with the WindowStartTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStartTo

`func (o *ListSandboxUsageReportsRequest) SetWindowStartTo(v time.Time)`

SetWindowStartTo sets WindowStartTo field to given value.

### HasWindowStartTo

`func (o *ListSandboxUsageReportsRequest) HasWindowStartTo() bool`

HasWindowStartTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


