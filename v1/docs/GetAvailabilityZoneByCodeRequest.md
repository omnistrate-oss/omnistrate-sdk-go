# GetAvailabilityZoneByCodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Code** | **string** | Cloud-provider native availability zone code | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetAvailabilityZoneByCodeRequest

`func NewGetAvailabilityZoneByCodeRequest(cloudProviderName string, code string, token string, ) *GetAvailabilityZoneByCodeRequest`

NewGetAvailabilityZoneByCodeRequest instantiates a new GetAvailabilityZoneByCodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAvailabilityZoneByCodeRequestWithDefaults

`func NewGetAvailabilityZoneByCodeRequestWithDefaults() *GetAvailabilityZoneByCodeRequest`

NewGetAvailabilityZoneByCodeRequestWithDefaults instantiates a new GetAvailabilityZoneByCodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *GetAvailabilityZoneByCodeRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *GetAvailabilityZoneByCodeRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *GetAvailabilityZoneByCodeRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetCode

`func (o *GetAvailabilityZoneByCodeRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetAvailabilityZoneByCodeRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetAvailabilityZoneByCodeRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetToken

`func (o *GetAvailabilityZoneByCodeRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetAvailabilityZoneByCodeRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetAvailabilityZoneByCodeRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


