# DeleteHelmPackageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | **string** | The chart name of the Helm package to delete | 
**ChartVersion** | **string** | The chart version of the Helm package to delete | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewDeleteHelmPackageRequest

`func NewDeleteHelmPackageRequest(chartName string, chartVersion string, token string, ) *DeleteHelmPackageRequest`

NewDeleteHelmPackageRequest instantiates a new DeleteHelmPackageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteHelmPackageRequestWithDefaults

`func NewDeleteHelmPackageRequestWithDefaults() *DeleteHelmPackageRequest`

NewDeleteHelmPackageRequestWithDefaults instantiates a new DeleteHelmPackageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *DeleteHelmPackageRequest) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *DeleteHelmPackageRequest) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *DeleteHelmPackageRequest) SetChartName(v string)`

SetChartName sets ChartName field to given value.


### GetChartVersion

`func (o *DeleteHelmPackageRequest) GetChartVersion() string`

GetChartVersion returns the ChartVersion field if non-nil, zero value otherwise.

### GetChartVersionOk

`func (o *DeleteHelmPackageRequest) GetChartVersionOk() (*string, bool)`

GetChartVersionOk returns a tuple with the ChartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartVersion

`func (o *DeleteHelmPackageRequest) SetChartVersion(v string)`

SetChartVersion sets ChartVersion field to given value.


### GetToken

`func (o *DeleteHelmPackageRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeleteHelmPackageRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeleteHelmPackageRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


