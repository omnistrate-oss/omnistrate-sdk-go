# AccountConfigSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of an Account Config | 
**Name** | **string** | The name of the account configuration. | 
**TargetAccountID** | **string** | The cloud-provider-specific account identifier, such as an AWS account ID, GCP project ID, Azure subscription ID, OCI tenancy ID, or Nebius tenant ID. | 

## Methods

### NewAccountConfigSummary

`func NewAccountConfigSummary(id string, name string, targetAccountID string, ) *AccountConfigSummary`

NewAccountConfigSummary instantiates a new AccountConfigSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountConfigSummaryWithDefaults

`func NewAccountConfigSummaryWithDefaults() *AccountConfigSummary`

NewAccountConfigSummaryWithDefaults instantiates a new AccountConfigSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountConfigSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountConfigSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountConfigSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AccountConfigSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountConfigSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountConfigSummary) SetName(v string)`

SetName sets Name field to given value.


### GetTargetAccountID

`func (o *AccountConfigSummary) GetTargetAccountID() string`

GetTargetAccountID returns the TargetAccountID field if non-nil, zero value otherwise.

### GetTargetAccountIDOk

`func (o *AccountConfigSummary) GetTargetAccountIDOk() (*string, bool)`

GetTargetAccountIDOk returns a tuple with the TargetAccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountID

`func (o *AccountConfigSummary) SetTargetAccountID(v string)`

SetTargetAccountID sets TargetAccountID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


