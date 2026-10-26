# ListSyncsFromWarehouseV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | [**[]SyncV1**](SyncV1.md) | A list that contains the latest syncs for the specified Warehouse. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSyncsFromWarehouseV1Output

`func NewListSyncsFromWarehouseV1Output(reports []SyncV1, pagination PaginationOutput, ) *ListSyncsFromWarehouseV1Output`

NewListSyncsFromWarehouseV1Output instantiates a new ListSyncsFromWarehouseV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncsFromWarehouseV1OutputWithDefaults

`func NewListSyncsFromWarehouseV1OutputWithDefaults() *ListSyncsFromWarehouseV1Output`

NewListSyncsFromWarehouseV1OutputWithDefaults instantiates a new ListSyncsFromWarehouseV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *ListSyncsFromWarehouseV1Output) GetReports() []SyncV1`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *ListSyncsFromWarehouseV1Output) GetReportsOk() (*[]SyncV1, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *ListSyncsFromWarehouseV1Output) SetReports(v []SyncV1)`

SetReports sets Reports field to given value.


### GetPagination

`func (o *ListSyncsFromWarehouseV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSyncsFromWarehouseV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSyncsFromWarehouseV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


