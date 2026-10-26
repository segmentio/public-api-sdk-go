# APICallSnapshotV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiCalls** | **string** | A bigint of the number of API calls in this snapshot. | 
**Timestamp** | **string** | Timestamp of this snapshot within the billing cycle in the ISO-8601 format. | 

## Methods

### NewAPICallSnapshotV1

`func NewAPICallSnapshotV1(apiCalls string, timestamp string, ) *APICallSnapshotV1`

NewAPICallSnapshotV1 instantiates a new APICallSnapshotV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPICallSnapshotV1WithDefaults

`func NewAPICallSnapshotV1WithDefaults() *APICallSnapshotV1`

NewAPICallSnapshotV1WithDefaults instantiates a new APICallSnapshotV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiCalls

`func (o *APICallSnapshotV1) GetApiCalls() string`

GetApiCalls returns the ApiCalls field if non-nil, zero value otherwise.

### GetApiCallsOk

`func (o *APICallSnapshotV1) GetApiCallsOk() (*string, bool)`

GetApiCallsOk returns a tuple with the ApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiCalls

`func (o *APICallSnapshotV1) SetApiCalls(v string)`

SetApiCalls sets ApiCalls field to given value.


### GetTimestamp

`func (o *APICallSnapshotV1) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *APICallSnapshotV1) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *APICallSnapshotV1) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


