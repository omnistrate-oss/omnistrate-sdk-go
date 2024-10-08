# DescribeIdentityProviderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Client ID of the Identity Provider | 
**ClientSecret** | Pointer to **string** | The Client Secret of the Identity Provider | [optional] 
**Id** | **string** | The Identity Provider ID | 
**IdentityProviderName** | **string** | The name of the Identity Provider | 
**Status** | **string** | The status of the Identity Provider | 

## Methods

### NewDescribeIdentityProviderResult

`func NewDescribeIdentityProviderResult(clientId string, id string, identityProviderName string, status string, ) *DescribeIdentityProviderResult`

NewDescribeIdentityProviderResult instantiates a new DescribeIdentityProviderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIdentityProviderResultWithDefaults

`func NewDescribeIdentityProviderResultWithDefaults() *DescribeIdentityProviderResult`

NewDescribeIdentityProviderResultWithDefaults instantiates a new DescribeIdentityProviderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DescribeIdentityProviderResult) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DescribeIdentityProviderResult) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DescribeIdentityProviderResult) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *DescribeIdentityProviderResult) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DescribeIdentityProviderResult) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DescribeIdentityProviderResult) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *DescribeIdentityProviderResult) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetId

`func (o *DescribeIdentityProviderResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeIdentityProviderResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeIdentityProviderResult) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityProviderName

`func (o *DescribeIdentityProviderResult) GetIdentityProviderName() string`

GetIdentityProviderName returns the IdentityProviderName field if non-nil, zero value otherwise.

### GetIdentityProviderNameOk

`func (o *DescribeIdentityProviderResult) GetIdentityProviderNameOk() (*string, bool)`

GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderName

`func (o *DescribeIdentityProviderResult) SetIdentityProviderName(v string)`

SetIdentityProviderName sets IdentityProviderName field to given value.


### GetStatus

`func (o *DescribeIdentityProviderResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeIdentityProviderResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeIdentityProviderResult) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


