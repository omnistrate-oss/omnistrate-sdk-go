# DescribeMarketplaceContractRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | Which marketplace channel a contract came from. SUGER reaches AWS, Azure and GCP buyers through one listing. SANDBOX is the simulated channel, and is a real member of this set rather than a test mode | [optional] 
**ExternalRef** | Pointer to **string** |  | [optional] 
**Id** | **string** | The Omnistrate contract identifier | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeMarketplaceContractRequest

`func NewDescribeMarketplaceContractRequest(id string, token string, ) *DescribeMarketplaceContractRequest`

NewDescribeMarketplaceContractRequest instantiates a new DescribeMarketplaceContractRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMarketplaceContractRequestWithDefaults

`func NewDescribeMarketplaceContractRequestWithDefaults() *DescribeMarketplaceContractRequest`

NewDescribeMarketplaceContractRequestWithDefaults instantiates a new DescribeMarketplaceContractRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DescribeMarketplaceContractRequest) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DescribeMarketplaceContractRequest) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DescribeMarketplaceContractRequest) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DescribeMarketplaceContractRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetExternalRef

`func (o *DescribeMarketplaceContractRequest) GetExternalRef() string`

GetExternalRef returns the ExternalRef field if non-nil, zero value otherwise.

### GetExternalRefOk

`func (o *DescribeMarketplaceContractRequest) GetExternalRefOk() (*string, bool)`

GetExternalRefOk returns a tuple with the ExternalRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRef

`func (o *DescribeMarketplaceContractRequest) SetExternalRef(v string)`

SetExternalRef sets ExternalRef field to given value.

### HasExternalRef

`func (o *DescribeMarketplaceContractRequest) HasExternalRef() bool`

HasExternalRef returns a boolean if a field has been set.

### GetId

`func (o *DescribeMarketplaceContractRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeMarketplaceContractRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeMarketplaceContractRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *DescribeMarketplaceContractRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeMarketplaceContractRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeMarketplaceContractRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


