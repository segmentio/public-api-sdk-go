# SyncLoadPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliverSuccessCount** | **string** | Counts the subset of records successfully delivered to the downstream Destination. | 
**DeliverFailureCount** | **string** | Counts the subset of records status that failed to be delivered by either receiving a permanent error (for example, 400 Bad Request) or by exhausting all retries for temporary errors (for example, 429 Too Many Requests). | 
**ErrorCode** | **string** | Error code indicates a fatal sync error code, if applicable. | 
**StartedAt** | **string** | Time that the load phase started. | 
**FinishedAt** | **string** | Time that the load phase finished. | 

## Methods

### NewSyncLoadPhase

`func NewSyncLoadPhase(deliverSuccessCount string, deliverFailureCount string, errorCode string, startedAt string, finishedAt string, ) *SyncLoadPhase`

NewSyncLoadPhase instantiates a new SyncLoadPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncLoadPhaseWithDefaults

`func NewSyncLoadPhaseWithDefaults() *SyncLoadPhase`

NewSyncLoadPhaseWithDefaults instantiates a new SyncLoadPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliverSuccessCount

`func (o *SyncLoadPhase) GetDeliverSuccessCount() string`

GetDeliverSuccessCount returns the DeliverSuccessCount field if non-nil, zero value otherwise.

### GetDeliverSuccessCountOk

`func (o *SyncLoadPhase) GetDeliverSuccessCountOk() (*string, bool)`

GetDeliverSuccessCountOk returns a tuple with the DeliverSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverSuccessCount

`func (o *SyncLoadPhase) SetDeliverSuccessCount(v string)`

SetDeliverSuccessCount sets DeliverSuccessCount field to given value.


### GetDeliverFailureCount

`func (o *SyncLoadPhase) GetDeliverFailureCount() string`

GetDeliverFailureCount returns the DeliverFailureCount field if non-nil, zero value otherwise.

### GetDeliverFailureCountOk

`func (o *SyncLoadPhase) GetDeliverFailureCountOk() (*string, bool)`

GetDeliverFailureCountOk returns a tuple with the DeliverFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverFailureCount

`func (o *SyncLoadPhase) SetDeliverFailureCount(v string)`

SetDeliverFailureCount sets DeliverFailureCount field to given value.


### GetErrorCode

`func (o *SyncLoadPhase) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyncLoadPhase) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SyncLoadPhase) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetStartedAt

`func (o *SyncLoadPhase) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SyncLoadPhase) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SyncLoadPhase) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *SyncLoadPhase) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SyncLoadPhase) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SyncLoadPhase) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


