# DescribeHelmPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package to describe | 
**ChartVersion** | **string** | The chart version of the Helm package to describe | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDescribeHelmPackageRequest

`func NewDescribeHelmPackageRequest(chartName string, chartVersion string, token string, ) *DescribeHelmPackageRequest`

NewDescribeHelmPackageRequest instantiates a new DescribeHelmPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeHelmPackageRequestWithDefaults

`func NewDescribeHelmPackageRequestWithDefaults() *DescribeHelmPackageRequest`

NewDescribeHelmPackageRequestWithDefaults instantiates a new DescribeHelmPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *DescribeHelmPackageRequest) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *DescribeHelmPackageRequest) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *DescribeHelmPackageRequest) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *DescribeHelmPackageRequest) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *DescribeHelmPackageRequest) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *DescribeHelmPackageRequest) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetToken

`func (o *DescribeHelmPackageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DescribeHelmPackageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DescribeHelmPackageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


