# SimulateContractEventRequest2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyerRef** | Pointer to **string** | The simulated buyer. Defaults to a value derived from the run, and supplying the same one twice is how an ISV rehearses a returning buyer accumulating a second contract against one organization | [optional] 
**Control** | **string** |  | 
**MarketplaceContractId** | Pointer to **string** | Which simulated contract to act on. Omitted for purchase, which is the control that creates one. Required for every other control, because there is otherwise no way to say which of several rehearsals is meant | [optional] 
**PlanRef** | Pointer to **string** | The plan to purchase or change to. Must be a key in the sandbox channel config&#39;s plan map, otherwise there is no product tier to fulfil against | [optional] 
**Quantity** | Pointer to **int64** | Seat or unit count, for purchase and quantity_change | [optional] 

## Methods

### NewSimulateContractEventRequest2

`func NewSimulateContractEventRequest2(control string, ) *SimulateContractEventRequest2`

NewSimulateContractEventRequest2 instantiates a new SimulateContractEventRequest2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimulateContractEventRequest2WithDefaults

`func NewSimulateContractEventRequest2WithDefaults() *SimulateContractEventRequest2`

NewSimulateContractEventRequest2WithDefaults instantiates a new SimulateContractEventRequest2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyerRef

`func (o *SimulateContractEventRequest2) GetBuyerRef() string`

GetBuyerRef returns the BuyerRef field if non-nil, zero value otherwise.

### GetBuyerRefOk

`func (o *SimulateContractEventRequest2) GetBuyerRefOk() (*string, bool)`

GetBuyerRefOk returns a tuple with the BuyerRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerRef

`func (o *SimulateContractEventRequest2) SetBuyerRef(v string)`

SetBuyerRef sets BuyerRef field to given value.

### HasBuyerRef

`func (o *SimulateContractEventRequest2) HasBuyerRef() bool`

HasBuyerRef returns a boolean if a field has been set.

### GetControl

`func (o *SimulateContractEventRequest2) GetControl() string`

GetControl returns the Control field if non-nil, zero value otherwise.

### GetControlOk

`func (o *SimulateContractEventRequest2) GetControlOk() (*string, bool)`

GetControlOk returns a tuple with the Control field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControl

`func (o *SimulateContractEventRequest2) SetControl(v string)`

SetControl sets Control field to given value.


### GetMarketplaceContractId

`func (o *SimulateContractEventRequest2) GetMarketplaceContractId() string`

GetMarketplaceContractId returns the MarketplaceContractId field if non-nil, zero value otherwise.

### GetMarketplaceContractIdOk

`func (o *SimulateContractEventRequest2) GetMarketplaceContractIdOk() (*string, bool)`

GetMarketplaceContractIdOk returns a tuple with the MarketplaceContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplaceContractId

`func (o *SimulateContractEventRequest2) SetMarketplaceContractId(v string)`

SetMarketplaceContractId sets MarketplaceContractId field to given value.

### HasMarketplaceContractId

`func (o *SimulateContractEventRequest2) HasMarketplaceContractId() bool`

HasMarketplaceContractId returns a boolean if a field has been set.

### GetPlanRef

`func (o *SimulateContractEventRequest2) GetPlanRef() string`

GetPlanRef returns the PlanRef field if non-nil, zero value otherwise.

### GetPlanRefOk

`func (o *SimulateContractEventRequest2) GetPlanRefOk() (*string, bool)`

GetPlanRefOk returns a tuple with the PlanRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanRef

`func (o *SimulateContractEventRequest2) SetPlanRef(v string)`

SetPlanRef sets PlanRef field to given value.

### HasPlanRef

`func (o *SimulateContractEventRequest2) HasPlanRef() bool`

HasPlanRef returns a boolean if a field has been set.

### GetQuantity

`func (o *SimulateContractEventRequest2) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SimulateContractEventRequest2) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SimulateContractEventRequest2) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SimulateContractEventRequest2) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


