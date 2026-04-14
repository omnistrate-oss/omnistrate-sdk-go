# SearchInventoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The search query. Supports prefixed searches such as &#39;service:&lt;value&gt;&#39;, &#39;user:&lt;value&gt;&#39;, &#39;subscription:&lt;value&gt;&#39;, &#39;deploymentcell:&lt;value&gt;&#39;, &#39;serviceplan:&lt;value&gt;&#39;, &#39;resource:&lt;value&gt;&#39;, &#39;serverlessproxy:&lt;value&gt;&#39;, &#39;resourceinstance:&lt;value&gt;&#39;, &#39;snapshot:&lt;value&gt;&#39;, &#39;notification:&lt;value&gt;&#39;, &#39;workflow:&lt;value&gt;&#39;, &#39;upgradepath:&lt;value&gt;&#39;, &#39;all:&lt;value&gt;&#39;. Without a prefix, performs a text search across resource instances, notifications, and workflows. | 
**Token** | **string** | JWT token used to perform authorization | 

## Methods

### NewSearchInventoryRequest

`func NewSearchInventoryRequest(query string, token string, ) *SearchInventoryRequest`

NewSearchInventoryRequest instantiates a new SearchInventoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchInventoryRequestWithDefaults

`func NewSearchInventoryRequestWithDefaults() *SearchInventoryRequest`

NewSearchInventoryRequestWithDefaults instantiates a new SearchInventoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *SearchInventoryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchInventoryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchInventoryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetToken

`func (o *SearchInventoryRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SearchInventoryRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SearchInventoryRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


