# InfrastructureTLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** |  | [optional] 
**HostsTruncation** | Pointer to [**CheckpointTruncation**](CheckpointTruncation.md) |  | [optional] 
**SecretName** | Pointer to **string** |  | [optional] 

## Methods

### NewInfrastructureTLS

`func NewInfrastructureTLS() *InfrastructureTLS`

NewInfrastructureTLS instantiates a new InfrastructureTLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfrastructureTLSWithDefaults

`func NewInfrastructureTLSWithDefaults() *InfrastructureTLS`

NewInfrastructureTLSWithDefaults instantiates a new InfrastructureTLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *InfrastructureTLS) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *InfrastructureTLS) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *InfrastructureTLS) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *InfrastructureTLS) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetHostsTruncation

`func (o *InfrastructureTLS) GetHostsTruncation() CheckpointTruncation`

GetHostsTruncation returns the HostsTruncation field if non-nil, zero value otherwise.

### GetHostsTruncationOk

`func (o *InfrastructureTLS) GetHostsTruncationOk() (*CheckpointTruncation, bool)`

GetHostsTruncationOk returns a tuple with the HostsTruncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsTruncation

`func (o *InfrastructureTLS) SetHostsTruncation(v CheckpointTruncation)`

SetHostsTruncation sets HostsTruncation field to given value.

### HasHostsTruncation

`func (o *InfrastructureTLS) HasHostsTruncation() bool`

HasHostsTruncation returns a boolean if a field has been set.

### GetSecretName

`func (o *InfrastructureTLS) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *InfrastructureTLS) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *InfrastructureTLS) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.

### HasSecretName

`func (o *InfrastructureTLS) HasSecretName() bool`

HasSecretName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


