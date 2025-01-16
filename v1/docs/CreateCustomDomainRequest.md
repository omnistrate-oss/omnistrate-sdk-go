# CreateCustomDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | **string** | The root domain of the domain to use as suffix | 
**Description** | **string** | The description for the domain | 
**Name** | **string** | The name of the custom domain | 
**Route53Configuration** | [**Route53Configuration**](Route53Configuration.md) |  | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewCreateCustomDomainRequest

`func NewCreateCustomDomainRequest(customDomain string, description string, name string, route53Configuration Route53Configuration, token string, ) *CreateCustomDomainRequest`

NewCreateCustomDomainRequest instantiates a new CreateCustomDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomDomainRequestWithDefaults

`func NewCreateCustomDomainRequestWithDefaults() *CreateCustomDomainRequest`

NewCreateCustomDomainRequestWithDefaults instantiates a new CreateCustomDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *CreateCustomDomainRequest) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *CreateCustomDomainRequest) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *CreateCustomDomainRequest) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.


### GetDescription

`func (o *CreateCustomDomainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCustomDomainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCustomDomainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateCustomDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomDomainRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRoute53Configuration

`func (o *CreateCustomDomainRequest) GetRoute53Configuration() Route53Configuration`

GetRoute53Configuration returns the Route53Configuration field if non-nil, zero value otherwise.

### GetRoute53ConfigurationOk

`func (o *CreateCustomDomainRequest) GetRoute53ConfigurationOk() (*Route53Configuration, bool)`

GetRoute53ConfigurationOk returns a tuple with the Route53Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute53Configuration

`func (o *CreateCustomDomainRequest) SetRoute53Configuration(v Route53Configuration)`

SetRoute53Configuration sets Route53Configuration field to given value.


### GetToken

`func (o *CreateCustomDomainRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateCustomDomainRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateCustomDomainRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


