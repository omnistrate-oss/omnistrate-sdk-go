# DescribeAccountConfigByGCPProjectIDRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GcpProjectID** | **string** | The GCP project ID | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeAccountConfigByGCPProjectIDRequest

`func NewDescribeAccountConfigByGCPProjectIDRequest(gcpProjectID string, token string, ) *DescribeAccountConfigByGCPProjectIDRequest`

NewDescribeAccountConfigByGCPProjectIDRequest instantiates a new DescribeAccountConfigByGCPProjectIDRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAccountConfigByGCPProjectIDRequestWithDefaults

`func NewDescribeAccountConfigByGCPProjectIDRequestWithDefaults() *DescribeAccountConfigByGCPProjectIDRequest`

NewDescribeAccountConfigByGCPProjectIDRequestWithDefaults instantiates a new DescribeAccountConfigByGCPProjectIDRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGcpProjectID

`func (o *DescribeAccountConfigByGCPProjectIDRequest) GetGcpProjectID() string`

GetGcpProjectID returns the GcpProjectID field if non-nil, zero value otherwise.

### GetGcpProjectIDOk

`func (o *DescribeAccountConfigByGCPProjectIDRequest) GetGcpProjectIDOk() (*string, bool)`

GetGcpProjectIDOk returns a tuple with the GcpProjectID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpProjectID

`func (o *DescribeAccountConfigByGCPProjectIDRequest) SetGcpProjectID(v string)`

SetGcpProjectID sets GcpProjectID field to given value.


### GetToken

`func (o *DescribeAccountConfigByGCPProjectIDRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeAccountConfigByGCPProjectIDRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeAccountConfigByGCPProjectIDRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


