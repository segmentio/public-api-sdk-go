# GetEventsVolumeFromWorkspaceV1Output

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Observability event volume path. | 
**Query** | [**GetEventsVolumeFromWorkspaceV1Query**](GetEventsVolumeFromWorkspaceV1Query.md) |  | 
**Result** | [**[]SourceEventVolumeV1**](SourceEventVolumeV1.md) | The resultant list of series broken down by the dimensions requested over the time frame requested and ordered by the total count of events in all series. Note: The limit of entries returned is 5000. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewGetEventsVolumeFromWorkspaceV1Output

`func NewGetEventsVolumeFromWorkspaceV1Output(path string, query GetEventsVolumeFromWorkspaceV1Query, result []SourceEventVolumeV1, ) *GetEventsVolumeFromWorkspaceV1Output`

NewGetEventsVolumeFromWorkspaceV1Output instantiates a new GetEventsVolumeFromWorkspaceV1Output object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventsVolumeFromWorkspaceV1OutputWithDefaults

`func NewGetEventsVolumeFromWorkspaceV1OutputWithDefaults() *GetEventsVolumeFromWorkspaceV1Output`

NewGetEventsVolumeFromWorkspaceV1OutputWithDefaults instantiates a new GetEventsVolumeFromWorkspaceV1Output object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetEventsVolumeFromWorkspaceV1Output) SetPath(v string)`

SetPath sets Path field to given value.


### GetQuery

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetQuery() GetEventsVolumeFromWorkspaceV1Query`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetQueryOk() (*GetEventsVolumeFromWorkspaceV1Query, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GetEventsVolumeFromWorkspaceV1Output) SetQuery(v GetEventsVolumeFromWorkspaceV1Query)`

SetQuery sets Query field to given value.


### GetResult

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetResult() []SourceEventVolumeV1`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetResultOk() (*[]SourceEventVolumeV1, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetEventsVolumeFromWorkspaceV1Output) SetResult(v []SourceEventVolumeV1)`

SetResult sets Result field to given value.


### GetPagination

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetEventsVolumeFromWorkspaceV1Output) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetEventsVolumeFromWorkspaceV1Output) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetEventsVolumeFromWorkspaceV1Output) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


