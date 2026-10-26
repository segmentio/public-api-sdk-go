# ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncStatuses** | [**[]ReverseETLSyncStatus**](ReverseETLSyncStatus.md) | The reverse ETL sync statuses that were looked up of the subscription id. | 
**Pagination** | Pointer to [**PaginationOutput**](PaginationOutput.md) |  | [optional] 

## Methods

### NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput

`func NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput(syncStatuses []ReverseETLSyncStatus, ) *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput`

NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutputWithDefaults

`func NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutputWithDefaults() *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput`

NewListReverseETLSyncStatusesFromModelAndSubscriptionIdOutputWithDefaults instantiates a new ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncStatuses

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetSyncStatuses() []ReverseETLSyncStatus`

GetSyncStatuses returns the SyncStatuses field if non-nil, zero value otherwise.

### GetSyncStatusesOk

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetSyncStatusesOk() (*[]ReverseETLSyncStatus, bool)`

GetSyncStatusesOk returns a tuple with the SyncStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatuses

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) SetSyncStatuses(v []ReverseETLSyncStatus)`

SetSyncStatuses sets SyncStatuses field to given value.


### GetPagination

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetPagination() PaginationOutput`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) GetPaginationOk() (*PaginationOutput, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) SetPagination(v PaginationOutput)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListReverseETLSyncStatusesFromModelAndSubscriptionIdOutput) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


