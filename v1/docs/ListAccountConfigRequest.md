# ListAccountConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAccountConfigRequest

`func NewListAccountConfigRequest(cloudProviderName string, token string, ) *ListAccountConfigRequest`

NewListAccountConfigRequest instantiates a new ListAccountConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountConfigRequestWithDefaults

`func NewListAccountConfigRequestWithDefaults() *ListAccountConfigRequest`

NewListAccountConfigRequestWithDefaults instantiates a new ListAccountConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *ListAccountConfigRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ListAccountConfigRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ListAccountConfigRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetToken

`func (o *ListAccountConfigRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAccountConfigRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAccountConfigRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


