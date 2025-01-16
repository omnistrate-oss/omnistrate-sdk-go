# DescribeCustomDomainResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomDomain** | **string** | The root domain of the domain to use as suffix | 
**Description** | **string** | The description for the domain | 
**Id** | **string** | ID of an Custom Domain | 
**Name** | **string** | The name of the custom domain | 
**Route53Configuration** | [**Route53ConfigurationDescription**](Route53ConfigurationDescription.md) |  | 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** | The status message of the custom domain | [optional] 

## Methods

### NewDescribeCustomDomainResult

`func NewDescribeCustomDomainResult(customDomain string, description string, id string, name string, route53Configuration Route53ConfigurationDescription, ) *DescribeCustomDomainResult`

NewDescribeCustomDomainResult instantiates a new DescribeCustomDomainResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCustomDomainResultWithDefaults

`func NewDescribeCustomDomainResultWithDefaults() *DescribeCustomDomainResult`

NewDescribeCustomDomainResultWithDefaults instantiates a new DescribeCustomDomainResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomDomain

`func (o *DescribeCustomDomainResult) GetCustomDomain() string`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *DescribeCustomDomainResult) GetCustomDomainOk() (*string, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *DescribeCustomDomainResult) SetCustomDomain(v string)`

SetCustomDomain sets CustomDomain field to given value.


### GetDescription

`func (o *DescribeCustomDomainResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DescribeCustomDomainResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DescribeCustomDomainResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *DescribeCustomDomainResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeCustomDomainResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeCustomDomainResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DescribeCustomDomainResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeCustomDomainResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeCustomDomainResult) SetName(v string)`

SetName sets Name field to given value.


### GetRoute53Configuration

`func (o *DescribeCustomDomainResult) GetRoute53Configuration() Route53ConfigurationDescription`

GetRoute53Configuration returns the Route53Configuration field if non-nil, zero value otherwise.

### GetRoute53ConfigurationOk

`func (o *DescribeCustomDomainResult) GetRoute53ConfigurationOk() (*Route53ConfigurationDescription, bool)`

GetRoute53ConfigurationOk returns a tuple with the Route53Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute53Configuration

`func (o *DescribeCustomDomainResult) SetRoute53Configuration(v Route53ConfigurationDescription)`

SetRoute53Configuration sets Route53Configuration field to given value.


### GetStatus

`func (o *DescribeCustomDomainResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeCustomDomainResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeCustomDomainResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeCustomDomainResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DescribeCustomDomainResult) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DescribeCustomDomainResult) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DescribeCustomDomainResult) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DescribeCustomDomainResult) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


