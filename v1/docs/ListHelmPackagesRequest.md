# ListHelmPackagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChartName** | Pointer to **string** | The chart name of the Helm package to filter by | [optional] 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewListHelmPackagesRequest

`func NewListHelmPackagesRequest(token string, ) *ListHelmPackagesRequest`

NewListHelmPackagesRequest instantiates a new ListHelmPackagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHelmPackagesRequestWithDefaults

`func NewListHelmPackagesRequestWithDefaults() *ListHelmPackagesRequest`

NewListHelmPackagesRequestWithDefaults instantiates a new ListHelmPackagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChartName

`func (o *ListHelmPackagesRequest) GetChartName() string`

GetChartName returns the ChartName field if non-nil, zero value otherwise.

### GetChartNameOk

`func (o *ListHelmPackagesRequest) GetChartNameOk() (*string, bool)`

GetChartNameOk returns a tuple with the ChartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartName

`func (o *ListHelmPackagesRequest) SetChartName(v string)`

SetChartName sets ChartName field to given value.

### HasChartName

`func (o *ListHelmPackagesRequest) HasChartName() bool`

HasChartName returns a boolean if a field has been set.

### GetToken

`func (o *ListHelmPackagesRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListHelmPackagesRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListHelmPackagesRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


