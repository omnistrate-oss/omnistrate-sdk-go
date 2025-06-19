# DescribeIdentityProviderTypeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportCustomEndpoints** | **bool** | Whether the Identity Provider supports custom endpoints | 
**SupportCustomScopes** | **bool** | Whether the Identity Provider supports custom scopes | 
**IdentityProviderName** | **string** | The name of the identity provider | 

## Methods

### NewDescribeIdentityProviderTypeResult

`func NewDescribeIdentityProviderTypeResult(supportCustomEndpoints bool, supportCustomScopes bool, identityProviderName string, ) *DescribeIdentityProviderTypeResult`

NewDescribeIdentityProviderTypeResult instantiates a new DescribeIdentityProviderTypeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIdentityProviderTypeResultWithDefaults

`func NewDescribeIdentityProviderTypeResultWithDefaults() *DescribeIdentityProviderTypeResult`

NewDescribeIdentityProviderTypeResultWithDefaults instantiates a new DescribeIdentityProviderTypeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportCustomEndpoints

`func (o *DescribeIdentityProviderTypeResult) GetSupportCustomEndpoints() bool`

GetSupportCustomEndpoints returns the SupportCustomEndpoints field if non-nil, zero value otherwise.

### GetSupportCustomEndpointsOk

`func (o *DescribeIdentityProviderTypeResult) GetSupportCustomEndpointsOk() (*bool, bool)`

GetSupportCustomEndpointsOk returns a tuple with the SupportCustomEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomEndpoints

`func (o *DescribeIdentityProviderTypeResult) SetSupportCustomEndpoints(v bool)`

SetSupportCustomEndpoints sets SupportCustomEndpoints field to given value.


### GetSupportCustomScopes

`func (o *DescribeIdentityProviderTypeResult) GetSupportCustomScopes() bool`

GetSupportCustomScopes returns the SupportCustomScopes field if non-nil, zero value otherwise.

### GetSupportCustomScopesOk

`func (o *DescribeIdentityProviderTypeResult) GetSupportCustomScopesOk() (*bool, bool)`

GetSupportCustomScopesOk returns a tuple with the SupportCustomScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomScopes

`func (o *DescribeIdentityProviderTypeResult) SetSupportCustomScopes(v bool)`

SetSupportCustomScopes sets SupportCustomScopes field to given value.


### GetIdentityProviderName

`func (o *DescribeIdentityProviderTypeResult) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *DescribeIdentityProviderTypeResult) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *DescribeIdentityProviderTypeResult) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


