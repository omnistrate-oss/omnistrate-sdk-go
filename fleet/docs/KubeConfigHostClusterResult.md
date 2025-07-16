# KubeConfigHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiServerEndpoint** | **string** | The API server endpoint of the host cluster | 
**CaDataBase64** | **string** | Base64 encoded CA data for the host cluster | 
**ClientCertificateDataBase64** | **string** | Base64 encoded client certificate data for the host cluster | 
**ClientKeyDataBase64** | **string** | Base64 encoded client key data for the host cluster | 
**Id** | **string** | ID of a Host Cluster | 
**ServiceAccountToken** | **string** | Service account token for the host cluster | 
**UserName** | **string** | The user name to use in the kubeconfig | 

## Methods

### NewKubeConfigHostClusterResult

`func NewKubeConfigHostClusterResult(apiServerEndpoint string, caDataBase64 string, clientCertificateDataBase64 string, clientKeyDataBase64 string, id string, serviceAccountToken string, userName string, ) *KubeConfigHostClusterResult`

NewKubeConfigHostClusterResult instantiates a new KubeConfigHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigHostClusterResultWithDefaults

`func NewKubeConfigHostClusterResultWithDefaults() *KubeConfigHostClusterResult`

NewKubeConfigHostClusterResultWithDefaults instantiates a new KubeConfigHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiServerEndpoint

`func (o *KubeConfigHostClusterResult) GetApiServerEndpoint() string`

GetApiServerEndpoint returns the ApiServerEndpoint field if non-nil, zero value otherwise.

### GetApiServerEndpointOk

`func (o *KubeConfigHostClusterResult) GetApiServerEndpointOk() (*string, bool)`

GetApiServerEndpointOk returns a tuple with the ApiServerEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerEndpoint

`func (o *KubeConfigHostClusterResult) SetApiServerEndpoint(v string)`

SetApiServerEndpoint sets ApiServerEndpoint field to given value.


### GetCaDataBase64

`func (o *KubeConfigHostClusterResult) GetCaDataBase64() string`

GetCaDataBase64 returns the CaDataBase64 field if non-nil, zero value otherwise.

### GetCaDataBase64Ok

`func (o *KubeConfigHostClusterResult) GetCaDataBase64Ok() (*string, bool)`

GetCaDataBase64Ok returns a tuple with the CaDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaDataBase64

`func (o *KubeConfigHostClusterResult) SetCaDataBase64(v string)`

SetCaDataBase64 sets CaDataBase64 field to given value.


### GetClientCertificateDataBase64

`func (o *KubeConfigHostClusterResult) GetClientCertificateDataBase64() string`

GetClientCertificateDataBase64 returns the ClientCertificateDataBase64 field if non-nil, zero value otherwise.

### GetClientCertificateDataBase64Ok

`func (o *KubeConfigHostClusterResult) GetClientCertificateDataBase64Ok() (*string, bool)`

GetClientCertificateDataBase64Ok returns a tuple with the ClientCertificateDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificateDataBase64

`func (o *KubeConfigHostClusterResult) SetClientCertificateDataBase64(v string)`

SetClientCertificateDataBase64 sets ClientCertificateDataBase64 field to given value.


### GetClientKeyDataBase64

`func (o *KubeConfigHostClusterResult) GetClientKeyDataBase64() string`

GetClientKeyDataBase64 returns the ClientKeyDataBase64 field if non-nil, zero value otherwise.

### GetClientKeyDataBase64Ok

`func (o *KubeConfigHostClusterResult) GetClientKeyDataBase64Ok() (*string, bool)`

GetClientKeyDataBase64Ok returns a tuple with the ClientKeyDataBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKeyDataBase64

`func (o *KubeConfigHostClusterResult) SetClientKeyDataBase64(v string)`

SetClientKeyDataBase64 sets ClientKeyDataBase64 field to given value.


### GetId

`func (o *KubeConfigHostClusterResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubeConfigHostClusterResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubeConfigHostClusterResult) SetId(v string)`

SetId sets Id field to given value.


### GetServiceAccountToken

`func (o *KubeConfigHostClusterResult) GetServiceAccountToken() string`

GetServiceAccountToken returns the ServiceAccountToken field if non-nil, zero value otherwise.

### GetServiceAccountTokenOk

`func (o *KubeConfigHostClusterResult) GetServiceAccountTokenOk() (*string, bool)`

GetServiceAccountTokenOk returns a tuple with the ServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountToken

`func (o *KubeConfigHostClusterResult) SetServiceAccountToken(v string)`

SetServiceAccountToken sets ServiceAccountToken field to given value.


### GetUserName

`func (o *KubeConfigHostClusterResult) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *KubeConfigHostClusterResult) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *KubeConfigHostClusterResult) SetUserName(v string)`

SetUserName sets UserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


