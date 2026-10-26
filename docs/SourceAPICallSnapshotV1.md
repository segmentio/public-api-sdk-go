# SourceAPICallSnapshotV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | The Source id. | 
**ApiCalls** | **string** | A bigint of the number of API calls in this snapshot. | 
**Timestamp** | **string** | Timestamp of this snapshot within the billing cycle in the ISO-8601 format. | 

## Methods

### NewSourceAPICallSnapshotV1

`func NewSourceAPICallSnapshotV1(sourceId string, apiCalls string, timestamp string, ) *SourceAPICallSnapshotV1`

NewSourceAPICallSnapshotV1 instantiates a new SourceAPICallSnapshotV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAPICallSnapshotV1WithDefaults

`func NewSourceAPICallSnapshotV1WithDefaults() *SourceAPICallSnapshotV1`

NewSourceAPICallSnapshotV1WithDefaults instantiates a new SourceAPICallSnapshotV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *SourceAPICallSnapshotV1) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceAPICallSnapshotV1) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceAPICallSnapshotV1) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetApiCalls

`func (o *SourceAPICallSnapshotV1) GetApiCalls() string`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *SourceAPICallSnapshotV1) GetApiCallsOk() (*string, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *SourceAPICallSnapshotV1) SetApiCalls(v string)`

SetApiCalls sets ApiCalls field to given value.


### GetTimestamp

`func (o *SourceAPICallSnapshotV1) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SourceAPICallSnapshotV1) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SourceAPICallSnapshotV1) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


