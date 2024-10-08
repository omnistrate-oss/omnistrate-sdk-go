# ListIdentityProvidersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviders** | [**[]DescribeIdentityProviderResult**](DescribeIdentityProviderResult.md) | The list of Identity Providers | 

## Methods

### NewListIdentityProvidersResult

`func NewListIdentityProvidersResult(identityProviders []DescribeIdentityProviderResult, ) *ListIdentityProvidersResult`

NewListIdentityProvidersResult instantiates a new ListIdentityProvidersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentityProvidersResultWithDefaults

`func NewListIdentityProvidersResultWithDefaults() *ListIdentityProvidersResult`

NewListIdentityProvidersResultWithDefaults instantiates a new ListIdentityProvidersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviders

`func (o *ListIdentityProvidersResult) GetIdentityProviders() []DescribeIdentityProviderResult`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *ListIdentityProvidersResult) GetIdentityProvidersOk() (*[]DescribeIdentityProviderResult, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *ListIdentityProvidersResult) SetIdentityProviders(v []DescribeIdentityProviderResult)`

SetIdentityProviders sets IdentityProviders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


