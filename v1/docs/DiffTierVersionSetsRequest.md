# DiffTierVersionSetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnotherVersion** | **string** | The target version to compare against. | 
**ProductTierId** | **string** | ID of a Product Tier | 
**ServiceId** | **string** | ID of a Service | 
**Token** | **string** | JWT token used to perform authorization | 
**Version** | **string** | The version number for the version set. | 

## Methods

### NewDiffTierVersionSetsRequest

`func NewDiffTierVersionSetsRequest(anotherVersion string, productTierId string, serviceId string, token string, version string, ) *DiffTierVersionSetsRequest`

NewDiffTierVersionSetsRequest instantiates a new DiffTierVersionSetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiffTierVersionSetsRequestWithDefaults

`func NewDiffTierVersionSetsRequestWithDefaults() *DiffTierVersionSetsRequest`

NewDiffTierVersionSetsRequestWithDefaults instantiates a new DiffTierVersionSetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnotherVersion

`func (o *DiffTierVersionSetsRequest) GetAnotherVersion() string`

GetAnotherVersion returns the AnotherVersion field if non-nil, zero value otherwise.

### GetAnotherVersionOk

`func (o *DiffTierVersionSetsRequest) GetAnotherVersionOk() (*string, bool)`

GetAnotherVersionOk returns a tuple with the AnotherVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnotherVersion

`func (o *DiffTierVersionSetsRequest) SetAnotherVersion(v string)`

SetAnotherVersion sets AnotherVersion field to given value.


### GetProductTierId

`func (o *DiffTierVersionSetsRequest) GetProductTierId() string`

GetProductTierId returns the ProductTierId field if non-nil, zero value otherwise.

### GetProductTierIdOk

`func (o *DiffTierVersionSetsRequest) GetProductTierIdOk() (*string, bool)`

GetProductTierIdOk returns a tuple with the ProductTierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTierId

`func (o *DiffTierVersionSetsRequest) SetProductTierId(v string)`

SetProductTierId sets ProductTierId field to given value.


### GetServiceId

`func (o *DiffTierVersionSetsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *DiffTierVersionSetsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *DiffTierVersionSetsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetToken

`func (o *DiffTierVersionSetsRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DiffTierVersionSetsRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DiffTierVersionSetsRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetVersion

`func (o *DiffTierVersionSetsRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DiffTierVersionSetsRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DiffTierVersionSetsRequest) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


