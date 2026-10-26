# ReverseETLManualSyncJobOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartedAt** | **string** | A datetime that indicates when the job was triggered. | 
**SyncId** | **string** | The unique id for the sync that was triggered. | 

## Methods

### NewReverseETLManualSyncJobOutput

`func NewReverseETLManualSyncJobOutput(startedAt string, syncId string, ) *ReverseETLManualSyncJobOutput`

NewReverseETLManualSyncJobOutput instantiates a new ReverseETLManualSyncJobOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReverseETLManualSyncJobOutputWithDefaults

`func NewReverseETLManualSyncJobOutputWithDefaults() *ReverseETLManualSyncJobOutput`

NewReverseETLManualSyncJobOutputWithDefaults instantiates a new ReverseETLManualSyncJobOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartedAt

`func (o *ReverseETLManualSyncJobOutput) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ReverseETLManualSyncJobOutput) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ReverseETLManualSyncJobOutput) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetSyncId

`func (o *ReverseETLManualSyncJobOutput) GetSyncId() string`

GetSyncId returns the SyncId field if non-nil, zero value otherwise.

### GetSyncIdOk

`func (o *ReverseETLManualSyncJobOutput) GetSyncIdOk() (*string, bool)`

GetSyncIdOk returns a tuple with the SyncId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncId

`func (o *ReverseETLManualSyncJobOutput) SetSyncId(v string)`

SetSyncId sets SyncId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


