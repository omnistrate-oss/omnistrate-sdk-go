# CreateCustomDomainRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | **string** | The root domain of the domain to use as suffix | 
**Description** | **string** | The description for the domain | 
**Name** | **string** | The name of the custom domain | 
**Route53Configuration** | [**Route53Configuration**](Route53Configuration.md) |  | 

## Methods

### NewCreateCustomDomainRequestBody

`func NewCreateCustomDomainRequestBody(customDomain string, description string, name string, route53Configuration Route53Configuration, ) *CreateCustomDomainRequestBody`

NewCreateCustomDomainRequestBody instantiates a new CreateCustomDomainRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomDomainRequestBodyWithDefaults

`func NewCreateCustomDomainRequestBodyWithDefaults() *CreateCustomDomainRequestBody`

NewCreateCustomDomainRequestBodyWithDefaults instantiates a new CreateCustomDomainRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *CreateCustomDomainRequestBody) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *CreateCustomDomainRequestBody) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *CreateCustomDomainRequestBody) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.


### GetDescription

`func (o *CreateCustomDomainRequestBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomDomainRequestBody) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomDomainRequestBody) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateCustomDomainRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomDomainRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomDomainRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetRoute53Configuration

`func (o *CreateCustomDomainRequestBody) GetRoute53Configuration() Route53Configuration`

GetRoute53Configuration returns the Route53Configuration field if non-nil, zero value otherwise.

### GetRoute53ConfigurationOk

`func (o *CreateCustomDomainRequestBody) GetRoute53ConfigurationOk() (*Route53Configuration, bool)`

GetRoute53ConfigurationOk returns a tuple with the Route53Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute53Configuration

`func (o *CreateCustomDomainRequestBody) SetRoute53Configuration(v Route53Configuration)`

SetRoute53Configuration sets Route53Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


