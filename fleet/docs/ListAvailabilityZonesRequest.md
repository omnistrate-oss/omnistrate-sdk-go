# ListAvailabilityZonesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderName** | **string** | Name of the Infra Provider | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListAvailabilityZonesRequest

`func NewListAvailabilityZonesRequest(cloudProviderName string, token string, ) *ListAvailabilityZonesRequest`

NewListAvailabilityZonesRequest instantiates a new ListAvailabilityZonesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailabilityZonesRequestWithDefaults

`func NewListAvailabilityZonesRequestWithDefaults() *ListAvailabilityZonesRequest`

NewListAvailabilityZonesRequestWithDefaults instantiates a new ListAvailabilityZonesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderName

`func (o *ListAvailabilityZonesRequest) GetCloudProviderName() string`

GetCloudProviderName returns the CloudProviderName field if non-nil, zero value otherwise.

### GetCloudProviderNameOk

`func (o *ListAvailabilityZonesRequest) GetCloudProviderNameOk() (*string, bool)`

GetCloudProviderNameOk returns a tuple with the CloudProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderName

`func (o *ListAvailabilityZonesRequest) SetCloudProviderName(v string)`

SetCloudProviderName sets CloudProviderName field to given value.


### GetToken

`func (o *ListAvailabilityZonesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListAvailabilityZonesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListAvailabilityZonesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


