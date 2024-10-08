# EmailConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **string** | Email address to send notifications to | 

## Methods

### NewEmailConfiguration

`func NewEmailConfiguration(to string, ) *EmailConfiguration`

NewEmailConfiguration instantiates a new EmailConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConfigurationWithDefaults

`func NewEmailConfigurationWithDefaults() *EmailConfiguration`

NewEmailConfigurationWithDefaults instantiates a new EmailConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *EmailConfiguration) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailConfiguration) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailConfiguration) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


