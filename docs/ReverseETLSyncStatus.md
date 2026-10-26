# ReverseETLSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncId** | **string** | The id of the sync. | 
**ModelId** | **string** | The Model id of the sync. | 
**SourceId** | **string** | The Source id of the sync. | 
**SyncStatus** | **string** | The status of the sync. It currently can be IN_PROGRESS, FAIL, SUCCESS, PARTIAL_SUCCESS. | 
**Duration** | **string** | The duration of the sync. | 
**StartedAt** | **string** | When the sync started. | 
**FinishedAt** | Pointer to **string** | When the sync started. | [optional] 
**ExtractPhase** | Pointer to [**SyncExtractPhase**](SyncExtractPhase.md) |  | [optional] 
**LoadPhase** | Pointer to [**SyncLoadPhase**](SyncLoadPhase.md) |  | [optional] 
**Error** | Pointer to **string** | Error message if applicable. | [optional] 
**ErrorCode** | Pointer to **string** | Error code indicates a fatal sync error code, if applicable. | [optional] 

## Methods

### NewReverseETLSyncStatus

`func NewReverseETLSyncStatus(syncId string, modelId string, sourceId string, syncStatus string, duration string, startedAt string, ) *ReverseETLSyncStatus`

NewReverseETLSyncStatus instantiates a new ReverseETLSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseETLSyncStatusWithDefaults

`func NewReverseETLSyncStatusWithDefaults() *ReverseETLSyncStatus`

NewReverseETLSyncStatusWithDefaults instantiates a new ReverseETLSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncId

`func (o *ReverseETLSyncStatus) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *ReverseETLSyncStatus) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *ReverseETLSyncStatus) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.


### GetModelId

`func (o *ReverseETLSyncStatus) GetModelId() string`

GetModelId returns the ModelId field if non-nil, zero value otherwise.

### GetModelIdOk

`func (o *ReverseETLSyncStatus) GetModelIdOk() (*string, bool)`

GetModelIdOk returns a tuple with the ModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelId

`func (o *ReverseETLSyncStatus) SetModelId(v string)`

SetModelId sets ModelId field to given value.


### GetSourceId

`func (o *ReverseETLSyncStatus) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ReverseETLSyncStatus) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ReverseETLSyncStatus) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSyncStatus

`func (o *ReverseETLSyncStatus) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *ReverseETLSyncStatus) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *ReverseETLSyncStatus) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.


### GetDuration

`func (o *ReverseETLSyncStatus) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ReverseETLSyncStatus) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ReverseETLSyncStatus) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetStartedAt

`func (o *ReverseETLSyncStatus) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ReverseETLSyncStatus) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ReverseETLSyncStatus) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *ReverseETLSyncStatus) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ReverseETLSyncStatus) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ReverseETLSyncStatus) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ReverseETLSyncStatus) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetExtractPhase

`func (o *ReverseETLSyncStatus) GetExtractPhase() SyncExtractPhase`

GetExtractPhase returns the ExtractPhase field if non-nil, zero value otherwise.

### GetExtractPhaseOk

`func (o *ReverseETLSyncStatus) GetExtractPhaseOk() (*SyncExtractPhase, bool)`

GetExtractPhaseOk returns a tuple with the ExtractPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractPhase

`func (o *ReverseETLSyncStatus) SetExtractPhase(v SyncExtractPhase)`

SetExtractPhase sets ExtractPhase field to given value.

### HasExtractPhase

`func (o *ReverseETLSyncStatus) HasExtractPhase() bool`

HasExtractPhase returns a boolean if a field has been set.

### GetLoadPhase

`func (o *ReverseETLSyncStatus) GetLoadPhase() SyncLoadPhase`

GetLoadPhase returns the LoadPhase field if non-nil, zero value otherwise.

### GetLoadPhaseOk

`func (o *ReverseETLSyncStatus) GetLoadPhaseOk() (*SyncLoadPhase, bool)`

GetLoadPhaseOk returns a tuple with the LoadPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadPhase

`func (o *ReverseETLSyncStatus) SetLoadPhase(v SyncLoadPhase)`

SetLoadPhase sets LoadPhase field to given value.

### HasLoadPhase

`func (o *ReverseETLSyncStatus) HasLoadPhase() bool`

HasLoadPhase returns a boolean if a field has been set.

### GetError

`func (o *ReverseETLSyncStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ReverseETLSyncStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ReverseETLSyncStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ReverseETLSyncStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *ReverseETLSyncStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ReverseETLSyncStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ReverseETLSyncStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ReverseETLSyncStatus) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


