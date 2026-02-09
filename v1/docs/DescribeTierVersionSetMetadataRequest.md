# DescribeTierVersionSetMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Version** | **string** | The version number for the specific version set. | 

## Methods

### NewDescribeTierVersionSetMetadataRequest

`func NewDescribeTierVersionSetMetadataRequest(productTierId string, serviceId string, token string, version string, ) *DescribeTierVersionSetMetadataRequest`

NewDescribeTierVersionSetMetadataRequest instantiates a new DescribeTierVersionSetMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTierVersionSetMetadataRequestWithDefaults

`func NewDescribeTierVersionSetMetadataRequestWithDefaults() *DescribeTierVersionSetMetadataRequest`

NewDescribeTierVersionSetMetadataRequestWithDefaults instantiates a new DescribeTierVersionSetMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductTierId

`func (o *DescribeTierVersionSetMetadataRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DescribeTierVersionSetMetadataRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DescribeTierVersionSetMetadataRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *DescribeTierVersionSetMetadataRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DescribeTierVersionSetMetadataRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DescribeTierVersionSetMetadataRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DescribeTierVersionSetMetadataRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeTierVersionSetMetadataRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeTierVersionSetMetadataRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVersion

`func (o *DescribeTierVersionSetMetadataRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DescribeTierVersionSetMetadataRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DescribeTierVersionSetMetadataRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


