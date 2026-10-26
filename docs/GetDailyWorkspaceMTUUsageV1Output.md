# GetDailyWorkspaceMTUUsageV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyWorkspaceMTUUsage** | [**[]MtuSnapshotV1**](MtuSnapshotV1.md) | The list of daily Workspace MTU count aggregates. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDailyWorkspaceMTUUsageV1Output

`func NewGetDailyWorkspaceMTUUsageV1Output(dailyWorkspaceMTUUsage []MtuSnapshotV1, pagination PaginationOutput, ) *GetDailyWorkspaceMTUUsageV1Output`

NewGetDailyWorkspaceMTUUsageV1Output instantiates a new GetDailyWorkspaceMTUUsageV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWorkspaceMTUUsageV1OutputWithDefaults

`func NewGetDailyWorkspaceMTUUsageV1OutputWithDefaults() *GetDailyWorkspaceMTUUsageV1Output`

NewGetDailyWorkspaceMTUUsageV1OutputWithDefaults instantiates a new GetDailyWorkspaceMTUUsageV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyWorkspaceMTUUsage

`func (o *GetDailyWorkspaceMTUUsageV1Output) GetDailyWorkspaceMTUUsage() []MtuSnapshotV1`

GetDailyWorkspaceMTUUsage returns the DailyWorkspaceMTUUsage field if non-nil, zero value otherwise.

### GetDailyWorkspaceMTUUsageOk

`func (o *GetDailyWorkspaceMTUUsageV1Output) GetDailyWorkspaceMTUUsageOk() (*[]MtuSnapshotV1, bool)`

GetDailyWorkspaceMTUUsageOk returns a tuple with the DailyWorkspaceMTUUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyWorkspaceMTUUsage

`func (o *GetDailyWorkspaceMTUUsageV1Output) SetDailyWorkspaceMTUUsage(v []MtuSnapshotV1)`

SetDailyWorkspaceMTUUsage sets DailyWorkspaceMTUUsage field to given value.


### GetPagination

`func (o *GetDailyWorkspaceMTUUsageV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDailyWorkspaceMTUUsageV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDailyWorkspaceMTUUsageV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


