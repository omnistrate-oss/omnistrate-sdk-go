# AccountConfigCloudNativeNetworkOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudNativeNetworkId** | **string** | The cloud provider network ID (e.g. AWS VPC ID) | 
**Import** | **bool** | True to import (mark READY for deployments), false to unimport (revert to AVAILABLE) | 

## Methods

### NewAccountConfigCloudNativeNetworkOperation

`func NewAccountConfigCloudNativeNetworkOperation(cloudNativeNetworkId string, import_ bool, ) *AccountConfigCloudNativeNetworkOperation`

NewAccountConfigCloudNativeNetworkOperation instantiates a new AccountConfigCloudNativeNetworkOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigCloudNativeNetworkOperationWithDefaults

`func NewAccountConfigCloudNativeNetworkOperationWithDefaults() *AccountConfigCloudNativeNetworkOperation`

NewAccountConfigCloudNativeNetworkOperationWithDefaults instantiates a new AccountConfigCloudNativeNetworkOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudNativeNetworkId

`func (o *AccountConfigCloudNativeNetworkOperation) GetCloudNativeNetworkId() string`

GetCloudNativeNetworkId returns the CloudNativeNetworkId field if non-nil, zero value otherwise.

### GetCloudNativeNetworkIdOk

`func (o *AccountConfigCloudNativeNetworkOperation) GetCloudNativeNetworkIdOk() (*string, bool)`

GetCloudNativeNetworkIdOk returns a tuple with the CloudNativeNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNativeNetworkId

`func (o *AccountConfigCloudNativeNetworkOperation) SetCloudNativeNetworkId(v string)`

SetCloudNativeNetworkId sets CloudNativeNetworkId field to given value.


### GetImport

`func (o *AccountConfigCloudNativeNetworkOperation) GetImport() bool`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *AccountConfigCloudNativeNetworkOperation) GetImportOk() (*bool, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *AccountConfigCloudNativeNetworkOperation) SetImport(v bool)`

SetImport sets Import field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


