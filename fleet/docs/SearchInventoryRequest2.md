# SearchInventoryRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**SearchInventoryFilters**](SearchInventoryFilters.md) |  | [optional] 
**Query** | **string** | The search query. Supports prefixed searches such as &#39;service:&lt;value&gt;&#39;, &#39;user:&lt;value&gt;&#39;, &#39;subscription:&lt;value&gt;&#39;, &#39;deploymentcell:&lt;value&gt;&#39;, &#39;serviceplan:&lt;value&gt;&#39;, &#39;resource:&lt;value&gt;&#39;, &#39;serverlessproxy:&lt;value&gt;&#39;, &#39;resourceinstance:&lt;value&gt;&#39;, &#39;snapshot:&lt;value&gt;&#39;, &#39;notification:&lt;value&gt;&#39;, &#39;workflow:&lt;value&gt;&#39;, &#39;upgradepath:&lt;value&gt;&#39;, &#39;all:&lt;value&gt;&#39;. Without a prefix, performs a text search across resource instances, notifications, and workflows. | 

## Methods

### NewSearchInventoryRequest2

`func NewSearchInventoryRequest2(query string, ) *SearchInventoryRequest2`

NewSearchInventoryRequest2 instantiates a new SearchInventoryRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchInventoryRequest2WithDefaults

`func NewSearchInventoryRequest2WithDefaults() *SearchInventoryRequest2`

NewSearchInventoryRequest2WithDefaults instantiates a new SearchInventoryRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *SearchInventoryRequest2) GetFilters() SearchInventoryFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *SearchInventoryRequest2) GetFiltersOk() (*SearchInventoryFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *SearchInventoryRequest2) SetFilters(v SearchInventoryFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *SearchInventoryRequest2) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetQuery

`func (o *SearchInventoryRequest2) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchInventoryRequest2) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchInventoryRequest2) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


