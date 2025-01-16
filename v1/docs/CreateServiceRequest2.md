# CreateServiceRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A brief description of the service | 
**Name** | **string** | Name of the Service | 
**ServiceLogoURL** | Pointer to **string** | The logo for the service | [optional] 

## Methods

### NewCreateServiceRequest2

`func NewCreateServiceRequest2(description string, name string, ) *CreateServiceRequest2`

NewCreateServiceRequest2 instantiates a new CreateServiceRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceRequest2WithDefaults

`func NewCreateServiceRequest2WithDefaults() *CreateServiceRequest2`

NewCreateServiceRequest2WithDefaults instantiates a new CreateServiceRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceRequest2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceRequest2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceRequest2) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateServiceRequest2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceRequest2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceRequest2) SetName(v string)`

SetName sets Name field to given value.


### GetServiceLogoURL

`func (o *CreateServiceRequest2) GetServiceLogoURL() string`

GetServiceLogoURL returns the ServiceLogoURL field if non-nil, zero value otherwise.

### GetServiceLogoURLOk

`func (o *CreateServiceRequest2) GetServiceLogoURLOk() (*string, bool)`

GetServiceLogoURLOk returns a tuple with the ServiceLogoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceLogoURL

`func (o *CreateServiceRequest2) SetServiceLogoURL(v string)`

SetServiceLogoURL sets ServiceLogoURL field to given value.

### HasServiceLogoURL

`func (o *CreateServiceRequest2) HasServiceLogoURL() bool`

HasServiceLogoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


