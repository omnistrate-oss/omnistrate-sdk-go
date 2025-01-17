# ListHelmPackageInstallationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostClusterID** | Pointer to **string** | ID of a Host Cluster | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListHelmPackageInstallationsRequest

`func NewListHelmPackageInstallationsRequest(token string, ) *ListHelmPackageInstallationsRequest`

NewListHelmPackageInstallationsRequest instantiates a new ListHelmPackageInstallationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHelmPackageInstallationsRequestWithDefaults

`func NewListHelmPackageInstallationsRequestWithDefaults() *ListHelmPackageInstallationsRequest`

NewListHelmPackageInstallationsRequestWithDefaults instantiates a new ListHelmPackageInstallationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostClusterID

`func (o *ListHelmPackageInstallationsRequest) GetHostClusterID() string`

GetHostClusterID returns the HostClusterID field if non-nil, zero value otherwise.

### GetHostClusterIDOk

`func (o *ListHelmPackageInstallationsRequest) GetHostClusterIDOk() (*string, bool)`

GetHostClusterIDOk returns a tuple with the HostClusterID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostClusterID

`func (o *ListHelmPackageInstallationsRequest) SetHostClusterID(v string)`

SetHostClusterID sets HostClusterID field to given value.

### HasHostClusterID

`func (o *ListHelmPackageInstallationsRequest) HasHostClusterID() bool`

HasHostClusterID returns a boolean if a field has been set.

### GetToken

`func (o *ListHelmPackageInstallationsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListHelmPackageInstallationsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListHelmPackageInstallationsRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


