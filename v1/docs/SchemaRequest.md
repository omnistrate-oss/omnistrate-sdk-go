# SchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the schema, e.g., &#39;compose&#39;, &#39;service-plan&#39;, &#39;x-omnistrate-mode-internal&#39; | 

## Methods

### NewSchemaRequest

`func NewSchemaRequest(type_ string, ) *SchemaRequest`

NewSchemaRequest instantiates a new SchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaRequestWithDefaults

`func NewSchemaRequestWithDefaults() *SchemaRequest`

NewSchemaRequestWithDefaults instantiates a new SchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


