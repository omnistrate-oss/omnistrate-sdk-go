# InfrastructureIngress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]InfrastructureAddress**](InfrastructureAddress.md) |  | [optional] 
**AddressesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**ClassName** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to [**InfrastructureObjectReference**](InfrastructureObjectReference.md) |  | [optional] 
**Rules** | Pointer to [**[]InfrastructureIngressRule**](InfrastructureIngressRule.md) |  | [optional] 
**RulesTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**Tls** | Pointer to [**[]InfrastructureTLS**](InfrastructureTLS.md) |  | [optional] 
**TlsTermination** | Pointer to **string** |  | [optional] 
**TlsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 

## Methods

### NewInfrastructureIngress

`func NewInfrastructureIngress() *InfrastructureIngress`

NewInfrastructureIngress instantiates a new InfrastructureIngress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureIngressWithDefaults

`func NewInfrastructureIngressWithDefaults() *InfrastructureIngress`

NewInfrastructureIngressWithDefaults instantiates a new InfrastructureIngress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *InfrastructureIngress) GetAddresses() []InfrastructureAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *InfrastructureIngress) GetAddressesOk() (*[]InfrastructureAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *InfrastructureIngress) SetAddresses(v []InfrastructureAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *InfrastructureIngress) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAddressesTruncation

`func (o *InfrastructureIngress) GetAddressesTruncation() CheckpointTruncation`

GetAddressesTruncation returns the AddressesTruncation field if non-nil, zero value otherwise.

### GetAddressesTruncationOk

`func (o *InfrastructureIngress) GetAddressesTruncationOk() (*CheckpointTruncation, bool)`

GetAddressesTruncationOk returns a tuple with the AddressesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesTruncation

`func (o *InfrastructureIngress) SetAddressesTruncation(v CheckpointTruncation)`

SetAddressesTruncation sets AddressesTruncation field to given value.

### HasAddressesTruncation

`func (o *InfrastructureIngress) HasAddressesTruncation() bool`

HasAddressesTruncation returns a boolean if a field has been set.

### GetClassName

`func (o *InfrastructureIngress) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *InfrastructureIngress) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *InfrastructureIngress) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *InfrastructureIngress) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetRef

`func (o *InfrastructureIngress) GetRef() InfrastructureObjectReference`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *InfrastructureIngress) GetRefOk() (*InfrastructureObjectReference, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *InfrastructureIngress) SetRef(v InfrastructureObjectReference)`

SetRef sets Ref field to given value.

### HasRef

`func (o *InfrastructureIngress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRules

`func (o *InfrastructureIngress) GetRules() []InfrastructureIngressRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InfrastructureIngress) GetRulesOk() (*[]InfrastructureIngressRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InfrastructureIngress) SetRules(v []InfrastructureIngressRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InfrastructureIngress) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetRulesTruncation

`func (o *InfrastructureIngress) GetRulesTruncation() CheckpointTruncation`

GetRulesTruncation returns the RulesTruncation field if non-nil, zero value otherwise.

### GetRulesTruncationOk

`func (o *InfrastructureIngress) GetRulesTruncationOk() (*CheckpointTruncation, bool)`

GetRulesTruncationOk returns a tuple with the RulesTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesTruncation

`func (o *InfrastructureIngress) SetRulesTruncation(v CheckpointTruncation)`

SetRulesTruncation sets RulesTruncation field to given value.

### HasRulesTruncation

`func (o *InfrastructureIngress) HasRulesTruncation() bool`

HasRulesTruncation returns a boolean if a field has been set.

### GetTls

`func (o *InfrastructureIngress) GetTls() []InfrastructureTLS`

GetTls returns the Tls field if non-nil, zero value otherwise.

### GetTlsOk

`func (o *InfrastructureIngress) GetTlsOk() (*[]InfrastructureTLS, bool)`

GetTlsOk returns a tuple with the Tls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTls

`func (o *InfrastructureIngress) SetTls(v []InfrastructureTLS)`

SetTls sets Tls field to given value.

### HasTls

`func (o *InfrastructureIngress) HasTls() bool`

HasTls returns a boolean if a field has been set.

### GetTlsTermination

`func (o *InfrastructureIngress) GetTlsTermination() string`

GetTlsTermination returns the TlsTermination field if non-nil, zero value otherwise.

### GetTlsTerminationOk

`func (o *InfrastructureIngress) GetTlsTerminationOk() (*string, bool)`

GetTlsTerminationOk returns a tuple with the TlsTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTermination

`func (o *InfrastructureIngress) SetTlsTermination(v string)`

SetTlsTermination sets TlsTermination field to given value.

### HasTlsTermination

`func (o *InfrastructureIngress) HasTlsTermination() bool`

HasTlsTermination returns a boolean if a field has been set.

### GetTlsTruncation

`func (o *InfrastructureIngress) GetTlsTruncation() CheckpointTruncation`

GetTlsTruncation returns the TlsTruncation field if non-nil, zero value otherwise.

### GetTlsTruncationOk

`func (o *InfrastructureIngress) GetTlsTruncationOk() (*CheckpointTruncation, bool)`

GetTlsTruncationOk returns a tuple with the TlsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTruncation

`func (o *InfrastructureIngress) SetTlsTruncation(v CheckpointTruncation)`

SetTlsTruncation sets TlsTruncation field to given value.

### HasTlsTruncation

`func (o *InfrastructureIngress) HasTlsTruncation() bool`

HasTlsTruncation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


