# GetDailyPerSourceAPICallsUsageV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyPerSourceAPICallsUsage** | [**[]SourceAPICallSnapshotV1**](SourceAPICallSnapshotV1.md) | The list of daily per Source API calls count aggregates. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDailyPerSourceAPICallsUsageV1Output

`func NewGetDailyPerSourceAPICallsUsageV1Output(dailyPerSourceAPICallsUsage []SourceAPICallSnapshotV1, pagination PaginationOutput, ) *GetDailyPerSourceAPICallsUsageV1Output`

NewGetDailyPerSourceAPICallsUsageV1Output instantiates a new GetDailyPerSourceAPICallsUsageV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyPerSourceAPICallsUsageV1OutputWithDefaults

`func NewGetDailyPerSourceAPICallsUsageV1OutputWithDefaults() *GetDailyPerSourceAPICallsUsageV1Output`

NewGetDailyPerSourceAPICallsUsageV1OutputWithDefaults instantiates a new GetDailyPerSourceAPICallsUsageV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyPerSourceAPICallsUsage

`func (o *GetDailyPerSourceAPICallsUsageV1Output) GetDailyPerSourceAPICallsUsage() []SourceAPICallSnapshotV1`

GetDailyPerSourceAPICallsUsage returns the DailyPerSourceAPICallsUsage field if non-nil, zero value otherwise.

### GetDailyPerSourceAPICallsUsageOk

`func (o *GetDailyPerSourceAPICallsUsageV1Output) GetDailyPerSourceAPICallsUsageOk() (*[]SourceAPICallSnapshotV1, bool)`

GetDailyPerSourceAPICallsUsageOk returns a tuple with the DailyPerSourceAPICallsUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyPerSourceAPICallsUsage

`func (o *GetDailyPerSourceAPICallsUsageV1Output) SetDailyPerSourceAPICallsUsage(v []SourceAPICallSnapshotV1)`

SetDailyPerSourceAPICallsUsage sets DailyPerSourceAPICallsUsage field to given value.


### GetPagination

`func (o *GetDailyPerSourceAPICallsUsageV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDailyPerSourceAPICallsUsageV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDailyPerSourceAPICallsUsageV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


