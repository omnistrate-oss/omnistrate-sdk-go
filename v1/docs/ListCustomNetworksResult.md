# ListCustomNetworksResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomNetworks** | [**[]CustomNetwork**](CustomNetwork.md) | List of custom networks | 

## Methods

### NewListCustomNetworksResult

`func NewListCustomNetworksResult(customNetworks []CustomNetwork, ) *ListCustomNetworksResult`

NewListCustomNetworksResult instantiates a new ListCustomNetworksResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomNetworksResultWithDefaults

`func NewListCustomNetworksResultWithDefaults() *ListCustomNetworksResult`

NewListCustomNetworksResultWithDefaults instantiates a new ListCustomNetworksResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomNetworks

`func (o *ListCustomNetworksResult) GetCustomNetworks() []CustomNetwork`

GetCustomNetworks returns the CustomNetworks field if non-nil, zero value otherwise.

### GetCustomNetworksOk

`func (o *ListCustomNetworksResult) GetCustomNetworksOk() (*[]CustomNetwork, bool)`

GetCustomNetworksOk returns a tuple with the CustomNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNetworks

`func (o *ListCustomNetworksResult) SetCustomNetworks(v []CustomNetwork)`

SetCustomNetworks sets CustomNetworks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


