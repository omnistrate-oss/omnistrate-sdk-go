# SaveHelmPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmPackage** | [**HelmPackage**](HelmPackage.md) |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSaveHelmPackageRequest

`func NewSaveHelmPackageRequest(helmPackage HelmPackage, token string, ) *SaveHelmPackageRequest`

NewSaveHelmPackageRequest instantiates a new SaveHelmPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveHelmPackageRequestWithDefaults

`func NewSaveHelmPackageRequestWithDefaults() *SaveHelmPackageRequest`

NewSaveHelmPackageRequestWithDefaults instantiates a new SaveHelmPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelmPackage

`func (o *SaveHelmPackageRequest) GetHelmPackage() HelmPackage`

GetHelmPackage returns the HelmPackage field if non-nil, zero value otherwise.

### GetHelmPackageOk

`func (o *SaveHelmPackageRequest) GetHelmPackageOk() (*HelmPackage, bool)`

GetHelmPackageOk returns a tuple with the HelmPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmPackage

`func (o *SaveHelmPackageRequest) SetHelmPackage(v HelmPackage)`

SetHelmPackage sets HelmPackage field to given value.


### GetToken

`func (o *SaveHelmPackageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SaveHelmPackageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SaveHelmPackageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


