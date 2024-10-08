# ListAssociatedResourcesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Infra Config ID to operate on | 
**ResourceIds** | **[]string** | The resources associated with the infra config | 
**ServiceId** | **string** | The service ID | 

## Methods

### NewListAssociatedResourcesResult

`func NewListAssociatedResourcesResult(id string, resourceIds []string, serviceId string, ) *ListAssociatedResourcesResult`

NewListAssociatedResourcesResult instantiates a new ListAssociatedResourcesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAssociatedResourcesResultWithDefaults

`func NewListAssociatedResourcesResultWithDefaults() *ListAssociatedResourcesResult`

NewListAssociatedResourcesResultWithDefaults instantiates a new ListAssociatedResourcesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAssociatedResourcesResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAssociatedResourcesResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAssociatedResourcesResult) SetId(v string)`

SetId sets Id field to given value.


### GetResourceIds

`func (o *ListAssociatedResourcesResult) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *ListAssociatedResourcesResult) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *ListAssociatedResourcesResult) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.


### GetServiceId

`func (o *ListAssociatedResourcesResult) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ListAssociatedResourcesResult) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ListAssociatedResourcesResult) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


