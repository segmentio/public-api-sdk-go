# GetDailyWorkspaceAPICallsUsageV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyWorkspaceAPICallsUsage** | [**[]APICallSnapshotV1**](APICallSnapshotV1.md) | The list of daily Workspace API calls count aggregates. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDailyWorkspaceAPICallsUsageV1Output

`func NewGetDailyWorkspaceAPICallsUsageV1Output(dailyWorkspaceAPICallsUsage []APICallSnapshotV1, pagination PaginationOutput, ) *GetDailyWorkspaceAPICallsUsageV1Output`

NewGetDailyWorkspaceAPICallsUsageV1Output instantiates a new GetDailyWorkspaceAPICallsUsageV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyWorkspaceAPICallsUsageV1OutputWithDefaults

`func NewGetDailyWorkspaceAPICallsUsageV1OutputWithDefaults() *GetDailyWorkspaceAPICallsUsageV1Output`

NewGetDailyWorkspaceAPICallsUsageV1OutputWithDefaults instantiates a new GetDailyWorkspaceAPICallsUsageV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyWorkspaceAPICallsUsage

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetDailyWorkspaceAPICallsUsage() []APICallSnapshotV1`

GetDailyWorkspaceAPICallsUsage returns the DailyWorkspaceAPICallsUsage field if non-nil, zero value otherwise.

### GetDailyWorkspaceAPICallsUsageOk

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetDailyWorkspaceAPICallsUsageOk() (*[]APICallSnapshotV1, bool)`

GetDailyWorkspaceAPICallsUsageOk returns a tuple with the DailyWorkspaceAPICallsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyWorkspaceAPICallsUsage

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) SetDailyWorkspaceAPICallsUsage(v []APICallSnapshotV1)`

SetDailyWorkspaceAPICallsUsage sets DailyWorkspaceAPICallsUsage field to given value.


### GetPagination

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDailyWorkspaceAPICallsUsageV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


