# GetCloudProviderByNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the CloudProvider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewGetCloudProviderByNameRequest

`func NewGetCloudProviderByNameRequest(name string, token string, ) *GetCloudProviderByNameRequest`

NewGetCloudProviderByNameRequest instantiates a new GetCloudProviderByNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCloudProviderByNameRequestWithDefaults

`func NewGetCloudProviderByNameRequestWithDefaults() *GetCloudProviderByNameRequest`

NewGetCloudProviderByNameRequestWithDefaults instantiates a new GetCloudProviderByNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetCloudProviderByNameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCloudProviderByNameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCloudProviderByNameRequest) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *GetCloudProviderByNameRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetCloudProviderByNameRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetCloudProviderByNameRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


