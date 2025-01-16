# GetRegionByCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Code** | **string** | Cloud-provider native region code | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetRegionByCodeRequest

`func NewGetRegionByCodeRequest(cloudProviderName string, code string, token string, ) *GetRegionByCodeRequest`

NewGetRegionByCodeRequest instantiates a new GetRegionByCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionByCodeRequestWithDefaults

`func NewGetRegionByCodeRequestWithDefaults() *GetRegionByCodeRequest`

NewGetRegionByCodeRequestWithDefaults instantiates a new GetRegionByCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *GetRegionByCodeRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *GetRegionByCodeRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *GetRegionByCodeRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCode

`func (o *GetRegionByCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetRegionByCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetRegionByCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetToken

`func (o *GetRegionByCodeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetRegionByCodeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetRegionByCodeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


