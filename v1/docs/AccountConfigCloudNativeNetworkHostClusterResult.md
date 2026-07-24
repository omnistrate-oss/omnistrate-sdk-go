# AccountConfigCloudNativeNetworkHostClusterResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EligibleToImport** | **bool** | Whether this host cluster can be imported into Omnistrate | 
**IneligibilityReason** | Pointer to **string** | The reason this host cluster is not eligible to import. Omitted when eligibleToImport is true. | [optional] 
**Name** | **string** | The cloud provider host cluster name | 

## Methods

### NewAccountConfigCloudNativeNetworkHostClusterResult

`func NewAccountConfigCloudNativeNetworkHostClusterResult(eligibleToImport bool, name string, ) *AccountConfigCloudNativeNetworkHostClusterResult`

NewAccountConfigCloudNativeNetworkHostClusterResult instantiates a new AccountConfigCloudNativeNetworkHostClusterResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigCloudNativeNetworkHostClusterResultWithDefaults

`func NewAccountConfigCloudNativeNetworkHostClusterResultWithDefaults() *AccountConfigCloudNativeNetworkHostClusterResult`

NewAccountConfigCloudNativeNetworkHostClusterResultWithDefaults instantiates a new AccountConfigCloudNativeNetworkHostClusterResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligibleToImport

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetEligibleToImport() bool`

GetEligibleToImport returns the EligibleToImport field if non-nil, zero value otherwise.

### GetEligibleToImportOk

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetEligibleToImportOk() (*bool, bool)`

GetEligibleToImportOk returns a tuple with the EligibleToImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleToImport

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) SetEligibleToImport(v bool)`

SetEligibleToImport sets EligibleToImport field to given value.


### GetIneligibilityReason

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetIneligibilityReason() string`

GetIneligibilityReason returns the IneligibilityReason field if non-nil, zero value otherwise.

### GetIneligibilityReasonOk

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetIneligibilityReasonOk() (*string, bool)`

GetIneligibilityReasonOk returns a tuple with the IneligibilityReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIneligibilityReason

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) SetIneligibilityReason(v string)`

SetIneligibilityReason sets IneligibilityReason field to given value.

### HasIneligibilityReason

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) HasIneligibilityReason() bool`

HasIneligibilityReason returns a boolean if a field has been set.

### GetName

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountConfigCloudNativeNetworkHostClusterResult) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


