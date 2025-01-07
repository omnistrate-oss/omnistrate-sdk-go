# CustomDNSEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ARecordTarget** | Pointer to **string** |  | [optional] 
**CnameTarget** | Pointer to **string** |  | [optional] 
**DnsHost** | Pointer to **string** |  | [optional] 
**DnsName** | Pointer to **string** | DEPRECATED: Use dnsHost instead | [optional] 
**Enabled** | **bool** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomDNSEndpoint

`func NewCustomDNSEndpoint(enabled bool, ) *CustomDNSEndpoint`

NewCustomDNSEndpoint instantiates a new CustomDNSEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDNSEndpointWithDefaults

`func NewCustomDNSEndpointWithDefaults() *CustomDNSEndpoint`

NewCustomDNSEndpointWithDefaults instantiates a new CustomDNSEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetARecordTarget

`func (o *CustomDNSEndpoint) GetARecordTarget() string`

GetARecordTarget returns the ARecordTarget field if non-nil, zero value otherwise.

### GetARecordTargetOk

`func (o *CustomDNSEndpoint) GetARecordTargetOk() (*string, bool)`

GetARecordTargetOk returns a tuple with the ARecordTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetARecordTarget

`func (o *CustomDNSEndpoint) SetARecordTarget(v string)`

SetARecordTarget sets ARecordTarget field to given value.

### HasARecordTarget

`func (o *CustomDNSEndpoint) HasARecordTarget() bool`

HasARecordTarget returns a boolean if a field has been set.

### GetCnameTarget

`func (o *CustomDNSEndpoint) GetCnameTarget() string`

GetCnameTarget returns the CnameTarget field if non-nil, zero value otherwise.

### GetCnameTargetOk

`func (o *CustomDNSEndpoint) GetCnameTargetOk() (*string, bool)`

GetCnameTargetOk returns a tuple with the CnameTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnameTarget

`func (o *CustomDNSEndpoint) SetCnameTarget(v string)`

SetCnameTarget sets CnameTarget field to given value.

### HasCnameTarget

`func (o *CustomDNSEndpoint) HasCnameTarget() bool`

HasCnameTarget returns a boolean if a field has been set.

### GetDnsHost

`func (o *CustomDNSEndpoint) GetDnsHost() string`

GetDnsHost returns the DnsHost field if non-nil, zero value otherwise.

### GetDnsHostOk

`func (o *CustomDNSEndpoint) GetDnsHostOk() (*string, bool)`

GetDnsHostOk returns a tuple with the DnsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsHost

`func (o *CustomDNSEndpoint) SetDnsHost(v string)`

SetDnsHost sets DnsHost field to given value.

### HasDnsHost

`func (o *CustomDNSEndpoint) HasDnsHost() bool`

HasDnsHost returns a boolean if a field has been set.

### GetDnsName

`func (o *CustomDNSEndpoint) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *CustomDNSEndpoint) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *CustomDNSEndpoint) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *CustomDNSEndpoint) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetEnabled

`func (o *CustomDNSEndpoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomDNSEndpoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomDNSEndpoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetName

`func (o *CustomDNSEndpoint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomDNSEndpoint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomDNSEndpoint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomDNSEndpoint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDNSEndpoint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDNSEndpoint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDNSEndpoint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDNSEndpoint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


