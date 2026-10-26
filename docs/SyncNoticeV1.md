# SyncNoticeV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | The severity of the notice. | 
**Message** | **string** | The human-readable message that describes the notice. | 
**CreatedAt** | **string** | The timestamp of this sync notice&#39;s creation. | 

## Methods

### NewSyncNoticeV1

`func NewSyncNoticeV1(level string, message string, createdAt string, ) *SyncNoticeV1`

NewSyncNoticeV1 instantiates a new SyncNoticeV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncNoticeV1WithDefaults

`func NewSyncNoticeV1WithDefaults() *SyncNoticeV1`

NewSyncNoticeV1WithDefaults instantiates a new SyncNoticeV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *SyncNoticeV1) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SyncNoticeV1) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SyncNoticeV1) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMessage

`func (o *SyncNoticeV1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyncNoticeV1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyncNoticeV1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCreatedAt

`func (o *SyncNoticeV1) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SyncNoticeV1) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SyncNoticeV1) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


