# ImportAccountConfigCloudNativeNetworkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) to import for deployments | 
**Id** | **string** | ID of an Account Config | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewImportAccountConfigCloudNativeNetworkRequest

`func NewImportAccountConfigCloudNativeNetworkRequest(cloudNativeNetworkId string, id string, token string, ) *ImportAccountConfigCloudNativeNetworkRequest`

NewImportAccountConfigCloudNativeNetworkRequest instantiates a new ImportAccountConfigCloudNativeNetworkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportAccountConfigCloudNativeNetworkRequestWithDefaults

`func NewImportAccountConfigCloudNativeNetworkRequestWithDefaults() *ImportAccountConfigCloudNativeNetworkRequest`

NewImportAccountConfigCloudNativeNetworkRequestWithDefaults instantiates a new ImportAccountConfigCloudNativeNetworkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *ImportAccountConfigCloudNativeNetworkRequest) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetId

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportAccountConfigCloudNativeNetworkRequest) SetId(v string)`

SetId sets Id field to given value.


### GetToken

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ImportAccountConfigCloudNativeNetworkRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ImportAccountConfigCloudNativeNetworkRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


