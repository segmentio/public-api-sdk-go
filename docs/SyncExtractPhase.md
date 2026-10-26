# SyncExtractPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedCount** | **string** | Counts the subset of records with status&#x3D;new, which indicates records that were created/inserted/added. | 
**UpdatedCount** | **string** | Counts the subset of records with status&#x3D;updated, which indicates records that were modified/updated. | 
**DeletedCount** | **string** | Counts the subset of records with status&#x3D;deleted, which indicates records that were deleted/removed. | 
**ExtractCount** | **string** | Counts the total number of records/rows handled by extract, across all statuses. | 
**ErrorCode** | **string** | Error code indicates a fatal sync error code, if applicable. | 
**StartedAt** | **string** | Time that the extract phase started. | 
**FinishedAt** | **string** | Time that the extract phase finished. | 

## Methods

### NewSyncExtractPhase

`func NewSyncExtractPhase(addedCount string, updatedCount string, deletedCount string, extractCount string, errorCode string, startedAt string, finishedAt string, ) *SyncExtractPhase`

NewSyncExtractPhase instantiates a new SyncExtractPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncExtractPhaseWithDefaults

`func NewSyncExtractPhaseWithDefaults() *SyncExtractPhase`

NewSyncExtractPhaseWithDefaults instantiates a new SyncExtractPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedCount

`func (o *SyncExtractPhase) GetAddedCount() string`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *SyncExtractPhase) GetAddedCountOk() (*string, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *SyncExtractPhase) SetAddedCount(v string)`

SetAddedCount sets AddedCount field to given value.


### GetUpdatedCount

`func (o *SyncExtractPhase) GetUpdatedCount() string`

GetUpdatedCount returns the UpdatedCount field if non-nil, zero value otherwise.

### GetUpdatedCountOk

`func (o *SyncExtractPhase) GetUpdatedCountOk() (*string, bool)`

GetUpdatedCountOk returns a tuple with the UpdatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedCount

`func (o *SyncExtractPhase) SetUpdatedCount(v string)`

SetUpdatedCount sets UpdatedCount field to given value.


### GetDeletedCount

`func (o *SyncExtractPhase) GetDeletedCount() string`

GetDeletedCount returns the DeletedCount field if non-nil, zero value otherwise.

### GetDeletedCountOk

`func (o *SyncExtractPhase) GetDeletedCountOk() (*string, bool)`

GetDeletedCountOk returns a tuple with the DeletedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedCount

`func (o *SyncExtractPhase) SetDeletedCount(v string)`

SetDeletedCount sets DeletedCount field to given value.


### GetExtractCount

`func (o *SyncExtractPhase) GetExtractCount() string`

GetExtractCount returns the ExtractCount field if non-nil, zero value otherwise.

### GetExtractCountOk

`func (o *SyncExtractPhase) GetExtractCountOk() (*string, bool)`

GetExtractCountOk returns a tuple with the ExtractCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractCount

`func (o *SyncExtractPhase) SetExtractCount(v string)`

SetExtractCount sets ExtractCount field to given value.


### GetErrorCode

`func (o *SyncExtractPhase) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyncExtractPhase) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SyncExtractPhase) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetStartedAt

`func (o *SyncExtractPhase) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SyncExtractPhase) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SyncExtractPhase) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *SyncExtractPhase) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SyncExtractPhase) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SyncExtractPhase) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


