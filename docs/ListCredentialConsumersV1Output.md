# ListCredentialConsumersV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouses** | [**[]WarehouseV1**](WarehouseV1.md) | The Warehouses that use this Credential. | 
**WarehousesPagination** | [**PaginationOutput**](PaginationOutput.md) |  | 
**Sources** | [**[]SourceV1**](SourceV1.md) | The Sources that use this Credential. | 
**SourcesPagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListCredentialConsumersV1Output

`func NewListCredentialConsumersV1Output(warehouses []WarehouseV1, warehousesPagination PaginationOutput, sources []SourceV1, sourcesPagination PaginationOutput, ) *ListCredentialConsumersV1Output`

NewListCredentialConsumersV1Output instantiates a new ListCredentialConsumersV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialConsumersV1OutputWithDefaults

`func NewListCredentialConsumersV1OutputWithDefaults() *ListCredentialConsumersV1Output`

NewListCredentialConsumersV1OutputWithDefaults instantiates a new ListCredentialConsumersV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouses

`func (o *ListCredentialConsumersV1Output) GetWarehouses() []WarehouseV1`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *ListCredentialConsumersV1Output) GetWarehousesOk() (*[]WarehouseV1, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *ListCredentialConsumersV1Output) SetWarehouses(v []WarehouseV1)`

SetWarehouses sets Warehouses field to given value.


### GetWarehousesPagination

`func (o *ListCredentialConsumersV1Output) GetWarehousesPagination() PaginationOutput`

GetWarehousesPagination returns the WarehousesPagination field if non-nil, zero value otherwise.

### GetWarehousesPaginationOk

`func (o *ListCredentialConsumersV1Output) GetWarehousesPaginationOk() (*PaginationOutput, bool)`

GetWarehousesPaginationOk returns a tuple with the WarehousesPagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousesPagination

`func (o *ListCredentialConsumersV1Output) SetWarehousesPagination(v PaginationOutput)`

SetWarehousesPagination sets WarehousesPagination field to given value.


### GetSources

`func (o *ListCredentialConsumersV1Output) GetSources() []SourceV1`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *ListCredentialConsumersV1Output) GetSourcesOk() (*[]SourceV1, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *ListCredentialConsumersV1Output) SetSources(v []SourceV1)`

SetSources sets Sources field to given value.


### GetSourcesPagination

`func (o *ListCredentialConsumersV1Output) GetSourcesPagination() PaginationOutput`

GetSourcesPagination returns the SourcesPagination field if non-nil, zero value otherwise.

### GetSourcesPaginationOk

`func (o *ListCredentialConsumersV1Output) GetSourcesPaginationOk() (*PaginationOutput, bool)`

GetSourcesPaginationOk returns a tuple with the SourcesPagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesPagination

`func (o *ListCredentialConsumersV1Output) SetSourcesPagination(v PaginationOutput)`

SetSourcesPagination sets SourcesPagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


