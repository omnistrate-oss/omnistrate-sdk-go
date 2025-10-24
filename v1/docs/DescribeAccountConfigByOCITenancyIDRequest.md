# DescribeAccountConfigByOCITenancyIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OciTenancyID** | **string** | The Tenancy OCID for Oracle Cloud Infrastructure | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccountConfigByOCITenancyIDRequest

`func NewDescribeAccountConfigByOCITenancyIDRequest(ociTenancyID string, token string, ) *DescribeAccountConfigByOCITenancyIDRequest`

NewDescribeAccountConfigByOCITenancyIDRequest instantiates a new DescribeAccountConfigByOCITenancyIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByOCITenancyIDRequestWithDefaults

`func NewDescribeAccountConfigByOCITenancyIDRequestWithDefaults() *DescribeAccountConfigByOCITenancyIDRequest`

NewDescribeAccountConfigByOCITenancyIDRequestWithDefaults instantiates a new DescribeAccountConfigByOCITenancyIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOciTenancyID

`func (o *DescribeAccountConfigByOCITenancyIDRequest) GetOciTenancyID() string`

GetOciTenancyID returns the OciTenancyID field if non-nil, zero value otherwise.

### GetOciTenancyIDOk

`func (o *DescribeAccountConfigByOCITenancyIDRequest) GetOciTenancyIDOk() (*string, bool)`

GetOciTenancyIDOk returns a tuple with the OciTenancyID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciTenancyID

`func (o *DescribeAccountConfigByOCITenancyIDRequest) SetOciTenancyID(v string)`

SetOciTenancyID sets OciTenancyID field to given value.


### GetToken

`func (o *DescribeAccountConfigByOCITenancyIDRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccountConfigByOCITenancyIDRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccountConfigByOCITenancyIDRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


