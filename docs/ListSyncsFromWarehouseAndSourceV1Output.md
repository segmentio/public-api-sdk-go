# ListSyncsFromWarehouseAndSourceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reports** | [**[]SyncV1**](SyncV1.md) | A list that contains the latest syncs for the specified Warehouse-source pair. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewListSyncsFromWarehouseAndSourceV1Output

`func NewListSyncsFromWarehouseAndSourceV1Output(reports []SyncV1, pagination PaginationOutput, ) *ListSyncsFromWarehouseAndSourceV1Output`

NewListSyncsFromWarehouseAndSourceV1Output instantiates a new ListSyncsFromWarehouseAndSourceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSyncsFromWarehouseAndSourceV1OutputWithDefaults

`func NewListSyncsFromWarehouseAndSourceV1OutputWithDefaults() *ListSyncsFromWarehouseAndSourceV1Output`

NewListSyncsFromWarehouseAndSourceV1OutputWithDefaults instantiates a new ListSyncsFromWarehouseAndSourceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReports

`func (o *ListSyncsFromWarehouseAndSourceV1Output) GetReports() []SyncV1`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *ListSyncsFromWarehouseAndSourceV1Output) GetReportsOk() (*[]SyncV1, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *ListSyncsFromWarehouseAndSourceV1Output) SetReports(v []SyncV1)`

SetReports sets Reports field to given value.


### GetPagination

`func (o *ListSyncsFromWarehouseAndSourceV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSyncsFromWarehouseAndSourceV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSyncsFromWarehouseAndSourceV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


