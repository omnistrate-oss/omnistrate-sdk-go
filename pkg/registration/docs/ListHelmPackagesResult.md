# ListHelmPackagesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HelmPackages** | Pointer to [**[]HelmPackage**](HelmPackage.md) | List of Helm packages | [optional] 

## Methods

### NewListHelmPackagesResult

`func NewListHelmPackagesResult() *ListHelmPackagesResult`

NewListHelmPackagesResult instantiates a new ListHelmPackagesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHelmPackagesResultWithDefaults

`func NewListHelmPackagesResultWithDefaults() *ListHelmPackagesResult`

NewListHelmPackagesResultWithDefaults instantiates a new ListHelmPackagesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHelmPackages

`func (o *ListHelmPackagesResult) GetHelmPackages() []HelmPackage`

GetHelmPackages returns the HelmPackages field if non-nil, zero value otherwise.

### GetHelmPackagesOk

`func (o *ListHelmPackagesResult) GetHelmPackagesOk() (*[]HelmPackage, bool)`

GetHelmPackagesOk returns a tuple with the HelmPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmPackages

`func (o *ListHelmPackagesResult) SetHelmPackages(v []HelmPackage)`

SetHelmPackages sets HelmPackages field to given value.

### HasHelmPackages

`func (o *ListHelmPackagesResult) HasHelmPackages() bool`

HasHelmPackages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


