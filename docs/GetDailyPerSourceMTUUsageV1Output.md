# GetDailyPerSourceMTUUsageV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyPerSourceMTUUsage** | [**[]UsersPerSourceSnapshotV1**](UsersPerSourceSnapshotV1.md) | The list of daily per Source MTU count aggregates. | 
**Pagination** | [**PaginationOutput**](PaginationOutput.md) |  | 

## Methods

### NewGetDailyPerSourceMTUUsageV1Output

`func NewGetDailyPerSourceMTUUsageV1Output(dailyPerSourceMTUUsage []UsersPerSourceSnapshotV1, pagination PaginationOutput, ) *GetDailyPerSourceMTUUsageV1Output`

NewGetDailyPerSourceMTUUsageV1Output instantiates a new GetDailyPerSourceMTUUsageV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyPerSourceMTUUsageV1OutputWithDefaults

`func NewGetDailyPerSourceMTUUsageV1OutputWithDefaults() *GetDailyPerSourceMTUUsageV1Output`

NewGetDailyPerSourceMTUUsageV1OutputWithDefaults instantiates a new GetDailyPerSourceMTUUsageV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyPerSourceMTUUsage

`func (o *GetDailyPerSourceMTUUsageV1Output) GetDailyPerSourceMTUUsage() []UsersPerSourceSnapshotV1`

GetDailyPerSourceMTUUsage returns the DailyPerSourceMTUUsage field if non-nil, zero value otherwise.

### GetDailyPerSourceMTUUsageOk

`func (o *GetDailyPerSourceMTUUsageV1Output) GetDailyPerSourceMTUUsageOk() (*[]UsersPerSourceSnapshotV1, bool)`

GetDailyPerSourceMTUUsageOk returns a tuple with the DailyPerSourceMTUUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyPerSourceMTUUsage

`func (o *GetDailyPerSourceMTUUsageV1Output) SetDailyPerSourceMTUUsage(v []UsersPerSourceSnapshotV1)`

SetDailyPerSourceMTUUsage sets DailyPerSourceMTUUsage field to given value.


### GetPagination

`func (o *GetDailyPerSourceMTUUsageV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetDailyPerSourceMTUUsageV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetDailyPerSourceMTUUsageV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


