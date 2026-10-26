# SyncV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The id of the Source loaded in the sync. | 
**Start** | **string** | The start time of the sync. | 
**End** | **NullableString** | The time the sync completed. Returns null if unfinished. | 
**Status** | **string** | The status of the sync. | 
**Duration** | **float32** | The duration of the sync in seconds. Returns the partial duration if the sync has not finished yet. | 
**HumanDuration** | **string** | The human-readable counterpart of &#x60;duration&#x60;. | 
**Count** | **float32** | The number of rows synced into the Warehouse. | 
**Notices** | [**[]SyncNoticeV1**](SyncNoticeV1.md) | Notices that contain the events that occurred during the sync. | 

## Methods

### NewSyncV1

`func NewSyncV1(sourceId string, start string, end NullableString, status string, duration float32, humanDuration string, count float32, notices []SyncNoticeV1, ) *SyncV1`

NewSyncV1 instantiates a new SyncV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncV1WithDefaults

`func NewSyncV1WithDefaults() *SyncV1`

NewSyncV1WithDefaults instantiates a new SyncV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SyncV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SyncV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SyncV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetStart

`func (o *SyncV1) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SyncV1) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SyncV1) SetStart(v string)`

SetStart sets Start field to given value.


### GetEnd

`func (o *SyncV1) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SyncV1) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SyncV1) SetEnd(v string)`

SetEnd sets End field to given value.


### SetEndNil

`func (o *SyncV1) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *SyncV1) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStatus

`func (o *SyncV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyncV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyncV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDuration

`func (o *SyncV1) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SyncV1) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SyncV1) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetHumanDuration

`func (o *SyncV1) GetHumanDuration() string`

GetHumanDuration returns the HumanDuration field if non-nil, zero value otherwise.

### GetHumanDurationOk

`func (o *SyncV1) GetHumanDurationOk() (*string, bool)`

GetHumanDurationOk returns a tuple with the HumanDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanDuration

`func (o *SyncV1) SetHumanDuration(v string)`

SetHumanDuration sets HumanDuration field to given value.


### GetCount

`func (o *SyncV1) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SyncV1) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SyncV1) SetCount(v float32)`

SetCount sets Count field to given value.


### GetNotices

`func (o *SyncV1) GetNotices() []SyncNoticeV1`

GetNotices returns the Notices field if non-nil, zero value otherwise.

### GetNoticesOk

`func (o *SyncV1) GetNoticesOk() (*[]SyncNoticeV1, bool)`

GetNoticesOk returns a tuple with the Notices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotices

`func (o *SyncV1) SetNotices(v []SyncNoticeV1)`

SetNotices sets Notices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


